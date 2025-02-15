/**
 * Copyright (c) F5, Inc.
 *
 * This source code is licensed under the Apache License, Version 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package sdk

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	filesSDK "github.com/nginx/agent/sdk/v2/files"
	"github.com/nginx/agent/sdk/v2/proto"
	"github.com/nginx/agent/sdk/v2/zip"

	crossplane "github.com/nginxinc/nginx-go-crossplane"
	log "github.com/sirupsen/logrus"
)

type DirectoryMap struct {
	paths map[string]*proto.Directory
}

func newDirectoryMap() *DirectoryMap {
	return &DirectoryMap{make(map[string]*proto.Directory)}
}

func (dm DirectoryMap) addDirectory(dir string) error {
	_, ok := dm.paths[dir]
	if !ok {
		info, err := os.Stat(dir)
		if err != nil {
			return fmt.Errorf("configs: could not read dir info(%s): %s", dir, err)
		}

		directory := &proto.Directory{
			Name:        dir,
			Mtime:       filesSDK.TimeConvert(info.ModTime()),
			Permissions: filesSDK.GetPermissions(info.Mode()),
			Size_:       info.Size(),
			Files:       make([]*proto.File, 0),
		}

		dm.paths[dir] = directory
	}
	return nil
}

func (dm DirectoryMap) appendFile(dir string, info fs.FileInfo) error {
	lineCount, err := filesSDK.GetLineCount(filepath.Join(dir, info.Name()))
	if err != nil {
		log.Debugf("Failed to get line count: %v", err)
	}

	fileProto := &proto.File{
		Name:        info.Name(),
		Lines:       int32(lineCount),
		Mtime:       filesSDK.TimeConvert(info.ModTime()),
		Permissions: filesSDK.GetPermissions(info.Mode()),
		Size_:       info.Size(),
	}

	return dm.appendFileWithProto(dir, fileProto)
}

func (dm DirectoryMap) appendFileWithProto(dir string, fileProto *proto.File) error {
	_, ok := dm.paths[dir]
	if !ok {
		err := dm.addDirectory(dir)

		if err != nil {
			return err
		}
	}

	dm.paths[dir].Files = append(dm.paths[dir].Files, fileProto)

	return nil
}

// GetNginxConfig parse the configFile into proto.NginxConfig payload, using the provided nginxID, and systemID for
// ConfigDescriptor in the NginxConfig. The allowedDirectories is used to allowlist the directories we include
// in the aux payload.
func GetNginxConfig(
	confFile,
	nginxId,
	systemId string,
	allowedDirectories map[string]struct{},
) (*proto.NginxConfig, error) {
	payload, err := crossplane.Parse(confFile,
		&crossplane.ParseOptions{
			SingleFile:         false,
			StopParsingOnError: true,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error reading config from %s, error: %s", confFile, err)
	}

	nginxConfig := &proto.NginxConfig{
		Action: proto.NginxConfigAction_RETURN,
		ConfigData: &proto.ConfigDescriptor{
			NginxId:  nginxId,
			SystemId: systemId,
		},
		Zconfig:      nil,
		Zaux:         nil,
		AccessLogs:   &proto.AccessLogs{AccessLog: make([]*proto.AccessLog, 0)},
		ErrorLogs:    &proto.ErrorLogs{ErrorLog: make([]*proto.ErrorLog, 0)},
		Ssl:          &proto.SslCertificates{SslCerts: make([]*proto.SslCertificate, 0)},
		DirectoryMap: &proto.DirectoryMap{Directories: make([]*proto.Directory, 0)},
	}

	err = updateNginxConfigFromPayload(confFile, payload, nginxConfig, allowedDirectories)
	if err != nil {
		return nil, fmt.Errorf("error assemble payload from %s, error: %s", confFile, err)
	}

	return nginxConfig, nil
}

// updateNginxConfigFromPayload updates config files from payload.
func updateNginxConfigFromPayload(
	confFile string,
	payload *crossplane.Payload,
	nginxConfig *proto.NginxConfig,
	allowedDirectories map[string]struct{},
) error {
	conf, err := zip.NewWriter(filepath.Dir(confFile))
	if err != nil {
		return fmt.Errorf("configs: could not create zip writer: %s", err)
	}
	aux, err := zip.NewWriter(filepath.Dir(confFile))
	if err != nil {
		return fmt.Errorf("configs: could not create auxillary zip writer: %s", err)
	}

	// cache the directory map, so we can look up using the base
	directoryMap := newDirectoryMap()
	formatMap := map[string]string{}  // map of accessLog/errorLog formats
	seen := make(map[string]struct{}) // local cache of seen files

	// Add files to the zipped config in a consistent order.
	if err = conf.AddFile(payload.Config[0].File); err != nil {
		return fmt.Errorf("configs: could not add conf(%s): %v", payload.Config[0].File, err)
	}

	rest := make([]crossplane.Config, len(payload.Config[1:]))
	copy(rest, payload.Config[1:])
	sort.Slice(rest, func(i, j int) bool {
		return rest[i].File < rest[j].File
	})
	for _, xpConf := range rest {
		if err = conf.AddFile(xpConf.File); err != nil {
			return fmt.Errorf("configs could not add conf file to archive: %s", err)
		}
	}

	// all files in the payload are config files
	var info fs.FileInfo
	for _, xpConf := range payload.Config {
		base := filepath.Dir(xpConf.File)

		info, err = os.Stat(xpConf.File)
		if err != nil {
			return fmt.Errorf("configs: could not read file info(%s): %s", xpConf.File, err)
		}

		if err := directoryMap.appendFile(base, info); err != nil {
			return err
		}

		err = updateNginxConfigFileConfig(xpConf, nginxConfig, filepath.Dir(confFile), aux, formatMap, seen, allowedDirectories, directoryMap)
		if err != nil {
			return fmt.Errorf("configs: failed to update nginx config: %s", err)
		}
	}

	nginxConfig.Zconfig, err = conf.Proto()
	if err != nil {
		return fmt.Errorf("configs: failed to get conf proto: %s", err)
	}

	if aux.FileLen() > 0 {
		nginxConfig.Zaux, err = aux.Proto()
		if err != nil {
			return fmt.Errorf("configs: failed to get aux proto: %s", err)
		}
	}

	setDirectoryMap(directoryMap, nginxConfig)

	return nil
}

func setDirectoryMap(directories *DirectoryMap, nginxConfig *proto.NginxConfig) {
	// empty the DirectoryMap first
	nginxConfig.DirectoryMap.Directories = nginxConfig.DirectoryMap.Directories[:0]
	keys := make([]string, 0, len(directories.paths))
	for k := range directories.paths {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		nginxConfig.DirectoryMap.Directories = append(nginxConfig.DirectoryMap.Directories, directories.paths[k])
	}
}

func updateNginxConfigFileConfig(
	conf crossplane.Config,
	nginxConfig *proto.NginxConfig,
	hostDir string,
	aux *zip.Writer,
	formatMap map[string]string,
	seen map[string]struct{},
	allowedDirectories map[string]struct{},
	directoryMap *DirectoryMap,
) error {
	err := CrossplaneConfigTraverse(&conf,
		func(parent *crossplane.Directive, directive *crossplane.Directive) (bool, error) {
			switch directive.Directive {
			case "log_format":
				if len(directive.Args) >= 2 {
					formatMap[directive.Args[0]] = strings.Join(directive.Args[1:], "")
				}
			case "root":
				if err := updateNginxConfigFileWithRoot(aux, directive.Args[0], seen, allowedDirectories, directoryMap); err != nil {
					return true, err
				}
			case "ssl_certificate", "proxy_ssl_certificate", "ssl_trusted_certificate":
				if err := updateNginxConfigWithCert(directive.Directive, directive.Args[0], nginxConfig, aux, hostDir, directoryMap, allowedDirectories); err != nil {
					return true, err
				}
			case "access_log":
				updateNginxConfigWithAccessLog(
					directive.Args[0],
					getAccessLogDirectiveFormat(directive),
					nginxConfig, formatMap, seen)
			case "error_log":
				updateNginxConfigWithErrorLog(
					directive.Args[0],
					getErrorLogDirectiveLevel(directive),
					nginxConfig, seen)
			}
			return true, nil
		})
	if err != nil {
		return err
	}
	return nil
}

func updateNginxConfigWithCert(
	directive string,
	file string,
	nginxConfig *proto.NginxConfig,
	aux *zip.Writer,
	rootDir string,
	directoryMap *DirectoryMap,
	allowedDirectories map[string]struct{},
) error {
	if strings.HasPrefix("$", file) {
		// variable loading, not actual cert file
		return nil
	}

	if !filepath.IsAbs(file) {
		file = filepath.Join(rootDir, file)
	}
	info, err := os.Stat(file)
	if err != nil {
		return err
	}

	isAllowed := false
	for dir := range allowedDirectories {
		if strings.HasPrefix(file, dir) {
			isAllowed = true
			break
		}
	}

	if directive == "ssl_certificate" || directive == "proxy_ssl_certificate" {
		cert, err := LoadCertificate(file)
		if err != nil {
			return fmt.Errorf("configs: could not load cert(%s): %s", file, err)
		}

		fingerprint := sha256.Sum256(cert.Raw)
		certProto := &proto.SslCertificate{
			FileName:           file,
			PublicKeyAlgorithm: cert.PublicKeyAlgorithm.String(),
			SignatureAlgorithm: cert.SignatureAlgorithm.String(),
			Issuer: &proto.CertificateName{
				CommonName:         cert.Issuer.CommonName,
				Country:            cert.Issuer.Country,
				Locality:           cert.Issuer.Locality,
				Organization:       cert.Issuer.Organization,
				OrganizationalUnit: cert.Issuer.OrganizationalUnit,
			},
			Subject: &proto.CertificateName{
				CommonName:         cert.Subject.CommonName,
				Country:            cert.Subject.Country,
				Locality:           cert.Subject.Locality,
				Organization:       cert.Subject.Organization,
				OrganizationalUnit: cert.Subject.OrganizationalUnit,
				State:              cert.Subject.Province,
			},
			Validity: &proto.CertificateDates{
				NotBefore: cert.NotBefore.Unix(),
				NotAfter:  cert.NotAfter.Unix(),
			},
			SubjAltNames:           cert.DNSNames,
			SerialNumber:           cert.SerialNumber.String(),
			OcspUrl:                cert.IssuingCertificateURL,
			SubjectKeyIdentifier:   convertToHexFormat(hex.EncodeToString(cert.SubjectKeyId)),
			Fingerprint:            convertToHexFormat(hex.EncodeToString(fingerprint[:])),
			FingerprintAlgorithm:   cert.SignatureAlgorithm.String(),
			Version:                int64(cert.Version),
			AuthorityKeyIdentifier: convertToHexFormat(hex.EncodeToString(cert.AuthorityKeyId)),
		}
		certProto.Mtime = filesSDK.TimeConvert(info.ModTime())
		certProto.Size_ = info.Size()

		nginxConfig.Ssl.SslCerts = append(nginxConfig.Ssl.SslCerts, certProto)
	}

	if !isAllowed {
		log.Infof("certs: %s outside allowed directory, not including in aux payloads", file)
		// we want the meta information, but skip putting the files into the aux contents
		return nil
	}
	if err := directoryMap.appendFile(filepath.Dir(file), info); err != nil {
		return err
	}

	if err := aux.AddFile(file); err != nil {
		return fmt.Errorf("configs: could not add cert to aux file writer: %s", err)
	}

	return nil
}

func getAccessLogDirectiveFormat(directive *crossplane.Directive) string {
	var format string
	if len(directive.Args) >= 2 {
		format = strings.ReplaceAll(directive.Args[1], "$", "")
	}
	return format
}

func getErrorLogDirectiveLevel(directive *crossplane.Directive) string {
	if len(directive.Args) >= 2 {
		return directive.Args[1]
	}
	return ""
}

func updateNginxConfigWithAccessLog(file string, format string, nginxConfig *proto.NginxConfig, formatMap map[string]string, seen map[string]struct{}) {
	if _, ok := seen[file]; ok {
		return
	}

	al := &proto.AccessLog{
		Name:     file,
		Readable: false,
	}

	info, err := os.Stat(file)
	if err == nil {
		// survivable error
		al.Readable = true
		al.Permissions = filesSDK.GetPermissions(info.Mode())
	}

	if formatMap[format] != "" {
		al.Format = formatMap[format]
	} else {
		al.Format = format
	}

	nginxConfig.AccessLogs.AccessLog = append(nginxConfig.AccessLogs.AccessLog, al)
	seen[file] = struct{}{}
}

func updateNginxConfigWithAccessLogPath(file string, nginxConfig *proto.NginxConfig, seen map[string]struct{}) {
	if _, ok := seen[file]; ok {
		return
	}
	al := &proto.AccessLog{
		Name: file,
	}

	nginxConfig.AccessLogs.AccessLog = append(nginxConfig.AccessLogs.AccessLog, al)
	seen[file] = struct{}{}
}

func updateNginxConfigWithErrorLog(
	file string,
	level string,
	nginxConfig *proto.NginxConfig,
	seen map[string]struct{},
) {
	if _, ok := seen[file]; ok {
		return
	}
	el := &proto.ErrorLog{
		Name:     file,
		LogLevel: level,
		Readable: false,
	}
	info, err := os.Stat(file)
	if err == nil {
		// survivable error
		el.Permissions = filesSDK.GetPermissions(info.Mode())
		el.Readable = true
	}

	nginxConfig.ErrorLogs.ErrorLog = append(nginxConfig.ErrorLogs.ErrorLog, el)
	seen[file] = struct{}{}
}

func updateNginxConfigWithErrorLogPath(
	file string,
	nginxConfig *proto.NginxConfig,
	seen map[string]struct{},
) {
	if _, ok := seen[file]; ok {
		return
	}
	el := &proto.ErrorLog{
		Name: file,
	}
	nginxConfig.ErrorLogs.ErrorLog = append(nginxConfig.ErrorLogs.ErrorLog, el)
	seen[file] = struct{}{}
}

// root directive, so we slurp up all the files in the directory
func updateNginxConfigFileWithRoot(
	aux *zip.Writer,
	dir string,
	seen map[string]struct{},
	allowedDirectories map[string]struct{},
	directoryMap *DirectoryMap,
) error {
	if _, ok := seen[dir]; ok {
		return nil
	}
	seen[dir] = struct{}{}
	if !allowedPath(dir, allowedDirectories) {
		log.Debugf("Directory %s, is not in the allowed directory list so it will be excluded. Please add the directory to config_dirs in nginx-agent.conf", dir)
		return nil
	}

	return filepath.WalkDir(dir,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if _, ok := seen[path]; ok {
				return nil
			}
			seen[path] = struct{}{}

			if d.IsDir() {
				if err := directoryMap.addDirectory(path); err != nil {
					return err
				}
				return nil
			}

			var info fs.FileInfo
			info, err = d.Info()
			if err != nil {
				return err
			}
			reader, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("could read file(%s): %s", path, err)
			}
			defer reader.Close()

			if err := directoryMap.appendFile(filepath.Dir(path), info); err != nil {
				return err
			}

			if err = aux.Add(path, info.Mode(), reader); err != nil {
				return fmt.Errorf("adding auxillary file error: %s", err)
			}

			return nil
		},
	)
}

func updateNginxConfigFileWithAuxFile(
	aux *zip.Writer,
	file string,
	config *proto.NginxConfig,
	seen map[string]struct{},
	allowedDirectories map[string]struct{},
	directoryMap *DirectoryMap,
	okIfFileNotExist bool,
) error {
	if _, ok := seen[file]; ok {
		return nil
	}
	if !allowedPath(file, allowedDirectories) {
		log.Warnf("Unable to retrieve the NAP aux file %s as it is not in the allowed directory list. Please add the directory to config_dirs in nginx-agent.conf.", file)
		return nil
	}

	info, err := os.Stat(file)
	if err != nil {
		if okIfFileNotExist {
			log.Debugf("Unable to retrieve the aux file %s.", file)
			return nil
		} else {
			return err
		}
	}

	if err := directoryMap.appendFile(filepath.Dir(file), info); err != nil {
		return err
	}

	if err := aux.AddFile(file); err != nil {
		return err
	}
	seen[file] = struct{}{}
	return nil
}

func GetNginxConfigFiles(config *proto.NginxConfig) (confFiles, auxFiles []*proto.File, err error) {
	if config.GetZconfig() == nil {
		return nil, nil, errors.New("config is empty")
	}

	confFiles, err = zip.UnPack(config.GetZconfig())
	if err != nil {
		return nil, nil, fmt.Errorf("unpack zipped config error: %s", err)
	}

	if aux := config.GetZaux(); aux != nil && len(aux.Contents) > 0 {
		auxFiles, err = zip.UnPack(aux)
		if err != nil {
			return nil, nil, fmt.Errorf("unpack zipped auxiliary error: %s", err)
		}
	}
	return confFiles, auxFiles, nil
}

// AddAuxfileToNginxConfig adds the specified newAuxFile to the Nginx Config cfg
func AddAuxfileToNginxConfig(
	confFile string,
	cfg *proto.NginxConfig,
	newAuxFile string,
	allowedDirectories map[string]struct{},
	okIfFileNotExist bool,
) (*proto.NginxConfig, error) {
	directoryMap := newDirectoryMap()
	for _, d := range cfg.DirectoryMap.Directories {
		for _, f := range d.Files {
			err := directoryMap.appendFileWithProto(d.Name, f)
			if err != nil {
				return nil, err
			}
		}
	}

	_, auxFiles, err := GetNginxConfigFiles(cfg)
	if err != nil {
		return nil, err
	}

	aux, err := zip.NewWriter(filepath.Dir(confFile))
	if err != nil {
		return nil, fmt.Errorf("configs: could not create auxillary zip writer: %s", err)
	}

	seen := make(map[string]struct{})
	for _, file := range auxFiles {
		seen[file.Name] = struct{}{}
		err = aux.AddFile(file.Name)
		if err != nil {
			return nil, err
		}
	}

	// add the aux file
	err = updateNginxConfigFileWithAuxFile(aux, newAuxFile, cfg, seen, allowedDirectories, directoryMap, okIfFileNotExist)
	if err != nil {
		return nil, fmt.Errorf("configs: failed to update nginx app protect metadata file: %s", err)
	}

	if aux.FileLen() > 0 {
		cfg.Zaux, err = aux.Proto()
		if err != nil {
			log.Errorf("configs: failed to get aux proto: %s", err)
			return nil, err
		}
	}

	setDirectoryMap(directoryMap, cfg)

	return cfg, nil
}

const (
	plusAPIDirective = "api"
	ossAPIDirective  = "stub_status"
	apiFormat        = "http://%s%s"
)

func parseStatusAPIEndpoints(parent *crossplane.Directive, current *crossplane.Directive) ([]string, []string) {
	var plusUrls []string
	var ossUrls []string
	// process from the location block
	if current.Directive != "location" {
		return plusUrls, ossUrls
	}

	for _, locChild := range current.Block {
		if locChild.Directive != plusAPIDirective && locChild.Directive != ossAPIDirective {
			continue
		}
		host := parseServerHost(parent)
		path := parseLocationPath(current)
		switch locChild.Directive {
		case plusAPIDirective:
			plusUrls = append(plusUrls, fmt.Sprintf(apiFormat, host, path))
		case ossAPIDirective:
			ossUrls = append(ossUrls, fmt.Sprintf(apiFormat, host, path))
		}
	}
	return plusUrls, ossUrls
}

func parseServerHost(parent *crossplane.Directive) string {
	listenPort := "80"
	serverName := "localhost"
	for _, dir := range parent.Block {
		switch dir.Directive {
		case "listen":
			host, port, err := net.SplitHostPort(dir.Args[0])
			if err == nil {
				if host != "*" && host != "::" {
					serverName = host
				}
				listenPort = port
			} else {
				if isPort(dir.Args[0]) {
					listenPort = dir.Args[0]
				} else {
					serverName = dir.Args[0]
				}
			}
		case "server_name":
			if dir.Args[0] == "_" {
				// default server
				continue
			}
			serverName = dir.Args[0]
		}
	}
	return fmt.Sprintf("%s:%s", serverName, listenPort)
}

func isPort(value string) bool {
	port, err := strconv.Atoi(value)
	return err == nil && port >= 1 && port <= 65535
}

func parseLocationPath(location *crossplane.Directive) string {
	path := "/"
	if len(location.Args) > 0 {
		path = location.Args[0]
	}
	return path
}

func statusAPICallback(parent *crossplane.Directive, current *crossplane.Directive) string {
	plusUrls, ossUrls := parseStatusAPIEndpoints(parent, current)

	for _, url := range plusUrls {
		if pingStatusAPIEndpoint(url) {
			log.Debugf("api at %q found", url)
			return url
		}
		log.Debugf("api at %q is not reachable", url)
	}

	for _, url := range ossUrls {
		if pingStatusAPIEndpoint(url) {
			log.Debugf("stub_status at %q found", url)
			return url
		}
		log.Debugf("stub_status at %q is not reachable", url)
	}

	return ""
}

// pingStatusAPIEndpoint ensures the statusAPI is reachable from the agent
func pingStatusAPIEndpoint(statusAPI string) bool {
	client := http.Client{Timeout: 50 * time.Millisecond}

	if _, err := client.Head(statusAPI); err != nil {
		return false
	}
	return true
}

func GetStatusApiInfo(confFile string) (statusApi string, err error) {
	payload, err := crossplane.Parse(confFile,
		&crossplane.ParseOptions{
			SingleFile:         false,
			StopParsingOnError: true,
			CombineConfigs:     true,
		},
	)
	if err != nil {
		return "", fmt.Errorf("error reading config from %s, error: %s", confFile, err)
	}

	for _, xpConf := range payload.Config {
		statusApi = CrossplaneConfigTraverseStr(&xpConf, statusAPICallback)
		if statusApi != "" {
			return statusApi, nil
		}
	}
	return "", errors.New("no status api reachable from the agent found")
}

func GetErrorAndAccessLogs(confFile string) (*proto.ErrorLogs, *proto.AccessLogs, error) {
	nginxConfig := &proto.NginxConfig{
		Action:       proto.NginxConfigAction_RETURN,
		ConfigData:   nil,
		Zconfig:      nil,
		Zaux:         nil,
		AccessLogs:   &proto.AccessLogs{AccessLog: make([]*proto.AccessLog, 0)},
		ErrorLogs:    &proto.ErrorLogs{ErrorLog: make([]*proto.ErrorLog, 0)},
		Ssl:          &proto.SslCertificates{SslCerts: make([]*proto.SslCertificate, 0)},
		DirectoryMap: &proto.DirectoryMap{Directories: make([]*proto.Directory, 0)},
	}

	payload, err := crossplane.Parse(confFile,
		&crossplane.ParseOptions{
			SingleFile:         false,
			StopParsingOnError: true,
		},
	)
	if err != nil {
		return nginxConfig.ErrorLogs, nginxConfig.AccessLogs, err
	}

	seen := make(map[string]struct{})
	for _, xpConf := range payload.Config {
		var err error
		err = CrossplaneConfigTraverse(&xpConf,
			func(parent *crossplane.Directive, current *crossplane.Directive) (bool, error) {
				switch current.Directive {
				case "access_log":
					updateNginxConfigWithAccessLogPath(current.Args[0], nginxConfig, seen)
				case "error_log":
					updateNginxConfigWithErrorLogPath(current.Args[0], nginxConfig, seen)
				}
				return true, nil
			})
		return nginxConfig.ErrorLogs, nginxConfig.AccessLogs, err
	}
	return nginxConfig.ErrorLogs, nginxConfig.AccessLogs, err
}

func GetErrorLogs(errorLogs *proto.ErrorLogs) []string {
	result := []string{}
	for _, log := range errorLogs.ErrorLog {
		result = append(result, log.Name)
	}
	return result
}

func GetAccessLogs(accessLogs *proto.AccessLogs) []string {
	result := []string{}
	for _, log := range accessLogs.AccessLog {
		result = append(result, log.Name)
	}
	return result
}

// allowedPath return true if the provided path has a prefix in the allowedDirectories, false otherwise. The
// path could be a filepath or directory.
func allowedPath(path string, allowedDirectories map[string]struct{}) bool {
	for d := range allowedDirectories {
		if strings.HasPrefix(path, d) {
			return true
		}
	}
	return false
}

func convertToHexFormat(hexString string) string {
	hexString = strings.ToUpper(hexString)
	formatted := ""
	for i := 0; i < len(hexString); i++ {
		if i > 0 && i%2 == 0 {
			formatted += ":"
		}
		formatted += string(hexString[i])
	}
	return formatted
}

func GetAppProtectPolicyAndSecurityLogFiles(cfg *proto.NginxConfig) ([]string, []string) {
	policyMap := make(map[string]bool)
	profileMap := make(map[string]bool)

	for _, directory := range cfg.GetDirectoryMap().GetDirectories() {
		for _, file := range directory.GetFiles() {
			confFile := path.Join(directory.GetName(), file.GetName())

			payload, err := crossplane.Parse(confFile,
				&crossplane.ParseOptions{
					SingleFile:         false,
					StopParsingOnError: true,
				},
			)

			if err != nil {
				continue
			}

			for _, conf := range payload.Config {
				err = CrossplaneConfigTraverse(&conf,
					func(parent *crossplane.Directive, directive *crossplane.Directive) (bool, error) {
						switch directive.Directive {
						case "app_protect_policy_file":
							if len(directive.Args) == 1 {
								_, policy := path.Split(directive.Args[0])
								policyMap[policy] = true
							}
						case "app_protect_security_log":
							if len(directive.Args) == 2 {
								_, profile := path.Split(directive.Args[0])
								profileMap[profile] = true
							}
						}
						return true, nil
					})
				if err != nil {
					continue
				}
			}
			if err != nil {
				continue
			}
		}
	}
	policies := []string{}
	for policy := range policyMap {
		policies = append(policies, policy)
	}

	profiles := []string{}
	for profile := range profileMap {
		profiles = append(profiles, profile)
	}

	return policies, profiles
}

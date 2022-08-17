package logger

import (
	"io"
	"os"
	"path"

	log "github.com/sirupsen/logrus"
)

const (
	defaultLogDir  = "/var/log/nginx-agent"
	defaultLogFile = "agent.log"
)

// SetLogLevel -
func SetLogLevel(level string) {
	if level == "" {
		return
	}
	l, err := log.ParseLevel(level)
	if err != nil {
		log.Errorf("Failed to set log level: %v", err)
		return
	}
	log.SetLevel(l)
	log.Warnf("Log level is %s", l)
}

// SetLogFile returns a file descriptor for the log file that must be handled by the caller
func SetLogFile(logFile string) *os.File {
	logPath := logFile
	if logFile == "" {
		logPath = path.Join(defaultLogDir, defaultLogFile)
	}

	fileInfo, err := os.Stat(logPath)
	if err != nil {
		log.Errorf("error reading log directory %v", err)
		return nil
	}

	if fileInfo.IsDir() {
		// is a directory
		logPath = path.Join(logPath, defaultLogFile)
	}

	logFileHandle, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Errorf("Failed to set log file, proceeding to log only to stdout/stderr: %v", err)
		return nil
	}
	log.SetOutput(io.MultiWriter(os.Stdout, logFileHandle))
	return logFileHandle
}

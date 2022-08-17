// Code generated by counterfeiter. DO NOT EDIT.
package core

import (
	"io/fs"
	"sync"

	"github.com/nginx/agent/sdk/v2/proto"
)

type FakeEnvironment struct {
	DiskDevicesStub        func() ([]string, error)
	diskDevicesMutex       sync.RWMutex
	diskDevicesArgsForCall []struct {
	}
	diskDevicesReturns struct {
		result1 []string
		result2 error
	}
	diskDevicesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	FileStatStub        func(string) (fs.FileInfo, error)
	fileStatMutex       sync.RWMutex
	fileStatArgsForCall []struct {
		arg1 string
	}
	fileStatReturns struct {
		result1 fs.FileInfo
		result2 error
	}
	fileStatReturnsOnCall map[int]struct {
		result1 fs.FileInfo
		result2 error
	}
	GetContainerIDStub        func() (string, error)
	getContainerIDMutex       sync.RWMutex
	getContainerIDArgsForCall []struct {
	}
	getContainerIDReturns struct {
		result1 string
		result2 error
	}
	getContainerIDReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetHostnameStub        func() string
	getHostnameMutex       sync.RWMutex
	getHostnameArgsForCall []struct {
	}
	getHostnameReturns struct {
		result1 string
	}
	getHostnameReturnsOnCall map[int]struct {
		result1 string
	}
	GetNetOverflowStub        func() (float64, error)
	getNetOverflowMutex       sync.RWMutex
	getNetOverflowArgsForCall []struct {
	}
	getNetOverflowReturns struct {
		result1 float64
		result2 error
	}
	getNetOverflowReturnsOnCall map[int]struct {
		result1 float64
		result2 error
	}
	GetSystemUUIDStub        func() string
	getSystemUUIDMutex       sync.RWMutex
	getSystemUUIDArgsForCall []struct {
	}
	getSystemUUIDReturns struct {
		result1 string
	}
	getSystemUUIDReturnsOnCall map[int]struct {
		result1 string
	}
	IsContainerStub        func() bool
	isContainerMutex       sync.RWMutex
	isContainerArgsForCall []struct {
	}
	isContainerReturns struct {
		result1 bool
	}
	isContainerReturnsOnCall map[int]struct {
		result1 bool
	}
	NewHostInfoStub        func(string, *[]string, string, bool) *proto.HostInfo
	newHostInfoMutex       sync.RWMutex
	newHostInfoArgsForCall []struct {
		arg1 string
		arg2 *[]string
		arg3 string
		arg4 bool
	}
	newHostInfoReturns struct {
		result1 *proto.HostInfo
	}
	newHostInfoReturnsOnCall map[int]struct {
		result1 *proto.HostInfo
	}
	ProcessesStub        func() []Process
	processesMutex       sync.RWMutex
	processesArgsForCall []struct {
	}
	processesReturns struct {
		result1 []Process
	}
	processesReturnsOnCall map[int]struct {
		result1 []Process
	}
	ReadDirectoryStub        func(string, string) ([]string, error)
	readDirectoryMutex       sync.RWMutex
	readDirectoryArgsForCall []struct {
		arg1 string
		arg2 string
	}
	readDirectoryReturns struct {
		result1 []string
		result2 error
	}
	readDirectoryReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	WriteFilesStub        func(ConfigApplyMarker, []*proto.File, string, map[string]struct{}) error
	writeFilesMutex       sync.RWMutex
	writeFilesArgsForCall []struct {
		arg1 ConfigApplyMarker
		arg2 []*proto.File
		arg3 string
		arg4 map[string]struct{}
	}
	writeFilesReturns struct {
		result1 error
	}
	writeFilesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEnvironment) DiskDevices() ([]string, error) {
	fake.diskDevicesMutex.Lock()
	ret, specificReturn := fake.diskDevicesReturnsOnCall[len(fake.diskDevicesArgsForCall)]
	fake.diskDevicesArgsForCall = append(fake.diskDevicesArgsForCall, struct {
	}{})
	stub := fake.DiskDevicesStub
	fakeReturns := fake.diskDevicesReturns
	fake.recordInvocation("DiskDevices", []interface{}{})
	fake.diskDevicesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEnvironment) DiskDevicesCallCount() int {
	fake.diskDevicesMutex.RLock()
	defer fake.diskDevicesMutex.RUnlock()
	return len(fake.diskDevicesArgsForCall)
}

func (fake *FakeEnvironment) DiskDevicesCalls(stub func() ([]string, error)) {
	fake.diskDevicesMutex.Lock()
	defer fake.diskDevicesMutex.Unlock()
	fake.DiskDevicesStub = stub
}

func (fake *FakeEnvironment) DiskDevicesReturns(result1 []string, result2 error) {
	fake.diskDevicesMutex.Lock()
	defer fake.diskDevicesMutex.Unlock()
	fake.DiskDevicesStub = nil
	fake.diskDevicesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvironment) DiskDevicesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.diskDevicesMutex.Lock()
	defer fake.diskDevicesMutex.Unlock()
	fake.DiskDevicesStub = nil
	if fake.diskDevicesReturnsOnCall == nil {
		fake.diskDevicesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.diskDevicesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvironment) FileStat(arg1 string) (fs.FileInfo, error) {
	fake.fileStatMutex.Lock()
	ret, specificReturn := fake.fileStatReturnsOnCall[len(fake.fileStatArgsForCall)]
	fake.fileStatArgsForCall = append(fake.fileStatArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.FileStatStub
	fakeReturns := fake.fileStatReturns
	fake.recordInvocation("FileStat", []interface{}{arg1})
	fake.fileStatMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEnvironment) FileStatCallCount() int {
	fake.fileStatMutex.RLock()
	defer fake.fileStatMutex.RUnlock()
	return len(fake.fileStatArgsForCall)
}

func (fake *FakeEnvironment) FileStatCalls(stub func(string) (fs.FileInfo, error)) {
	fake.fileStatMutex.Lock()
	defer fake.fileStatMutex.Unlock()
	fake.FileStatStub = stub
}

func (fake *FakeEnvironment) FileStatArgsForCall(i int) string {
	fake.fileStatMutex.RLock()
	defer fake.fileStatMutex.RUnlock()
	argsForCall := fake.fileStatArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEnvironment) FileStatReturns(result1 fs.FileInfo, result2 error) {
	fake.fileStatMutex.Lock()
	defer fake.fileStatMutex.Unlock()
	fake.FileStatStub = nil
	fake.fileStatReturns = struct {
		result1 fs.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvironment) FileStatReturnsOnCall(i int, result1 fs.FileInfo, result2 error) {
	fake.fileStatMutex.Lock()
	defer fake.fileStatMutex.Unlock()
	fake.FileStatStub = nil
	if fake.fileStatReturnsOnCall == nil {
		fake.fileStatReturnsOnCall = make(map[int]struct {
			result1 fs.FileInfo
			result2 error
		})
	}
	fake.fileStatReturnsOnCall[i] = struct {
		result1 fs.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvironment) GetContainerID() (string, error) {
	fake.getContainerIDMutex.Lock()
	ret, specificReturn := fake.getContainerIDReturnsOnCall[len(fake.getContainerIDArgsForCall)]
	fake.getContainerIDArgsForCall = append(fake.getContainerIDArgsForCall, struct {
	}{})
	stub := fake.GetContainerIDStub
	fakeReturns := fake.getContainerIDReturns
	fake.recordInvocation("GetContainerID", []interface{}{})
	fake.getContainerIDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEnvironment) GetContainerIDCallCount() int {
	fake.getContainerIDMutex.RLock()
	defer fake.getContainerIDMutex.RUnlock()
	return len(fake.getContainerIDArgsForCall)
}

func (fake *FakeEnvironment) GetContainerIDCalls(stub func() (string, error)) {
	fake.getContainerIDMutex.Lock()
	defer fake.getContainerIDMutex.Unlock()
	fake.GetContainerIDStub = stub
}

func (fake *FakeEnvironment) GetContainerIDReturns(result1 string, result2 error) {
	fake.getContainerIDMutex.Lock()
	defer fake.getContainerIDMutex.Unlock()
	fake.GetContainerIDStub = nil
	fake.getContainerIDReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvironment) GetContainerIDReturnsOnCall(i int, result1 string, result2 error) {
	fake.getContainerIDMutex.Lock()
	defer fake.getContainerIDMutex.Unlock()
	fake.GetContainerIDStub = nil
	if fake.getContainerIDReturnsOnCall == nil {
		fake.getContainerIDReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getContainerIDReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvironment) GetHostname() string {
	fake.getHostnameMutex.Lock()
	ret, specificReturn := fake.getHostnameReturnsOnCall[len(fake.getHostnameArgsForCall)]
	fake.getHostnameArgsForCall = append(fake.getHostnameArgsForCall, struct {
	}{})
	stub := fake.GetHostnameStub
	fakeReturns := fake.getHostnameReturns
	fake.recordInvocation("GetHostname", []interface{}{})
	fake.getHostnameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeEnvironment) GetHostnameCallCount() int {
	fake.getHostnameMutex.RLock()
	defer fake.getHostnameMutex.RUnlock()
	return len(fake.getHostnameArgsForCall)
}

func (fake *FakeEnvironment) GetHostnameCalls(stub func() string) {
	fake.getHostnameMutex.Lock()
	defer fake.getHostnameMutex.Unlock()
	fake.GetHostnameStub = stub
}

func (fake *FakeEnvironment) GetHostnameReturns(result1 string) {
	fake.getHostnameMutex.Lock()
	defer fake.getHostnameMutex.Unlock()
	fake.GetHostnameStub = nil
	fake.getHostnameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) GetHostnameReturnsOnCall(i int, result1 string) {
	fake.getHostnameMutex.Lock()
	defer fake.getHostnameMutex.Unlock()
	fake.GetHostnameStub = nil
	if fake.getHostnameReturnsOnCall == nil {
		fake.getHostnameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getHostnameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) GetNetOverflow() (float64, error) {
	fake.getNetOverflowMutex.Lock()
	ret, specificReturn := fake.getNetOverflowReturnsOnCall[len(fake.getNetOverflowArgsForCall)]
	fake.getNetOverflowArgsForCall = append(fake.getNetOverflowArgsForCall, struct {
	}{})
	stub := fake.GetNetOverflowStub
	fakeReturns := fake.getNetOverflowReturns
	fake.recordInvocation("GetNetOverflow", []interface{}{})
	fake.getNetOverflowMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEnvironment) GetNetOverflowCallCount() int {
	fake.getNetOverflowMutex.RLock()
	defer fake.getNetOverflowMutex.RUnlock()
	return len(fake.getNetOverflowArgsForCall)
}

func (fake *FakeEnvironment) GetNetOverflowCalls(stub func() (float64, error)) {
	fake.getNetOverflowMutex.Lock()
	defer fake.getNetOverflowMutex.Unlock()
	fake.GetNetOverflowStub = stub
}

func (fake *FakeEnvironment) GetNetOverflowReturns(result1 float64, result2 error) {
	fake.getNetOverflowMutex.Lock()
	defer fake.getNetOverflowMutex.Unlock()
	fake.GetNetOverflowStub = nil
	fake.getNetOverflowReturns = struct {
		result1 float64
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvironment) GetNetOverflowReturnsOnCall(i int, result1 float64, result2 error) {
	fake.getNetOverflowMutex.Lock()
	defer fake.getNetOverflowMutex.Unlock()
	fake.GetNetOverflowStub = nil
	if fake.getNetOverflowReturnsOnCall == nil {
		fake.getNetOverflowReturnsOnCall = make(map[int]struct {
			result1 float64
			result2 error
		})
	}
	fake.getNetOverflowReturnsOnCall[i] = struct {
		result1 float64
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvironment) GetSystemUUID() string {
	fake.getSystemUUIDMutex.Lock()
	ret, specificReturn := fake.getSystemUUIDReturnsOnCall[len(fake.getSystemUUIDArgsForCall)]
	fake.getSystemUUIDArgsForCall = append(fake.getSystemUUIDArgsForCall, struct {
	}{})
	stub := fake.GetSystemUUIDStub
	fakeReturns := fake.getSystemUUIDReturns
	fake.recordInvocation("GetSystemUUID", []interface{}{})
	fake.getSystemUUIDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeEnvironment) GetSystemUUIDCallCount() int {
	fake.getSystemUUIDMutex.RLock()
	defer fake.getSystemUUIDMutex.RUnlock()
	return len(fake.getSystemUUIDArgsForCall)
}

func (fake *FakeEnvironment) GetSystemUUIDCalls(stub func() string) {
	fake.getSystemUUIDMutex.Lock()
	defer fake.getSystemUUIDMutex.Unlock()
	fake.GetSystemUUIDStub = stub
}

func (fake *FakeEnvironment) GetSystemUUIDReturns(result1 string) {
	fake.getSystemUUIDMutex.Lock()
	defer fake.getSystemUUIDMutex.Unlock()
	fake.GetSystemUUIDStub = nil
	fake.getSystemUUIDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) GetSystemUUIDReturnsOnCall(i int, result1 string) {
	fake.getSystemUUIDMutex.Lock()
	defer fake.getSystemUUIDMutex.Unlock()
	fake.GetSystemUUIDStub = nil
	if fake.getSystemUUIDReturnsOnCall == nil {
		fake.getSystemUUIDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getSystemUUIDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeEnvironment) IsContainer() bool {
	fake.isContainerMutex.Lock()
	ret, specificReturn := fake.isContainerReturnsOnCall[len(fake.isContainerArgsForCall)]
	fake.isContainerArgsForCall = append(fake.isContainerArgsForCall, struct {
	}{})
	stub := fake.IsContainerStub
	fakeReturns := fake.isContainerReturns
	fake.recordInvocation("IsContainer", []interface{}{})
	fake.isContainerMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeEnvironment) IsContainerCallCount() int {
	fake.isContainerMutex.RLock()
	defer fake.isContainerMutex.RUnlock()
	return len(fake.isContainerArgsForCall)
}

func (fake *FakeEnvironment) IsContainerCalls(stub func() bool) {
	fake.isContainerMutex.Lock()
	defer fake.isContainerMutex.Unlock()
	fake.IsContainerStub = stub
}

func (fake *FakeEnvironment) IsContainerReturns(result1 bool) {
	fake.isContainerMutex.Lock()
	defer fake.isContainerMutex.Unlock()
	fake.IsContainerStub = nil
	fake.isContainerReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeEnvironment) IsContainerReturnsOnCall(i int, result1 bool) {
	fake.isContainerMutex.Lock()
	defer fake.isContainerMutex.Unlock()
	fake.IsContainerStub = nil
	if fake.isContainerReturnsOnCall == nil {
		fake.isContainerReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isContainerReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeEnvironment) NewHostInfo(arg1 string, arg2 *[]string, arg3 string, arg4 bool) *proto.HostInfo {
	fake.newHostInfoMutex.Lock()
	ret, specificReturn := fake.newHostInfoReturnsOnCall[len(fake.newHostInfoArgsForCall)]
	fake.newHostInfoArgsForCall = append(fake.newHostInfoArgsForCall, struct {
		arg1 string
		arg2 *[]string
		arg3 string
		arg4 bool
	}{arg1, arg2, arg3, arg4})
	stub := fake.NewHostInfoStub
	fakeReturns := fake.newHostInfoReturns
	fake.recordInvocation("NewHostInfo", []interface{}{arg1, arg2, arg3, arg4})
	fake.newHostInfoMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeEnvironment) NewHostInfoCallCount() int {
	fake.newHostInfoMutex.RLock()
	defer fake.newHostInfoMutex.RUnlock()
	return len(fake.newHostInfoArgsForCall)
}

func (fake *FakeEnvironment) NewHostInfoCalls(stub func(string, *[]string, string, bool) *proto.HostInfo) {
	fake.newHostInfoMutex.Lock()
	defer fake.newHostInfoMutex.Unlock()
	fake.NewHostInfoStub = stub
}

func (fake *FakeEnvironment) NewHostInfoArgsForCall(i int) (string, *[]string, string, bool) {
	fake.newHostInfoMutex.RLock()
	defer fake.newHostInfoMutex.RUnlock()
	argsForCall := fake.newHostInfoArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeEnvironment) NewHostInfoReturns(result1 *proto.HostInfo) {
	fake.newHostInfoMutex.Lock()
	defer fake.newHostInfoMutex.Unlock()
	fake.NewHostInfoStub = nil
	fake.newHostInfoReturns = struct {
		result1 *proto.HostInfo
	}{result1}
}

func (fake *FakeEnvironment) NewHostInfoReturnsOnCall(i int, result1 *proto.HostInfo) {
	fake.newHostInfoMutex.Lock()
	defer fake.newHostInfoMutex.Unlock()
	fake.NewHostInfoStub = nil
	if fake.newHostInfoReturnsOnCall == nil {
		fake.newHostInfoReturnsOnCall = make(map[int]struct {
			result1 *proto.HostInfo
		})
	}
	fake.newHostInfoReturnsOnCall[i] = struct {
		result1 *proto.HostInfo
	}{result1}
}

func (fake *FakeEnvironment) Processes() []Process {
	fake.processesMutex.Lock()
	ret, specificReturn := fake.processesReturnsOnCall[len(fake.processesArgsForCall)]
	fake.processesArgsForCall = append(fake.processesArgsForCall, struct {
	}{})
	stub := fake.ProcessesStub
	fakeReturns := fake.processesReturns
	fake.recordInvocation("Processes", []interface{}{})
	fake.processesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeEnvironment) ProcessesCallCount() int {
	fake.processesMutex.RLock()
	defer fake.processesMutex.RUnlock()
	return len(fake.processesArgsForCall)
}

func (fake *FakeEnvironment) ProcessesCalls(stub func() []Process) {
	fake.processesMutex.Lock()
	defer fake.processesMutex.Unlock()
	fake.ProcessesStub = stub
}

func (fake *FakeEnvironment) ProcessesReturns(result1 []Process) {
	fake.processesMutex.Lock()
	defer fake.processesMutex.Unlock()
	fake.ProcessesStub = nil
	fake.processesReturns = struct {
		result1 []Process
	}{result1}
}

func (fake *FakeEnvironment) ProcessesReturnsOnCall(i int, result1 []Process) {
	fake.processesMutex.Lock()
	defer fake.processesMutex.Unlock()
	fake.ProcessesStub = nil
	if fake.processesReturnsOnCall == nil {
		fake.processesReturnsOnCall = make(map[int]struct {
			result1 []Process
		})
	}
	fake.processesReturnsOnCall[i] = struct {
		result1 []Process
	}{result1}
}

func (fake *FakeEnvironment) ReadDirectory(arg1 string, arg2 string) ([]string, error) {
	fake.readDirectoryMutex.Lock()
	ret, specificReturn := fake.readDirectoryReturnsOnCall[len(fake.readDirectoryArgsForCall)]
	fake.readDirectoryArgsForCall = append(fake.readDirectoryArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.ReadDirectoryStub
	fakeReturns := fake.readDirectoryReturns
	fake.recordInvocation("ReadDirectory", []interface{}{arg1, arg2})
	fake.readDirectoryMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEnvironment) ReadDirectoryCallCount() int {
	fake.readDirectoryMutex.RLock()
	defer fake.readDirectoryMutex.RUnlock()
	return len(fake.readDirectoryArgsForCall)
}

func (fake *FakeEnvironment) ReadDirectoryCalls(stub func(string, string) ([]string, error)) {
	fake.readDirectoryMutex.Lock()
	defer fake.readDirectoryMutex.Unlock()
	fake.ReadDirectoryStub = stub
}

func (fake *FakeEnvironment) ReadDirectoryArgsForCall(i int) (string, string) {
	fake.readDirectoryMutex.RLock()
	defer fake.readDirectoryMutex.RUnlock()
	argsForCall := fake.readDirectoryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEnvironment) ReadDirectoryReturns(result1 []string, result2 error) {
	fake.readDirectoryMutex.Lock()
	defer fake.readDirectoryMutex.Unlock()
	fake.ReadDirectoryStub = nil
	fake.readDirectoryReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvironment) ReadDirectoryReturnsOnCall(i int, result1 []string, result2 error) {
	fake.readDirectoryMutex.Lock()
	defer fake.readDirectoryMutex.Unlock()
	fake.ReadDirectoryStub = nil
	if fake.readDirectoryReturnsOnCall == nil {
		fake.readDirectoryReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.readDirectoryReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvironment) WriteFiles(arg1 ConfigApplyMarker, arg2 []*proto.File, arg3 string, arg4 map[string]struct{}) error {
	var arg2Copy []*proto.File
	if arg2 != nil {
		arg2Copy = make([]*proto.File, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.writeFilesMutex.Lock()
	ret, specificReturn := fake.writeFilesReturnsOnCall[len(fake.writeFilesArgsForCall)]
	fake.writeFilesArgsForCall = append(fake.writeFilesArgsForCall, struct {
		arg1 ConfigApplyMarker
		arg2 []*proto.File
		arg3 string
		arg4 map[string]struct{}
	}{arg1, arg2Copy, arg3, arg4})
	stub := fake.WriteFilesStub
	fakeReturns := fake.writeFilesReturns
	fake.recordInvocation("WriteFiles", []interface{}{arg1, arg2Copy, arg3, arg4})
	fake.writeFilesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeEnvironment) WriteFilesCallCount() int {
	fake.writeFilesMutex.RLock()
	defer fake.writeFilesMutex.RUnlock()
	return len(fake.writeFilesArgsForCall)
}

func (fake *FakeEnvironment) WriteFilesCalls(stub func(ConfigApplyMarker, []*proto.File, string, map[string]struct{}) error) {
	fake.writeFilesMutex.Lock()
	defer fake.writeFilesMutex.Unlock()
	fake.WriteFilesStub = stub
}

func (fake *FakeEnvironment) WriteFilesArgsForCall(i int) (ConfigApplyMarker, []*proto.File, string, map[string]struct{}) {
	fake.writeFilesMutex.RLock()
	defer fake.writeFilesMutex.RUnlock()
	argsForCall := fake.writeFilesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeEnvironment) WriteFilesReturns(result1 error) {
	fake.writeFilesMutex.Lock()
	defer fake.writeFilesMutex.Unlock()
	fake.WriteFilesStub = nil
	fake.writeFilesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEnvironment) WriteFilesReturnsOnCall(i int, result1 error) {
	fake.writeFilesMutex.Lock()
	defer fake.writeFilesMutex.Unlock()
	fake.WriteFilesStub = nil
	if fake.writeFilesReturnsOnCall == nil {
		fake.writeFilesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeFilesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEnvironment) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.diskDevicesMutex.RLock()
	defer fake.diskDevicesMutex.RUnlock()
	fake.fileStatMutex.RLock()
	defer fake.fileStatMutex.RUnlock()
	fake.getContainerIDMutex.RLock()
	defer fake.getContainerIDMutex.RUnlock()
	fake.getHostnameMutex.RLock()
	defer fake.getHostnameMutex.RUnlock()
	fake.getNetOverflowMutex.RLock()
	defer fake.getNetOverflowMutex.RUnlock()
	fake.getSystemUUIDMutex.RLock()
	defer fake.getSystemUUIDMutex.RUnlock()
	fake.isContainerMutex.RLock()
	defer fake.isContainerMutex.RUnlock()
	fake.newHostInfoMutex.RLock()
	defer fake.newHostInfoMutex.RUnlock()
	fake.processesMutex.RLock()
	defer fake.processesMutex.RUnlock()
	fake.readDirectoryMutex.RLock()
	defer fake.readDirectoryMutex.RUnlock()
	fake.writeFilesMutex.RLock()
	defer fake.writeFilesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEnvironment) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ Environment = new(FakeEnvironment)

// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.ibm.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
)

type LparAPI struct {
	GetLparPropertiesStub        func(string) (*zhmcclient.LparProperties, error)
	getLparPropertiesMutex       sync.RWMutex
	getLparPropertiesArgsForCall []struct {
		arg1 string
	}
	getLparPropertiesReturns struct {
		result1 *zhmcclient.LparProperties
		result2 error
	}
	getLparPropertiesReturnsOnCall map[int]struct {
		result1 *zhmcclient.LparProperties
		result2 error
	}
	ListLPARsStub        func(string, map[string]string) ([]zhmcclient.LPAR, error)
	listLPARsMutex       sync.RWMutex
	listLPARsArgsForCall []struct {
		arg1 string
		arg2 map[string]string
	}
	listLPARsReturns struct {
		result1 []zhmcclient.LPAR
		result2 error
	}
	listLPARsReturnsOnCall map[int]struct {
		result1 []zhmcclient.LPAR
		result2 error
	}
	ListNicsStub        func(string) ([]string, error)
	listNicsMutex       sync.RWMutex
	listNicsArgsForCall []struct {
		arg1 string
	}
	listNicsReturns struct {
		result1 []string
		result2 error
	}
	listNicsReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	MountIsoImageStub        func(string, []byte, string, string) error
	mountIsoImageMutex       sync.RWMutex
	mountIsoImageArgsForCall []struct {
		arg1 string
		arg2 []byte
		arg3 string
		arg4 string
	}
	mountIsoImageReturns struct {
		result1 error
	}
	mountIsoImageReturnsOnCall map[int]struct {
		result1 error
	}
	StartLPARStub        func(string) (string, error)
	startLPARMutex       sync.RWMutex
	startLPARArgsForCall []struct {
		arg1 string
	}
	startLPARReturns struct {
		result1 string
		result2 error
	}
	startLPARReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	StopLPARStub        func(string) (string, error)
	stopLPARMutex       sync.RWMutex
	stopLPARArgsForCall []struct {
		arg1 string
	}
	stopLPARReturns struct {
		result1 string
		result2 error
	}
	stopLPARReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	UnmountIsoImageStub        func(string) error
	unmountIsoImageMutex       sync.RWMutex
	unmountIsoImageArgsForCall []struct {
		arg1 string
	}
	unmountIsoImageReturns struct {
		result1 error
	}
	unmountIsoImageReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateLparPropertiesStub        func(string, *zhmcclient.LparProperties) error
	updateLparPropertiesMutex       sync.RWMutex
	updateLparPropertiesArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.LparProperties
	}
	updateLparPropertiesReturns struct {
		result1 error
	}
	updateLparPropertiesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *LparAPI) GetLparProperties(arg1 string) (*zhmcclient.LparProperties, error) {
	fake.getLparPropertiesMutex.Lock()
	ret, specificReturn := fake.getLparPropertiesReturnsOnCall[len(fake.getLparPropertiesArgsForCall)]
	fake.getLparPropertiesArgsForCall = append(fake.getLparPropertiesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetLparPropertiesStub
	fakeReturns := fake.getLparPropertiesReturns
	fake.recordInvocation("GetLparProperties", []interface{}{arg1})
	fake.getLparPropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *LparAPI) GetLparPropertiesCallCount() int {
	fake.getLparPropertiesMutex.RLock()
	defer fake.getLparPropertiesMutex.RUnlock()
	return len(fake.getLparPropertiesArgsForCall)
}

func (fake *LparAPI) GetLparPropertiesCalls(stub func(string) (*zhmcclient.LparProperties, error)) {
	fake.getLparPropertiesMutex.Lock()
	defer fake.getLparPropertiesMutex.Unlock()
	fake.GetLparPropertiesStub = stub
}

func (fake *LparAPI) GetLparPropertiesArgsForCall(i int) string {
	fake.getLparPropertiesMutex.RLock()
	defer fake.getLparPropertiesMutex.RUnlock()
	argsForCall := fake.getLparPropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *LparAPI) GetLparPropertiesReturns(result1 *zhmcclient.LparProperties, result2 error) {
	fake.getLparPropertiesMutex.Lock()
	defer fake.getLparPropertiesMutex.Unlock()
	fake.GetLparPropertiesStub = nil
	fake.getLparPropertiesReturns = struct {
		result1 *zhmcclient.LparProperties
		result2 error
	}{result1, result2}
}

func (fake *LparAPI) GetLparPropertiesReturnsOnCall(i int, result1 *zhmcclient.LparProperties, result2 error) {
	fake.getLparPropertiesMutex.Lock()
	defer fake.getLparPropertiesMutex.Unlock()
	fake.GetLparPropertiesStub = nil
	if fake.getLparPropertiesReturnsOnCall == nil {
		fake.getLparPropertiesReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.LparProperties
			result2 error
		})
	}
	fake.getLparPropertiesReturnsOnCall[i] = struct {
		result1 *zhmcclient.LparProperties
		result2 error
	}{result1, result2}
}

func (fake *LparAPI) ListLPARs(arg1 string, arg2 map[string]string) ([]zhmcclient.LPAR, error) {
	fake.listLPARsMutex.Lock()
	ret, specificReturn := fake.listLPARsReturnsOnCall[len(fake.listLPARsArgsForCall)]
	fake.listLPARsArgsForCall = append(fake.listLPARsArgsForCall, struct {
		arg1 string
		arg2 map[string]string
	}{arg1, arg2})
	stub := fake.ListLPARsStub
	fakeReturns := fake.listLPARsReturns
	fake.recordInvocation("ListLPARs", []interface{}{arg1, arg2})
	fake.listLPARsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *LparAPI) ListLPARsCallCount() int {
	fake.listLPARsMutex.RLock()
	defer fake.listLPARsMutex.RUnlock()
	return len(fake.listLPARsArgsForCall)
}

func (fake *LparAPI) ListLPARsCalls(stub func(string, map[string]string) ([]zhmcclient.LPAR, error)) {
	fake.listLPARsMutex.Lock()
	defer fake.listLPARsMutex.Unlock()
	fake.ListLPARsStub = stub
}

func (fake *LparAPI) ListLPARsArgsForCall(i int) (string, map[string]string) {
	fake.listLPARsMutex.RLock()
	defer fake.listLPARsMutex.RUnlock()
	argsForCall := fake.listLPARsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *LparAPI) ListLPARsReturns(result1 []zhmcclient.LPAR, result2 error) {
	fake.listLPARsMutex.Lock()
	defer fake.listLPARsMutex.Unlock()
	fake.ListLPARsStub = nil
	fake.listLPARsReturns = struct {
		result1 []zhmcclient.LPAR
		result2 error
	}{result1, result2}
}

func (fake *LparAPI) ListLPARsReturnsOnCall(i int, result1 []zhmcclient.LPAR, result2 error) {
	fake.listLPARsMutex.Lock()
	defer fake.listLPARsMutex.Unlock()
	fake.ListLPARsStub = nil
	if fake.listLPARsReturnsOnCall == nil {
		fake.listLPARsReturnsOnCall = make(map[int]struct {
			result1 []zhmcclient.LPAR
			result2 error
		})
	}
	fake.listLPARsReturnsOnCall[i] = struct {
		result1 []zhmcclient.LPAR
		result2 error
	}{result1, result2}
}

func (fake *LparAPI) ListNics(arg1 string) ([]string, error) {
	fake.listNicsMutex.Lock()
	ret, specificReturn := fake.listNicsReturnsOnCall[len(fake.listNicsArgsForCall)]
	fake.listNicsArgsForCall = append(fake.listNicsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ListNicsStub
	fakeReturns := fake.listNicsReturns
	fake.recordInvocation("ListNics", []interface{}{arg1})
	fake.listNicsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *LparAPI) ListNicsCallCount() int {
	fake.listNicsMutex.RLock()
	defer fake.listNicsMutex.RUnlock()
	return len(fake.listNicsArgsForCall)
}

func (fake *LparAPI) ListNicsCalls(stub func(string) ([]string, error)) {
	fake.listNicsMutex.Lock()
	defer fake.listNicsMutex.Unlock()
	fake.ListNicsStub = stub
}

func (fake *LparAPI) ListNicsArgsForCall(i int) string {
	fake.listNicsMutex.RLock()
	defer fake.listNicsMutex.RUnlock()
	argsForCall := fake.listNicsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *LparAPI) ListNicsReturns(result1 []string, result2 error) {
	fake.listNicsMutex.Lock()
	defer fake.listNicsMutex.Unlock()
	fake.ListNicsStub = nil
	fake.listNicsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *LparAPI) ListNicsReturnsOnCall(i int, result1 []string, result2 error) {
	fake.listNicsMutex.Lock()
	defer fake.listNicsMutex.Unlock()
	fake.ListNicsStub = nil
	if fake.listNicsReturnsOnCall == nil {
		fake.listNicsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listNicsReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *LparAPI) MountIsoImage(arg1 string, arg2 []byte, arg3 string, arg4 string) error {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.mountIsoImageMutex.Lock()
	ret, specificReturn := fake.mountIsoImageReturnsOnCall[len(fake.mountIsoImageArgsForCall)]
	fake.mountIsoImageArgsForCall = append(fake.mountIsoImageArgsForCall, struct {
		arg1 string
		arg2 []byte
		arg3 string
		arg4 string
	}{arg1, arg2Copy, arg3, arg4})
	stub := fake.MountIsoImageStub
	fakeReturns := fake.mountIsoImageReturns
	fake.recordInvocation("MountIsoImage", []interface{}{arg1, arg2Copy, arg3, arg4})
	fake.mountIsoImageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *LparAPI) MountIsoImageCallCount() int {
	fake.mountIsoImageMutex.RLock()
	defer fake.mountIsoImageMutex.RUnlock()
	return len(fake.mountIsoImageArgsForCall)
}

func (fake *LparAPI) MountIsoImageCalls(stub func(string, []byte, string, string) error) {
	fake.mountIsoImageMutex.Lock()
	defer fake.mountIsoImageMutex.Unlock()
	fake.MountIsoImageStub = stub
}

func (fake *LparAPI) MountIsoImageArgsForCall(i int) (string, []byte, string, string) {
	fake.mountIsoImageMutex.RLock()
	defer fake.mountIsoImageMutex.RUnlock()
	argsForCall := fake.mountIsoImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *LparAPI) MountIsoImageReturns(result1 error) {
	fake.mountIsoImageMutex.Lock()
	defer fake.mountIsoImageMutex.Unlock()
	fake.MountIsoImageStub = nil
	fake.mountIsoImageReturns = struct {
		result1 error
	}{result1}
}

func (fake *LparAPI) MountIsoImageReturnsOnCall(i int, result1 error) {
	fake.mountIsoImageMutex.Lock()
	defer fake.mountIsoImageMutex.Unlock()
	fake.MountIsoImageStub = nil
	if fake.mountIsoImageReturnsOnCall == nil {
		fake.mountIsoImageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.mountIsoImageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *LparAPI) StartLPAR(arg1 string) (string, error) {
	fake.startLPARMutex.Lock()
	ret, specificReturn := fake.startLPARReturnsOnCall[len(fake.startLPARArgsForCall)]
	fake.startLPARArgsForCall = append(fake.startLPARArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.StartLPARStub
	fakeReturns := fake.startLPARReturns
	fake.recordInvocation("StartLPAR", []interface{}{arg1})
	fake.startLPARMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *LparAPI) StartLPARCallCount() int {
	fake.startLPARMutex.RLock()
	defer fake.startLPARMutex.RUnlock()
	return len(fake.startLPARArgsForCall)
}

func (fake *LparAPI) StartLPARCalls(stub func(string) (string, error)) {
	fake.startLPARMutex.Lock()
	defer fake.startLPARMutex.Unlock()
	fake.StartLPARStub = stub
}

func (fake *LparAPI) StartLPARArgsForCall(i int) string {
	fake.startLPARMutex.RLock()
	defer fake.startLPARMutex.RUnlock()
	argsForCall := fake.startLPARArgsForCall[i]
	return argsForCall.arg1
}

func (fake *LparAPI) StartLPARReturns(result1 string, result2 error) {
	fake.startLPARMutex.Lock()
	defer fake.startLPARMutex.Unlock()
	fake.StartLPARStub = nil
	fake.startLPARReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *LparAPI) StartLPARReturnsOnCall(i int, result1 string, result2 error) {
	fake.startLPARMutex.Lock()
	defer fake.startLPARMutex.Unlock()
	fake.StartLPARStub = nil
	if fake.startLPARReturnsOnCall == nil {
		fake.startLPARReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.startLPARReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *LparAPI) StopLPAR(arg1 string) (string, error) {
	fake.stopLPARMutex.Lock()
	ret, specificReturn := fake.stopLPARReturnsOnCall[len(fake.stopLPARArgsForCall)]
	fake.stopLPARArgsForCall = append(fake.stopLPARArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.StopLPARStub
	fakeReturns := fake.stopLPARReturns
	fake.recordInvocation("StopLPAR", []interface{}{arg1})
	fake.stopLPARMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *LparAPI) StopLPARCallCount() int {
	fake.stopLPARMutex.RLock()
	defer fake.stopLPARMutex.RUnlock()
	return len(fake.stopLPARArgsForCall)
}

func (fake *LparAPI) StopLPARCalls(stub func(string) (string, error)) {
	fake.stopLPARMutex.Lock()
	defer fake.stopLPARMutex.Unlock()
	fake.StopLPARStub = stub
}

func (fake *LparAPI) StopLPARArgsForCall(i int) string {
	fake.stopLPARMutex.RLock()
	defer fake.stopLPARMutex.RUnlock()
	argsForCall := fake.stopLPARArgsForCall[i]
	return argsForCall.arg1
}

func (fake *LparAPI) StopLPARReturns(result1 string, result2 error) {
	fake.stopLPARMutex.Lock()
	defer fake.stopLPARMutex.Unlock()
	fake.StopLPARStub = nil
	fake.stopLPARReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *LparAPI) StopLPARReturnsOnCall(i int, result1 string, result2 error) {
	fake.stopLPARMutex.Lock()
	defer fake.stopLPARMutex.Unlock()
	fake.StopLPARStub = nil
	if fake.stopLPARReturnsOnCall == nil {
		fake.stopLPARReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.stopLPARReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *LparAPI) UnmountIsoImage(arg1 string) error {
	fake.unmountIsoImageMutex.Lock()
	ret, specificReturn := fake.unmountIsoImageReturnsOnCall[len(fake.unmountIsoImageArgsForCall)]
	fake.unmountIsoImageArgsForCall = append(fake.unmountIsoImageArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.UnmountIsoImageStub
	fakeReturns := fake.unmountIsoImageReturns
	fake.recordInvocation("UnmountIsoImage", []interface{}{arg1})
	fake.unmountIsoImageMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *LparAPI) UnmountIsoImageCallCount() int {
	fake.unmountIsoImageMutex.RLock()
	defer fake.unmountIsoImageMutex.RUnlock()
	return len(fake.unmountIsoImageArgsForCall)
}

func (fake *LparAPI) UnmountIsoImageCalls(stub func(string) error) {
	fake.unmountIsoImageMutex.Lock()
	defer fake.unmountIsoImageMutex.Unlock()
	fake.UnmountIsoImageStub = stub
}

func (fake *LparAPI) UnmountIsoImageArgsForCall(i int) string {
	fake.unmountIsoImageMutex.RLock()
	defer fake.unmountIsoImageMutex.RUnlock()
	argsForCall := fake.unmountIsoImageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *LparAPI) UnmountIsoImageReturns(result1 error) {
	fake.unmountIsoImageMutex.Lock()
	defer fake.unmountIsoImageMutex.Unlock()
	fake.UnmountIsoImageStub = nil
	fake.unmountIsoImageReturns = struct {
		result1 error
	}{result1}
}

func (fake *LparAPI) UnmountIsoImageReturnsOnCall(i int, result1 error) {
	fake.unmountIsoImageMutex.Lock()
	defer fake.unmountIsoImageMutex.Unlock()
	fake.UnmountIsoImageStub = nil
	if fake.unmountIsoImageReturnsOnCall == nil {
		fake.unmountIsoImageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unmountIsoImageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *LparAPI) UpdateLparProperties(arg1 string, arg2 *zhmcclient.LparProperties) error {
	fake.updateLparPropertiesMutex.Lock()
	ret, specificReturn := fake.updateLparPropertiesReturnsOnCall[len(fake.updateLparPropertiesArgsForCall)]
	fake.updateLparPropertiesArgsForCall = append(fake.updateLparPropertiesArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.LparProperties
	}{arg1, arg2})
	stub := fake.UpdateLparPropertiesStub
	fakeReturns := fake.updateLparPropertiesReturns
	fake.recordInvocation("UpdateLparProperties", []interface{}{arg1, arg2})
	fake.updateLparPropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *LparAPI) UpdateLparPropertiesCallCount() int {
	fake.updateLparPropertiesMutex.RLock()
	defer fake.updateLparPropertiesMutex.RUnlock()
	return len(fake.updateLparPropertiesArgsForCall)
}

func (fake *LparAPI) UpdateLparPropertiesCalls(stub func(string, *zhmcclient.LparProperties) error) {
	fake.updateLparPropertiesMutex.Lock()
	defer fake.updateLparPropertiesMutex.Unlock()
	fake.UpdateLparPropertiesStub = stub
}

func (fake *LparAPI) UpdateLparPropertiesArgsForCall(i int) (string, *zhmcclient.LparProperties) {
	fake.updateLparPropertiesMutex.RLock()
	defer fake.updateLparPropertiesMutex.RUnlock()
	argsForCall := fake.updateLparPropertiesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *LparAPI) UpdateLparPropertiesReturns(result1 error) {
	fake.updateLparPropertiesMutex.Lock()
	defer fake.updateLparPropertiesMutex.Unlock()
	fake.UpdateLparPropertiesStub = nil
	fake.updateLparPropertiesReturns = struct {
		result1 error
	}{result1}
}

func (fake *LparAPI) UpdateLparPropertiesReturnsOnCall(i int, result1 error) {
	fake.updateLparPropertiesMutex.Lock()
	defer fake.updateLparPropertiesMutex.Unlock()
	fake.UpdateLparPropertiesStub = nil
	if fake.updateLparPropertiesReturnsOnCall == nil {
		fake.updateLparPropertiesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateLparPropertiesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *LparAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getLparPropertiesMutex.RLock()
	defer fake.getLparPropertiesMutex.RUnlock()
	fake.listLPARsMutex.RLock()
	defer fake.listLPARsMutex.RUnlock()
	fake.listNicsMutex.RLock()
	defer fake.listNicsMutex.RUnlock()
	fake.mountIsoImageMutex.RLock()
	defer fake.mountIsoImageMutex.RUnlock()
	fake.startLPARMutex.RLock()
	defer fake.startLPARMutex.RUnlock()
	fake.stopLPARMutex.RLock()
	defer fake.stopLPARMutex.RUnlock()
	fake.unmountIsoImageMutex.RLock()
	defer fake.unmountIsoImageMutex.RUnlock()
	fake.updateLparPropertiesMutex.RLock()
	defer fake.updateLparPropertiesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *LparAPI) recordInvocation(key string, args []interface{}) {
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

var _ zhmcclient.LparAPI = new(LparAPI)

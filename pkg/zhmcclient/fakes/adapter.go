// Copyright 2021-2023 IBM Corp. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fakes

import (
	"sync"

	"github.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
)

type AdapterAPI struct {
	CreateHipersocketStub        func(string, *zhmcclient.HipersocketPayload) (string, int, *zhmcclient.HmcError)
	createHipersocketMutex       sync.RWMutex
	createHipersocketArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.HipersocketPayload
	}
	createHipersocketReturns struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}
	createHipersocketReturnsOnCall map[int]struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}
	DeleteHipersocketStub        func(string) (int, *zhmcclient.HmcError)
	deleteHipersocketMutex       sync.RWMutex
	deleteHipersocketArgsForCall []struct {
		arg1 string
	}
	deleteHipersocketReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	deleteHipersocketReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	GetAdapterPropertiesStub        func(string) (*zhmcclient.AdapterProperties, int, *zhmcclient.HmcError)
	getAdapterPropertiesMutex       sync.RWMutex
	getAdapterPropertiesArgsForCall []struct {
		arg1 string
	}
	getAdapterPropertiesReturns struct {
		result1 *zhmcclient.AdapterProperties
		result2 int
		result3 *zhmcclient.HmcError
	}
	getAdapterPropertiesReturnsOnCall map[int]struct {
		result1 *zhmcclient.AdapterProperties
		result2 int
		result3 *zhmcclient.HmcError
	}
	GetNetworkAdapterPortPropertiesStub        func(string) (*zhmcclient.NetworkAdapterPort, int, *zhmcclient.HmcError)
	getNetworkAdapterPortPropertiesMutex       sync.RWMutex
	getNetworkAdapterPortPropertiesArgsForCall []struct {
		arg1 string
	}
	getNetworkAdapterPortPropertiesReturns struct {
		result1 *zhmcclient.NetworkAdapterPort
		result2 int
		result3 *zhmcclient.HmcError
	}
	getNetworkAdapterPortPropertiesReturnsOnCall map[int]struct {
		result1 *zhmcclient.NetworkAdapterPort
		result2 int
		result3 *zhmcclient.HmcError
	}
	GetStorageAdapterPortPropertiesStub        func(string) (*zhmcclient.StorageAdapterPort, int, *zhmcclient.HmcError)
	getStorageAdapterPortPropertiesMutex       sync.RWMutex
	getStorageAdapterPortPropertiesArgsForCall []struct {
		arg1 string
	}
	getStorageAdapterPortPropertiesReturns struct {
		result1 *zhmcclient.StorageAdapterPort
		result2 int
		result3 *zhmcclient.HmcError
	}
	getStorageAdapterPortPropertiesReturnsOnCall map[int]struct {
		result1 *zhmcclient.StorageAdapterPort
		result2 int
		result3 *zhmcclient.HmcError
	}
	ListAdaptersStub        func(string, map[string]string) ([]zhmcclient.Adapter, int, *zhmcclient.HmcError)
	listAdaptersMutex       sync.RWMutex
	listAdaptersArgsForCall []struct {
		arg1 string
		arg2 map[string]string
	}
	listAdaptersReturns struct {
		result1 []zhmcclient.Adapter
		result2 int
		result3 *zhmcclient.HmcError
	}
	listAdaptersReturnsOnCall map[int]struct {
		result1 []zhmcclient.Adapter
		result2 int
		result3 *zhmcclient.HmcError
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *AdapterAPI) CreateHipersocket(arg1 string, arg2 *zhmcclient.HipersocketPayload) (string, int, *zhmcclient.HmcError) {
	fake.createHipersocketMutex.Lock()
	ret, specificReturn := fake.createHipersocketReturnsOnCall[len(fake.createHipersocketArgsForCall)]
	fake.createHipersocketArgsForCall = append(fake.createHipersocketArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.HipersocketPayload
	}{arg1, arg2})
	stub := fake.CreateHipersocketStub
	fakeReturns := fake.createHipersocketReturns
	fake.recordInvocation("CreateHipersocket", []interface{}{arg1, arg2})
	fake.createHipersocketMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *AdapterAPI) CreateHipersocketCallCount() int {
	fake.createHipersocketMutex.RLock()
	defer fake.createHipersocketMutex.RUnlock()
	return len(fake.createHipersocketArgsForCall)
}

func (fake *AdapterAPI) CreateHipersocketCalls(stub func(string, *zhmcclient.HipersocketPayload) (string, int, *zhmcclient.HmcError)) {
	fake.createHipersocketMutex.Lock()
	defer fake.createHipersocketMutex.Unlock()
	fake.CreateHipersocketStub = stub
}

func (fake *AdapterAPI) CreateHipersocketArgsForCall(i int) (string, *zhmcclient.HipersocketPayload) {
	fake.createHipersocketMutex.RLock()
	defer fake.createHipersocketMutex.RUnlock()
	argsForCall := fake.createHipersocketArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *AdapterAPI) CreateHipersocketReturns(result1 string, result2 int, result3 *zhmcclient.HmcError) {
	fake.createHipersocketMutex.Lock()
	defer fake.createHipersocketMutex.Unlock()
	fake.CreateHipersocketStub = nil
	fake.createHipersocketReturns = struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *AdapterAPI) CreateHipersocketReturnsOnCall(i int, result1 string, result2 int, result3 *zhmcclient.HmcError) {
	fake.createHipersocketMutex.Lock()
	defer fake.createHipersocketMutex.Unlock()
	fake.CreateHipersocketStub = nil
	if fake.createHipersocketReturnsOnCall == nil {
		fake.createHipersocketReturnsOnCall = make(map[int]struct {
			result1 string
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.createHipersocketReturnsOnCall[i] = struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *AdapterAPI) DeleteHipersocket(arg1 string) (int, *zhmcclient.HmcError) {
	fake.deleteHipersocketMutex.Lock()
	ret, specificReturn := fake.deleteHipersocketReturnsOnCall[len(fake.deleteHipersocketArgsForCall)]
	fake.deleteHipersocketArgsForCall = append(fake.deleteHipersocketArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteHipersocketStub
	fakeReturns := fake.deleteHipersocketReturns
	fake.recordInvocation("DeleteHipersocket", []interface{}{arg1})
	fake.deleteHipersocketMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AdapterAPI) DeleteHipersocketCallCount() int {
	fake.deleteHipersocketMutex.RLock()
	defer fake.deleteHipersocketMutex.RUnlock()
	return len(fake.deleteHipersocketArgsForCall)
}

func (fake *AdapterAPI) DeleteHipersocketCalls(stub func(string) (int, *zhmcclient.HmcError)) {
	fake.deleteHipersocketMutex.Lock()
	defer fake.deleteHipersocketMutex.Unlock()
	fake.DeleteHipersocketStub = stub
}

func (fake *AdapterAPI) DeleteHipersocketArgsForCall(i int) string {
	fake.deleteHipersocketMutex.RLock()
	defer fake.deleteHipersocketMutex.RUnlock()
	argsForCall := fake.deleteHipersocketArgsForCall[i]
	return argsForCall.arg1
}

func (fake *AdapterAPI) DeleteHipersocketReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.deleteHipersocketMutex.Lock()
	defer fake.deleteHipersocketMutex.Unlock()
	fake.DeleteHipersocketStub = nil
	fake.deleteHipersocketReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *AdapterAPI) DeleteHipersocketReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.deleteHipersocketMutex.Lock()
	defer fake.deleteHipersocketMutex.Unlock()
	fake.DeleteHipersocketStub = nil
	if fake.deleteHipersocketReturnsOnCall == nil {
		fake.deleteHipersocketReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.deleteHipersocketReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *AdapterAPI) GetAdapterProperties(arg1 string) (*zhmcclient.AdapterProperties, int, *zhmcclient.HmcError) {
	fake.getAdapterPropertiesMutex.Lock()
	ret, specificReturn := fake.getAdapterPropertiesReturnsOnCall[len(fake.getAdapterPropertiesArgsForCall)]
	fake.getAdapterPropertiesArgsForCall = append(fake.getAdapterPropertiesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetAdapterPropertiesStub
	fakeReturns := fake.getAdapterPropertiesReturns
	fake.recordInvocation("GetAdapterProperties", []interface{}{arg1})
	fake.getAdapterPropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *AdapterAPI) GetAdapterPropertiesCallCount() int {
	fake.getAdapterPropertiesMutex.RLock()
	defer fake.getAdapterPropertiesMutex.RUnlock()
	return len(fake.getAdapterPropertiesArgsForCall)
}

func (fake *AdapterAPI) GetAdapterPropertiesCalls(stub func(string) (*zhmcclient.AdapterProperties, int, *zhmcclient.HmcError)) {
	fake.getAdapterPropertiesMutex.Lock()
	defer fake.getAdapterPropertiesMutex.Unlock()
	fake.GetAdapterPropertiesStub = stub
}

func (fake *AdapterAPI) GetAdapterPropertiesArgsForCall(i int) string {
	fake.getAdapterPropertiesMutex.RLock()
	defer fake.getAdapterPropertiesMutex.RUnlock()
	argsForCall := fake.getAdapterPropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *AdapterAPI) GetAdapterPropertiesReturns(result1 *zhmcclient.AdapterProperties, result2 int, result3 *zhmcclient.HmcError) {
	fake.getAdapterPropertiesMutex.Lock()
	defer fake.getAdapterPropertiesMutex.Unlock()
	fake.GetAdapterPropertiesStub = nil
	fake.getAdapterPropertiesReturns = struct {
		result1 *zhmcclient.AdapterProperties
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *AdapterAPI) GetAdapterPropertiesReturnsOnCall(i int, result1 *zhmcclient.AdapterProperties, result2 int, result3 *zhmcclient.HmcError) {
	fake.getAdapterPropertiesMutex.Lock()
	defer fake.getAdapterPropertiesMutex.Unlock()
	fake.GetAdapterPropertiesStub = nil
	if fake.getAdapterPropertiesReturnsOnCall == nil {
		fake.getAdapterPropertiesReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.AdapterProperties
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.getAdapterPropertiesReturnsOnCall[i] = struct {
		result1 *zhmcclient.AdapterProperties
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *AdapterAPI) GetNetworkAdapterPortProperties(arg1 string) (*zhmcclient.NetworkAdapterPort, int, *zhmcclient.HmcError) {
	fake.getNetworkAdapterPortPropertiesMutex.Lock()
	ret, specificReturn := fake.getNetworkAdapterPortPropertiesReturnsOnCall[len(fake.getNetworkAdapterPortPropertiesArgsForCall)]
	fake.getNetworkAdapterPortPropertiesArgsForCall = append(fake.getNetworkAdapterPortPropertiesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetNetworkAdapterPortPropertiesStub
	fakeReturns := fake.getNetworkAdapterPortPropertiesReturns
	fake.recordInvocation("GetNetworkAdapterPortProperties", []interface{}{arg1})
	fake.getNetworkAdapterPortPropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *AdapterAPI) GetNetworkAdapterPortPropertiesCallCount() int {
	fake.getNetworkAdapterPortPropertiesMutex.RLock()
	defer fake.getNetworkAdapterPortPropertiesMutex.RUnlock()
	return len(fake.getNetworkAdapterPortPropertiesArgsForCall)
}

func (fake *AdapterAPI) GetNetworkAdapterPortPropertiesCalls(stub func(string) (*zhmcclient.NetworkAdapterPort, int, *zhmcclient.HmcError)) {
	fake.getNetworkAdapterPortPropertiesMutex.Lock()
	defer fake.getNetworkAdapterPortPropertiesMutex.Unlock()
	fake.GetNetworkAdapterPortPropertiesStub = stub
}

func (fake *AdapterAPI) GetNetworkAdapterPortPropertiesArgsForCall(i int) string {
	fake.getNetworkAdapterPortPropertiesMutex.RLock()
	defer fake.getNetworkAdapterPortPropertiesMutex.RUnlock()
	argsForCall := fake.getNetworkAdapterPortPropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *AdapterAPI) GetNetworkAdapterPortPropertiesReturns(result1 *zhmcclient.NetworkAdapterPort, result2 int, result3 *zhmcclient.HmcError) {
	fake.getNetworkAdapterPortPropertiesMutex.Lock()
	defer fake.getNetworkAdapterPortPropertiesMutex.Unlock()
	fake.GetNetworkAdapterPortPropertiesStub = nil
	fake.getNetworkAdapterPortPropertiesReturns = struct {
		result1 *zhmcclient.NetworkAdapterPort
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *AdapterAPI) GetNetworkAdapterPortPropertiesReturnsOnCall(i int, result1 *zhmcclient.NetworkAdapterPort, result2 int, result3 *zhmcclient.HmcError) {
	fake.getNetworkAdapterPortPropertiesMutex.Lock()
	defer fake.getNetworkAdapterPortPropertiesMutex.Unlock()
	fake.GetNetworkAdapterPortPropertiesStub = nil
	if fake.getNetworkAdapterPortPropertiesReturnsOnCall == nil {
		fake.getNetworkAdapterPortPropertiesReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.NetworkAdapterPort
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.getNetworkAdapterPortPropertiesReturnsOnCall[i] = struct {
		result1 *zhmcclient.NetworkAdapterPort
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *AdapterAPI) GetStorageAdapterPortProperties(arg1 string) (*zhmcclient.StorageAdapterPort, int, *zhmcclient.HmcError) {
	fake.getStorageAdapterPortPropertiesMutex.Lock()
	ret, specificReturn := fake.getStorageAdapterPortPropertiesReturnsOnCall[len(fake.getStorageAdapterPortPropertiesArgsForCall)]
	fake.getStorageAdapterPortPropertiesArgsForCall = append(fake.getStorageAdapterPortPropertiesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetStorageAdapterPortPropertiesStub
	fakeReturns := fake.getStorageAdapterPortPropertiesReturns
	fake.recordInvocation("GetStorageAdapterPortProperties", []interface{}{arg1})
	fake.getStorageAdapterPortPropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *AdapterAPI) GetStorageAdapterPortPropertiesCallCount() int {
	fake.getStorageAdapterPortPropertiesMutex.RLock()
	defer fake.getStorageAdapterPortPropertiesMutex.RUnlock()
	return len(fake.getStorageAdapterPortPropertiesArgsForCall)
}

func (fake *AdapterAPI) GetStorageAdapterPortPropertiesCalls(stub func(string) (*zhmcclient.StorageAdapterPort, int, *zhmcclient.HmcError)) {
	fake.getStorageAdapterPortPropertiesMutex.Lock()
	defer fake.getStorageAdapterPortPropertiesMutex.Unlock()
	fake.GetStorageAdapterPortPropertiesStub = stub
}

func (fake *AdapterAPI) GetStorageAdapterPortPropertiesArgsForCall(i int) string {
	fake.getStorageAdapterPortPropertiesMutex.RLock()
	defer fake.getStorageAdapterPortPropertiesMutex.RUnlock()
	argsForCall := fake.getStorageAdapterPortPropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *AdapterAPI) GetStorageAdapterPortPropertiesReturns(result1 *zhmcclient.StorageAdapterPort, result2 int, result3 *zhmcclient.HmcError) {
	fake.getStorageAdapterPortPropertiesMutex.Lock()
	defer fake.getStorageAdapterPortPropertiesMutex.Unlock()
	fake.GetStorageAdapterPortPropertiesStub = nil
	fake.getStorageAdapterPortPropertiesReturns = struct {
		result1 *zhmcclient.StorageAdapterPort
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *AdapterAPI) GetStorageAdapterPortPropertiesReturnsOnCall(i int, result1 *zhmcclient.StorageAdapterPort, result2 int, result3 *zhmcclient.HmcError) {
	fake.getStorageAdapterPortPropertiesMutex.Lock()
	defer fake.getStorageAdapterPortPropertiesMutex.Unlock()
	fake.GetStorageAdapterPortPropertiesStub = nil
	if fake.getStorageAdapterPortPropertiesReturnsOnCall == nil {
		fake.getStorageAdapterPortPropertiesReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.StorageAdapterPort
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.getStorageAdapterPortPropertiesReturnsOnCall[i] = struct {
		result1 *zhmcclient.StorageAdapterPort
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *AdapterAPI) ListAdapters(arg1 string, arg2 map[string]string) ([]zhmcclient.Adapter, int, *zhmcclient.HmcError) {
	fake.listAdaptersMutex.Lock()
	ret, specificReturn := fake.listAdaptersReturnsOnCall[len(fake.listAdaptersArgsForCall)]
	fake.listAdaptersArgsForCall = append(fake.listAdaptersArgsForCall, struct {
		arg1 string
		arg2 map[string]string
	}{arg1, arg2})
	stub := fake.ListAdaptersStub
	fakeReturns := fake.listAdaptersReturns
	fake.recordInvocation("ListAdapters", []interface{}{arg1, arg2})
	fake.listAdaptersMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *AdapterAPI) ListAdaptersCallCount() int {
	fake.listAdaptersMutex.RLock()
	defer fake.listAdaptersMutex.RUnlock()
	return len(fake.listAdaptersArgsForCall)
}

func (fake *AdapterAPI) ListAdaptersCalls(stub func(string, map[string]string) ([]zhmcclient.Adapter, int, *zhmcclient.HmcError)) {
	fake.listAdaptersMutex.Lock()
	defer fake.listAdaptersMutex.Unlock()
	fake.ListAdaptersStub = stub
}

func (fake *AdapterAPI) ListAdaptersArgsForCall(i int) (string, map[string]string) {
	fake.listAdaptersMutex.RLock()
	defer fake.listAdaptersMutex.RUnlock()
	argsForCall := fake.listAdaptersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *AdapterAPI) ListAdaptersReturns(result1 []zhmcclient.Adapter, result2 int, result3 *zhmcclient.HmcError) {
	fake.listAdaptersMutex.Lock()
	defer fake.listAdaptersMutex.Unlock()
	fake.ListAdaptersStub = nil
	fake.listAdaptersReturns = struct {
		result1 []zhmcclient.Adapter
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *AdapterAPI) ListAdaptersReturnsOnCall(i int, result1 []zhmcclient.Adapter, result2 int, result3 *zhmcclient.HmcError) {
	fake.listAdaptersMutex.Lock()
	defer fake.listAdaptersMutex.Unlock()
	fake.ListAdaptersStub = nil
	if fake.listAdaptersReturnsOnCall == nil {
		fake.listAdaptersReturnsOnCall = make(map[int]struct {
			result1 []zhmcclient.Adapter
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.listAdaptersReturnsOnCall[i] = struct {
		result1 []zhmcclient.Adapter
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *AdapterAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createHipersocketMutex.RLock()
	defer fake.createHipersocketMutex.RUnlock()
	fake.deleteHipersocketMutex.RLock()
	defer fake.deleteHipersocketMutex.RUnlock()
	fake.getAdapterPropertiesMutex.RLock()
	defer fake.getAdapterPropertiesMutex.RUnlock()
	fake.getNetworkAdapterPortPropertiesMutex.RLock()
	defer fake.getNetworkAdapterPortPropertiesMutex.RUnlock()
	fake.getStorageAdapterPortPropertiesMutex.RLock()
	defer fake.getStorageAdapterPortPropertiesMutex.RUnlock()
	fake.listAdaptersMutex.RLock()
	defer fake.listAdaptersMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *AdapterAPI) recordInvocation(key string, args []interface{}) {
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

var _ zhmcclient.AdapterAPI = new(AdapterAPI)

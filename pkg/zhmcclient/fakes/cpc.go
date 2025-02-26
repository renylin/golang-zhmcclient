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

type CpcAPI struct {
	ListCPCsStub        func(map[string]string) ([]zhmcclient.CPC, int, *zhmcclient.HmcError)
	listCPCsMutex       sync.RWMutex
	listCPCsArgsForCall []struct {
		arg1 map[string]string
	}
	listCPCsReturns struct {
		result1 []zhmcclient.CPC
		result2 int
		result3 *zhmcclient.HmcError
	}
	listCPCsReturnsOnCall map[int]struct {
		result1 []zhmcclient.CPC
		result2 int
		result3 *zhmcclient.HmcError
	}
	GetCPCPropertiesStub        func(string) (*zhmcclient.CPCProperties, int, *zhmcclient.HmcError)
	getCPCPropertiesMutex       sync.RWMutex
	getCPCPropertiesArgsForCall []struct {
		arg1 string
	}
	getCPCPropertiesReturns struct {
		result1 *zhmcclient.CPCProperties
		result2 int
		result3 *zhmcclient.HmcError
	}
	getCPCPropertiesReturnsOnCall map[int]struct {
		result1 *zhmcclient.CPCProperties
		result2 int
		result3 *zhmcclient.HmcError
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CpcAPI) ListCPCs(arg1 map[string]string) ([]zhmcclient.CPC, int, *zhmcclient.HmcError) {
	fake.listCPCsMutex.Lock()
	ret, specificReturn := fake.listCPCsReturnsOnCall[len(fake.listCPCsArgsForCall)]
	fake.listCPCsArgsForCall = append(fake.listCPCsArgsForCall, struct {
		arg1 map[string]string
	}{arg1})
	stub := fake.ListCPCsStub
	fakeReturns := fake.listCPCsReturns
	fake.recordInvocation("ListCPCs", []interface{}{arg1})
	fake.listCPCsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *CpcAPI) ListCPCsCallCount() int {
	fake.listCPCsMutex.RLock()
	defer fake.listCPCsMutex.RUnlock()
	return len(fake.listCPCsArgsForCall)
}

func (fake *CpcAPI) ListCPCsCalls(stub func(map[string]string) ([]zhmcclient.CPC, int, *zhmcclient.HmcError)) {
	fake.listCPCsMutex.Lock()
	defer fake.listCPCsMutex.Unlock()
	fake.ListCPCsStub = stub
}

func (fake *CpcAPI) ListCPCsArgsForCall(i int) map[string]string {
	fake.listCPCsMutex.RLock()
	defer fake.listCPCsMutex.RUnlock()
	argsForCall := fake.listCPCsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *CpcAPI) ListCPCsReturns(result1 []zhmcclient.CPC, result2 int, result3 *zhmcclient.HmcError) {
	fake.listCPCsMutex.Lock()
	defer fake.listCPCsMutex.Unlock()
	fake.ListCPCsStub = nil
	fake.listCPCsReturns = struct {
		result1 []zhmcclient.CPC
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *CpcAPI) ListCPCsReturnsOnCall(i int, result1 []zhmcclient.CPC, result2 int, result3 *zhmcclient.HmcError) {
	fake.listCPCsMutex.Lock()
	defer fake.listCPCsMutex.Unlock()
	fake.ListCPCsStub = nil
	if fake.listCPCsReturnsOnCall == nil {
		fake.listCPCsReturnsOnCall = make(map[int]struct {
			result1 []zhmcclient.CPC
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.listCPCsReturnsOnCall[i] = struct {
		result1 []zhmcclient.CPC
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *CpcAPI) GetCPCProperties(arg1 string) (*zhmcclient.CPCProperties, int, *zhmcclient.HmcError) {
	fake.getCPCPropertiesMutex.Lock()
	ret, specificReturn := fake.getCPCPropertiesReturnsOnCall[len(fake.getCPCPropertiesArgsForCall)]
	fake.getCPCPropertiesArgsForCall = append(fake.getCPCPropertiesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetCPCPropertiesStub
	fakeReturns := fake.getCPCPropertiesReturns
	fake.recordInvocation("GetCPCProperties", []interface{}{arg1})
	fake.getCPCPropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *CpcAPI) GetCPCPropertiesCallCount() int {
	fake.getCPCPropertiesMutex.RLock()
	defer fake.getCPCPropertiesMutex.RUnlock()
	return len(fake.getCPCPropertiesArgsForCall)
}

func (fake *CpcAPI) GetCPCPropertiesCalls(stub func(string) (*zhmcclient.CPCProperties, int, *zhmcclient.HmcError)) {
	fake.getCPCPropertiesMutex.Lock()
	defer fake.getCPCPropertiesMutex.Unlock()
	fake.GetCPCPropertiesStub = stub
}

func (fake *CpcAPI) GetCPCPropertiesArgsForCall(i int) string {
	fake.getCPCPropertiesMutex.RLock()
	defer fake.getCPCPropertiesMutex.RUnlock()
	argsForCall := fake.getCPCPropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *CpcAPI) GetCPCPropertiesReturns(result1 *zhmcclient.CPCProperties, result2 int, result3 *zhmcclient.HmcError) {
	fake.getCPCPropertiesMutex.Lock()
	defer fake.getCPCPropertiesMutex.Unlock()
	fake.GetCPCPropertiesStub = nil
	fake.getCPCPropertiesReturns = struct {
		result1 *zhmcclient.CPCProperties
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *CpcAPI) GetCPCPropertiesReturnsOnCall(i int, result1 *zhmcclient.CPCProperties, result2 int, result3 *zhmcclient.HmcError) {
	fake.getCPCPropertiesMutex.Lock()
	defer fake.getCPCPropertiesMutex.Unlock()
	fake.GetCPCPropertiesStub = nil
	if fake.getCPCPropertiesReturnsOnCall == nil {
		fake.getCPCPropertiesReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.CPCProperties
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.getCPCPropertiesReturnsOnCall[i] = struct {
		result1 *zhmcclient.CPCProperties
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *CpcAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listCPCsMutex.RLock()
	defer fake.listCPCsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CpcAPI) recordInvocation(key string, args []interface{}) {
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

var _ zhmcclient.CpcAPI = new(CpcAPI)

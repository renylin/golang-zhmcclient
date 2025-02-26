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

type JobAPI struct {
	CancelJobStub        func(string) (int, *zhmcclient.HmcError)
	cancelJobMutex       sync.RWMutex
	cancelJobArgsForCall []struct {
		arg1 string
	}
	cancelJobReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	cancelJobReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	DeleteJobStub        func(string) (int, *zhmcclient.HmcError)
	deleteJobMutex       sync.RWMutex
	deleteJobArgsForCall []struct {
		arg1 string
	}
	deleteJobReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	deleteJobReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	QueryJobStub        func(string) (*zhmcclient.Job, int, *zhmcclient.HmcError)
	queryJobMutex       sync.RWMutex
	queryJobArgsForCall []struct {
		arg1 string
	}
	queryJobReturns struct {
		result1 *zhmcclient.Job
		result2 int
		result3 *zhmcclient.HmcError
	}
	queryJobReturnsOnCall map[int]struct {
		result1 *zhmcclient.Job
		result2 int
		result3 *zhmcclient.HmcError
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *JobAPI) CancelJob(arg1 string) (int, *zhmcclient.HmcError) {
	fake.cancelJobMutex.Lock()
	ret, specificReturn := fake.cancelJobReturnsOnCall[len(fake.cancelJobArgsForCall)]
	fake.cancelJobArgsForCall = append(fake.cancelJobArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.CancelJobStub
	fakeReturns := fake.cancelJobReturns
	fake.recordInvocation("CancelJob", []interface{}{arg1})
	fake.cancelJobMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *JobAPI) CancelJobCallCount() int {
	fake.cancelJobMutex.RLock()
	defer fake.cancelJobMutex.RUnlock()
	return len(fake.cancelJobArgsForCall)
}

func (fake *JobAPI) CancelJobCalls(stub func(string) (int, *zhmcclient.HmcError)) {
	fake.cancelJobMutex.Lock()
	defer fake.cancelJobMutex.Unlock()
	fake.CancelJobStub = stub
}

func (fake *JobAPI) CancelJobArgsForCall(i int) string {
	fake.cancelJobMutex.RLock()
	defer fake.cancelJobMutex.RUnlock()
	argsForCall := fake.cancelJobArgsForCall[i]
	return argsForCall.arg1
}

func (fake *JobAPI) CancelJobReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.cancelJobMutex.Lock()
	defer fake.cancelJobMutex.Unlock()
	fake.CancelJobStub = nil
	fake.cancelJobReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *JobAPI) CancelJobReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.cancelJobMutex.Lock()
	defer fake.cancelJobMutex.Unlock()
	fake.CancelJobStub = nil
	if fake.cancelJobReturnsOnCall == nil {
		fake.cancelJobReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.cancelJobReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *JobAPI) DeleteJob(arg1 string) (int, *zhmcclient.HmcError) {
	fake.deleteJobMutex.Lock()
	ret, specificReturn := fake.deleteJobReturnsOnCall[len(fake.deleteJobArgsForCall)]
	fake.deleteJobArgsForCall = append(fake.deleteJobArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteJobStub
	fakeReturns := fake.deleteJobReturns
	fake.recordInvocation("DeleteJob", []interface{}{arg1})
	fake.deleteJobMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *JobAPI) DeleteJobCallCount() int {
	fake.deleteJobMutex.RLock()
	defer fake.deleteJobMutex.RUnlock()
	return len(fake.deleteJobArgsForCall)
}

func (fake *JobAPI) DeleteJobCalls(stub func(string) (int, *zhmcclient.HmcError)) {
	fake.deleteJobMutex.Lock()
	defer fake.deleteJobMutex.Unlock()
	fake.DeleteJobStub = stub
}

func (fake *JobAPI) DeleteJobArgsForCall(i int) string {
	fake.deleteJobMutex.RLock()
	defer fake.deleteJobMutex.RUnlock()
	argsForCall := fake.deleteJobArgsForCall[i]
	return argsForCall.arg1
}

func (fake *JobAPI) DeleteJobReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.deleteJobMutex.Lock()
	defer fake.deleteJobMutex.Unlock()
	fake.DeleteJobStub = nil
	fake.deleteJobReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *JobAPI) DeleteJobReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.deleteJobMutex.Lock()
	defer fake.deleteJobMutex.Unlock()
	fake.DeleteJobStub = nil
	if fake.deleteJobReturnsOnCall == nil {
		fake.deleteJobReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.deleteJobReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *JobAPI) QueryJob(arg1 string) (*zhmcclient.Job, int, *zhmcclient.HmcError) {
	fake.queryJobMutex.Lock()
	ret, specificReturn := fake.queryJobReturnsOnCall[len(fake.queryJobArgsForCall)]
	fake.queryJobArgsForCall = append(fake.queryJobArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.QueryJobStub
	fakeReturns := fake.queryJobReturns
	fake.recordInvocation("QueryJob", []interface{}{arg1})
	fake.queryJobMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *JobAPI) QueryJobCallCount() int {
	fake.queryJobMutex.RLock()
	defer fake.queryJobMutex.RUnlock()
	return len(fake.queryJobArgsForCall)
}

func (fake *JobAPI) QueryJobCalls(stub func(string) (*zhmcclient.Job, int, *zhmcclient.HmcError)) {
	fake.queryJobMutex.Lock()
	defer fake.queryJobMutex.Unlock()
	fake.QueryJobStub = stub
}

func (fake *JobAPI) QueryJobArgsForCall(i int) string {
	fake.queryJobMutex.RLock()
	defer fake.queryJobMutex.RUnlock()
	argsForCall := fake.queryJobArgsForCall[i]
	return argsForCall.arg1
}

func (fake *JobAPI) QueryJobReturns(result1 *zhmcclient.Job, result2 int, result3 *zhmcclient.HmcError) {
	fake.queryJobMutex.Lock()
	defer fake.queryJobMutex.Unlock()
	fake.QueryJobStub = nil
	fake.queryJobReturns = struct {
		result1 *zhmcclient.Job
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *JobAPI) QueryJobReturnsOnCall(i int, result1 *zhmcclient.Job, result2 int, result3 *zhmcclient.HmcError) {
	fake.queryJobMutex.Lock()
	defer fake.queryJobMutex.Unlock()
	fake.QueryJobStub = nil
	if fake.queryJobReturnsOnCall == nil {
		fake.queryJobReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.Job
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.queryJobReturnsOnCall[i] = struct {
		result1 *zhmcclient.Job
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *JobAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cancelJobMutex.RLock()
	defer fake.cancelJobMutex.RUnlock()
	fake.deleteJobMutex.RLock()
	defer fake.deleteJobMutex.RUnlock()
	fake.queryJobMutex.RLock()
	defer fake.queryJobMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *JobAPI) recordInvocation(key string, args []interface{}) {
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

var _ zhmcclient.JobAPI = new(JobAPI)

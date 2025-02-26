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

type ZhmcAPI struct {
	AttachStorageGroupToPartitionStub        func(string, *zhmcclient.StorageGroupPayload) (int, *zhmcclient.HmcError)
	attachStorageGroupToPartitionMutex       sync.RWMutex
	attachStorageGroupToPartitionArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.StorageGroupPayload
	}
	attachStorageGroupToPartitionReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	attachStorageGroupToPartitionReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
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
	CreateLPARStub        func(string, *zhmcclient.LparProperties) (string, int, *zhmcclient.HmcError)
	createLPARMutex       sync.RWMutex
	createLPARArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.LparProperties
	}
	createLPARReturns struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}
	createLPARReturnsOnCall map[int]struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}
	CreateNicStub        func(string, *zhmcclient.NIC) (string, int, *zhmcclient.HmcError)
	createNicMutex       sync.RWMutex
	createNicArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.NIC
	}
	createNicReturns struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}
	createNicReturnsOnCall map[int]struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}
	CreateStorageGroupsStub        func(string, *zhmcclient.CreateStorageGroupProperties) (*zhmcclient.StorageGroupCreateResponse, int, *zhmcclient.HmcError)
	createStorageGroupsMutex       sync.RWMutex
	createStorageGroupsArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.CreateStorageGroupProperties
	}
	createStorageGroupsReturns struct {
		result1 *zhmcclient.StorageGroupCreateResponse
		result2 int
		result3 *zhmcclient.HmcError
	}
	createStorageGroupsReturnsOnCall map[int]struct {
		result1 *zhmcclient.StorageGroupCreateResponse
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
	DeleteLPARStub        func(string) (int, *zhmcclient.HmcError)
	deleteLPARMutex       sync.RWMutex
	deleteLPARArgsForCall []struct {
		arg1 string
	}
	deleteLPARReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	deleteLPARReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	DeleteNicStub        func(string) (int, *zhmcclient.HmcError)
	deleteNicMutex       sync.RWMutex
	deleteNicArgsForCall []struct {
		arg1 string
	}
	deleteNicReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	deleteNicReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	DeleteStorageGroupStub        func(string) (int, *zhmcclient.HmcError)
	deleteStorageGroupMutex       sync.RWMutex
	deleteStorageGroupArgsForCall []struct {
		arg1 string
	}
	deleteStorageGroupReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	deleteStorageGroupReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	DetachStorageGroupToPartitionStub        func(string, *zhmcclient.StorageGroupPayload) (int, *zhmcclient.HmcError)
	detachStorageGroupToPartitionMutex       sync.RWMutex
	detachStorageGroupToPartitionArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.StorageGroupPayload
	}
	detachStorageGroupToPartitionReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	detachStorageGroupToPartitionReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	FetchAsciiConsoleURIStub        func(string, *zhmcclient.AsciiConsoleURIPayload) (*zhmcclient.AsciiConsoleURIResponse, int, *zhmcclient.HmcError)
	fetchAsciiConsoleURIMutex       sync.RWMutex
	fetchAsciiConsoleURIArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.AsciiConsoleURIPayload
	}
	fetchAsciiConsoleURIReturns struct {
		result1 *zhmcclient.AsciiConsoleURIResponse
		result2 int
		result3 *zhmcclient.HmcError
	}
	fetchAsciiConsoleURIReturnsOnCall map[int]struct {
		result1 *zhmcclient.AsciiConsoleURIResponse
		result2 int
		result3 *zhmcclient.HmcError
	}
	FulfillStorageGroupStub        func(string, *zhmcclient.StorageGroupProperties) (int, *zhmcclient.HmcError)
	fulfillStorageGroupMutex       sync.RWMutex
	fulfillStorageGroupArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.StorageGroupProperties
	}
	fulfillStorageGroupReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	fulfillStorageGroupReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
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
	GetLparPropertiesStub        func(string) (*zhmcclient.LparProperties, int, *zhmcclient.HmcError)
	getLparPropertiesMutex       sync.RWMutex
	getLparPropertiesArgsForCall []struct {
		arg1 string
	}
	getLparPropertiesReturns struct {
		result1 *zhmcclient.LparProperties
		result2 int
		result3 *zhmcclient.HmcError
	}
	getLparPropertiesReturnsOnCall map[int]struct {
		result1 *zhmcclient.LparProperties
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
	GetNicPropertiesStub        func(string) (*zhmcclient.NIC, int, *zhmcclient.HmcError)
	getNicPropertiesMutex       sync.RWMutex
	getNicPropertiesArgsForCall []struct {
		arg1 string
	}
	getNicPropertiesReturns struct {
		result1 *zhmcclient.NIC
		result2 int
		result3 *zhmcclient.HmcError
	}
	getNicPropertiesReturnsOnCall map[int]struct {
		result1 *zhmcclient.NIC
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
	GetStorageGroupPartitionsStub        func(string, map[string]string) (*zhmcclient.StorageGroupPartitions, int, *zhmcclient.HmcError)
	getStorageGroupPartitionsMutex       sync.RWMutex
	getStorageGroupPartitionsArgsForCall []struct {
		arg1 string
		arg2 map[string]string
	}
	getStorageGroupPartitionsReturns struct {
		result1 *zhmcclient.StorageGroupPartitions
		result2 int
		result3 *zhmcclient.HmcError
	}
	getStorageGroupPartitionsReturnsOnCall map[int]struct {
		result1 *zhmcclient.StorageGroupPartitions
		result2 int
		result3 *zhmcclient.HmcError
	}
	GetStorageGroupPropertiesStub        func(string) (*zhmcclient.StorageGroupProperties, int, *zhmcclient.HmcError)
	getStorageGroupPropertiesMutex       sync.RWMutex
	getStorageGroupPropertiesArgsForCall []struct {
		arg1 string
	}
	getStorageGroupPropertiesReturns struct {
		result1 *zhmcclient.StorageGroupProperties
		result2 int
		result3 *zhmcclient.HmcError
	}
	getStorageGroupPropertiesReturnsOnCall map[int]struct {
		result1 *zhmcclient.StorageGroupProperties
		result2 int
		result3 *zhmcclient.HmcError
	}
	GetStorageVolumePropertiesStub        func(string) (*zhmcclient.StorageVolume, int, *zhmcclient.HmcError)
	getStorageVolumePropertiesMutex       sync.RWMutex
	getStorageVolumePropertiesArgsForCall []struct {
		arg1 string
	}
	getStorageVolumePropertiesReturns struct {
		result1 *zhmcclient.StorageVolume
		result2 int
		result3 *zhmcclient.HmcError
	}
	getStorageVolumePropertiesReturnsOnCall map[int]struct {
		result1 *zhmcclient.StorageVolume
		result2 int
		result3 *zhmcclient.HmcError
	}
	GetVirtualSwitchPropertiesStub        func(string) (*zhmcclient.VirtualSwitchProperties, int, *zhmcclient.HmcError)
	getVirtualSwitchPropertiesMutex       sync.RWMutex
	getVirtualSwitchPropertiesArgsForCall []struct {
		arg1 string
	}
	getVirtualSwitchPropertiesReturns struct {
		result1 *zhmcclient.VirtualSwitchProperties
		result2 int
		result3 *zhmcclient.HmcError
	}
	getVirtualSwitchPropertiesReturnsOnCall map[int]struct {
		result1 *zhmcclient.VirtualSwitchProperties
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
	ListLPARsStub        func(string, map[string]string) ([]zhmcclient.LPAR, int, *zhmcclient.HmcError)
	listLPARsMutex       sync.RWMutex
	listLPARsArgsForCall []struct {
		arg1 string
		arg2 map[string]string
	}
	listLPARsReturns struct {
		result1 []zhmcclient.LPAR
		result2 int
		result3 *zhmcclient.HmcError
	}
	listLPARsReturnsOnCall map[int]struct {
		result1 []zhmcclient.LPAR
		result2 int
		result3 *zhmcclient.HmcError
	}
	ListNicsStub        func(string) ([]string, int, *zhmcclient.HmcError)
	listNicsMutex       sync.RWMutex
	listNicsArgsForCall []struct {
		arg1 string
	}
	listNicsReturns struct {
		result1 []string
		result2 int
		result3 *zhmcclient.HmcError
	}
	listNicsReturnsOnCall map[int]struct {
		result1 []string
		result2 int
		result3 *zhmcclient.HmcError
	}
	ListStorageGroupsStub        func(string, string) ([]zhmcclient.StorageGroup, int, *zhmcclient.HmcError)
	listStorageGroupsMutex       sync.RWMutex
	listStorageGroupsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	listStorageGroupsReturns struct {
		result1 []zhmcclient.StorageGroup
		result2 int
		result3 *zhmcclient.HmcError
	}
	listStorageGroupsReturnsOnCall map[int]struct {
		result1 []zhmcclient.StorageGroup
		result2 int
		result3 *zhmcclient.HmcError
	}
	ListStorageVolumesStub        func(string) ([]zhmcclient.StorageVolume, int, *zhmcclient.HmcError)
	listStorageVolumesMutex       sync.RWMutex
	listStorageVolumesArgsForCall []struct {
		arg1 string
	}
	listStorageVolumesReturns struct {
		result1 []zhmcclient.StorageVolume
		result2 int
		result3 *zhmcclient.HmcError
	}
	listStorageVolumesReturnsOnCall map[int]struct {
		result1 []zhmcclient.StorageVolume
		result2 int
		result3 *zhmcclient.HmcError
	}
	ListVirtualSwitchesStub        func(string, map[string]string) ([]zhmcclient.VirtualSwitch, int, *zhmcclient.HmcError)
	listVirtualSwitchesMutex       sync.RWMutex
	listVirtualSwitchesArgsForCall []struct {
		arg1 string
		arg2 map[string]string
	}
	listVirtualSwitchesReturns struct {
		result1 []zhmcclient.VirtualSwitch
		result2 int
		result3 *zhmcclient.HmcError
	}
	listVirtualSwitchesReturnsOnCall map[int]struct {
		result1 []zhmcclient.VirtualSwitch
		result2 int
		result3 *zhmcclient.HmcError
	}
	MountIsoImageStub        func(string, string, string) (int, *zhmcclient.HmcError)
	mountIsoImageMutex       sync.RWMutex
	mountIsoImageArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	mountIsoImageReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	mountIsoImageReturnsOnCall map[int]struct {
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
	StartLPARStub        func(string) (string, int, *zhmcclient.HmcError)
	startLPARMutex       sync.RWMutex
	startLPARArgsForCall []struct {
		arg1 string
	}
	startLPARReturns struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}
	startLPARReturnsOnCall map[int]struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}
	StopLPARStub        func(string) (string, int, *zhmcclient.HmcError)
	stopLPARMutex       sync.RWMutex
	stopLPARArgsForCall []struct {
		arg1 string
	}
	stopLPARReturns struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}
	stopLPARReturnsOnCall map[int]struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}
	UnmountIsoImageStub        func(string) (int, *zhmcclient.HmcError)
	unmountIsoImageMutex       sync.RWMutex
	unmountIsoImageArgsForCall []struct {
		arg1 string
	}
	unmountIsoImageReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	unmountIsoImageReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	UpdateLparPropertiesStub        func(string, *zhmcclient.LparProperties) (int, *zhmcclient.HmcError)
	updateLparPropertiesMutex       sync.RWMutex
	updateLparPropertiesArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.LparProperties
	}
	updateLparPropertiesReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	updateLparPropertiesReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	UpdateNicPropertiesStub        func(string, *zhmcclient.NIC) (int, *zhmcclient.HmcError)
	updateNicPropertiesMutex       sync.RWMutex
	updateNicPropertiesArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.NIC
	}
	updateNicPropertiesReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	updateNicPropertiesReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	UpdateStorageGroupPropertiesStub        func(string, *zhmcclient.StorageGroupProperties) (int, *zhmcclient.HmcError)
	updateStorageGroupPropertiesMutex       sync.RWMutex
	updateStorageGroupPropertiesArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.StorageGroupProperties
	}
	updateStorageGroupPropertiesReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	updateStorageGroupPropertiesReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ZhmcAPI) AttachStorageGroupToPartition(arg1 string, arg2 *zhmcclient.StorageGroupPayload) (int, *zhmcclient.HmcError) {
	fake.attachStorageGroupToPartitionMutex.Lock()
	ret, specificReturn := fake.attachStorageGroupToPartitionReturnsOnCall[len(fake.attachStorageGroupToPartitionArgsForCall)]
	fake.attachStorageGroupToPartitionArgsForCall = append(fake.attachStorageGroupToPartitionArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.StorageGroupPayload
	}{arg1, arg2})
	stub := fake.AttachStorageGroupToPartitionStub
	fakeReturns := fake.attachStorageGroupToPartitionReturns
	fake.recordInvocation("AttachStorageGroupToPartition", []interface{}{arg1, arg2})
	fake.attachStorageGroupToPartitionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ZhmcAPI) AttachStorageGroupToPartitionCallCount() int {
	fake.attachStorageGroupToPartitionMutex.RLock()
	defer fake.attachStorageGroupToPartitionMutex.RUnlock()
	return len(fake.attachStorageGroupToPartitionArgsForCall)
}

func (fake *ZhmcAPI) AttachStorageGroupToPartitionCalls(stub func(string, *zhmcclient.StorageGroupPayload) (int, *zhmcclient.HmcError)) {
	fake.attachStorageGroupToPartitionMutex.Lock()
	defer fake.attachStorageGroupToPartitionMutex.Unlock()
	fake.AttachStorageGroupToPartitionStub = stub
}

func (fake *ZhmcAPI) AttachStorageGroupToPartitionArgsForCall(i int) (string, *zhmcclient.StorageGroupPayload) {
	fake.attachStorageGroupToPartitionMutex.RLock()
	defer fake.attachStorageGroupToPartitionMutex.RUnlock()
	argsForCall := fake.attachStorageGroupToPartitionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ZhmcAPI) AttachStorageGroupToPartitionReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.attachStorageGroupToPartitionMutex.Lock()
	defer fake.attachStorageGroupToPartitionMutex.Unlock()
	fake.AttachStorageGroupToPartitionStub = nil
	fake.attachStorageGroupToPartitionReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) AttachStorageGroupToPartitionReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.attachStorageGroupToPartitionMutex.Lock()
	defer fake.attachStorageGroupToPartitionMutex.Unlock()
	fake.AttachStorageGroupToPartitionStub = nil
	if fake.attachStorageGroupToPartitionReturnsOnCall == nil {
		fake.attachStorageGroupToPartitionReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.attachStorageGroupToPartitionReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) CancelJob(arg1 string) (int, *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) CancelJobCallCount() int {
	fake.cancelJobMutex.RLock()
	defer fake.cancelJobMutex.RUnlock()
	return len(fake.cancelJobArgsForCall)
}

func (fake *ZhmcAPI) CancelJobCalls(stub func(string) (int, *zhmcclient.HmcError)) {
	fake.cancelJobMutex.Lock()
	defer fake.cancelJobMutex.Unlock()
	fake.CancelJobStub = stub
}

func (fake *ZhmcAPI) CancelJobArgsForCall(i int) string {
	fake.cancelJobMutex.RLock()
	defer fake.cancelJobMutex.RUnlock()
	argsForCall := fake.cancelJobArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) CancelJobReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.cancelJobMutex.Lock()
	defer fake.cancelJobMutex.Unlock()
	fake.CancelJobStub = nil
	fake.cancelJobReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) CancelJobReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) CreateHipersocket(arg1 string, arg2 *zhmcclient.HipersocketPayload) (string, int, *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) CreateHipersocketCallCount() int {
	fake.createHipersocketMutex.RLock()
	defer fake.createHipersocketMutex.RUnlock()
	return len(fake.createHipersocketArgsForCall)
}

func (fake *ZhmcAPI) CreateHipersocketCalls(stub func(string, *zhmcclient.HipersocketPayload) (string, int, *zhmcclient.HmcError)) {
	fake.createHipersocketMutex.Lock()
	defer fake.createHipersocketMutex.Unlock()
	fake.CreateHipersocketStub = stub
}

func (fake *ZhmcAPI) CreateHipersocketArgsForCall(i int) (string, *zhmcclient.HipersocketPayload) {
	fake.createHipersocketMutex.RLock()
	defer fake.createHipersocketMutex.RUnlock()
	argsForCall := fake.createHipersocketArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ZhmcAPI) CreateHipersocketReturns(result1 string, result2 int, result3 *zhmcclient.HmcError) {
	fake.createHipersocketMutex.Lock()
	defer fake.createHipersocketMutex.Unlock()
	fake.CreateHipersocketStub = nil
	fake.createHipersocketReturns = struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) CreateHipersocketReturnsOnCall(i int, result1 string, result2 int, result3 *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) CreateLPAR(arg1 string, arg2 *zhmcclient.LparProperties) (string, int, *zhmcclient.HmcError) {
	fake.createLPARMutex.Lock()
	ret, specificReturn := fake.createLPARReturnsOnCall[len(fake.createLPARArgsForCall)]
	fake.createLPARArgsForCall = append(fake.createLPARArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.LparProperties
	}{arg1, arg2})
	stub := fake.CreateLPARStub
	fakeReturns := fake.createLPARReturns
	fake.recordInvocation("CreateLPAR", []interface{}{arg1, arg2})
	fake.createLPARMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ZhmcAPI) CreateLPARCallCount() int {
	fake.createLPARMutex.RLock()
	defer fake.createLPARMutex.RUnlock()
	return len(fake.createLPARArgsForCall)
}

func (fake *ZhmcAPI) CreateLPARCalls(stub func(string, *zhmcclient.LparProperties) (string, int, *zhmcclient.HmcError)) {
	fake.createLPARMutex.Lock()
	defer fake.createLPARMutex.Unlock()
	fake.CreateLPARStub = stub
}

func (fake *ZhmcAPI) CreateLPARArgsForCall(i int) (string, *zhmcclient.LparProperties) {
	fake.createLPARMutex.RLock()
	defer fake.createLPARMutex.RUnlock()
	argsForCall := fake.createLPARArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ZhmcAPI) CreateLPARReturns(result1 string, result2 int, result3 *zhmcclient.HmcError) {
	fake.createLPARMutex.Lock()
	defer fake.createLPARMutex.Unlock()
	fake.CreateLPARStub = nil
	fake.createLPARReturns = struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) CreateLPARReturnsOnCall(i int, result1 string, result2 int, result3 *zhmcclient.HmcError) {
	fake.createLPARMutex.Lock()
	defer fake.createLPARMutex.Unlock()
	fake.CreateLPARStub = nil
	if fake.createLPARReturnsOnCall == nil {
		fake.createLPARReturnsOnCall = make(map[int]struct {
			result1 string
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.createLPARReturnsOnCall[i] = struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) CreateNic(arg1 string, arg2 *zhmcclient.NIC) (string, int, *zhmcclient.HmcError) {
	fake.createNicMutex.Lock()
	ret, specificReturn := fake.createNicReturnsOnCall[len(fake.createNicArgsForCall)]
	fake.createNicArgsForCall = append(fake.createNicArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.NIC
	}{arg1, arg2})
	stub := fake.CreateNicStub
	fakeReturns := fake.createNicReturns
	fake.recordInvocation("CreateNic", []interface{}{arg1, arg2})
	fake.createNicMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ZhmcAPI) CreateNicCallCount() int {
	fake.createNicMutex.RLock()
	defer fake.createNicMutex.RUnlock()
	return len(fake.createNicArgsForCall)
}

func (fake *ZhmcAPI) CreateNicCalls(stub func(string, *zhmcclient.NIC) (string, int, *zhmcclient.HmcError)) {
	fake.createNicMutex.Lock()
	defer fake.createNicMutex.Unlock()
	fake.CreateNicStub = stub
}

func (fake *ZhmcAPI) CreateNicArgsForCall(i int) (string, *zhmcclient.NIC) {
	fake.createNicMutex.RLock()
	defer fake.createNicMutex.RUnlock()
	argsForCall := fake.createNicArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ZhmcAPI) CreateNicReturns(result1 string, result2 int, result3 *zhmcclient.HmcError) {
	fake.createNicMutex.Lock()
	defer fake.createNicMutex.Unlock()
	fake.CreateNicStub = nil
	fake.createNicReturns = struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) CreateNicReturnsOnCall(i int, result1 string, result2 int, result3 *zhmcclient.HmcError) {
	fake.createNicMutex.Lock()
	defer fake.createNicMutex.Unlock()
	fake.CreateNicStub = nil
	if fake.createNicReturnsOnCall == nil {
		fake.createNicReturnsOnCall = make(map[int]struct {
			result1 string
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.createNicReturnsOnCall[i] = struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) CreateStorageGroups(arg1 string, arg2 *zhmcclient.CreateStorageGroupProperties) (*zhmcclient.StorageGroupCreateResponse, int, *zhmcclient.HmcError) {
	fake.createStorageGroupsMutex.Lock()
	ret, specificReturn := fake.createStorageGroupsReturnsOnCall[len(fake.createStorageGroupsArgsForCall)]
	fake.createStorageGroupsArgsForCall = append(fake.createStorageGroupsArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.CreateStorageGroupProperties
	}{arg1, arg2})
	stub := fake.CreateStorageGroupsStub
	fakeReturns := fake.createStorageGroupsReturns
	fake.recordInvocation("CreateStorageGroups", []interface{}{arg1, arg2})
	fake.createStorageGroupsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ZhmcAPI) CreateStorageGroupsCallCount() int {
	fake.createStorageGroupsMutex.RLock()
	defer fake.createStorageGroupsMutex.RUnlock()
	return len(fake.createStorageGroupsArgsForCall)
}

func (fake *ZhmcAPI) CreateStorageGroupsCalls(stub func(string, *zhmcclient.CreateStorageGroupProperties) (*zhmcclient.StorageGroupCreateResponse, int, *zhmcclient.HmcError)) {
	fake.createStorageGroupsMutex.Lock()
	defer fake.createStorageGroupsMutex.Unlock()
	fake.CreateStorageGroupsStub = stub
}

func (fake *ZhmcAPI) CreateStorageGroupsArgsForCall(i int) (string, *zhmcclient.CreateStorageGroupProperties) {
	fake.createStorageGroupsMutex.RLock()
	defer fake.createStorageGroupsMutex.RUnlock()
	argsForCall := fake.createStorageGroupsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ZhmcAPI) CreateStorageGroupsReturns(result1 *zhmcclient.StorageGroupCreateResponse, result2 int, result3 *zhmcclient.HmcError) {
	fake.createStorageGroupsMutex.Lock()
	defer fake.createStorageGroupsMutex.Unlock()
	fake.CreateStorageGroupsStub = nil
	fake.createStorageGroupsReturns = struct {
		result1 *zhmcclient.StorageGroupCreateResponse
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) CreateStorageGroupsReturnsOnCall(i int, result1 *zhmcclient.StorageGroupCreateResponse, result2 int, result3 *zhmcclient.HmcError) {
	fake.createStorageGroupsMutex.Lock()
	defer fake.createStorageGroupsMutex.Unlock()
	fake.CreateStorageGroupsStub = nil
	if fake.createStorageGroupsReturnsOnCall == nil {
		fake.createStorageGroupsReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.StorageGroupCreateResponse
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.createStorageGroupsReturnsOnCall[i] = struct {
		result1 *zhmcclient.StorageGroupCreateResponse
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) DeleteHipersocket(arg1 string) (int, *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) DeleteHipersocketCallCount() int {
	fake.deleteHipersocketMutex.RLock()
	defer fake.deleteHipersocketMutex.RUnlock()
	return len(fake.deleteHipersocketArgsForCall)
}

func (fake *ZhmcAPI) DeleteHipersocketCalls(stub func(string) (int, *zhmcclient.HmcError)) {
	fake.deleteHipersocketMutex.Lock()
	defer fake.deleteHipersocketMutex.Unlock()
	fake.DeleteHipersocketStub = stub
}

func (fake *ZhmcAPI) DeleteHipersocketArgsForCall(i int) string {
	fake.deleteHipersocketMutex.RLock()
	defer fake.deleteHipersocketMutex.RUnlock()
	argsForCall := fake.deleteHipersocketArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) DeleteHipersocketReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.deleteHipersocketMutex.Lock()
	defer fake.deleteHipersocketMutex.Unlock()
	fake.DeleteHipersocketStub = nil
	fake.deleteHipersocketReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) DeleteHipersocketReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) DeleteJob(arg1 string) (int, *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) DeleteJobCallCount() int {
	fake.deleteJobMutex.RLock()
	defer fake.deleteJobMutex.RUnlock()
	return len(fake.deleteJobArgsForCall)
}

func (fake *ZhmcAPI) DeleteJobCalls(stub func(string) (int, *zhmcclient.HmcError)) {
	fake.deleteJobMutex.Lock()
	defer fake.deleteJobMutex.Unlock()
	fake.DeleteJobStub = stub
}

func (fake *ZhmcAPI) DeleteJobArgsForCall(i int) string {
	fake.deleteJobMutex.RLock()
	defer fake.deleteJobMutex.RUnlock()
	argsForCall := fake.deleteJobArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) DeleteJobReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.deleteJobMutex.Lock()
	defer fake.deleteJobMutex.Unlock()
	fake.DeleteJobStub = nil
	fake.deleteJobReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) DeleteJobReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) DeleteLPAR(arg1 string) (int, *zhmcclient.HmcError) {
	fake.deleteLPARMutex.Lock()
	ret, specificReturn := fake.deleteLPARReturnsOnCall[len(fake.deleteLPARArgsForCall)]
	fake.deleteLPARArgsForCall = append(fake.deleteLPARArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteLPARStub
	fakeReturns := fake.deleteLPARReturns
	fake.recordInvocation("DeleteLPAR", []interface{}{arg1})
	fake.deleteLPARMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ZhmcAPI) DeleteLPARCallCount() int {
	fake.deleteLPARMutex.RLock()
	defer fake.deleteLPARMutex.RUnlock()
	return len(fake.deleteLPARArgsForCall)
}

func (fake *ZhmcAPI) DeleteLPARCalls(stub func(string) (int, *zhmcclient.HmcError)) {
	fake.deleteLPARMutex.Lock()
	defer fake.deleteLPARMutex.Unlock()
	fake.DeleteLPARStub = stub
}

func (fake *ZhmcAPI) DeleteLPARArgsForCall(i int) string {
	fake.deleteLPARMutex.RLock()
	defer fake.deleteLPARMutex.RUnlock()
	argsForCall := fake.deleteLPARArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) DeleteLPARReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.deleteLPARMutex.Lock()
	defer fake.deleteLPARMutex.Unlock()
	fake.DeleteLPARStub = nil
	fake.deleteLPARReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) DeleteLPARReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.deleteLPARMutex.Lock()
	defer fake.deleteLPARMutex.Unlock()
	fake.DeleteLPARStub = nil
	if fake.deleteLPARReturnsOnCall == nil {
		fake.deleteLPARReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.deleteLPARReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) DeleteNic(arg1 string) (int, *zhmcclient.HmcError) {
	fake.deleteNicMutex.Lock()
	ret, specificReturn := fake.deleteNicReturnsOnCall[len(fake.deleteNicArgsForCall)]
	fake.deleteNicArgsForCall = append(fake.deleteNicArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteNicStub
	fakeReturns := fake.deleteNicReturns
	fake.recordInvocation("DeleteNic", []interface{}{arg1})
	fake.deleteNicMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ZhmcAPI) DeleteNicCallCount() int {
	fake.deleteNicMutex.RLock()
	defer fake.deleteNicMutex.RUnlock()
	return len(fake.deleteNicArgsForCall)
}

func (fake *ZhmcAPI) DeleteNicCalls(stub func(string) (int, *zhmcclient.HmcError)) {
	fake.deleteNicMutex.Lock()
	defer fake.deleteNicMutex.Unlock()
	fake.DeleteNicStub = stub
}

func (fake *ZhmcAPI) DeleteNicArgsForCall(i int) string {
	fake.deleteNicMutex.RLock()
	defer fake.deleteNicMutex.RUnlock()
	argsForCall := fake.deleteNicArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) DeleteNicReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.deleteNicMutex.Lock()
	defer fake.deleteNicMutex.Unlock()
	fake.DeleteNicStub = nil
	fake.deleteNicReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) DeleteNicReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.deleteNicMutex.Lock()
	defer fake.deleteNicMutex.Unlock()
	fake.DeleteNicStub = nil
	if fake.deleteNicReturnsOnCall == nil {
		fake.deleteNicReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.deleteNicReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) DeleteStorageGroup(arg1 string) (int, *zhmcclient.HmcError) {
	fake.deleteStorageGroupMutex.Lock()
	ret, specificReturn := fake.deleteStorageGroupReturnsOnCall[len(fake.deleteStorageGroupArgsForCall)]
	fake.deleteStorageGroupArgsForCall = append(fake.deleteStorageGroupArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteStorageGroupStub
	fakeReturns := fake.deleteStorageGroupReturns
	fake.recordInvocation("DeleteStorageGroup", []interface{}{arg1})
	fake.deleteStorageGroupMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ZhmcAPI) DeleteStorageGroupCallCount() int {
	fake.deleteStorageGroupMutex.RLock()
	defer fake.deleteStorageGroupMutex.RUnlock()
	return len(fake.deleteStorageGroupArgsForCall)
}

func (fake *ZhmcAPI) DeleteStorageGroupCalls(stub func(string) (int, *zhmcclient.HmcError)) {
	fake.deleteStorageGroupMutex.Lock()
	defer fake.deleteStorageGroupMutex.Unlock()
	fake.DeleteStorageGroupStub = stub
}

func (fake *ZhmcAPI) DeleteStorageGroupArgsForCall(i int) string {
	fake.deleteStorageGroupMutex.RLock()
	defer fake.deleteStorageGroupMutex.RUnlock()
	argsForCall := fake.deleteStorageGroupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) DeleteStorageGroupReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.deleteStorageGroupMutex.Lock()
	defer fake.deleteStorageGroupMutex.Unlock()
	fake.DeleteStorageGroupStub = nil
	fake.deleteStorageGroupReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) DeleteStorageGroupReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.deleteStorageGroupMutex.Lock()
	defer fake.deleteStorageGroupMutex.Unlock()
	fake.DeleteStorageGroupStub = nil
	if fake.deleteStorageGroupReturnsOnCall == nil {
		fake.deleteStorageGroupReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.deleteStorageGroupReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) DetachStorageGroupToPartition(arg1 string, arg2 *zhmcclient.StorageGroupPayload) (int, *zhmcclient.HmcError) {
	fake.detachStorageGroupToPartitionMutex.Lock()
	ret, specificReturn := fake.detachStorageGroupToPartitionReturnsOnCall[len(fake.detachStorageGroupToPartitionArgsForCall)]
	fake.detachStorageGroupToPartitionArgsForCall = append(fake.detachStorageGroupToPartitionArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.StorageGroupPayload
	}{arg1, arg2})
	stub := fake.DetachStorageGroupToPartitionStub
	fakeReturns := fake.detachStorageGroupToPartitionReturns
	fake.recordInvocation("DetachStorageGroupToPartition", []interface{}{arg1, arg2})
	fake.detachStorageGroupToPartitionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ZhmcAPI) DetachStorageGroupToPartitionCallCount() int {
	fake.detachStorageGroupToPartitionMutex.RLock()
	defer fake.detachStorageGroupToPartitionMutex.RUnlock()
	return len(fake.detachStorageGroupToPartitionArgsForCall)
}

func (fake *ZhmcAPI) DetachStorageGroupToPartitionCalls(stub func(string, *zhmcclient.StorageGroupPayload) (int, *zhmcclient.HmcError)) {
	fake.detachStorageGroupToPartitionMutex.Lock()
	defer fake.detachStorageGroupToPartitionMutex.Unlock()
	fake.DetachStorageGroupToPartitionStub = stub
}

func (fake *ZhmcAPI) DetachStorageGroupToPartitionArgsForCall(i int) (string, *zhmcclient.StorageGroupPayload) {
	fake.detachStorageGroupToPartitionMutex.RLock()
	defer fake.detachStorageGroupToPartitionMutex.RUnlock()
	argsForCall := fake.detachStorageGroupToPartitionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ZhmcAPI) DetachStorageGroupToPartitionReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.detachStorageGroupToPartitionMutex.Lock()
	defer fake.detachStorageGroupToPartitionMutex.Unlock()
	fake.DetachStorageGroupToPartitionStub = nil
	fake.detachStorageGroupToPartitionReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) DetachStorageGroupToPartitionReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.detachStorageGroupToPartitionMutex.Lock()
	defer fake.detachStorageGroupToPartitionMutex.Unlock()
	fake.DetachStorageGroupToPartitionStub = nil
	if fake.detachStorageGroupToPartitionReturnsOnCall == nil {
		fake.detachStorageGroupToPartitionReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.detachStorageGroupToPartitionReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) FetchAsciiConsoleURI(arg1 string, arg2 *zhmcclient.AsciiConsoleURIPayload) (*zhmcclient.AsciiConsoleURIResponse, int, *zhmcclient.HmcError) {
	fake.fetchAsciiConsoleURIMutex.Lock()
	ret, specificReturn := fake.fetchAsciiConsoleURIReturnsOnCall[len(fake.fetchAsciiConsoleURIArgsForCall)]
	fake.fetchAsciiConsoleURIArgsForCall = append(fake.fetchAsciiConsoleURIArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.AsciiConsoleURIPayload
	}{arg1, arg2})
	stub := fake.FetchAsciiConsoleURIStub
	fakeReturns := fake.fetchAsciiConsoleURIReturns
	fake.recordInvocation("FetchAsciiConsoleURI", []interface{}{arg1, arg2})
	fake.fetchAsciiConsoleURIMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ZhmcAPI) FetchAsciiConsoleURICallCount() int {
	fake.fetchAsciiConsoleURIMutex.RLock()
	defer fake.fetchAsciiConsoleURIMutex.RUnlock()
	return len(fake.fetchAsciiConsoleURIArgsForCall)
}

func (fake *ZhmcAPI) FetchAsciiConsoleURICalls(stub func(string, *zhmcclient.AsciiConsoleURIPayload) (*zhmcclient.AsciiConsoleURIResponse, int, *zhmcclient.HmcError)) {
	fake.fetchAsciiConsoleURIMutex.Lock()
	defer fake.fetchAsciiConsoleURIMutex.Unlock()
	fake.FetchAsciiConsoleURIStub = stub
}

func (fake *ZhmcAPI) FetchAsciiConsoleURIArgsForCall(i int) (string, *zhmcclient.AsciiConsoleURIPayload) {
	fake.fetchAsciiConsoleURIMutex.RLock()
	defer fake.fetchAsciiConsoleURIMutex.RUnlock()
	argsForCall := fake.fetchAsciiConsoleURIArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ZhmcAPI) FetchAsciiConsoleURIReturns(result1 *zhmcclient.AsciiConsoleURIResponse, result2 int, result3 *zhmcclient.HmcError) {
	fake.fetchAsciiConsoleURIMutex.Lock()
	defer fake.fetchAsciiConsoleURIMutex.Unlock()
	fake.FetchAsciiConsoleURIStub = nil
	fake.fetchAsciiConsoleURIReturns = struct {
		result1 *zhmcclient.AsciiConsoleURIResponse
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) FetchAsciiConsoleURIReturnsOnCall(i int, result1 *zhmcclient.AsciiConsoleURIResponse, result2 int, result3 *zhmcclient.HmcError) {
	fake.fetchAsciiConsoleURIMutex.Lock()
	defer fake.fetchAsciiConsoleURIMutex.Unlock()
	fake.FetchAsciiConsoleURIStub = nil
	if fake.fetchAsciiConsoleURIReturnsOnCall == nil {
		fake.fetchAsciiConsoleURIReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.AsciiConsoleURIResponse
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.fetchAsciiConsoleURIReturnsOnCall[i] = struct {
		result1 *zhmcclient.AsciiConsoleURIResponse
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) FulfillStorageGroup(arg1 string, arg2 *zhmcclient.StorageGroupProperties) (int, *zhmcclient.HmcError) {
	fake.fulfillStorageGroupMutex.Lock()
	ret, specificReturn := fake.fulfillStorageGroupReturnsOnCall[len(fake.fulfillStorageGroupArgsForCall)]
	fake.fulfillStorageGroupArgsForCall = append(fake.fulfillStorageGroupArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.StorageGroupProperties
	}{arg1, arg2})
	stub := fake.FulfillStorageGroupStub
	fakeReturns := fake.fulfillStorageGroupReturns
	fake.recordInvocation("FulfillStorageGroup", []interface{}{arg1, arg2})
	fake.fulfillStorageGroupMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ZhmcAPI) FulfillStorageGroupCallCount() int {
	fake.fulfillStorageGroupMutex.RLock()
	defer fake.fulfillStorageGroupMutex.RUnlock()
	return len(fake.fulfillStorageGroupArgsForCall)
}

func (fake *ZhmcAPI) FulfillStorageGroupCalls(stub func(string, *zhmcclient.StorageGroupProperties) (int, *zhmcclient.HmcError)) {
	fake.fulfillStorageGroupMutex.Lock()
	defer fake.fulfillStorageGroupMutex.Unlock()
	fake.FulfillStorageGroupStub = stub
}

func (fake *ZhmcAPI) FulfillStorageGroupArgsForCall(i int) (string, *zhmcclient.StorageGroupProperties) {
	fake.fulfillStorageGroupMutex.RLock()
	defer fake.fulfillStorageGroupMutex.RUnlock()
	argsForCall := fake.fulfillStorageGroupArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ZhmcAPI) FulfillStorageGroupReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.fulfillStorageGroupMutex.Lock()
	defer fake.fulfillStorageGroupMutex.Unlock()
	fake.FulfillStorageGroupStub = nil
	fake.fulfillStorageGroupReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) FulfillStorageGroupReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.fulfillStorageGroupMutex.Lock()
	defer fake.fulfillStorageGroupMutex.Unlock()
	fake.FulfillStorageGroupStub = nil
	if fake.fulfillStorageGroupReturnsOnCall == nil {
		fake.fulfillStorageGroupReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.fulfillStorageGroupReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) GetAdapterProperties(arg1 string) (*zhmcclient.AdapterProperties, int, *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) GetAdapterPropertiesCallCount() int {
	fake.getAdapterPropertiesMutex.RLock()
	defer fake.getAdapterPropertiesMutex.RUnlock()
	return len(fake.getAdapterPropertiesArgsForCall)
}

func (fake *ZhmcAPI) GetAdapterPropertiesCalls(stub func(string) (*zhmcclient.AdapterProperties, int, *zhmcclient.HmcError)) {
	fake.getAdapterPropertiesMutex.Lock()
	defer fake.getAdapterPropertiesMutex.Unlock()
	fake.GetAdapterPropertiesStub = stub
}

func (fake *ZhmcAPI) GetAdapterPropertiesArgsForCall(i int) string {
	fake.getAdapterPropertiesMutex.RLock()
	defer fake.getAdapterPropertiesMutex.RUnlock()
	argsForCall := fake.getAdapterPropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) GetAdapterPropertiesReturns(result1 *zhmcclient.AdapterProperties, result2 int, result3 *zhmcclient.HmcError) {
	fake.getAdapterPropertiesMutex.Lock()
	defer fake.getAdapterPropertiesMutex.Unlock()
	fake.GetAdapterPropertiesStub = nil
	fake.getAdapterPropertiesReturns = struct {
		result1 *zhmcclient.AdapterProperties
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) GetAdapterPropertiesReturnsOnCall(i int, result1 *zhmcclient.AdapterProperties, result2 int, result3 *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) GetLparProperties(arg1 string) (*zhmcclient.LparProperties, int, *zhmcclient.HmcError) {
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
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ZhmcAPI) GetLparPropertiesCallCount() int {
	fake.getLparPropertiesMutex.RLock()
	defer fake.getLparPropertiesMutex.RUnlock()
	return len(fake.getLparPropertiesArgsForCall)
}

func (fake *ZhmcAPI) GetLparPropertiesCalls(stub func(string) (*zhmcclient.LparProperties, int, *zhmcclient.HmcError)) {
	fake.getLparPropertiesMutex.Lock()
	defer fake.getLparPropertiesMutex.Unlock()
	fake.GetLparPropertiesStub = stub
}

func (fake *ZhmcAPI) GetLparPropertiesArgsForCall(i int) string {
	fake.getLparPropertiesMutex.RLock()
	defer fake.getLparPropertiesMutex.RUnlock()
	argsForCall := fake.getLparPropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) GetLparPropertiesReturns(result1 *zhmcclient.LparProperties, result2 int, result3 *zhmcclient.HmcError) {
	fake.getLparPropertiesMutex.Lock()
	defer fake.getLparPropertiesMutex.Unlock()
	fake.GetLparPropertiesStub = nil
	fake.getLparPropertiesReturns = struct {
		result1 *zhmcclient.LparProperties
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) GetLparPropertiesReturnsOnCall(i int, result1 *zhmcclient.LparProperties, result2 int, result3 *zhmcclient.HmcError) {
	fake.getLparPropertiesMutex.Lock()
	defer fake.getLparPropertiesMutex.Unlock()
	fake.GetLparPropertiesStub = nil
	if fake.getLparPropertiesReturnsOnCall == nil {
		fake.getLparPropertiesReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.LparProperties
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.getLparPropertiesReturnsOnCall[i] = struct {
		result1 *zhmcclient.LparProperties
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) GetNetworkAdapterPortProperties(arg1 string) (*zhmcclient.NetworkAdapterPort, int, *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) GetNetworkAdapterPortPropertiesCallCount() int {
	fake.getNetworkAdapterPortPropertiesMutex.RLock()
	defer fake.getNetworkAdapterPortPropertiesMutex.RUnlock()
	return len(fake.getNetworkAdapterPortPropertiesArgsForCall)
}

func (fake *ZhmcAPI) GetNetworkAdapterPortPropertiesCalls(stub func(string) (*zhmcclient.NetworkAdapterPort, int, *zhmcclient.HmcError)) {
	fake.getNetworkAdapterPortPropertiesMutex.Lock()
	defer fake.getNetworkAdapterPortPropertiesMutex.Unlock()
	fake.GetNetworkAdapterPortPropertiesStub = stub
}

func (fake *ZhmcAPI) GetNetworkAdapterPortPropertiesArgsForCall(i int) string {
	fake.getNetworkAdapterPortPropertiesMutex.RLock()
	defer fake.getNetworkAdapterPortPropertiesMutex.RUnlock()
	argsForCall := fake.getNetworkAdapterPortPropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) GetNetworkAdapterPortPropertiesReturns(result1 *zhmcclient.NetworkAdapterPort, result2 int, result3 *zhmcclient.HmcError) {
	fake.getNetworkAdapterPortPropertiesMutex.Lock()
	defer fake.getNetworkAdapterPortPropertiesMutex.Unlock()
	fake.GetNetworkAdapterPortPropertiesStub = nil
	fake.getNetworkAdapterPortPropertiesReturns = struct {
		result1 *zhmcclient.NetworkAdapterPort
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) GetNetworkAdapterPortPropertiesReturnsOnCall(i int, result1 *zhmcclient.NetworkAdapterPort, result2 int, result3 *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) GetNicProperties(arg1 string) (*zhmcclient.NIC, int, *zhmcclient.HmcError) {
	fake.getNicPropertiesMutex.Lock()
	ret, specificReturn := fake.getNicPropertiesReturnsOnCall[len(fake.getNicPropertiesArgsForCall)]
	fake.getNicPropertiesArgsForCall = append(fake.getNicPropertiesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetNicPropertiesStub
	fakeReturns := fake.getNicPropertiesReturns
	fake.recordInvocation("GetNicProperties", []interface{}{arg1})
	fake.getNicPropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ZhmcAPI) GetNicPropertiesCallCount() int {
	fake.getNicPropertiesMutex.RLock()
	defer fake.getNicPropertiesMutex.RUnlock()
	return len(fake.getNicPropertiesArgsForCall)
}

func (fake *ZhmcAPI) GetNicPropertiesCalls(stub func(string) (*zhmcclient.NIC, int, *zhmcclient.HmcError)) {
	fake.getNicPropertiesMutex.Lock()
	defer fake.getNicPropertiesMutex.Unlock()
	fake.GetNicPropertiesStub = stub
}

func (fake *ZhmcAPI) GetNicPropertiesArgsForCall(i int) string {
	fake.getNicPropertiesMutex.RLock()
	defer fake.getNicPropertiesMutex.RUnlock()
	argsForCall := fake.getNicPropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) GetNicPropertiesReturns(result1 *zhmcclient.NIC, result2 int, result3 *zhmcclient.HmcError) {
	fake.getNicPropertiesMutex.Lock()
	defer fake.getNicPropertiesMutex.Unlock()
	fake.GetNicPropertiesStub = nil
	fake.getNicPropertiesReturns = struct {
		result1 *zhmcclient.NIC
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) GetNicPropertiesReturnsOnCall(i int, result1 *zhmcclient.NIC, result2 int, result3 *zhmcclient.HmcError) {
	fake.getNicPropertiesMutex.Lock()
	defer fake.getNicPropertiesMutex.Unlock()
	fake.GetNicPropertiesStub = nil
	if fake.getNicPropertiesReturnsOnCall == nil {
		fake.getNicPropertiesReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.NIC
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.getNicPropertiesReturnsOnCall[i] = struct {
		result1 *zhmcclient.NIC
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) GetCPCProperties(arg1 string) (*zhmcclient.CPCProperties, int, *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) GetCPCPropertiesCallCount() int {
	fake.getCPCPropertiesMutex.RLock()
	defer fake.getCPCPropertiesMutex.RUnlock()
	return len(fake.getCPCPropertiesArgsForCall)
}

func (fake *ZhmcAPI) GetCPCPropertiesCalls(stub func(string) (*zhmcclient.CPCProperties, int, *zhmcclient.HmcError)) {
	fake.getCPCPropertiesMutex.Lock()
	defer fake.getCPCPropertiesMutex.Unlock()
	fake.GetCPCPropertiesStub = stub
}

func (fake *ZhmcAPI) GetCPCPropertiesArgsForCall(i int) string {
	fake.getCPCPropertiesMutex.RLock()
	defer fake.getCPCPropertiesMutex.RUnlock()
	argsForCall := fake.getCPCPropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) GetCPCPropertiesReturns(result1 *zhmcclient.CPCProperties, result2 int, result3 *zhmcclient.HmcError) {
	fake.getCPCPropertiesMutex.Lock()
	defer fake.getCPCPropertiesMutex.Unlock()
	fake.GetCPCPropertiesStub = nil
	fake.getCPCPropertiesReturns = struct {
		result1 *zhmcclient.CPCProperties
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) GetCPCPropertiesReturnsOnCall(i int, result1 *zhmcclient.CPCProperties, result2 int, result3 *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) GetStorageAdapterPortProperties(arg1 string) (*zhmcclient.StorageAdapterPort, int, *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) GetStorageAdapterPortPropertiesCallCount() int {
	fake.getStorageAdapterPortPropertiesMutex.RLock()
	defer fake.getStorageAdapterPortPropertiesMutex.RUnlock()
	return len(fake.getStorageAdapterPortPropertiesArgsForCall)
}

func (fake *ZhmcAPI) GetStorageAdapterPortPropertiesCalls(stub func(string) (*zhmcclient.StorageAdapterPort, int, *zhmcclient.HmcError)) {
	fake.getStorageAdapterPortPropertiesMutex.Lock()
	defer fake.getStorageAdapterPortPropertiesMutex.Unlock()
	fake.GetStorageAdapterPortPropertiesStub = stub
}

func (fake *ZhmcAPI) GetStorageAdapterPortPropertiesArgsForCall(i int) string {
	fake.getStorageAdapterPortPropertiesMutex.RLock()
	defer fake.getStorageAdapterPortPropertiesMutex.RUnlock()
	argsForCall := fake.getStorageAdapterPortPropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) GetStorageAdapterPortPropertiesReturns(result1 *zhmcclient.StorageAdapterPort, result2 int, result3 *zhmcclient.HmcError) {
	fake.getStorageAdapterPortPropertiesMutex.Lock()
	defer fake.getStorageAdapterPortPropertiesMutex.Unlock()
	fake.GetStorageAdapterPortPropertiesStub = nil
	fake.getStorageAdapterPortPropertiesReturns = struct {
		result1 *zhmcclient.StorageAdapterPort
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) GetStorageAdapterPortPropertiesReturnsOnCall(i int, result1 *zhmcclient.StorageAdapterPort, result2 int, result3 *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) GetStorageGroupPartitions(arg1 string, arg2 map[string]string) (*zhmcclient.StorageGroupPartitions, int, *zhmcclient.HmcError) {
	fake.getStorageGroupPartitionsMutex.Lock()
	ret, specificReturn := fake.getStorageGroupPartitionsReturnsOnCall[len(fake.getStorageGroupPartitionsArgsForCall)]
	fake.getStorageGroupPartitionsArgsForCall = append(fake.getStorageGroupPartitionsArgsForCall, struct {
		arg1 string
		arg2 map[string]string
	}{arg1, arg2})
	stub := fake.GetStorageGroupPartitionsStub
	fakeReturns := fake.getStorageGroupPartitionsReturns
	fake.recordInvocation("GetStorageGroupPartitions", []interface{}{arg1, arg2})
	fake.getStorageGroupPartitionsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ZhmcAPI) GetStorageGroupPartitionsCallCount() int {
	fake.getStorageGroupPartitionsMutex.RLock()
	defer fake.getStorageGroupPartitionsMutex.RUnlock()
	return len(fake.getStorageGroupPartitionsArgsForCall)
}

func (fake *ZhmcAPI) GetStorageGroupPartitionsCalls(stub func(string, map[string]string) (*zhmcclient.StorageGroupPartitions, int, *zhmcclient.HmcError)) {
	fake.getStorageGroupPartitionsMutex.Lock()
	defer fake.getStorageGroupPartitionsMutex.Unlock()
	fake.GetStorageGroupPartitionsStub = stub
}

func (fake *ZhmcAPI) GetStorageGroupPartitionsArgsForCall(i int) (string, map[string]string) {
	fake.getStorageGroupPartitionsMutex.RLock()
	defer fake.getStorageGroupPartitionsMutex.RUnlock()
	argsForCall := fake.getStorageGroupPartitionsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ZhmcAPI) GetStorageGroupPartitionsReturns(result1 *zhmcclient.StorageGroupPartitions, result2 int, result3 *zhmcclient.HmcError) {
	fake.getStorageGroupPartitionsMutex.Lock()
	defer fake.getStorageGroupPartitionsMutex.Unlock()
	fake.GetStorageGroupPartitionsStub = nil
	fake.getStorageGroupPartitionsReturns = struct {
		result1 *zhmcclient.StorageGroupPartitions
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) GetStorageGroupPartitionsReturnsOnCall(i int, result1 *zhmcclient.StorageGroupPartitions, result2 int, result3 *zhmcclient.HmcError) {
	fake.getStorageGroupPartitionsMutex.Lock()
	defer fake.getStorageGroupPartitionsMutex.Unlock()
	fake.GetStorageGroupPartitionsStub = nil
	if fake.getStorageGroupPartitionsReturnsOnCall == nil {
		fake.getStorageGroupPartitionsReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.StorageGroupPartitions
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.getStorageGroupPartitionsReturnsOnCall[i] = struct {
		result1 *zhmcclient.StorageGroupPartitions
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) GetStorageGroupProperties(arg1 string) (*zhmcclient.StorageGroupProperties, int, *zhmcclient.HmcError) {
	fake.getStorageGroupPropertiesMutex.Lock()
	ret, specificReturn := fake.getStorageGroupPropertiesReturnsOnCall[len(fake.getStorageGroupPropertiesArgsForCall)]
	fake.getStorageGroupPropertiesArgsForCall = append(fake.getStorageGroupPropertiesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetStorageGroupPropertiesStub
	fakeReturns := fake.getStorageGroupPropertiesReturns
	fake.recordInvocation("GetStorageGroupProperties", []interface{}{arg1})
	fake.getStorageGroupPropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ZhmcAPI) GetStorageGroupPropertiesCallCount() int {
	fake.getStorageGroupPropertiesMutex.RLock()
	defer fake.getStorageGroupPropertiesMutex.RUnlock()
	return len(fake.getStorageGroupPropertiesArgsForCall)
}

func (fake *ZhmcAPI) GetStorageGroupPropertiesCalls(stub func(string) (*zhmcclient.StorageGroupProperties, int, *zhmcclient.HmcError)) {
	fake.getStorageGroupPropertiesMutex.Lock()
	defer fake.getStorageGroupPropertiesMutex.Unlock()
	fake.GetStorageGroupPropertiesStub = stub
}

func (fake *ZhmcAPI) GetStorageGroupPropertiesArgsForCall(i int) string {
	fake.getStorageGroupPropertiesMutex.RLock()
	defer fake.getStorageGroupPropertiesMutex.RUnlock()
	argsForCall := fake.getStorageGroupPropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) GetStorageGroupPropertiesReturns(result1 *zhmcclient.StorageGroupProperties, result2 int, result3 *zhmcclient.HmcError) {
	fake.getStorageGroupPropertiesMutex.Lock()
	defer fake.getStorageGroupPropertiesMutex.Unlock()
	fake.GetStorageGroupPropertiesStub = nil
	fake.getStorageGroupPropertiesReturns = struct {
		result1 *zhmcclient.StorageGroupProperties
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) GetStorageGroupPropertiesReturnsOnCall(i int, result1 *zhmcclient.StorageGroupProperties, result2 int, result3 *zhmcclient.HmcError) {
	fake.getStorageGroupPropertiesMutex.Lock()
	defer fake.getStorageGroupPropertiesMutex.Unlock()
	fake.GetStorageGroupPropertiesStub = nil
	if fake.getStorageGroupPropertiesReturnsOnCall == nil {
		fake.getStorageGroupPropertiesReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.StorageGroupProperties
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.getStorageGroupPropertiesReturnsOnCall[i] = struct {
		result1 *zhmcclient.StorageGroupProperties
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) GetStorageVolumeProperties(arg1 string) (*zhmcclient.StorageVolume, int, *zhmcclient.HmcError) {
	fake.getStorageVolumePropertiesMutex.Lock()
	ret, specificReturn := fake.getStorageVolumePropertiesReturnsOnCall[len(fake.getStorageVolumePropertiesArgsForCall)]
	fake.getStorageVolumePropertiesArgsForCall = append(fake.getStorageVolumePropertiesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetStorageVolumePropertiesStub
	fakeReturns := fake.getStorageVolumePropertiesReturns
	fake.recordInvocation("GetStorageVolumeProperties", []interface{}{arg1})
	fake.getStorageVolumePropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ZhmcAPI) GetStorageVolumePropertiesCallCount() int {
	fake.getStorageVolumePropertiesMutex.RLock()
	defer fake.getStorageVolumePropertiesMutex.RUnlock()
	return len(fake.getStorageVolumePropertiesArgsForCall)
}

func (fake *ZhmcAPI) GetStorageVolumePropertiesCalls(stub func(string) (*zhmcclient.StorageVolume, int, *zhmcclient.HmcError)) {
	fake.getStorageVolumePropertiesMutex.Lock()
	defer fake.getStorageVolumePropertiesMutex.Unlock()
	fake.GetStorageVolumePropertiesStub = stub
}

func (fake *ZhmcAPI) GetStorageVolumePropertiesArgsForCall(i int) string {
	fake.getStorageVolumePropertiesMutex.RLock()
	defer fake.getStorageVolumePropertiesMutex.RUnlock()
	argsForCall := fake.getStorageVolumePropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) GetStorageVolumePropertiesReturns(result1 *zhmcclient.StorageVolume, result2 int, result3 *zhmcclient.HmcError) {
	fake.getStorageVolumePropertiesMutex.Lock()
	defer fake.getStorageVolumePropertiesMutex.Unlock()
	fake.GetStorageVolumePropertiesStub = nil
	fake.getStorageVolumePropertiesReturns = struct {
		result1 *zhmcclient.StorageVolume
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) GetStorageVolumePropertiesReturnsOnCall(i int, result1 *zhmcclient.StorageVolume, result2 int, result3 *zhmcclient.HmcError) {
	fake.getStorageVolumePropertiesMutex.Lock()
	defer fake.getStorageVolumePropertiesMutex.Unlock()
	fake.GetStorageVolumePropertiesStub = nil
	if fake.getStorageVolumePropertiesReturnsOnCall == nil {
		fake.getStorageVolumePropertiesReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.StorageVolume
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.getStorageVolumePropertiesReturnsOnCall[i] = struct {
		result1 *zhmcclient.StorageVolume
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) GetVirtualSwitchProperties(arg1 string) (*zhmcclient.VirtualSwitchProperties, int, *zhmcclient.HmcError) {
	fake.getVirtualSwitchPropertiesMutex.Lock()
	ret, specificReturn := fake.getVirtualSwitchPropertiesReturnsOnCall[len(fake.getVirtualSwitchPropertiesArgsForCall)]
	fake.getVirtualSwitchPropertiesArgsForCall = append(fake.getVirtualSwitchPropertiesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetVirtualSwitchPropertiesStub
	fakeReturns := fake.getVirtualSwitchPropertiesReturns
	fake.recordInvocation("GetVirtualSwitchProperties", []interface{}{arg1})
	fake.getVirtualSwitchPropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ZhmcAPI) GetVirtualSwitchPropertiesCallCount() int {
	fake.getVirtualSwitchPropertiesMutex.RLock()
	defer fake.getVirtualSwitchPropertiesMutex.RUnlock()
	return len(fake.getVirtualSwitchPropertiesArgsForCall)
}

func (fake *ZhmcAPI) GetVirtualSwitchPropertiesCalls(stub func(string) (*zhmcclient.VirtualSwitchProperties, int, *zhmcclient.HmcError)) {
	fake.getVirtualSwitchPropertiesMutex.Lock()
	defer fake.getVirtualSwitchPropertiesMutex.Unlock()
	fake.GetVirtualSwitchPropertiesStub = stub
}

func (fake *ZhmcAPI) GetVirtualSwitchPropertiesArgsForCall(i int) string {
	fake.getVirtualSwitchPropertiesMutex.RLock()
	defer fake.getVirtualSwitchPropertiesMutex.RUnlock()
	argsForCall := fake.getVirtualSwitchPropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) GetVirtualSwitchPropertiesReturns(result1 *zhmcclient.VirtualSwitchProperties, result2 int, result3 *zhmcclient.HmcError) {
	fake.getVirtualSwitchPropertiesMutex.Lock()
	defer fake.getVirtualSwitchPropertiesMutex.Unlock()
	fake.GetVirtualSwitchPropertiesStub = nil
	fake.getVirtualSwitchPropertiesReturns = struct {
		result1 *zhmcclient.VirtualSwitchProperties
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) GetVirtualSwitchPropertiesReturnsOnCall(i int, result1 *zhmcclient.VirtualSwitchProperties, result2 int, result3 *zhmcclient.HmcError) {
	fake.getVirtualSwitchPropertiesMutex.Lock()
	defer fake.getVirtualSwitchPropertiesMutex.Unlock()
	fake.GetVirtualSwitchPropertiesStub = nil
	if fake.getVirtualSwitchPropertiesReturnsOnCall == nil {
		fake.getVirtualSwitchPropertiesReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.VirtualSwitchProperties
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.getVirtualSwitchPropertiesReturnsOnCall[i] = struct {
		result1 *zhmcclient.VirtualSwitchProperties
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) ListAdapters(arg1 string, arg2 map[string]string) ([]zhmcclient.Adapter, int, *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) ListAdaptersCallCount() int {
	fake.listAdaptersMutex.RLock()
	defer fake.listAdaptersMutex.RUnlock()
	return len(fake.listAdaptersArgsForCall)
}

func (fake *ZhmcAPI) ListAdaptersCalls(stub func(string, map[string]string) ([]zhmcclient.Adapter, int, *zhmcclient.HmcError)) {
	fake.listAdaptersMutex.Lock()
	defer fake.listAdaptersMutex.Unlock()
	fake.ListAdaptersStub = stub
}

func (fake *ZhmcAPI) ListAdaptersArgsForCall(i int) (string, map[string]string) {
	fake.listAdaptersMutex.RLock()
	defer fake.listAdaptersMutex.RUnlock()
	argsForCall := fake.listAdaptersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ZhmcAPI) ListAdaptersReturns(result1 []zhmcclient.Adapter, result2 int, result3 *zhmcclient.HmcError) {
	fake.listAdaptersMutex.Lock()
	defer fake.listAdaptersMutex.Unlock()
	fake.ListAdaptersStub = nil
	fake.listAdaptersReturns = struct {
		result1 []zhmcclient.Adapter
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) ListAdaptersReturnsOnCall(i int, result1 []zhmcclient.Adapter, result2 int, result3 *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) ListCPCs(arg1 map[string]string) ([]zhmcclient.CPC, int, *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) ListCPCsCallCount() int {
	fake.listCPCsMutex.RLock()
	defer fake.listCPCsMutex.RUnlock()
	return len(fake.listCPCsArgsForCall)
}

func (fake *ZhmcAPI) ListCPCsCalls(stub func(map[string]string) ([]zhmcclient.CPC, int, *zhmcclient.HmcError)) {
	fake.listCPCsMutex.Lock()
	defer fake.listCPCsMutex.Unlock()
	fake.ListCPCsStub = stub
}

func (fake *ZhmcAPI) ListCPCsArgsForCall(i int) map[string]string {
	fake.listCPCsMutex.RLock()
	defer fake.listCPCsMutex.RUnlock()
	argsForCall := fake.listCPCsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) ListCPCsReturns(result1 []zhmcclient.CPC, result2 int, result3 *zhmcclient.HmcError) {
	fake.listCPCsMutex.Lock()
	defer fake.listCPCsMutex.Unlock()
	fake.ListCPCsStub = nil
	fake.listCPCsReturns = struct {
		result1 []zhmcclient.CPC
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) ListCPCsReturnsOnCall(i int, result1 []zhmcclient.CPC, result2 int, result3 *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) ListLPARs(arg1 string, arg2 map[string]string) ([]zhmcclient.LPAR, int, *zhmcclient.HmcError) {
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
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ZhmcAPI) ListLPARsCallCount() int {
	fake.listLPARsMutex.RLock()
	defer fake.listLPARsMutex.RUnlock()
	return len(fake.listLPARsArgsForCall)
}

func (fake *ZhmcAPI) ListLPARsCalls(stub func(string, map[string]string) ([]zhmcclient.LPAR, int, *zhmcclient.HmcError)) {
	fake.listLPARsMutex.Lock()
	defer fake.listLPARsMutex.Unlock()
	fake.ListLPARsStub = stub
}

func (fake *ZhmcAPI) ListLPARsArgsForCall(i int) (string, map[string]string) {
	fake.listLPARsMutex.RLock()
	defer fake.listLPARsMutex.RUnlock()
	argsForCall := fake.listLPARsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ZhmcAPI) ListLPARsReturns(result1 []zhmcclient.LPAR, result2 int, result3 *zhmcclient.HmcError) {
	fake.listLPARsMutex.Lock()
	defer fake.listLPARsMutex.Unlock()
	fake.ListLPARsStub = nil
	fake.listLPARsReturns = struct {
		result1 []zhmcclient.LPAR
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) ListLPARsReturnsOnCall(i int, result1 []zhmcclient.LPAR, result2 int, result3 *zhmcclient.HmcError) {
	fake.listLPARsMutex.Lock()
	defer fake.listLPARsMutex.Unlock()
	fake.ListLPARsStub = nil
	if fake.listLPARsReturnsOnCall == nil {
		fake.listLPARsReturnsOnCall = make(map[int]struct {
			result1 []zhmcclient.LPAR
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.listLPARsReturnsOnCall[i] = struct {
		result1 []zhmcclient.LPAR
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) ListNics(arg1 string) ([]string, int, *zhmcclient.HmcError) {
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
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ZhmcAPI) ListNicsCallCount() int {
	fake.listNicsMutex.RLock()
	defer fake.listNicsMutex.RUnlock()
	return len(fake.listNicsArgsForCall)
}

func (fake *ZhmcAPI) ListNicsCalls(stub func(string) ([]string, int, *zhmcclient.HmcError)) {
	fake.listNicsMutex.Lock()
	defer fake.listNicsMutex.Unlock()
	fake.ListNicsStub = stub
}

func (fake *ZhmcAPI) ListNicsArgsForCall(i int) string {
	fake.listNicsMutex.RLock()
	defer fake.listNicsMutex.RUnlock()
	argsForCall := fake.listNicsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) ListNicsReturns(result1 []string, result2 int, result3 *zhmcclient.HmcError) {
	fake.listNicsMutex.Lock()
	defer fake.listNicsMutex.Unlock()
	fake.ListNicsStub = nil
	fake.listNicsReturns = struct {
		result1 []string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) ListNicsReturnsOnCall(i int, result1 []string, result2 int, result3 *zhmcclient.HmcError) {
	fake.listNicsMutex.Lock()
	defer fake.listNicsMutex.Unlock()
	fake.ListNicsStub = nil
	if fake.listNicsReturnsOnCall == nil {
		fake.listNicsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.listNicsReturnsOnCall[i] = struct {
		result1 []string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) ListStorageGroups(arg1 string, arg2 string) ([]zhmcclient.StorageGroup, int, *zhmcclient.HmcError) {
	fake.listStorageGroupsMutex.Lock()
	ret, specificReturn := fake.listStorageGroupsReturnsOnCall[len(fake.listStorageGroupsArgsForCall)]
	fake.listStorageGroupsArgsForCall = append(fake.listStorageGroupsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.ListStorageGroupsStub
	fakeReturns := fake.listStorageGroupsReturns
	fake.recordInvocation("ListStorageGroups", []interface{}{arg1, arg2})
	fake.listStorageGroupsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ZhmcAPI) ListStorageGroupsCallCount() int {
	fake.listStorageGroupsMutex.RLock()
	defer fake.listStorageGroupsMutex.RUnlock()
	return len(fake.listStorageGroupsArgsForCall)
}

func (fake *ZhmcAPI) ListStorageGroupsCalls(stub func(string, string) ([]zhmcclient.StorageGroup, int, *zhmcclient.HmcError)) {
	fake.listStorageGroupsMutex.Lock()
	defer fake.listStorageGroupsMutex.Unlock()
	fake.ListStorageGroupsStub = stub
}

func (fake *ZhmcAPI) ListStorageGroupsArgsForCall(i int) (string, string) {
	fake.listStorageGroupsMutex.RLock()
	defer fake.listStorageGroupsMutex.RUnlock()
	argsForCall := fake.listStorageGroupsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ZhmcAPI) ListStorageGroupsReturns(result1 []zhmcclient.StorageGroup, result2 int, result3 *zhmcclient.HmcError) {
	fake.listStorageGroupsMutex.Lock()
	defer fake.listStorageGroupsMutex.Unlock()
	fake.ListStorageGroupsStub = nil
	fake.listStorageGroupsReturns = struct {
		result1 []zhmcclient.StorageGroup
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) ListStorageGroupsReturnsOnCall(i int, result1 []zhmcclient.StorageGroup, result2 int, result3 *zhmcclient.HmcError) {
	fake.listStorageGroupsMutex.Lock()
	defer fake.listStorageGroupsMutex.Unlock()
	fake.ListStorageGroupsStub = nil
	if fake.listStorageGroupsReturnsOnCall == nil {
		fake.listStorageGroupsReturnsOnCall = make(map[int]struct {
			result1 []zhmcclient.StorageGroup
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.listStorageGroupsReturnsOnCall[i] = struct {
		result1 []zhmcclient.StorageGroup
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) ListStorageVolumes(arg1 string) ([]zhmcclient.StorageVolume, int, *zhmcclient.HmcError) {
	fake.listStorageVolumesMutex.Lock()
	ret, specificReturn := fake.listStorageVolumesReturnsOnCall[len(fake.listStorageVolumesArgsForCall)]
	fake.listStorageVolumesArgsForCall = append(fake.listStorageVolumesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ListStorageVolumesStub
	fakeReturns := fake.listStorageVolumesReturns
	fake.recordInvocation("ListStorageVolumes", []interface{}{arg1})
	fake.listStorageVolumesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ZhmcAPI) ListStorageVolumesCallCount() int {
	fake.listStorageVolumesMutex.RLock()
	defer fake.listStorageVolumesMutex.RUnlock()
	return len(fake.listStorageVolumesArgsForCall)
}

func (fake *ZhmcAPI) ListStorageVolumesCalls(stub func(string) ([]zhmcclient.StorageVolume, int, *zhmcclient.HmcError)) {
	fake.listStorageVolumesMutex.Lock()
	defer fake.listStorageVolumesMutex.Unlock()
	fake.ListStorageVolumesStub = stub
}

func (fake *ZhmcAPI) ListStorageVolumesArgsForCall(i int) string {
	fake.listStorageVolumesMutex.RLock()
	defer fake.listStorageVolumesMutex.RUnlock()
	argsForCall := fake.listStorageVolumesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) ListStorageVolumesReturns(result1 []zhmcclient.StorageVolume, result2 int, result3 *zhmcclient.HmcError) {
	fake.listStorageVolumesMutex.Lock()
	defer fake.listStorageVolumesMutex.Unlock()
	fake.ListStorageVolumesStub = nil
	fake.listStorageVolumesReturns = struct {
		result1 []zhmcclient.StorageVolume
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) ListStorageVolumesReturnsOnCall(i int, result1 []zhmcclient.StorageVolume, result2 int, result3 *zhmcclient.HmcError) {
	fake.listStorageVolumesMutex.Lock()
	defer fake.listStorageVolumesMutex.Unlock()
	fake.ListStorageVolumesStub = nil
	if fake.listStorageVolumesReturnsOnCall == nil {
		fake.listStorageVolumesReturnsOnCall = make(map[int]struct {
			result1 []zhmcclient.StorageVolume
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.listStorageVolumesReturnsOnCall[i] = struct {
		result1 []zhmcclient.StorageVolume
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) ListVirtualSwitches(arg1 string, arg2 map[string]string) ([]zhmcclient.VirtualSwitch, int, *zhmcclient.HmcError) {
	fake.listVirtualSwitchesMutex.Lock()
	ret, specificReturn := fake.listVirtualSwitchesReturnsOnCall[len(fake.listVirtualSwitchesArgsForCall)]
	fake.listVirtualSwitchesArgsForCall = append(fake.listVirtualSwitchesArgsForCall, struct {
		arg1 string
		arg2 map[string]string
	}{arg1, arg2})
	stub := fake.ListVirtualSwitchesStub
	fakeReturns := fake.listVirtualSwitchesReturns
	fake.recordInvocation("ListVirtualSwitches", []interface{}{arg1, arg2})
	fake.listVirtualSwitchesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ZhmcAPI) ListVirtualSwitchesCallCount() int {
	fake.listVirtualSwitchesMutex.RLock()
	defer fake.listVirtualSwitchesMutex.RUnlock()
	return len(fake.listVirtualSwitchesArgsForCall)
}

func (fake *ZhmcAPI) ListVirtualSwitchesCalls(stub func(string, map[string]string) ([]zhmcclient.VirtualSwitch, int, *zhmcclient.HmcError)) {
	fake.listVirtualSwitchesMutex.Lock()
	defer fake.listVirtualSwitchesMutex.Unlock()
	fake.ListVirtualSwitchesStub = stub
}

func (fake *ZhmcAPI) ListVirtualSwitchesArgsForCall(i int) (string, map[string]string) {
	fake.listVirtualSwitchesMutex.RLock()
	defer fake.listVirtualSwitchesMutex.RUnlock()
	argsForCall := fake.listVirtualSwitchesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ZhmcAPI) ListVirtualSwitchesReturns(result1 []zhmcclient.VirtualSwitch, result2 int, result3 *zhmcclient.HmcError) {
	fake.listVirtualSwitchesMutex.Lock()
	defer fake.listVirtualSwitchesMutex.Unlock()
	fake.ListVirtualSwitchesStub = nil
	fake.listVirtualSwitchesReturns = struct {
		result1 []zhmcclient.VirtualSwitch
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) ListVirtualSwitchesReturnsOnCall(i int, result1 []zhmcclient.VirtualSwitch, result2 int, result3 *zhmcclient.HmcError) {
	fake.listVirtualSwitchesMutex.Lock()
	defer fake.listVirtualSwitchesMutex.Unlock()
	fake.ListVirtualSwitchesStub = nil
	if fake.listVirtualSwitchesReturnsOnCall == nil {
		fake.listVirtualSwitchesReturnsOnCall = make(map[int]struct {
			result1 []zhmcclient.VirtualSwitch
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.listVirtualSwitchesReturnsOnCall[i] = struct {
		result1 []zhmcclient.VirtualSwitch
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) MountIsoImage(arg1 string, arg2 string, arg3 string) (int, *zhmcclient.HmcError) {
	fake.mountIsoImageMutex.Lock()
	ret, specificReturn := fake.mountIsoImageReturnsOnCall[len(fake.mountIsoImageArgsForCall)]
	fake.mountIsoImageArgsForCall = append(fake.mountIsoImageArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.MountIsoImageStub
	fakeReturns := fake.mountIsoImageReturns
	fake.recordInvocation("MountIsoImage", []interface{}{arg1, arg2, arg3})
	fake.mountIsoImageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ZhmcAPI) MountIsoImageCallCount() int {
	fake.mountIsoImageMutex.RLock()
	defer fake.mountIsoImageMutex.RUnlock()
	return len(fake.mountIsoImageArgsForCall)
}

func (fake *ZhmcAPI) MountIsoImageCalls(stub func(string, string, string) (int, *zhmcclient.HmcError)) {
	fake.mountIsoImageMutex.Lock()
	defer fake.mountIsoImageMutex.Unlock()
	fake.MountIsoImageStub = stub
}

func (fake *ZhmcAPI) MountIsoImageArgsForCall(i int) (string, string, string) {
	fake.mountIsoImageMutex.RLock()
	defer fake.mountIsoImageMutex.RUnlock()
	argsForCall := fake.mountIsoImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ZhmcAPI) MountIsoImageReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.mountIsoImageMutex.Lock()
	defer fake.mountIsoImageMutex.Unlock()
	fake.MountIsoImageStub = nil
	fake.mountIsoImageReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) MountIsoImageReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.mountIsoImageMutex.Lock()
	defer fake.mountIsoImageMutex.Unlock()
	fake.MountIsoImageStub = nil
	if fake.mountIsoImageReturnsOnCall == nil {
		fake.mountIsoImageReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.mountIsoImageReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) QueryJob(arg1 string) (*zhmcclient.Job, int, *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) QueryJobCallCount() int {
	fake.queryJobMutex.RLock()
	defer fake.queryJobMutex.RUnlock()
	return len(fake.queryJobArgsForCall)
}

func (fake *ZhmcAPI) QueryJobCalls(stub func(string) (*zhmcclient.Job, int, *zhmcclient.HmcError)) {
	fake.queryJobMutex.Lock()
	defer fake.queryJobMutex.Unlock()
	fake.QueryJobStub = stub
}

func (fake *ZhmcAPI) QueryJobArgsForCall(i int) string {
	fake.queryJobMutex.RLock()
	defer fake.queryJobMutex.RUnlock()
	argsForCall := fake.queryJobArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) QueryJobReturns(result1 *zhmcclient.Job, result2 int, result3 *zhmcclient.HmcError) {
	fake.queryJobMutex.Lock()
	defer fake.queryJobMutex.Unlock()
	fake.QueryJobStub = nil
	fake.queryJobReturns = struct {
		result1 *zhmcclient.Job
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) QueryJobReturnsOnCall(i int, result1 *zhmcclient.Job, result2 int, result3 *zhmcclient.HmcError) {
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

func (fake *ZhmcAPI) StartLPAR(arg1 string) (string, int, *zhmcclient.HmcError) {
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
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ZhmcAPI) StartLPARCallCount() int {
	fake.startLPARMutex.RLock()
	defer fake.startLPARMutex.RUnlock()
	return len(fake.startLPARArgsForCall)
}

func (fake *ZhmcAPI) StartLPARCalls(stub func(string) (string, int, *zhmcclient.HmcError)) {
	fake.startLPARMutex.Lock()
	defer fake.startLPARMutex.Unlock()
	fake.StartLPARStub = stub
}

func (fake *ZhmcAPI) StartLPARArgsForCall(i int) string {
	fake.startLPARMutex.RLock()
	defer fake.startLPARMutex.RUnlock()
	argsForCall := fake.startLPARArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) StartLPARReturns(result1 string, result2 int, result3 *zhmcclient.HmcError) {
	fake.startLPARMutex.Lock()
	defer fake.startLPARMutex.Unlock()
	fake.StartLPARStub = nil
	fake.startLPARReturns = struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) StartLPARReturnsOnCall(i int, result1 string, result2 int, result3 *zhmcclient.HmcError) {
	fake.startLPARMutex.Lock()
	defer fake.startLPARMutex.Unlock()
	fake.StartLPARStub = nil
	if fake.startLPARReturnsOnCall == nil {
		fake.startLPARReturnsOnCall = make(map[int]struct {
			result1 string
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.startLPARReturnsOnCall[i] = struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) StopLPAR(arg1 string) (string, int, *zhmcclient.HmcError) {
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
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ZhmcAPI) StopLPARCallCount() int {
	fake.stopLPARMutex.RLock()
	defer fake.stopLPARMutex.RUnlock()
	return len(fake.stopLPARArgsForCall)
}

func (fake *ZhmcAPI) StopLPARCalls(stub func(string) (string, int, *zhmcclient.HmcError)) {
	fake.stopLPARMutex.Lock()
	defer fake.stopLPARMutex.Unlock()
	fake.StopLPARStub = stub
}

func (fake *ZhmcAPI) StopLPARArgsForCall(i int) string {
	fake.stopLPARMutex.RLock()
	defer fake.stopLPARMutex.RUnlock()
	argsForCall := fake.stopLPARArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) StopLPARReturns(result1 string, result2 int, result3 *zhmcclient.HmcError) {
	fake.stopLPARMutex.Lock()
	defer fake.stopLPARMutex.Unlock()
	fake.StopLPARStub = nil
	fake.stopLPARReturns = struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) StopLPARReturnsOnCall(i int, result1 string, result2 int, result3 *zhmcclient.HmcError) {
	fake.stopLPARMutex.Lock()
	defer fake.stopLPARMutex.Unlock()
	fake.StopLPARStub = nil
	if fake.stopLPARReturnsOnCall == nil {
		fake.stopLPARReturnsOnCall = make(map[int]struct {
			result1 string
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.stopLPARReturnsOnCall[i] = struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *ZhmcAPI) UnmountIsoImage(arg1 string) (int, *zhmcclient.HmcError) {
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
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ZhmcAPI) UnmountIsoImageCallCount() int {
	fake.unmountIsoImageMutex.RLock()
	defer fake.unmountIsoImageMutex.RUnlock()
	return len(fake.unmountIsoImageArgsForCall)
}

func (fake *ZhmcAPI) UnmountIsoImageCalls(stub func(string) (int, *zhmcclient.HmcError)) {
	fake.unmountIsoImageMutex.Lock()
	defer fake.unmountIsoImageMutex.Unlock()
	fake.UnmountIsoImageStub = stub
}

func (fake *ZhmcAPI) UnmountIsoImageArgsForCall(i int) string {
	fake.unmountIsoImageMutex.RLock()
	defer fake.unmountIsoImageMutex.RUnlock()
	argsForCall := fake.unmountIsoImageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ZhmcAPI) UnmountIsoImageReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.unmountIsoImageMutex.Lock()
	defer fake.unmountIsoImageMutex.Unlock()
	fake.UnmountIsoImageStub = nil
	fake.unmountIsoImageReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) UnmountIsoImageReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.unmountIsoImageMutex.Lock()
	defer fake.unmountIsoImageMutex.Unlock()
	fake.UnmountIsoImageStub = nil
	if fake.unmountIsoImageReturnsOnCall == nil {
		fake.unmountIsoImageReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.unmountIsoImageReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) UpdateLparProperties(arg1 string, arg2 *zhmcclient.LparProperties) (int, *zhmcclient.HmcError) {
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
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ZhmcAPI) UpdateLparPropertiesCallCount() int {
	fake.updateLparPropertiesMutex.RLock()
	defer fake.updateLparPropertiesMutex.RUnlock()
	return len(fake.updateLparPropertiesArgsForCall)
}

func (fake *ZhmcAPI) UpdateLparPropertiesCalls(stub func(string, *zhmcclient.LparProperties) (int, *zhmcclient.HmcError)) {
	fake.updateLparPropertiesMutex.Lock()
	defer fake.updateLparPropertiesMutex.Unlock()
	fake.UpdateLparPropertiesStub = stub
}

func (fake *ZhmcAPI) UpdateLparPropertiesArgsForCall(i int) (string, *zhmcclient.LparProperties) {
	fake.updateLparPropertiesMutex.RLock()
	defer fake.updateLparPropertiesMutex.RUnlock()
	argsForCall := fake.updateLparPropertiesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ZhmcAPI) UpdateLparPropertiesReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.updateLparPropertiesMutex.Lock()
	defer fake.updateLparPropertiesMutex.Unlock()
	fake.UpdateLparPropertiesStub = nil
	fake.updateLparPropertiesReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) UpdateLparPropertiesReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.updateLparPropertiesMutex.Lock()
	defer fake.updateLparPropertiesMutex.Unlock()
	fake.UpdateLparPropertiesStub = nil
	if fake.updateLparPropertiesReturnsOnCall == nil {
		fake.updateLparPropertiesReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.updateLparPropertiesReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) UpdateNicProperties(arg1 string, arg2 *zhmcclient.NIC) (int, *zhmcclient.HmcError) {
	fake.updateNicPropertiesMutex.Lock()
	ret, specificReturn := fake.updateNicPropertiesReturnsOnCall[len(fake.updateNicPropertiesArgsForCall)]
	fake.updateNicPropertiesArgsForCall = append(fake.updateNicPropertiesArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.NIC
	}{arg1, arg2})
	stub := fake.UpdateNicPropertiesStub
	fakeReturns := fake.updateNicPropertiesReturns
	fake.recordInvocation("UpdateNicProperties", []interface{}{arg1, arg2})
	fake.updateNicPropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ZhmcAPI) UpdateNicPropertiesCallCount() int {
	fake.updateNicPropertiesMutex.RLock()
	defer fake.updateNicPropertiesMutex.RUnlock()
	return len(fake.updateNicPropertiesArgsForCall)
}

func (fake *ZhmcAPI) UpdateNicPropertiesCalls(stub func(string, *zhmcclient.NIC) (int, *zhmcclient.HmcError)) {
	fake.updateNicPropertiesMutex.Lock()
	defer fake.updateNicPropertiesMutex.Unlock()
	fake.UpdateNicPropertiesStub = stub
}

func (fake *ZhmcAPI) UpdateNicPropertiesArgsForCall(i int) (string, *zhmcclient.NIC) {
	fake.updateNicPropertiesMutex.RLock()
	defer fake.updateNicPropertiesMutex.RUnlock()
	argsForCall := fake.updateNicPropertiesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ZhmcAPI) UpdateNicPropertiesReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.updateNicPropertiesMutex.Lock()
	defer fake.updateNicPropertiesMutex.Unlock()
	fake.UpdateNicPropertiesStub = nil
	fake.updateNicPropertiesReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) UpdateNicPropertiesReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.updateNicPropertiesMutex.Lock()
	defer fake.updateNicPropertiesMutex.Unlock()
	fake.UpdateNicPropertiesStub = nil
	if fake.updateNicPropertiesReturnsOnCall == nil {
		fake.updateNicPropertiesReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.updateNicPropertiesReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) UpdateStorageGroupProperties(arg1 string, arg2 *zhmcclient.StorageGroupProperties) (int, *zhmcclient.HmcError) {
	fake.updateStorageGroupPropertiesMutex.Lock()
	ret, specificReturn := fake.updateStorageGroupPropertiesReturnsOnCall[len(fake.updateStorageGroupPropertiesArgsForCall)]
	fake.updateStorageGroupPropertiesArgsForCall = append(fake.updateStorageGroupPropertiesArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.StorageGroupProperties
	}{arg1, arg2})
	stub := fake.UpdateStorageGroupPropertiesStub
	fakeReturns := fake.updateStorageGroupPropertiesReturns
	fake.recordInvocation("UpdateStorageGroupProperties", []interface{}{arg1, arg2})
	fake.updateStorageGroupPropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ZhmcAPI) UpdateStorageGroupPropertiesCallCount() int {
	fake.updateStorageGroupPropertiesMutex.RLock()
	defer fake.updateStorageGroupPropertiesMutex.RUnlock()
	return len(fake.updateStorageGroupPropertiesArgsForCall)
}

func (fake *ZhmcAPI) UpdateStorageGroupPropertiesCalls(stub func(string, *zhmcclient.StorageGroupProperties) (int, *zhmcclient.HmcError)) {
	fake.updateStorageGroupPropertiesMutex.Lock()
	defer fake.updateStorageGroupPropertiesMutex.Unlock()
	fake.UpdateStorageGroupPropertiesStub = stub
}

func (fake *ZhmcAPI) UpdateStorageGroupPropertiesArgsForCall(i int) (string, *zhmcclient.StorageGroupProperties) {
	fake.updateStorageGroupPropertiesMutex.RLock()
	defer fake.updateStorageGroupPropertiesMutex.RUnlock()
	argsForCall := fake.updateStorageGroupPropertiesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ZhmcAPI) UpdateStorageGroupPropertiesReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.updateStorageGroupPropertiesMutex.Lock()
	defer fake.updateStorageGroupPropertiesMutex.Unlock()
	fake.UpdateStorageGroupPropertiesStub = nil
	fake.updateStorageGroupPropertiesReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) UpdateStorageGroupPropertiesReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.updateStorageGroupPropertiesMutex.Lock()
	defer fake.updateStorageGroupPropertiesMutex.Unlock()
	fake.UpdateStorageGroupPropertiesStub = nil
	if fake.updateStorageGroupPropertiesReturnsOnCall == nil {
		fake.updateStorageGroupPropertiesReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.updateStorageGroupPropertiesReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *ZhmcAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.attachStorageGroupToPartitionMutex.RLock()
	defer fake.attachStorageGroupToPartitionMutex.RUnlock()
	fake.cancelJobMutex.RLock()
	defer fake.cancelJobMutex.RUnlock()
	fake.createHipersocketMutex.RLock()
	defer fake.createHipersocketMutex.RUnlock()
	fake.createLPARMutex.RLock()
	defer fake.createLPARMutex.RUnlock()
	fake.createNicMutex.RLock()
	defer fake.createNicMutex.RUnlock()
	fake.createStorageGroupsMutex.RLock()
	defer fake.createStorageGroupsMutex.RUnlock()
	fake.deleteHipersocketMutex.RLock()
	defer fake.deleteHipersocketMutex.RUnlock()
	fake.deleteJobMutex.RLock()
	defer fake.deleteJobMutex.RUnlock()
	fake.deleteLPARMutex.RLock()
	defer fake.deleteLPARMutex.RUnlock()
	fake.deleteNicMutex.RLock()
	defer fake.deleteNicMutex.RUnlock()
	fake.deleteStorageGroupMutex.RLock()
	defer fake.deleteStorageGroupMutex.RUnlock()
	fake.detachStorageGroupToPartitionMutex.RLock()
	defer fake.detachStorageGroupToPartitionMutex.RUnlock()
	fake.fetchAsciiConsoleURIMutex.RLock()
	defer fake.fetchAsciiConsoleURIMutex.RUnlock()
	fake.fulfillStorageGroupMutex.RLock()
	defer fake.fulfillStorageGroupMutex.RUnlock()
	fake.getAdapterPropertiesMutex.RLock()
	defer fake.getAdapterPropertiesMutex.RUnlock()
	fake.getLparPropertiesMutex.RLock()
	defer fake.getLparPropertiesMutex.RUnlock()
	fake.getNetworkAdapterPortPropertiesMutex.RLock()
	defer fake.getNetworkAdapterPortPropertiesMutex.RUnlock()
	fake.getNicPropertiesMutex.RLock()
	defer fake.getNicPropertiesMutex.RUnlock()
	fake.getStorageAdapterPortPropertiesMutex.RLock()
	defer fake.getStorageAdapterPortPropertiesMutex.RUnlock()
	fake.getStorageGroupPartitionsMutex.RLock()
	defer fake.getStorageGroupPartitionsMutex.RUnlock()
	fake.getStorageGroupPropertiesMutex.RLock()
	defer fake.getStorageGroupPropertiesMutex.RUnlock()
	fake.getStorageVolumePropertiesMutex.RLock()
	defer fake.getStorageVolumePropertiesMutex.RUnlock()
	fake.getVirtualSwitchPropertiesMutex.RLock()
	defer fake.getVirtualSwitchPropertiesMutex.RUnlock()
	fake.listAdaptersMutex.RLock()
	defer fake.listAdaptersMutex.RUnlock()
	fake.listCPCsMutex.RLock()
	defer fake.listCPCsMutex.RUnlock()
	fake.listLPARsMutex.RLock()
	defer fake.listLPARsMutex.RUnlock()
	fake.listNicsMutex.RLock()
	defer fake.listNicsMutex.RUnlock()
	fake.listStorageGroupsMutex.RLock()
	defer fake.listStorageGroupsMutex.RUnlock()
	fake.listStorageVolumesMutex.RLock()
	defer fake.listStorageVolumesMutex.RUnlock()
	fake.listVirtualSwitchesMutex.RLock()
	defer fake.listVirtualSwitchesMutex.RUnlock()
	fake.mountIsoImageMutex.RLock()
	defer fake.mountIsoImageMutex.RUnlock()
	fake.queryJobMutex.RLock()
	defer fake.queryJobMutex.RUnlock()
	fake.startLPARMutex.RLock()
	defer fake.startLPARMutex.RUnlock()
	fake.stopLPARMutex.RLock()
	defer fake.stopLPARMutex.RUnlock()
	fake.unmountIsoImageMutex.RLock()
	defer fake.unmountIsoImageMutex.RUnlock()
	fake.updateLparPropertiesMutex.RLock()
	defer fake.updateLparPropertiesMutex.RUnlock()
	fake.updateNicPropertiesMutex.RLock()
	defer fake.updateNicPropertiesMutex.RUnlock()
	fake.updateStorageGroupPropertiesMutex.RLock()
	defer fake.updateStorageGroupPropertiesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ZhmcAPI) recordInvocation(key string, args []interface{}) {
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

var _ zhmcclient.ZhmcAPI = new(ZhmcAPI)

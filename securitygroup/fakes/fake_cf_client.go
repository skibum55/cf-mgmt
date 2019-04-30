// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
	securitygroup "github.com/pivotalservices/cf-mgmt/securitygroup"
)

type FakeCFClient struct {
	BindRunningSecGroupStub        func(string) error
	bindRunningSecGroupMutex       sync.RWMutex
	bindRunningSecGroupArgsForCall []struct {
		arg1 string
	}
	bindRunningSecGroupReturns struct {
		result1 error
	}
	bindRunningSecGroupReturnsOnCall map[int]struct {
		result1 error
	}
	BindSecGroupStub        func(string, string) error
	bindSecGroupMutex       sync.RWMutex
	bindSecGroupArgsForCall []struct {
		arg1 string
		arg2 string
	}
	bindSecGroupReturns struct {
		result1 error
	}
	bindSecGroupReturnsOnCall map[int]struct {
		result1 error
	}
	BindStagingSecGroupStub        func(string) error
	bindStagingSecGroupMutex       sync.RWMutex
	bindStagingSecGroupArgsForCall []struct {
		arg1 string
	}
	bindStagingSecGroupReturns struct {
		result1 error
	}
	bindStagingSecGroupReturnsOnCall map[int]struct {
		result1 error
	}
	CreateSecGroupStub        func(string, []cfclient.SecGroupRule, []string) (*cfclient.SecGroup, error)
	createSecGroupMutex       sync.RWMutex
	createSecGroupArgsForCall []struct {
		arg1 string
		arg2 []cfclient.SecGroupRule
		arg3 []string
	}
	createSecGroupReturns struct {
		result1 *cfclient.SecGroup
		result2 error
	}
	createSecGroupReturnsOnCall map[int]struct {
		result1 *cfclient.SecGroup
		result2 error
	}
	GetSecGroupStub        func(string) (*cfclient.SecGroup, error)
	getSecGroupMutex       sync.RWMutex
	getSecGroupArgsForCall []struct {
		arg1 string
	}
	getSecGroupReturns struct {
		result1 *cfclient.SecGroup
		result2 error
	}
	getSecGroupReturnsOnCall map[int]struct {
		result1 *cfclient.SecGroup
		result2 error
	}
	ListSecGroupsStub        func() ([]cfclient.SecGroup, error)
	listSecGroupsMutex       sync.RWMutex
	listSecGroupsArgsForCall []struct {
	}
	listSecGroupsReturns struct {
		result1 []cfclient.SecGroup
		result2 error
	}
	listSecGroupsReturnsOnCall map[int]struct {
		result1 []cfclient.SecGroup
		result2 error
	}
	ListSpaceSecGroupsStub        func(string) ([]cfclient.SecGroup, error)
	listSpaceSecGroupsMutex       sync.RWMutex
	listSpaceSecGroupsArgsForCall []struct {
		arg1 string
	}
	listSpaceSecGroupsReturns struct {
		result1 []cfclient.SecGroup
		result2 error
	}
	listSpaceSecGroupsReturnsOnCall map[int]struct {
		result1 []cfclient.SecGroup
		result2 error
	}
	UnbindRunningSecGroupStub        func(string) error
	unbindRunningSecGroupMutex       sync.RWMutex
	unbindRunningSecGroupArgsForCall []struct {
		arg1 string
	}
	unbindRunningSecGroupReturns struct {
		result1 error
	}
	unbindRunningSecGroupReturnsOnCall map[int]struct {
		result1 error
	}
	UnbindSecGroupStub        func(string, string) error
	unbindSecGroupMutex       sync.RWMutex
	unbindSecGroupArgsForCall []struct {
		arg1 string
		arg2 string
	}
	unbindSecGroupReturns struct {
		result1 error
	}
	unbindSecGroupReturnsOnCall map[int]struct {
		result1 error
	}
	UnbindStagingSecGroupStub        func(string) error
	unbindStagingSecGroupMutex       sync.RWMutex
	unbindStagingSecGroupArgsForCall []struct {
		arg1 string
	}
	unbindStagingSecGroupReturns struct {
		result1 error
	}
	unbindStagingSecGroupReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateSecGroupStub        func(string, string, []cfclient.SecGroupRule, []string) (*cfclient.SecGroup, error)
	updateSecGroupMutex       sync.RWMutex
	updateSecGroupArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []cfclient.SecGroupRule
		arg4 []string
	}
	updateSecGroupReturns struct {
		result1 *cfclient.SecGroup
		result2 error
	}
	updateSecGroupReturnsOnCall map[int]struct {
		result1 *cfclient.SecGroup
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCFClient) BindRunningSecGroup(arg1 string) error {
	fake.bindRunningSecGroupMutex.Lock()
	ret, specificReturn := fake.bindRunningSecGroupReturnsOnCall[len(fake.bindRunningSecGroupArgsForCall)]
	fake.bindRunningSecGroupArgsForCall = append(fake.bindRunningSecGroupArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("BindRunningSecGroup", []interface{}{arg1})
	fake.bindRunningSecGroupMutex.Unlock()
	if fake.BindRunningSecGroupStub != nil {
		return fake.BindRunningSecGroupStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.bindRunningSecGroupReturns
	return fakeReturns.result1
}

func (fake *FakeCFClient) BindRunningSecGroupCallCount() int {
	fake.bindRunningSecGroupMutex.RLock()
	defer fake.bindRunningSecGroupMutex.RUnlock()
	return len(fake.bindRunningSecGroupArgsForCall)
}

func (fake *FakeCFClient) BindRunningSecGroupCalls(stub func(string) error) {
	fake.bindRunningSecGroupMutex.Lock()
	defer fake.bindRunningSecGroupMutex.Unlock()
	fake.BindRunningSecGroupStub = stub
}

func (fake *FakeCFClient) BindRunningSecGroupArgsForCall(i int) string {
	fake.bindRunningSecGroupMutex.RLock()
	defer fake.bindRunningSecGroupMutex.RUnlock()
	argsForCall := fake.bindRunningSecGroupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCFClient) BindRunningSecGroupReturns(result1 error) {
	fake.bindRunningSecGroupMutex.Lock()
	defer fake.bindRunningSecGroupMutex.Unlock()
	fake.BindRunningSecGroupStub = nil
	fake.bindRunningSecGroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) BindRunningSecGroupReturnsOnCall(i int, result1 error) {
	fake.bindRunningSecGroupMutex.Lock()
	defer fake.bindRunningSecGroupMutex.Unlock()
	fake.BindRunningSecGroupStub = nil
	if fake.bindRunningSecGroupReturnsOnCall == nil {
		fake.bindRunningSecGroupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.bindRunningSecGroupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) BindSecGroup(arg1 string, arg2 string) error {
	fake.bindSecGroupMutex.Lock()
	ret, specificReturn := fake.bindSecGroupReturnsOnCall[len(fake.bindSecGroupArgsForCall)]
	fake.bindSecGroupArgsForCall = append(fake.bindSecGroupArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("BindSecGroup", []interface{}{arg1, arg2})
	fake.bindSecGroupMutex.Unlock()
	if fake.BindSecGroupStub != nil {
		return fake.BindSecGroupStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.bindSecGroupReturns
	return fakeReturns.result1
}

func (fake *FakeCFClient) BindSecGroupCallCount() int {
	fake.bindSecGroupMutex.RLock()
	defer fake.bindSecGroupMutex.RUnlock()
	return len(fake.bindSecGroupArgsForCall)
}

func (fake *FakeCFClient) BindSecGroupCalls(stub func(string, string) error) {
	fake.bindSecGroupMutex.Lock()
	defer fake.bindSecGroupMutex.Unlock()
	fake.BindSecGroupStub = stub
}

func (fake *FakeCFClient) BindSecGroupArgsForCall(i int) (string, string) {
	fake.bindSecGroupMutex.RLock()
	defer fake.bindSecGroupMutex.RUnlock()
	argsForCall := fake.bindSecGroupArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCFClient) BindSecGroupReturns(result1 error) {
	fake.bindSecGroupMutex.Lock()
	defer fake.bindSecGroupMutex.Unlock()
	fake.BindSecGroupStub = nil
	fake.bindSecGroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) BindSecGroupReturnsOnCall(i int, result1 error) {
	fake.bindSecGroupMutex.Lock()
	defer fake.bindSecGroupMutex.Unlock()
	fake.BindSecGroupStub = nil
	if fake.bindSecGroupReturnsOnCall == nil {
		fake.bindSecGroupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.bindSecGroupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) BindStagingSecGroup(arg1 string) error {
	fake.bindStagingSecGroupMutex.Lock()
	ret, specificReturn := fake.bindStagingSecGroupReturnsOnCall[len(fake.bindStagingSecGroupArgsForCall)]
	fake.bindStagingSecGroupArgsForCall = append(fake.bindStagingSecGroupArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("BindStagingSecGroup", []interface{}{arg1})
	fake.bindStagingSecGroupMutex.Unlock()
	if fake.BindStagingSecGroupStub != nil {
		return fake.BindStagingSecGroupStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.bindStagingSecGroupReturns
	return fakeReturns.result1
}

func (fake *FakeCFClient) BindStagingSecGroupCallCount() int {
	fake.bindStagingSecGroupMutex.RLock()
	defer fake.bindStagingSecGroupMutex.RUnlock()
	return len(fake.bindStagingSecGroupArgsForCall)
}

func (fake *FakeCFClient) BindStagingSecGroupCalls(stub func(string) error) {
	fake.bindStagingSecGroupMutex.Lock()
	defer fake.bindStagingSecGroupMutex.Unlock()
	fake.BindStagingSecGroupStub = stub
}

func (fake *FakeCFClient) BindStagingSecGroupArgsForCall(i int) string {
	fake.bindStagingSecGroupMutex.RLock()
	defer fake.bindStagingSecGroupMutex.RUnlock()
	argsForCall := fake.bindStagingSecGroupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCFClient) BindStagingSecGroupReturns(result1 error) {
	fake.bindStagingSecGroupMutex.Lock()
	defer fake.bindStagingSecGroupMutex.Unlock()
	fake.BindStagingSecGroupStub = nil
	fake.bindStagingSecGroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) BindStagingSecGroupReturnsOnCall(i int, result1 error) {
	fake.bindStagingSecGroupMutex.Lock()
	defer fake.bindStagingSecGroupMutex.Unlock()
	fake.BindStagingSecGroupStub = nil
	if fake.bindStagingSecGroupReturnsOnCall == nil {
		fake.bindStagingSecGroupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.bindStagingSecGroupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) CreateSecGroup(arg1 string, arg2 []cfclient.SecGroupRule, arg3 []string) (*cfclient.SecGroup, error) {
	var arg2Copy []cfclient.SecGroupRule
	if arg2 != nil {
		arg2Copy = make([]cfclient.SecGroupRule, len(arg2))
		copy(arg2Copy, arg2)
	}
	var arg3Copy []string
	if arg3 != nil {
		arg3Copy = make([]string, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.createSecGroupMutex.Lock()
	ret, specificReturn := fake.createSecGroupReturnsOnCall[len(fake.createSecGroupArgsForCall)]
	fake.createSecGroupArgsForCall = append(fake.createSecGroupArgsForCall, struct {
		arg1 string
		arg2 []cfclient.SecGroupRule
		arg3 []string
	}{arg1, arg2Copy, arg3Copy})
	fake.recordInvocation("CreateSecGroup", []interface{}{arg1, arg2Copy, arg3Copy})
	fake.createSecGroupMutex.Unlock()
	if fake.CreateSecGroupStub != nil {
		return fake.CreateSecGroupStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createSecGroupReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) CreateSecGroupCallCount() int {
	fake.createSecGroupMutex.RLock()
	defer fake.createSecGroupMutex.RUnlock()
	return len(fake.createSecGroupArgsForCall)
}

func (fake *FakeCFClient) CreateSecGroupCalls(stub func(string, []cfclient.SecGroupRule, []string) (*cfclient.SecGroup, error)) {
	fake.createSecGroupMutex.Lock()
	defer fake.createSecGroupMutex.Unlock()
	fake.CreateSecGroupStub = stub
}

func (fake *FakeCFClient) CreateSecGroupArgsForCall(i int) (string, []cfclient.SecGroupRule, []string) {
	fake.createSecGroupMutex.RLock()
	defer fake.createSecGroupMutex.RUnlock()
	argsForCall := fake.createSecGroupArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCFClient) CreateSecGroupReturns(result1 *cfclient.SecGroup, result2 error) {
	fake.createSecGroupMutex.Lock()
	defer fake.createSecGroupMutex.Unlock()
	fake.CreateSecGroupStub = nil
	fake.createSecGroupReturns = struct {
		result1 *cfclient.SecGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) CreateSecGroupReturnsOnCall(i int, result1 *cfclient.SecGroup, result2 error) {
	fake.createSecGroupMutex.Lock()
	defer fake.createSecGroupMutex.Unlock()
	fake.CreateSecGroupStub = nil
	if fake.createSecGroupReturnsOnCall == nil {
		fake.createSecGroupReturnsOnCall = make(map[int]struct {
			result1 *cfclient.SecGroup
			result2 error
		})
	}
	fake.createSecGroupReturnsOnCall[i] = struct {
		result1 *cfclient.SecGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) GetSecGroup(arg1 string) (*cfclient.SecGroup, error) {
	fake.getSecGroupMutex.Lock()
	ret, specificReturn := fake.getSecGroupReturnsOnCall[len(fake.getSecGroupArgsForCall)]
	fake.getSecGroupArgsForCall = append(fake.getSecGroupArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetSecGroup", []interface{}{arg1})
	fake.getSecGroupMutex.Unlock()
	if fake.GetSecGroupStub != nil {
		return fake.GetSecGroupStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getSecGroupReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) GetSecGroupCallCount() int {
	fake.getSecGroupMutex.RLock()
	defer fake.getSecGroupMutex.RUnlock()
	return len(fake.getSecGroupArgsForCall)
}

func (fake *FakeCFClient) GetSecGroupCalls(stub func(string) (*cfclient.SecGroup, error)) {
	fake.getSecGroupMutex.Lock()
	defer fake.getSecGroupMutex.Unlock()
	fake.GetSecGroupStub = stub
}

func (fake *FakeCFClient) GetSecGroupArgsForCall(i int) string {
	fake.getSecGroupMutex.RLock()
	defer fake.getSecGroupMutex.RUnlock()
	argsForCall := fake.getSecGroupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCFClient) GetSecGroupReturns(result1 *cfclient.SecGroup, result2 error) {
	fake.getSecGroupMutex.Lock()
	defer fake.getSecGroupMutex.Unlock()
	fake.GetSecGroupStub = nil
	fake.getSecGroupReturns = struct {
		result1 *cfclient.SecGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) GetSecGroupReturnsOnCall(i int, result1 *cfclient.SecGroup, result2 error) {
	fake.getSecGroupMutex.Lock()
	defer fake.getSecGroupMutex.Unlock()
	fake.GetSecGroupStub = nil
	if fake.getSecGroupReturnsOnCall == nil {
		fake.getSecGroupReturnsOnCall = make(map[int]struct {
			result1 *cfclient.SecGroup
			result2 error
		})
	}
	fake.getSecGroupReturnsOnCall[i] = struct {
		result1 *cfclient.SecGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) ListSecGroups() ([]cfclient.SecGroup, error) {
	fake.listSecGroupsMutex.Lock()
	ret, specificReturn := fake.listSecGroupsReturnsOnCall[len(fake.listSecGroupsArgsForCall)]
	fake.listSecGroupsArgsForCall = append(fake.listSecGroupsArgsForCall, struct {
	}{})
	fake.recordInvocation("ListSecGroups", []interface{}{})
	fake.listSecGroupsMutex.Unlock()
	if fake.ListSecGroupsStub != nil {
		return fake.ListSecGroupsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listSecGroupsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) ListSecGroupsCallCount() int {
	fake.listSecGroupsMutex.RLock()
	defer fake.listSecGroupsMutex.RUnlock()
	return len(fake.listSecGroupsArgsForCall)
}

func (fake *FakeCFClient) ListSecGroupsCalls(stub func() ([]cfclient.SecGroup, error)) {
	fake.listSecGroupsMutex.Lock()
	defer fake.listSecGroupsMutex.Unlock()
	fake.ListSecGroupsStub = stub
}

func (fake *FakeCFClient) ListSecGroupsReturns(result1 []cfclient.SecGroup, result2 error) {
	fake.listSecGroupsMutex.Lock()
	defer fake.listSecGroupsMutex.Unlock()
	fake.ListSecGroupsStub = nil
	fake.listSecGroupsReturns = struct {
		result1 []cfclient.SecGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) ListSecGroupsReturnsOnCall(i int, result1 []cfclient.SecGroup, result2 error) {
	fake.listSecGroupsMutex.Lock()
	defer fake.listSecGroupsMutex.Unlock()
	fake.ListSecGroupsStub = nil
	if fake.listSecGroupsReturnsOnCall == nil {
		fake.listSecGroupsReturnsOnCall = make(map[int]struct {
			result1 []cfclient.SecGroup
			result2 error
		})
	}
	fake.listSecGroupsReturnsOnCall[i] = struct {
		result1 []cfclient.SecGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) ListSpaceSecGroups(arg1 string) ([]cfclient.SecGroup, error) {
	fake.listSpaceSecGroupsMutex.Lock()
	ret, specificReturn := fake.listSpaceSecGroupsReturnsOnCall[len(fake.listSpaceSecGroupsArgsForCall)]
	fake.listSpaceSecGroupsArgsForCall = append(fake.listSpaceSecGroupsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ListSpaceSecGroups", []interface{}{arg1})
	fake.listSpaceSecGroupsMutex.Unlock()
	if fake.ListSpaceSecGroupsStub != nil {
		return fake.ListSpaceSecGroupsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listSpaceSecGroupsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) ListSpaceSecGroupsCallCount() int {
	fake.listSpaceSecGroupsMutex.RLock()
	defer fake.listSpaceSecGroupsMutex.RUnlock()
	return len(fake.listSpaceSecGroupsArgsForCall)
}

func (fake *FakeCFClient) ListSpaceSecGroupsCalls(stub func(string) ([]cfclient.SecGroup, error)) {
	fake.listSpaceSecGroupsMutex.Lock()
	defer fake.listSpaceSecGroupsMutex.Unlock()
	fake.ListSpaceSecGroupsStub = stub
}

func (fake *FakeCFClient) ListSpaceSecGroupsArgsForCall(i int) string {
	fake.listSpaceSecGroupsMutex.RLock()
	defer fake.listSpaceSecGroupsMutex.RUnlock()
	argsForCall := fake.listSpaceSecGroupsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCFClient) ListSpaceSecGroupsReturns(result1 []cfclient.SecGroup, result2 error) {
	fake.listSpaceSecGroupsMutex.Lock()
	defer fake.listSpaceSecGroupsMutex.Unlock()
	fake.ListSpaceSecGroupsStub = nil
	fake.listSpaceSecGroupsReturns = struct {
		result1 []cfclient.SecGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) ListSpaceSecGroupsReturnsOnCall(i int, result1 []cfclient.SecGroup, result2 error) {
	fake.listSpaceSecGroupsMutex.Lock()
	defer fake.listSpaceSecGroupsMutex.Unlock()
	fake.ListSpaceSecGroupsStub = nil
	if fake.listSpaceSecGroupsReturnsOnCall == nil {
		fake.listSpaceSecGroupsReturnsOnCall = make(map[int]struct {
			result1 []cfclient.SecGroup
			result2 error
		})
	}
	fake.listSpaceSecGroupsReturnsOnCall[i] = struct {
		result1 []cfclient.SecGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) UnbindRunningSecGroup(arg1 string) error {
	fake.unbindRunningSecGroupMutex.Lock()
	ret, specificReturn := fake.unbindRunningSecGroupReturnsOnCall[len(fake.unbindRunningSecGroupArgsForCall)]
	fake.unbindRunningSecGroupArgsForCall = append(fake.unbindRunningSecGroupArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("UnbindRunningSecGroup", []interface{}{arg1})
	fake.unbindRunningSecGroupMutex.Unlock()
	if fake.UnbindRunningSecGroupStub != nil {
		return fake.UnbindRunningSecGroupStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.unbindRunningSecGroupReturns
	return fakeReturns.result1
}

func (fake *FakeCFClient) UnbindRunningSecGroupCallCount() int {
	fake.unbindRunningSecGroupMutex.RLock()
	defer fake.unbindRunningSecGroupMutex.RUnlock()
	return len(fake.unbindRunningSecGroupArgsForCall)
}

func (fake *FakeCFClient) UnbindRunningSecGroupCalls(stub func(string) error) {
	fake.unbindRunningSecGroupMutex.Lock()
	defer fake.unbindRunningSecGroupMutex.Unlock()
	fake.UnbindRunningSecGroupStub = stub
}

func (fake *FakeCFClient) UnbindRunningSecGroupArgsForCall(i int) string {
	fake.unbindRunningSecGroupMutex.RLock()
	defer fake.unbindRunningSecGroupMutex.RUnlock()
	argsForCall := fake.unbindRunningSecGroupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCFClient) UnbindRunningSecGroupReturns(result1 error) {
	fake.unbindRunningSecGroupMutex.Lock()
	defer fake.unbindRunningSecGroupMutex.Unlock()
	fake.UnbindRunningSecGroupStub = nil
	fake.unbindRunningSecGroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) UnbindRunningSecGroupReturnsOnCall(i int, result1 error) {
	fake.unbindRunningSecGroupMutex.Lock()
	defer fake.unbindRunningSecGroupMutex.Unlock()
	fake.UnbindRunningSecGroupStub = nil
	if fake.unbindRunningSecGroupReturnsOnCall == nil {
		fake.unbindRunningSecGroupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unbindRunningSecGroupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) UnbindSecGroup(arg1 string, arg2 string) error {
	fake.unbindSecGroupMutex.Lock()
	ret, specificReturn := fake.unbindSecGroupReturnsOnCall[len(fake.unbindSecGroupArgsForCall)]
	fake.unbindSecGroupArgsForCall = append(fake.unbindSecGroupArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("UnbindSecGroup", []interface{}{arg1, arg2})
	fake.unbindSecGroupMutex.Unlock()
	if fake.UnbindSecGroupStub != nil {
		return fake.UnbindSecGroupStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.unbindSecGroupReturns
	return fakeReturns.result1
}

func (fake *FakeCFClient) UnbindSecGroupCallCount() int {
	fake.unbindSecGroupMutex.RLock()
	defer fake.unbindSecGroupMutex.RUnlock()
	return len(fake.unbindSecGroupArgsForCall)
}

func (fake *FakeCFClient) UnbindSecGroupCalls(stub func(string, string) error) {
	fake.unbindSecGroupMutex.Lock()
	defer fake.unbindSecGroupMutex.Unlock()
	fake.UnbindSecGroupStub = stub
}

func (fake *FakeCFClient) UnbindSecGroupArgsForCall(i int) (string, string) {
	fake.unbindSecGroupMutex.RLock()
	defer fake.unbindSecGroupMutex.RUnlock()
	argsForCall := fake.unbindSecGroupArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCFClient) UnbindSecGroupReturns(result1 error) {
	fake.unbindSecGroupMutex.Lock()
	defer fake.unbindSecGroupMutex.Unlock()
	fake.UnbindSecGroupStub = nil
	fake.unbindSecGroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) UnbindSecGroupReturnsOnCall(i int, result1 error) {
	fake.unbindSecGroupMutex.Lock()
	defer fake.unbindSecGroupMutex.Unlock()
	fake.UnbindSecGroupStub = nil
	if fake.unbindSecGroupReturnsOnCall == nil {
		fake.unbindSecGroupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unbindSecGroupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) UnbindStagingSecGroup(arg1 string) error {
	fake.unbindStagingSecGroupMutex.Lock()
	ret, specificReturn := fake.unbindStagingSecGroupReturnsOnCall[len(fake.unbindStagingSecGroupArgsForCall)]
	fake.unbindStagingSecGroupArgsForCall = append(fake.unbindStagingSecGroupArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("UnbindStagingSecGroup", []interface{}{arg1})
	fake.unbindStagingSecGroupMutex.Unlock()
	if fake.UnbindStagingSecGroupStub != nil {
		return fake.UnbindStagingSecGroupStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.unbindStagingSecGroupReturns
	return fakeReturns.result1
}

func (fake *FakeCFClient) UnbindStagingSecGroupCallCount() int {
	fake.unbindStagingSecGroupMutex.RLock()
	defer fake.unbindStagingSecGroupMutex.RUnlock()
	return len(fake.unbindStagingSecGroupArgsForCall)
}

func (fake *FakeCFClient) UnbindStagingSecGroupCalls(stub func(string) error) {
	fake.unbindStagingSecGroupMutex.Lock()
	defer fake.unbindStagingSecGroupMutex.Unlock()
	fake.UnbindStagingSecGroupStub = stub
}

func (fake *FakeCFClient) UnbindStagingSecGroupArgsForCall(i int) string {
	fake.unbindStagingSecGroupMutex.RLock()
	defer fake.unbindStagingSecGroupMutex.RUnlock()
	argsForCall := fake.unbindStagingSecGroupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCFClient) UnbindStagingSecGroupReturns(result1 error) {
	fake.unbindStagingSecGroupMutex.Lock()
	defer fake.unbindStagingSecGroupMutex.Unlock()
	fake.UnbindStagingSecGroupStub = nil
	fake.unbindStagingSecGroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) UnbindStagingSecGroupReturnsOnCall(i int, result1 error) {
	fake.unbindStagingSecGroupMutex.Lock()
	defer fake.unbindStagingSecGroupMutex.Unlock()
	fake.UnbindStagingSecGroupStub = nil
	if fake.unbindStagingSecGroupReturnsOnCall == nil {
		fake.unbindStagingSecGroupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unbindStagingSecGroupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFClient) UpdateSecGroup(arg1 string, arg2 string, arg3 []cfclient.SecGroupRule, arg4 []string) (*cfclient.SecGroup, error) {
	var arg3Copy []cfclient.SecGroupRule
	if arg3 != nil {
		arg3Copy = make([]cfclient.SecGroupRule, len(arg3))
		copy(arg3Copy, arg3)
	}
	var arg4Copy []string
	if arg4 != nil {
		arg4Copy = make([]string, len(arg4))
		copy(arg4Copy, arg4)
	}
	fake.updateSecGroupMutex.Lock()
	ret, specificReturn := fake.updateSecGroupReturnsOnCall[len(fake.updateSecGroupArgsForCall)]
	fake.updateSecGroupArgsForCall = append(fake.updateSecGroupArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []cfclient.SecGroupRule
		arg4 []string
	}{arg1, arg2, arg3Copy, arg4Copy})
	fake.recordInvocation("UpdateSecGroup", []interface{}{arg1, arg2, arg3Copy, arg4Copy})
	fake.updateSecGroupMutex.Unlock()
	if fake.UpdateSecGroupStub != nil {
		return fake.UpdateSecGroupStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateSecGroupReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) UpdateSecGroupCallCount() int {
	fake.updateSecGroupMutex.RLock()
	defer fake.updateSecGroupMutex.RUnlock()
	return len(fake.updateSecGroupArgsForCall)
}

func (fake *FakeCFClient) UpdateSecGroupCalls(stub func(string, string, []cfclient.SecGroupRule, []string) (*cfclient.SecGroup, error)) {
	fake.updateSecGroupMutex.Lock()
	defer fake.updateSecGroupMutex.Unlock()
	fake.UpdateSecGroupStub = stub
}

func (fake *FakeCFClient) UpdateSecGroupArgsForCall(i int) (string, string, []cfclient.SecGroupRule, []string) {
	fake.updateSecGroupMutex.RLock()
	defer fake.updateSecGroupMutex.RUnlock()
	argsForCall := fake.updateSecGroupArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeCFClient) UpdateSecGroupReturns(result1 *cfclient.SecGroup, result2 error) {
	fake.updateSecGroupMutex.Lock()
	defer fake.updateSecGroupMutex.Unlock()
	fake.UpdateSecGroupStub = nil
	fake.updateSecGroupReturns = struct {
		result1 *cfclient.SecGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) UpdateSecGroupReturnsOnCall(i int, result1 *cfclient.SecGroup, result2 error) {
	fake.updateSecGroupMutex.Lock()
	defer fake.updateSecGroupMutex.Unlock()
	fake.UpdateSecGroupStub = nil
	if fake.updateSecGroupReturnsOnCall == nil {
		fake.updateSecGroupReturnsOnCall = make(map[int]struct {
			result1 *cfclient.SecGroup
			result2 error
		})
	}
	fake.updateSecGroupReturnsOnCall[i] = struct {
		result1 *cfclient.SecGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bindRunningSecGroupMutex.RLock()
	defer fake.bindRunningSecGroupMutex.RUnlock()
	fake.bindSecGroupMutex.RLock()
	defer fake.bindSecGroupMutex.RUnlock()
	fake.bindStagingSecGroupMutex.RLock()
	defer fake.bindStagingSecGroupMutex.RUnlock()
	fake.createSecGroupMutex.RLock()
	defer fake.createSecGroupMutex.RUnlock()
	fake.getSecGroupMutex.RLock()
	defer fake.getSecGroupMutex.RUnlock()
	fake.listSecGroupsMutex.RLock()
	defer fake.listSecGroupsMutex.RUnlock()
	fake.listSpaceSecGroupsMutex.RLock()
	defer fake.listSpaceSecGroupsMutex.RUnlock()
	fake.unbindRunningSecGroupMutex.RLock()
	defer fake.unbindRunningSecGroupMutex.RUnlock()
	fake.unbindSecGroupMutex.RLock()
	defer fake.unbindSecGroupMutex.RUnlock()
	fake.unbindStagingSecGroupMutex.RLock()
	defer fake.unbindStagingSecGroupMutex.RUnlock()
	fake.updateSecGroupMutex.RLock()
	defer fake.updateSecGroupMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCFClient) recordInvocation(key string, args []interface{}) {
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

var _ securitygroup.CFClient = new(FakeCFClient)

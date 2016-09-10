// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/pivotalservices/cf-mgmt/cloudcontroller (interfaces: Manager)

package mock_cloudcontroller

import (
	gomock "github.com/golang/mock/gomock"
	cloudcontroller "github.com/pivotalservices/cf-mgmt/cloudcontroller"
)

// Mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *_MockManagerRecorder
}

// Recorder for MockManager (not exported)
type _MockManagerRecorder struct {
	mock *MockManager
}

func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &_MockManagerRecorder{mock}
	return mock
}

func (_m *MockManager) EXPECT() *_MockManagerRecorder {
	return _m.recorder
}

func (_m *MockManager) AddUserToOrg(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "AddUserToOrg", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) AddUserToOrg(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddUserToOrg", arg0, arg1)
}

func (_m *MockManager) AddUserToOrgRole(_param0 string, _param1 string, _param2 string) error {
	ret := _m.ctrl.Call(_m, "AddUserToOrgRole", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) AddUserToOrgRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddUserToOrgRole", arg0, arg1, arg2)
}

func (_m *MockManager) AddUserToSpaceRole(_param0 string, _param1 string, _param2 string) error {
	ret := _m.ctrl.Call(_m, "AddUserToSpaceRole", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) AddUserToSpaceRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddUserToSpaceRole", arg0, arg1, arg2)
}

func (_m *MockManager) AssignQuotaToOrg(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "AssignQuotaToOrg", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) AssignQuotaToOrg(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AssignQuotaToOrg", arg0, arg1)
}

func (_m *MockManager) AssignQuotaToSpace(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "AssignQuotaToSpace", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) AssignQuotaToSpace(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AssignQuotaToSpace", arg0, arg1)
}

func (_m *MockManager) AssignSecurityGroupToSpace(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "AssignSecurityGroupToSpace", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) AssignSecurityGroupToSpace(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AssignSecurityGroupToSpace", arg0, arg1)
}

func (_m *MockManager) CreateOrg(_param0 string) error {
	ret := _m.ctrl.Call(_m, "CreateOrg", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) CreateOrg(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateOrg", arg0)
}

func (_m *MockManager) CreateQuota(_param0 string, _param1 int, _param2 int, _param3 int, _param4 int, _param5 bool) (string, error) {
	ret := _m.ctrl.Call(_m, "CreateQuota", _param0, _param1, _param2, _param3, _param4, _param5)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManagerRecorder) CreateQuota(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateQuota", arg0, arg1, arg2, arg3, arg4, arg5)
}

func (_m *MockManager) CreateSecurityGroup(_param0 string, _param1 string) (string, error) {
	ret := _m.ctrl.Call(_m, "CreateSecurityGroup", _param0, _param1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManagerRecorder) CreateSecurityGroup(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateSecurityGroup", arg0, arg1)
}

func (_m *MockManager) CreateSpace(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "CreateSpace", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) CreateSpace(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateSpace", arg0, arg1)
}

func (_m *MockManager) CreateSpaceQuota(_param0 string, _param1 string, _param2 int, _param3 int, _param4 int, _param5 int, _param6 bool) (string, error) {
	ret := _m.ctrl.Call(_m, "CreateSpaceQuota", _param0, _param1, _param2, _param3, _param4, _param5, _param6)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManagerRecorder) CreateSpaceQuota(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateSpaceQuota", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

func (_m *MockManager) ListOrgs() ([]*cloudcontroller.Org, error) {
	ret := _m.ctrl.Call(_m, "ListOrgs")
	ret0, _ := ret[0].([]*cloudcontroller.Org)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManagerRecorder) ListOrgs() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListOrgs")
}

func (_m *MockManager) ListQuotas() (map[string]string, error) {
	ret := _m.ctrl.Call(_m, "ListQuotas")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManagerRecorder) ListQuotas() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListQuotas")
}

func (_m *MockManager) ListSecurityGroups() (map[string]string, error) {
	ret := _m.ctrl.Call(_m, "ListSecurityGroups")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManagerRecorder) ListSecurityGroups() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListSecurityGroups")
}

func (_m *MockManager) ListSpaceQuotas(_param0 string) (map[string]string, error) {
	ret := _m.ctrl.Call(_m, "ListSpaceQuotas", _param0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManagerRecorder) ListSpaceQuotas(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListSpaceQuotas", arg0)
}

func (_m *MockManager) ListSpaces(_param0 string) ([]cloudcontroller.Space, error) {
	ret := _m.ctrl.Call(_m, "ListSpaces", _param0)
	ret0, _ := ret[0].([]cloudcontroller.Space)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManagerRecorder) ListSpaces(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListSpaces", arg0)
}

func (_m *MockManager) UpdateQuota(_param0 string, _param1 string, _param2 int, _param3 int, _param4 int, _param5 int, _param6 bool) error {
	ret := _m.ctrl.Call(_m, "UpdateQuota", _param0, _param1, _param2, _param3, _param4, _param5, _param6)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) UpdateQuota(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateQuota", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

func (_m *MockManager) UpdateSecurityGroup(_param0 string, _param1 string, _param2 string) error {
	ret := _m.ctrl.Call(_m, "UpdateSecurityGroup", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) UpdateSecurityGroup(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateSecurityGroup", arg0, arg1, arg2)
}

func (_m *MockManager) UpdateSpaceQuota(_param0 string, _param1 string, _param2 string, _param3 int, _param4 int, _param5 int, _param6 int, _param7 bool) error {
	ret := _m.ctrl.Call(_m, "UpdateSpaceQuota", _param0, _param1, _param2, _param3, _param4, _param5, _param6, _param7)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) UpdateSpaceQuota(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateSpaceQuota", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

func (_m *MockManager) UpdateSpaceSSH(_param0 bool, _param1 string) error {
	ret := _m.ctrl.Call(_m, "UpdateSpaceSSH", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) UpdateSpaceSSH(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateSpaceSSH", arg0, arg1)
}

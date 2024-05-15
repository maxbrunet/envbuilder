// Code generated by MockGen. DO NOT EDIT.
// Source: remount.go
//
// Generated by this command:
//
//	mockgen -source=remount.go -package=ebutil -destination=mock_mounter_test.go -write_generate_directive
//

// Package ebutil is a generated GoMock package.
package ebutil

import (
	os "os"
	reflect "reflect"

	procfs "github.com/prometheus/procfs"
	gomock "go.uber.org/mock/gomock"
)

//go:generate mockgen -source=remount.go -package=ebutil -destination=mock_mounter_test.go -write_generate_directive

// Mockmounter is a mock of mounter interface.
type Mockmounter struct {
	ctrl     *gomock.Controller
	recorder *MockmounterMockRecorder
}

// MockmounterMockRecorder is the mock recorder for Mockmounter.
type MockmounterMockRecorder struct {
	mock *Mockmounter
}

// NewMockmounter creates a new mock instance.
func NewMockmounter(ctrl *gomock.Controller) *Mockmounter {
	mock := &Mockmounter{ctrl: ctrl}
	mock.recorder = &MockmounterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockmounter) EXPECT() *MockmounterMockRecorder {
	return m.recorder
}

// GetMounts mocks base method.
func (m *Mockmounter) GetMounts() ([]*procfs.MountInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMounts")
	ret0, _ := ret[0].([]*procfs.MountInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMounts indicates an expected call of GetMounts.
func (mr *MockmounterMockRecorder) GetMounts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMounts", reflect.TypeOf((*Mockmounter)(nil).GetMounts))
}

// MkdirAll mocks base method.
func (m *Mockmounter) MkdirAll(arg0 string, arg1 os.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MkdirAll", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MkdirAll indicates an expected call of MkdirAll.
func (mr *MockmounterMockRecorder) MkdirAll(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MkdirAll", reflect.TypeOf((*Mockmounter)(nil).MkdirAll), arg0, arg1)
}

// Mount mocks base method.
func (m *Mockmounter) Mount(arg0, arg1, arg2 string, arg3 uintptr, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mount", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mount indicates an expected call of Mount.
func (mr *MockmounterMockRecorder) Mount(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mount", reflect.TypeOf((*Mockmounter)(nil).Mount), arg0, arg1, arg2, arg3, arg4)
}

// Unmount mocks base method.
func (m *Mockmounter) Unmount(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmount indicates an expected call of Unmount.
func (mr *MockmounterMockRecorder) Unmount(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmount", reflect.TypeOf((*Mockmounter)(nil).Unmount), arg0, arg1)
}
// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/napptive/catalog-manager/internal/pkg/storage (interfaces: StorageManager)

// Package catalog_manager is a generated GoMock package.
package catalog_manager

import (
	gomock "github.com/golang/mock/gomock"
	entities "github.com/napptive/catalog-manager/internal/pkg/entities"
	reflect "reflect"
)

// MockStorageManager is a mock of StorageManager interface
type MockStorageManager struct {
	ctrl     *gomock.Controller
	recorder *MockStorageManagerMockRecorder
}

// MockStorageManagerMockRecorder is the mock recorder for MockStorageManager
type MockStorageManagerMockRecorder struct {
	mock *MockStorageManager
}

// NewMockStorageManager creates a new mock instance
func NewMockStorageManager(ctrl *gomock.Controller) *MockStorageManager {
	mock := &MockStorageManager{ctrl: ctrl}
	mock.recorder = &MockStorageManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorageManager) EXPECT() *MockStorageManagerMockRecorder {
	return m.recorder
}

// ApplicationExists mocks base method
func (m *MockStorageManager) ApplicationExists(arg0, arg1, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationExists", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationExists indicates an expected call of ApplicationExists
func (mr *MockStorageManagerMockRecorder) ApplicationExists(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationExists", reflect.TypeOf((*MockStorageManager)(nil).ApplicationExists), arg0, arg1, arg2)
}

// CreateRepository mocks base method
func (m *MockStorageManager) CreateRepository(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRepository", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRepository indicates an expected call of CreateRepository
func (mr *MockStorageManagerMockRecorder) CreateRepository(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRepository", reflect.TypeOf((*MockStorageManager)(nil).CreateRepository), arg0)
}

// GetApplication mocks base method
func (m *MockStorageManager) GetApplication(arg0, arg1, arg2 string) ([]*entities.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplication", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*entities.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplication indicates an expected call of GetApplication
func (mr *MockStorageManagerMockRecorder) GetApplication(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplication", reflect.TypeOf((*MockStorageManager)(nil).GetApplication), arg0, arg1, arg2)
}

// RemoveApplication mocks base method
func (m *MockStorageManager) RemoveApplication(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveApplication", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveApplication indicates an expected call of RemoveApplication
func (mr *MockStorageManagerMockRecorder) RemoveApplication(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveApplication", reflect.TypeOf((*MockStorageManager)(nil).RemoveApplication), arg0, arg1, arg2)
}

// RemoveRepository mocks base method
func (m *MockStorageManager) RemoveRepository(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRepository", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRepository indicates an expected call of RemoveRepository
func (mr *MockStorageManagerMockRecorder) RemoveRepository(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRepository", reflect.TypeOf((*MockStorageManager)(nil).RemoveRepository), arg0)
}

// RepositoryExists mocks base method
func (m *MockStorageManager) RepositoryExists(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RepositoryExists", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RepositoryExists indicates an expected call of RepositoryExists
func (mr *MockStorageManagerMockRecorder) RepositoryExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RepositoryExists", reflect.TypeOf((*MockStorageManager)(nil).RepositoryExists), arg0)
}

// StoreApplication mocks base method
func (m *MockStorageManager) StoreApplication(arg0, arg1, arg2 string, arg3 []*entities.FileInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreApplication", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreApplication indicates an expected call of StoreApplication
func (mr *MockStorageManagerMockRecorder) StoreApplication(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreApplication", reflect.TypeOf((*MockStorageManager)(nil).StoreApplication), arg0, arg1, arg2, arg3)
}

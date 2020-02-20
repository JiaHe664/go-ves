// Code generated by MockGen. DO NOT EDIT.
// Source: ../model/internal/abstraction/session-account.go

// Package mock is a generated GoMock package.
package mock

import (
	abstraction "github.com/Myriad-Dreamin/go-ves/ves/model/internal/abstraction"
	database "github.com/Myriad-Dreamin/go-ves/ves/model/internal/database"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// SessionAccountDB is a mock of SessionAccountDB interface
type SessionAccountDB struct {
	ctrl     *gomock.Controller
	recorder *SessionAccountDBMockRecorder
}

// SessionAccountDBMockRecorder is the mock recorder for SessionAccountDB
type SessionAccountDBMockRecorder struct {
	mock *SessionAccountDB
}

// NewSessionAccountDB creates a new mock instance
func NewSessionAccountDB(ctrl *gomock.Controller) *SessionAccountDB {
	mock := &SessionAccountDB{ctrl: ctrl}
	mock.recorder = &SessionAccountDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *SessionAccountDB) EXPECT() *SessionAccountDBMockRecorder {
	return m.recorder
}

// GetTraits mocks base method
func (m *SessionAccountDB) GetTraits() abstraction.ORMTraits {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTraits")
	ret0, _ := ret[0].(abstraction.ORMTraits)
	return ret0
}

// GetTraits indicates an expected call of GetTraits
func (mr *SessionAccountDBMockRecorder) GetTraits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTraits", reflect.TypeOf((*SessionAccountDB)(nil).GetTraits))
}

// Create mocks base method
func (m *SessionAccountDB) Create(sa *database.SessionAccount) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", sa)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *SessionAccountDBMockRecorder) Create(sa interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*SessionAccountDB)(nil).Create), sa)
}

// Delete mocks base method
func (m *SessionAccountDB) Delete(sa *database.SessionAccount) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", sa)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *SessionAccountDBMockRecorder) Delete(sa interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*SessionAccountDB)(nil).Delete), sa)
}

// Find mocks base method
func (m *SessionAccountDB) Find(sa *database.SessionAccount) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", sa)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *SessionAccountDBMockRecorder) Find(sa interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*SessionAccountDB)(nil).Find), sa)
}

// UpdateAcknowledged mocks base method
func (m *SessionAccountDB) UpdateAcknowledged(sa *database.SessionAccount) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAcknowledged", sa)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAcknowledged indicates an expected call of UpdateAcknowledged
func (mr *SessionAccountDBMockRecorder) UpdateAcknowledged(sa interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAcknowledged", reflect.TypeOf((*SessionAccountDB)(nil).UpdateAcknowledged), sa)
}

// GetAcknowledged mocks base method
func (m *SessionAccountDB) GetAcknowledged(guid string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAcknowledged", guid)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAcknowledged indicates an expected call of GetAcknowledged
func (mr *SessionAccountDBMockRecorder) GetAcknowledged(guid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAcknowledged", reflect.TypeOf((*SessionAccountDB)(nil).GetAcknowledged), guid)
}

// GetTotal mocks base method
func (m *SessionAccountDB) GetTotal(guid string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotal", guid)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotal indicates an expected call of GetTotal
func (mr *SessionAccountDBMockRecorder) GetTotal(guid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotal", reflect.TypeOf((*SessionAccountDB)(nil).GetTotal), guid)
}

// ID mocks base method
func (m *SessionAccountDB) ID(id string) ([]database.SessionAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID", id)
	ret0, _ := ret[0].([]database.SessionAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ID indicates an expected call of ID
func (mr *SessionAccountDBMockRecorder) ID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*SessionAccountDB)(nil).ID), id)
}

// Filter mocks base method
func (m *SessionAccountDB) Filter(f *database.SessionAccountFilter) ([]database.SessionAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filter", f)
	ret0, _ := ret[0].([]database.SessionAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filter indicates an expected call of Filter
func (mr *SessionAccountDBMockRecorder) Filter(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*SessionAccountDB)(nil).Filter), f)
}

// FilterI mocks base method
func (m *SessionAccountDB) FilterI(f interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterI", f)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterI indicates an expected call of FilterI
func (mr *SessionAccountDBMockRecorder) FilterI(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterI", reflect.TypeOf((*SessionAccountDB)(nil).FilterI), f)
}

// Query mocks base method
func (m *SessionAccountDB) Query(opts ...abstraction.SessionAccountQueryOption) ([]database.SessionAccount, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].([]database.SessionAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query
func (mr *SessionAccountDBMockRecorder) Query(opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*SessionAccountDB)(nil).Query), opts...)
}

// Scan mocks base method
func (m *SessionAccountDB) Scan(obj interface{}, opts ...abstraction.SessionAccountQueryOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{obj}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scan", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scan indicates an expected call of Scan
func (mr *SessionAccountDBMockRecorder) Scan(obj interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{obj}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*SessionAccountDB)(nil).Scan), varargs...)
}

// MockSessionAccountQueryOption is a mock of SessionAccountQueryOption interface
type MockSessionAccountQueryOption struct {
	ctrl     *gomock.Controller
	recorder *MockSessionAccountQueryOptionMockRecorder
}

// MockSessionAccountQueryOptionMockRecorder is the mock recorder for MockSessionAccountQueryOption
type MockSessionAccountQueryOptionMockRecorder struct {
	mock *MockSessionAccountQueryOption
}

// NewMockSessionAccountQueryOption creates a new mock instance
func NewMockSessionAccountQueryOption(ctrl *gomock.Controller) *MockSessionAccountQueryOption {
	mock := &MockSessionAccountQueryOption{ctrl: ctrl}
	mock.recorder = &MockSessionAccountQueryOptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSessionAccountQueryOption) EXPECT() *MockSessionAccountQueryOptionMockRecorder {
	return m.recorder
}

// implementsSessionAccountQuery mocks base method
func (m *MockSessionAccountQueryOption) implementsSessionAccountQuery() abstraction.SessionAccountQueryOption {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "implementsSessionAccountQuery")
	ret0, _ := ret[0].(abstraction.SessionAccountQueryOption)
	return ret0
}

// implementsSessionAccountQuery indicates an expected call of implementsSessionAccountQuery
func (mr *MockSessionAccountQueryOptionMockRecorder) implementsSessionAccountQuery() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "implementsSessionAccountQuery", reflect.TypeOf((*MockSessionAccountQueryOption)(nil).implementsSessionAccountQuery))
}
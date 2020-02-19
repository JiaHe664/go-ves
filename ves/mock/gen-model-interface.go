// Code generated by MockGen. DO NOT EDIT.
// Source: ../control/gen-model-interface.go

// Package mock is a generated GoMock package.
package mock

import (
	uiptypes "github.com/HyperService-Consortium/go-uip/uiptypes"
	gorm_crud_dao "github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
	types "github.com/Myriad-Dreamin/go-ves/types"
	dblayer "github.com/Myriad-Dreamin/go-ves/ves/model/internal/db-layer"
	gomock "github.com/golang/mock/gomock"
	gorm "github.com/jinzhu/gorm"
	reflect "reflect"
)

// SessionDB is a mock of SessionDBI interface
type SessionDB struct {
	ctrl     *gomock.Controller
	recorder *SessionDBMockRecorder
}

// SessionDBMockRecorder is the mock recorder for SessionDB
type SessionDBMockRecorder struct {
	mock *SessionDB
}

// NewSessionDB creates a new mock instance
func NewSessionDB(ctrl *gomock.Controller) *SessionDB {
	mock := &SessionDB{ctrl: ctrl}
	mock.recorder = &SessionDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *SessionDB) EXPECT() *SessionDBMockRecorder {
	return m.recorder
}

// Filter mocks base method
func (m *SessionDB) Filter(arg0 *gorm_crud_dao.Filter) ([]dblayer.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filter", arg0)
	ret0, _ := ret[0].([]dblayer.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filter indicates an expected call of Filter
func (mr *SessionDBMockRecorder) Filter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*SessionDB)(nil).Filter), arg0)
}

// FilterI mocks base method
func (m *SessionDB) FilterI(arg0 interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterI", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterI indicates an expected call of FilterI
func (mr *SessionDBMockRecorder) FilterI(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterI", reflect.TypeOf((*SessionDB)(nil).FilterI), arg0)
}

// ID mocks base method
func (m *SessionDB) ID(arg0 uint) (*dblayer.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID", arg0)
	ret0, _ := ret[0].(*dblayer.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ID indicates an expected call of ID
func (mr *SessionDBMockRecorder) ID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*SessionDB)(nil).ID), arg0)
}

// ID_ mocks base method
func (m *SessionDB) ID_(arg0 *gorm.DB, arg1 uint) (*dblayer.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID_", arg0, arg1)
	ret0, _ := ret[0].(*dblayer.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ID_ indicates an expected call of ID_
func (mr *SessionDBMockRecorder) ID_(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID_", reflect.TypeOf((*SessionDB)(nil).ID_), arg0, arg1)
}

// QueryChain mocks base method
func (m *SessionDB) QueryChain() *dblayer.SessionQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryChain")
	ret0, _ := ret[0].(*dblayer.SessionQuery)
	return ret0
}

// QueryChain indicates an expected call of QueryChain
func (mr *SessionDBMockRecorder) QueryChain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryChain", reflect.TypeOf((*SessionDB)(nil).QueryChain))
}

// QueryGUID mocks base method
func (m *SessionDB) QueryGUID(arg0 string) (*dblayer.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryGUID", arg0)
	ret0, _ := ret[0].(*dblayer.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryGUID indicates an expected call of QueryGUID
func (mr *SessionDBMockRecorder) QueryGUID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryGUID", reflect.TypeOf((*SessionDB)(nil).QueryGUID), arg0)
}

// QueryGUIDByBytes mocks base method
func (m *SessionDB) QueryGUIDByBytes(arg0 []uint8) (*dblayer.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryGUIDByBytes", arg0)
	ret0, _ := ret[0].(*dblayer.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryGUIDByBytes indicates an expected call of QueryGUIDByBytes
func (mr *SessionDBMockRecorder) QueryGUIDByBytes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryGUIDByBytes", reflect.TypeOf((*SessionDB)(nil).QueryGUIDByBytes), arg0)
}

// SessionAccountDB is a mock of SessionAccountDBI interface
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

// Filter mocks base method
func (m *SessionAccountDB) Filter(arg0 *gorm_crud_dao.Filter) ([]dblayer.SessionAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filter", arg0)
	ret0, _ := ret[0].([]dblayer.SessionAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filter indicates an expected call of Filter
func (mr *SessionAccountDBMockRecorder) Filter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*SessionAccountDB)(nil).Filter), arg0)
}

// FilterI mocks base method
func (m *SessionAccountDB) FilterI(arg0 interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterI", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterI indicates an expected call of FilterI
func (mr *SessionAccountDBMockRecorder) FilterI(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterI", reflect.TypeOf((*SessionAccountDB)(nil).FilterI), arg0)
}

// GetAcknowledged mocks base method
func (m *SessionAccountDB) GetAcknowledged(arg0 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAcknowledged", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAcknowledged indicates an expected call of GetAcknowledged
func (mr *SessionAccountDBMockRecorder) GetAcknowledged(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAcknowledged", reflect.TypeOf((*SessionAccountDB)(nil).GetAcknowledged), arg0)
}

// GetTotal mocks base method
func (m *SessionAccountDB) GetTotal(arg0 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotal", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotal indicates an expected call of GetTotal
func (mr *SessionAccountDBMockRecorder) GetTotal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotal", reflect.TypeOf((*SessionAccountDB)(nil).GetTotal), arg0)
}

// ID mocks base method
func (m *SessionAccountDB) ID(arg0 string) ([]dblayer.SessionAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID", arg0)
	ret0, _ := ret[0].([]dblayer.SessionAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ID indicates an expected call of ID
func (mr *SessionAccountDBMockRecorder) ID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*SessionAccountDB)(nil).ID), arg0)
}

// ID_ mocks base method
func (m *SessionAccountDB) ID_(arg0 *gorm.DB, arg1 string) (*dblayer.SessionAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID_", arg0, arg1)
	ret0, _ := ret[0].(*dblayer.SessionAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ID_ indicates an expected call of ID_
func (mr *SessionAccountDBMockRecorder) ID_(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID_", reflect.TypeOf((*SessionAccountDB)(nil).ID_), arg0, arg1)
}

// QueryChain mocks base method
func (m *SessionAccountDB) QueryChain() *dblayer.SessionAccountQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryChain")
	ret0, _ := ret[0].(*dblayer.SessionAccountQuery)
	return ret0
}

// QueryChain indicates an expected call of QueryChain
func (mr *SessionAccountDBMockRecorder) QueryChain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryChain", reflect.TypeOf((*SessionAccountDB)(nil).QueryChain))
}

// SessionFSet is a mock of SessionFSetI interface
type SessionFSet struct {
	ctrl     *gomock.Controller
	recorder *SessionFSetMockRecorder
}

// SessionFSetMockRecorder is the mock recorder for SessionFSet
type SessionFSetMockRecorder struct {
	mock *SessionFSet
}

// NewSessionFSet creates a new mock instance
func NewSessionFSet(ctrl *gomock.Controller) *SessionFSet {
	mock := &SessionFSet{ctrl: ctrl}
	mock.recorder = &SessionFSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *SessionFSet) EXPECT() *SessionFSetMockRecorder {
	return m.recorder
}

// AckForInit mocks base method
func (m *SessionFSet) AckForInit(arg0 *dblayer.Session, arg1 uiptypes.Account, arg2 uiptypes.Signature) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AckForInit", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AckForInit indicates an expected call of AckForInit
func (mr *SessionFSetMockRecorder) AckForInit(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AckForInit", reflect.TypeOf((*SessionFSet)(nil).AckForInit), arg0, arg1, arg2)
}

// FindTransaction mocks base method
func (m *SessionFSet) FindTransaction(arg0 []uint8, arg1 int64) ([]uint8, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTransaction", arg0, arg1)
	ret0, _ := ret[0].([]uint8)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTransaction indicates an expected call of FindTransaction
func (mr *SessionFSetMockRecorder) FindTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTransaction", reflect.TypeOf((*SessionFSet)(nil).FindTransaction), arg0, arg1)
}

// GetAccounts mocks base method
func (m *SessionFSet) GetAccounts(arg0 *dblayer.Session) ([]uiptypes.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccounts", arg0)
	ret0, _ := ret[0].([]uiptypes.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccounts indicates an expected call of GetAccounts
func (mr *SessionFSetMockRecorder) GetAccounts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*SessionFSet)(nil).GetAccounts), arg0)
}

// GetAckCount mocks base method
func (m *SessionFSet) GetAckCount(arg0 *dblayer.Session) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAckCount", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAckCount indicates an expected call of GetAckCount
func (mr *SessionFSetMockRecorder) GetAckCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAckCount", reflect.TypeOf((*SessionFSet)(nil).GetAckCount), arg0)
}

// GetTransactingTransaction mocks base method
func (m *SessionFSet) GetTransactingTransaction(arg0 *dblayer.Session) ([]uint8, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactingTransaction", arg0)
	ret0, _ := ret[0].([]uint8)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactingTransaction indicates an expected call of GetTransactingTransaction
func (mr *SessionFSetMockRecorder) GetTransactingTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactingTransaction", reflect.TypeOf((*SessionFSet)(nil).GetTransactingTransaction), arg0)
}

// InitSessionInfo mocks base method
func (m *SessionFSet) InitSessionInfo(arg0 []uint8, arg1 []*uiptypes.TransactionIntent, arg2 []*dblayer.SessionAccount) (*dblayer.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitSessionInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(*dblayer.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitSessionInfo indicates an expected call of InitSessionInfo
func (mr *SessionFSetMockRecorder) InitSessionInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitSessionInfo", reflect.TypeOf((*SessionFSet)(nil).InitSessionInfo), arg0, arg1, arg2)
}

// NotifyAttestation mocks base method
func (m *SessionFSet) NotifyAttestation(arg0 types.NSBInterface, arg1 uiptypes.BlockChainInterface, arg2 uiptypes.Attestation) (interface{}, interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyAttestation", arg0, arg1, arg2)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// NotifyAttestation indicates an expected call of NotifyAttestation
func (mr *SessionFSetMockRecorder) NotifyAttestation(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyAttestation", reflect.TypeOf((*SessionFSet)(nil).NotifyAttestation), arg0, arg1, arg2)
}

// ProcessAttestation mocks base method
func (m *SessionFSet) ProcessAttestation(arg0 types.NSBInterface, arg1 uiptypes.BlockChainInterface, arg2 uiptypes.Attestation) (interface{}, interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessAttestation", arg0, arg1, arg2)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ProcessAttestation indicates an expected call of ProcessAttestation
func (mr *SessionFSetMockRecorder) ProcessAttestation(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessAttestation", reflect.TypeOf((*SessionFSet)(nil).ProcessAttestation), arg0, arg1, arg2)
}

// SyncFromISC mocks base method
func (m *SessionFSet) SyncFromISC() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncFromISC")
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncFromISC indicates an expected call of SyncFromISC
func (mr *SessionFSetMockRecorder) SyncFromISC() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncFromISC", reflect.TypeOf((*SessionFSet)(nil).SyncFromISC))
}

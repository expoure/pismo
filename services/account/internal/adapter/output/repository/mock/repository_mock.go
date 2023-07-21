// Code generated by MockGen. DO NOT EDIT.
// Source: internal/application/port/output/account_port.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	money "github.com/Rhymond/go-money"
	domain "github.com/expoure/pismo/account/internal/application/domain"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockAccountRepositoryPort is a mock of AccountRepositoryPort interface.
type MockAccountRepositoryPort struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepositoryPortMockRecorder
}

// MockAccountRepositoryPortMockRecorder is the mock recorder for MockAccountRepositoryPort.
type MockAccountRepositoryPortMockRecorder struct {
	mock *MockAccountRepositoryPort
}

// NewMockAccountRepositoryPort creates a new mock instance.
func NewMockAccountRepositoryPort(ctrl *gomock.Controller) *MockAccountRepositoryPort {
	mock := &MockAccountRepositoryPort{ctrl: ctrl}
	mock.recorder = &MockAccountRepositoryPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRepositoryPort) EXPECT() *MockAccountRepositoryPortMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockAccountRepositoryPort) CreateAccount(accountDomain domain.AccountDomain) (*domain.AccountDomain, *error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", accountDomain)
	ret0, _ := ret[0].(*domain.AccountDomain)
	ret1, _ := ret[1].(*error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockAccountRepositoryPortMockRecorder) CreateAccount(accountDomain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAccountRepositoryPort)(nil).CreateAccount), accountDomain)
}

// FindAccountBalanceByID mocks base method.
func (m *MockAccountRepositoryPort) FindAccountBalanceByID(id uuid.UUID) (*money.Money, *error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAccountBalanceByID", id)
	ret0, _ := ret[0].(*money.Money)
	ret1, _ := ret[1].(*error)
	return ret0, ret1
}

// FindAccountBalanceByID indicates an expected call of FindAccountBalanceByID.
func (mr *MockAccountRepositoryPortMockRecorder) FindAccountBalanceByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAccountBalanceByID", reflect.TypeOf((*MockAccountRepositoryPort)(nil).FindAccountBalanceByID), id)
}

// FindAccountByDocumentNumber mocks base method.
func (m *MockAccountRepositoryPort) FindAccountByDocumentNumber(documentNumber string) (*domain.AccountDomain, *error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAccountByDocumentNumber", documentNumber)
	ret0, _ := ret[0].(*domain.AccountDomain)
	ret1, _ := ret[1].(*error)
	return ret0, ret1
}

// FindAccountByDocumentNumber indicates an expected call of FindAccountByDocumentNumber.
func (mr *MockAccountRepositoryPortMockRecorder) FindAccountByDocumentNumber(documentNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAccountByDocumentNumber", reflect.TypeOf((*MockAccountRepositoryPort)(nil).FindAccountByDocumentNumber), documentNumber)
}

// FindAccountByID mocks base method.
func (m *MockAccountRepositoryPort) FindAccountByID(id uuid.UUID) (*domain.AccountDomain, *error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAccountByID", id)
	ret0, _ := ret[0].(*domain.AccountDomain)
	ret1, _ := ret[1].(*error)
	return ret0, ret1
}

// FindAccountByID indicates an expected call of FindAccountByID.
func (mr *MockAccountRepositoryPortMockRecorder) FindAccountByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAccountByID", reflect.TypeOf((*MockAccountRepositoryPort)(nil).FindAccountByID), id)
}

// UpdateAccountBalanceByID mocks base method.
func (m *MockAccountRepositoryPort) UpdateAccountBalanceByID(id uuid.UUID, transactionAmount int64) (*money.Money, *error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccountBalanceByID", id, transactionAmount)
	ret0, _ := ret[0].(*money.Money)
	ret1, _ := ret[1].(*error)
	return ret0, ret1
}

// UpdateAccountBalanceByID indicates an expected call of UpdateAccountBalanceByID.
func (mr *MockAccountRepositoryPortMockRecorder) UpdateAccountBalanceByID(id, transactionAmount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccountBalanceByID", reflect.TypeOf((*MockAccountRepositoryPort)(nil).UpdateAccountBalanceByID), id, transactionAmount)
}

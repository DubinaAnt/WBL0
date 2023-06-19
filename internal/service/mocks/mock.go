// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"
	model "WBL0/internal/model"

	gomock "github.com/golang/mock/gomock"
)

// MockOrder is a mock of Order interface.
type MockOrder struct {
	ctrl     *gomock.Controller
	recorder *MockOrderMockRecorder
}

// MockOrderMockRecorder is the mock recorder for MockOrder.
type MockOrderMockRecorder struct {
	mock *MockOrder
}

// NewMockOrder creates a new mock instance.
func NewMockOrder(ctrl *gomock.Controller) *MockOrder {
	mock := &MockOrder{ctrl: ctrl}
	mock.recorder = &MockOrderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrder) EXPECT() *MockOrderMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockOrder) CreateOrder(order model.Order) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", order)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderMockRecorder) CreateOrder(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrder)(nil).CreateOrder), order)
}

// GetOrder mocks base method.
func (m *MockOrder) GetOrder(uid string) (model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", uid)
	ret0, _ := ret[0].(model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockOrderMockRecorder) GetOrder(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockOrder)(nil).GetOrder), uid)
}

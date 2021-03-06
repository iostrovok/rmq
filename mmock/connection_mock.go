// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/iostrovok/rmq (interfaces: Connection)

// Package mmock is a generated GoMock package.
package mmock

import (
	gomock "github.com/golang/mock/gomock"
	rmq "github.com/iostrovok/rmq"
	reflect "reflect"
)

// MockConnection is a mock of Connection interface
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// CollectStats mocks base method
func (m *MockConnection) CollectStats(arg0 []string) rmq.Stats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CollectStats", arg0)
	ret0, _ := ret[0].(rmq.Stats)
	return ret0
}

// CollectStats indicates an expected call of CollectStats
func (mr *MockConnectionMockRecorder) CollectStats(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CollectStats", reflect.TypeOf((*MockConnection)(nil).CollectStats), arg0)
}

// GetOpenQueues mocks base method
func (m *MockConnection) GetOpenQueues() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOpenQueues")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetOpenQueues indicates an expected call of GetOpenQueues
func (mr *MockConnectionMockRecorder) GetOpenQueues() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpenQueues", reflect.TypeOf((*MockConnection)(nil).GetOpenQueues))
}

// OpenQueue mocks base method
func (m *MockConnection) OpenQueue(arg0 string) rmq.Queue {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenQueue", arg0)
	ret0, _ := ret[0].(rmq.Queue)
	return ret0
}

// OpenQueue indicates an expected call of OpenQueue
func (mr *MockConnectionMockRecorder) OpenQueue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenQueue", reflect.TypeOf((*MockConnection)(nil).OpenQueue), arg0)
}

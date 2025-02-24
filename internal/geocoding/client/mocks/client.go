// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	client "calindra/internal/geocoding/client"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockGeoCodingClient is a mock of GeoCodingClient interface
type MockGeoCodingClient struct {
	ctrl     *gomock.Controller
	recorder *MockGeoCodingClientMockRecorder
}

// MockGeoCodingClientMockRecorder is the mock recorder for MockGeoCodingClient
type MockGeoCodingClientMockRecorder struct {
	mock *MockGeoCodingClient
}

// NewMockGeoCodingClient creates a new mock instance
func NewMockGeoCodingClient(ctrl *gomock.Controller) *MockGeoCodingClient {
	mock := &MockGeoCodingClient{ctrl: ctrl}
	mock.recorder = &MockGeoCodingClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGeoCodingClient) EXPECT() *MockGeoCodingClientMockRecorder {
	return m.recorder
}

// FindAddress mocks base method
func (m *MockGeoCodingClient) FindAddress(address string) (*client.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAddress", address)
	ret0, _ := ret[0].(*client.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAddress indicates an expected call of FindAddress
func (mr *MockGeoCodingClientMockRecorder) FindAddress(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAddress", reflect.TypeOf((*MockGeoCodingClient)(nil).FindAddress), address)
}

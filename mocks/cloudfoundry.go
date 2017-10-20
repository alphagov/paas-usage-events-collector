// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/alphagov/paas-usage-events-collector/cloudfoundry (interfaces: Client,UsageEventsAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	cloudfoundry "github.com/alphagov/paas-usage-events-collector/cloudfoundry"
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
	time "time"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockClient) Get(arg0 string) (*http.Response, error) {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockClientMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClient)(nil).Get), arg0)
}

// MockUsageEventsAPI is a mock of UsageEventsAPI interface
type MockUsageEventsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockUsageEventsAPIMockRecorder
}

// MockUsageEventsAPIMockRecorder is the mock recorder for MockUsageEventsAPI
type MockUsageEventsAPIMockRecorder struct {
	mock *MockUsageEventsAPI
}

// NewMockUsageEventsAPI creates a new mock instance
func NewMockUsageEventsAPI(ctrl *gomock.Controller) *MockUsageEventsAPI {
	mock := &MockUsageEventsAPI{ctrl: ctrl}
	mock.recorder = &MockUsageEventsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsageEventsAPI) EXPECT() *MockUsageEventsAPIMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockUsageEventsAPI) Get(arg0 string, arg1 int, arg2 time.Duration) (*cloudfoundry.UsageEventList, error) {
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*cloudfoundry.UsageEventList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockUsageEventsAPIMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUsageEventsAPI)(nil).Get), arg0, arg1, arg2)
}

// Type mocks base method
func (m *MockUsageEventsAPI) Type() string {
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(string)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockUsageEventsAPIMockRecorder) Type() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockUsageEventsAPI)(nil).Type))
}

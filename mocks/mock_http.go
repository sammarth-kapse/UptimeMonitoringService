// Code generated by MockGen. DO NOT EDIT.
// Source: UptimeMonitoringService/httpRequests (interfaces: HttpController)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockHttpController is a mock of HttpController interface
type MockHttpController struct {
	ctrl     *gomock.Controller
	recorder *MockHttpControllerMockRecorder
}

// MockHttpControllerMockRecorder is the mock recorder for MockHttpController
type MockHttpControllerMockRecorder struct {
	mock *MockHttpController
}

// NewMockHttpController creates a new mock instance
func NewMockHttpController(ctrl *gomock.Controller) *MockHttpController {
	mock := &MockHttpController{ctrl: ctrl}
	mock.recorder = &MockHttpControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHttpController) EXPECT() *MockHttpControllerMockRecorder {
	return m.recorder
}

// MakeHTTPGetRequest mocks base method
func (m *MockHttpController) MakeHTTPGetRequest(arg0 int, arg1 string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeHTTPGetRequest", arg0, arg1)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeHTTPGetRequest indicates an expected call of MakeHTTPGetRequest
func (mr *MockHttpControllerMockRecorder) MakeHTTPGetRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeHTTPGetRequest", reflect.TypeOf((*MockHttpController)(nil).MakeHTTPGetRequest), arg0, arg1)
}
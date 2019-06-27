// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1 (interfaces: ValidatorServiceServer,ValidatorService_WaitForActivationServer)

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1"
	metadata "google.golang.org/grpc/metadata"
)

// MockValidatorServiceServer is a mock of ValidatorServiceServer interface
type MockValidatorServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorServiceServerMockRecorder
}

// MockValidatorServiceServerMockRecorder is the mock recorder for MockValidatorServiceServer
type MockValidatorServiceServerMockRecorder struct {
	mock *MockValidatorServiceServer
}

// NewMockValidatorServiceServer creates a new mock instance
func NewMockValidatorServiceServer(ctrl *gomock.Controller) *MockValidatorServiceServer {
	mock := &MockValidatorServiceServer{ctrl: ctrl}
	mock.recorder = &MockValidatorServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidatorServiceServer) EXPECT() *MockValidatorServiceServerMockRecorder {
	return m.recorder
}

// CommitteeAssignment mocks base method
func (m *MockValidatorServiceServer) CommitteeAssignment(arg0 context.Context, arg1 *v1.CommitteeAssignmentsRequest) (*v1.CommitteeAssignmentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitteeAssignment", arg0, arg1)
	ret0, _ := ret[0].(*v1.CommitteeAssignmentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommitteeAssignment indicates an expected call of CommitteeAssignment
func (mr *MockValidatorServiceServerMockRecorder) CommitteeAssignment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitteeAssignment", reflect.TypeOf((*MockValidatorServiceServer)(nil).CommitteeAssignment), arg0, arg1)
}

// ExitedValidators mocks base method
func (m *MockValidatorServiceServer) ExitedValidators(arg0 context.Context, arg1 *v1.ExitedValidatorsRequest) (*v1.ExitedValidatorsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExitedValidators", arg0, arg1)
	ret0, _ := ret[0].(*v1.ExitedValidatorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExitedValidators indicates an expected call of ExitedValidators
func (mr *MockValidatorServiceServerMockRecorder) ExitedValidators(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExitedValidators", reflect.TypeOf((*MockValidatorServiceServer)(nil).ExitedValidators), arg0, arg1)
}

// ValidatorIndex mocks base method
func (m *MockValidatorServiceServer) ValidatorIndex(arg0 context.Context, arg1 *v1.ValidatorIndexRequest) (*v1.ValidatorIndexResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatorIndex", arg0, arg1)
	ret0, _ := ret[0].(*v1.ValidatorIndexResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidatorIndex indicates an expected call of ValidatorIndex
func (mr *MockValidatorServiceServerMockRecorder) ValidatorIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorIndex", reflect.TypeOf((*MockValidatorServiceServer)(nil).ValidatorIndex), arg0, arg1)
}

// ValidatorPerformance mocks base method
func (m *MockValidatorServiceServer) ValidatorPerformance(arg0 context.Context, arg1 *v1.ValidatorPerformanceRequest) (*v1.ValidatorPerformanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatorPerformance", arg0, arg1)
	ret0, _ := ret[0].(*v1.ValidatorPerformanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidatorPerformance indicates an expected call of ValidatorPerformance
func (mr *MockValidatorServiceServerMockRecorder) ValidatorPerformance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorPerformance", reflect.TypeOf((*MockValidatorServiceServer)(nil).ValidatorPerformance), arg0, arg1)
}

// ValidatorStatus mocks base method
func (m *MockValidatorServiceServer) ValidatorStatus(arg0 context.Context, arg1 *v1.ValidatorIndexRequest) (*v1.ValidatorStatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatorStatus", arg0, arg1)
	ret0, _ := ret[0].(*v1.ValidatorStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidatorStatus indicates an expected call of ValidatorStatus
func (mr *MockValidatorServiceServerMockRecorder) ValidatorStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorStatus", reflect.TypeOf((*MockValidatorServiceServer)(nil).ValidatorStatus), arg0, arg1)
}

// WaitForActivation mocks base method
func (m *MockValidatorServiceServer) WaitForActivation(arg0 *v1.ValidatorActivationRequest, arg1 v1.ValidatorService_WaitForActivationServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForActivation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForActivation indicates an expected call of WaitForActivation
func (mr *MockValidatorServiceServerMockRecorder) WaitForActivation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForActivation", reflect.TypeOf((*MockValidatorServiceServer)(nil).WaitForActivation), arg0, arg1)
}

// MockValidatorService_WaitForActivationServer is a mock of ValidatorService_WaitForActivationServer interface
type MockValidatorService_WaitForActivationServer struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorService_WaitForActivationServerMockRecorder
}

// MockValidatorService_WaitForActivationServerMockRecorder is the mock recorder for MockValidatorService_WaitForActivationServer
type MockValidatorService_WaitForActivationServerMockRecorder struct {
	mock *MockValidatorService_WaitForActivationServer
}

// NewMockValidatorService_WaitForActivationServer creates a new mock instance
func NewMockValidatorService_WaitForActivationServer(ctrl *gomock.Controller) *MockValidatorService_WaitForActivationServer {
	mock := &MockValidatorService_WaitForActivationServer{ctrl: ctrl}
	mock.recorder = &MockValidatorService_WaitForActivationServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidatorService_WaitForActivationServer) EXPECT() *MockValidatorService_WaitForActivationServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockValidatorService_WaitForActivationServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockValidatorService_WaitForActivationServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockValidatorService_WaitForActivationServer)(nil).Context))
}

// RecvMsg mocks base method
func (m *MockValidatorService_WaitForActivationServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockValidatorService_WaitForActivationServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockValidatorService_WaitForActivationServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockValidatorService_WaitForActivationServer) Send(arg0 *v1.ValidatorActivationResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockValidatorService_WaitForActivationServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockValidatorService_WaitForActivationServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockValidatorService_WaitForActivationServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockValidatorService_WaitForActivationServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockValidatorService_WaitForActivationServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockValidatorService_WaitForActivationServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockValidatorService_WaitForActivationServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockValidatorService_WaitForActivationServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockValidatorService_WaitForActivationServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockValidatorService_WaitForActivationServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockValidatorService_WaitForActivationServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockValidatorService_WaitForActivationServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockValidatorService_WaitForActivationServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockValidatorService_WaitForActivationServer)(nil).SetTrailer), arg0)
}

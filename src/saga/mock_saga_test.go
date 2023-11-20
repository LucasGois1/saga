// Code generated by MockGen. DO NOT EDIT.
// Source: ./entities.go

// Package saga_test is a generated GoMock package.
package saga_test

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	saga "github.com/LucasGois1/saga-pattern/src/saga"
)

// MockSaga is a mock of Saga interface.
type MockSaga struct {
	ctrl     *gomock.Controller
	recorder *MockSagaMockRecorder
}

// MockSagaMockRecorder is the mock recorder for MockSaga.
type MockSagaMockRecorder struct {
	mock *MockSaga
}

// NewMockSaga creates a new mock instance.
func NewMockSaga(ctrl *gomock.Controller) *MockSaga {
	mock := &MockSaga{ctrl: ctrl}
	mock.recorder = &MockSagaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSaga) EXPECT() *MockSagaMockRecorder {
	return m.recorder
}

// Compensate mocks base method.
func (m *MockSaga) Compensate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compensate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Compensate indicates an expected call of Compensate.
func (mr *MockSagaMockRecorder) Compensate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compensate", reflect.TypeOf((*MockSaga)(nil).Compensate))
}

// Execute mocks base method.
func (m *MockSaga) Execute() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute")
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockSagaMockRecorder) Execute() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockSaga)(nil).Execute))
}

// Identifier mocks base method.
func (m *MockSaga) Identifier() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Identifier")
	ret0, _ := ret[0].(string)
	return ret0
}

// Identifier indicates an expected call of Identifier.
func (mr *MockSagaMockRecorder) Identifier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identifier", reflect.TypeOf((*MockSaga)(nil).Identifier))
}

// MockIBuilder is a mock of IBuilder interface.
type MockIBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockIBuilderMockRecorder
}

// MockIBuilderMockRecorder is the mock recorder for MockIBuilder.
type MockIBuilderMockRecorder struct {
	mock *MockIBuilder
}

// NewMockIBuilder creates a new mock instance.
func NewMockIBuilder(ctrl *gomock.Controller) *MockIBuilder {
	mock := &MockIBuilder{ctrl: ctrl}
	mock.recorder = &MockIBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBuilder) EXPECT() *MockIBuilderMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockIBuilder) Build() saga.Saga {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build")
	ret0, _ := ret[0].(saga.Saga)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockIBuilderMockRecorder) Build() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockIBuilder)(nil).Build))
}

// MockSagaBuilder is a mock of SagaBuilder interface.
type MockSagaBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockSagaBuilderMockRecorder
}

// MockSagaBuilderMockRecorder is the mock recorder for MockSagaBuilder.
type MockSagaBuilderMockRecorder struct {
	mock *MockSagaBuilder
}

// NewMockSagaBuilder creates a new mock instance.
func NewMockSagaBuilder(ctrl *gomock.Controller) *MockSagaBuilder {
	mock := &MockSagaBuilder{ctrl: ctrl}
	mock.recorder = &MockSagaBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSagaBuilder) EXPECT() *MockSagaBuilderMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockSagaBuilder) Build() saga.Saga {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build")
	ret0, _ := ret[0].(saga.Saga)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockSagaBuilderMockRecorder) Build() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockSagaBuilder)(nil).Build))
}

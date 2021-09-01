// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/eks-anywhere/pkg/addonmanager/addonclients (interfaces: Flux)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	types "github.com/aws/eks-anywhere/pkg/types"
	gomock "github.com/golang/mock/gomock"
)

// MockFlux is a mock of Flux interface.
type MockFlux struct {
	ctrl     *gomock.Controller
	recorder *MockFluxMockRecorder
}

// MockFluxMockRecorder is the mock recorder for MockFlux.
type MockFluxMockRecorder struct {
	mock *MockFlux
}

// NewMockFlux creates a new mock instance.
func NewMockFlux(ctrl *gomock.Controller) *MockFlux {
	mock := &MockFlux{ctrl: ctrl}
	mock.recorder = &MockFluxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFlux) EXPECT() *MockFluxMockRecorder {
	return m.recorder
}

// BootstrapToolkitsComponents mocks base method.
func (m *MockFlux) BootstrapToolkitsComponents(arg0 context.Context, arg1 *types.Cluster, arg2 *v1alpha1.GitOpsConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BootstrapToolkitsComponents", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// BootstrapToolkitsComponents indicates an expected call of BootstrapToolkitsComponents.
func (mr *MockFluxMockRecorder) BootstrapToolkitsComponents(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BootstrapToolkitsComponents", reflect.TypeOf((*MockFlux)(nil).BootstrapToolkitsComponents), arg0, arg1, arg2)
}

// ForceReconcileGitRepo mocks base method.
func (m *MockFlux) ForceReconcileGitRepo(arg0 context.Context, arg1 *types.Cluster, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceReconcileGitRepo", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForceReconcileGitRepo indicates an expected call of ForceReconcileGitRepo.
func (mr *MockFluxMockRecorder) ForceReconcileGitRepo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceReconcileGitRepo", reflect.TypeOf((*MockFlux)(nil).ForceReconcileGitRepo), arg0, arg1, arg2)
}

// PauseKustomization mocks base method.
func (m *MockFlux) PauseKustomization(arg0 context.Context, arg1 *types.Cluster, arg2 *v1alpha1.GitOpsConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PauseKustomization", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PauseKustomization indicates an expected call of PauseKustomization.
func (mr *MockFluxMockRecorder) PauseKustomization(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PauseKustomization", reflect.TypeOf((*MockFlux)(nil).PauseKustomization), arg0, arg1, arg2)
}

// ResumeKustomization mocks base method.
func (m *MockFlux) ResumeKustomization(arg0 context.Context, arg1 *types.Cluster, arg2 *v1alpha1.GitOpsConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResumeKustomization", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResumeKustomization indicates an expected call of ResumeKustomization.
func (mr *MockFluxMockRecorder) ResumeKustomization(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResumeKustomization", reflect.TypeOf((*MockFlux)(nil).ResumeKustomization), arg0, arg1, arg2)
}

// UninstallToolkitsComponents mocks base method.
func (m *MockFlux) UninstallToolkitsComponents(arg0 context.Context, arg1 *types.Cluster, arg2 *v1alpha1.GitOpsConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UninstallToolkitsComponents", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UninstallToolkitsComponents indicates an expected call of UninstallToolkitsComponents.
func (mr *MockFluxMockRecorder) UninstallToolkitsComponents(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UninstallToolkitsComponents", reflect.TypeOf((*MockFlux)(nil).UninstallToolkitsComponents), arg0, arg1, arg2)
}
// Code generated by MockGen. DO NOT EDIT.
// Source: group_member.go
//
// Generated by this command:
//
//	mockgen -source=group_member.go -destination=../../test/mock/repository/group_member_mock.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	model "backend/domain/model"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockGroupMember is a mock of GroupMember interface.
type MockGroupMember struct {
	ctrl     *gomock.Controller
	recorder *MockGroupMemberMockRecorder
	isgomock struct{}
}

// MockGroupMemberMockRecorder is the mock recorder for MockGroupMember.
type MockGroupMemberMockRecorder struct {
	mock *MockGroupMember
}

// NewMockGroupMember creates a new mock instance.
func NewMockGroupMember(ctrl *gomock.Controller) *MockGroupMember {
	mock := &MockGroupMember{ctrl: ctrl}
	mock.recorder = &MockGroupMemberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupMember) EXPECT() *MockGroupMemberMockRecorder {
	return m.recorder
}

// CountByGroupID mocks base method.
func (m *MockGroupMember) CountByGroupID(ctx context.Context, groupID model.GroupID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountByGroupID", ctx, groupID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountByGroupID indicates an expected call of CountByGroupID.
func (mr *MockGroupMemberMockRecorder) CountByGroupID(ctx, groupID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountByGroupID", reflect.TypeOf((*MockGroupMember)(nil).CountByGroupID), ctx, groupID)
}

// FindByAccountID mocks base method.
func (m *MockGroupMember) FindByAccountID(ctx context.Context, accountID model.AccountID) ([]model.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByAccountID", ctx, accountID)
	ret0, _ := ret[0].([]model.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByAccountID indicates an expected call of FindByAccountID.
func (mr *MockGroupMemberMockRecorder) FindByAccountID(ctx, accountID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByAccountID", reflect.TypeOf((*MockGroupMember)(nil).FindByAccountID), ctx, accountID)
}

// FindByGroupID mocks base method.
func (m *MockGroupMember) FindByGroupID(ctx context.Context, groupID model.GroupID) ([]model.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByGroupID", ctx, groupID)
	ret0, _ := ret[0].([]model.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByGroupID indicates an expected call of FindByGroupID.
func (mr *MockGroupMemberMockRecorder) FindByGroupID(ctx, groupID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByGroupID", reflect.TypeOf((*MockGroupMember)(nil).FindByGroupID), ctx, groupID)
}

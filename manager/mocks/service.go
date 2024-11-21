// Code generated by mockery v2.43.2. DO NOT EDIT.

// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	manager "github.com/ultravioletrs/cocos/manager"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// FetchAttestationPolicy provides a mock function with given fields: ctx, computationID
func (_m *Service) FetchAttestationPolicy(ctx context.Context, computationID string) ([]byte, error) {
	ret := _m.Called(ctx, computationID)

	if len(ret) == 0 {
		panic("no return value specified for FetchAttestationPolicy")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]byte, error)); ok {
		return rf(ctx, computationID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []byte); ok {
		r0 = rf(ctx, computationID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, computationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReportBrokenConnection provides a mock function with given fields: addr
func (_m *Service) ReportBrokenConnection(addr string) {
	_m.Called(addr)
}

// ReturnSVMInfo provides a mock function with given fields: ctx
func (_m *Service) ReturnSVMInfo(ctx context.Context) (string, int, string, string) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ReturnSVMInfo")
	}

	var r0 string
	var r1 int
	var r2 string
	var r3 string
	if rf, ok := ret.Get(0).(func(context.Context) (string, int, string, string)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) int); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context) string); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Get(2).(string)
	}

	if rf, ok := ret.Get(3).(func(context.Context) string); ok {
		r3 = rf(ctx)
	} else {
		r3 = ret.Get(3).(string)
	}

	return r0, r1, r2, r3
}

// Run provides a mock function with given fields: ctx, c
func (_m *Service) Run(ctx context.Context, c *manager.ComputationRunReq) (string, error) {
	ret := _m.Called(ctx, c)

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *manager.ComputationRunReq) (string, error)); ok {
		return rf(ctx, c)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *manager.ComputationRunReq) string); ok {
		r0 = rf(ctx, c)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *manager.ComputationRunReq) error); ok {
		r1 = rf(ctx, c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stop provides a mock function with given fields: ctx, computationID
func (_m *Service) Stop(ctx context.Context, computationID string) error {
	ret := _m.Called(ctx, computationID)

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, computationID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	agent "github.com/ultravioletrs/cocos/agent"

	metadata "google.golang.org/grpc/metadata"

	mock "github.com/stretchr/testify/mock"
)

// AgentService_IMAMeasurementsClient is an autogenerated mock type for the AgentService_IMAMeasurementsClient type
type AgentService_IMAMeasurementsClient[Res interface{}] struct {
	mock.Mock
}

type AgentService_IMAMeasurementsClient_Expecter[Res interface{}] struct {
	mock *mock.Mock
}

func (_m *AgentService_IMAMeasurementsClient[Res]) EXPECT() *AgentService_IMAMeasurementsClient_Expecter[Res] {
	return &AgentService_IMAMeasurementsClient_Expecter[Res]{mock: &_m.Mock}
}

// CloseSend provides a mock function with no fields
func (_m *AgentService_IMAMeasurementsClient[Res]) CloseSend() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CloseSend")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AgentService_IMAMeasurementsClient_CloseSend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseSend'
type AgentService_IMAMeasurementsClient_CloseSend_Call[Res interface{}] struct {
	*mock.Call
}

// CloseSend is a helper method to define mock.On call
func (_e *AgentService_IMAMeasurementsClient_Expecter[Res]) CloseSend() *AgentService_IMAMeasurementsClient_CloseSend_Call[Res] {
	return &AgentService_IMAMeasurementsClient_CloseSend_Call[Res]{Call: _e.mock.On("CloseSend")}
}

func (_c *AgentService_IMAMeasurementsClient_CloseSend_Call[Res]) Run(run func()) *AgentService_IMAMeasurementsClient_CloseSend_Call[Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AgentService_IMAMeasurementsClient_CloseSend_Call[Res]) Return(_a0 error) *AgentService_IMAMeasurementsClient_CloseSend_Call[Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AgentService_IMAMeasurementsClient_CloseSend_Call[Res]) RunAndReturn(run func() error) *AgentService_IMAMeasurementsClient_CloseSend_Call[Res] {
	_c.Call.Return(run)
	return _c
}

// Context provides a mock function with no fields
func (_m *AgentService_IMAMeasurementsClient[Res]) Context() context.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Context")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// AgentService_IMAMeasurementsClient_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type AgentService_IMAMeasurementsClient_Context_Call[Res interface{}] struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *AgentService_IMAMeasurementsClient_Expecter[Res]) Context() *AgentService_IMAMeasurementsClient_Context_Call[Res] {
	return &AgentService_IMAMeasurementsClient_Context_Call[Res]{Call: _e.mock.On("Context")}
}

func (_c *AgentService_IMAMeasurementsClient_Context_Call[Res]) Run(run func()) *AgentService_IMAMeasurementsClient_Context_Call[Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AgentService_IMAMeasurementsClient_Context_Call[Res]) Return(_a0 context.Context) *AgentService_IMAMeasurementsClient_Context_Call[Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AgentService_IMAMeasurementsClient_Context_Call[Res]) RunAndReturn(run func() context.Context) *AgentService_IMAMeasurementsClient_Context_Call[Res] {
	_c.Call.Return(run)
	return _c
}

// Header provides a mock function with no fields
func (_m *AgentService_IMAMeasurementsClient[Res]) Header() (metadata.MD, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Header")
	}

	var r0 metadata.MD
	var r1 error
	if rf, ok := ret.Get(0).(func() (metadata.MD, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AgentService_IMAMeasurementsClient_Header_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Header'
type AgentService_IMAMeasurementsClient_Header_Call[Res interface{}] struct {
	*mock.Call
}

// Header is a helper method to define mock.On call
func (_e *AgentService_IMAMeasurementsClient_Expecter[Res]) Header() *AgentService_IMAMeasurementsClient_Header_Call[Res] {
	return &AgentService_IMAMeasurementsClient_Header_Call[Res]{Call: _e.mock.On("Header")}
}

func (_c *AgentService_IMAMeasurementsClient_Header_Call[Res]) Run(run func()) *AgentService_IMAMeasurementsClient_Header_Call[Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AgentService_IMAMeasurementsClient_Header_Call[Res]) Return(_a0 metadata.MD, _a1 error) *AgentService_IMAMeasurementsClient_Header_Call[Res] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AgentService_IMAMeasurementsClient_Header_Call[Res]) RunAndReturn(run func() (metadata.MD, error)) *AgentService_IMAMeasurementsClient_Header_Call[Res] {
	_c.Call.Return(run)
	return _c
}

// Recv provides a mock function with no fields
func (_m *AgentService_IMAMeasurementsClient[Res]) Recv() (*agent.IMAMeasurementsResponse, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Recv")
	}

	var r0 *agent.IMAMeasurementsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() (*agent.IMAMeasurementsResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *agent.IMAMeasurementsResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*agent.IMAMeasurementsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AgentService_IMAMeasurementsClient_Recv_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Recv'
type AgentService_IMAMeasurementsClient_Recv_Call[Res interface{}] struct {
	*mock.Call
}

// Recv is a helper method to define mock.On call
func (_e *AgentService_IMAMeasurementsClient_Expecter[Res]) Recv() *AgentService_IMAMeasurementsClient_Recv_Call[Res] {
	return &AgentService_IMAMeasurementsClient_Recv_Call[Res]{Call: _e.mock.On("Recv")}
}

func (_c *AgentService_IMAMeasurementsClient_Recv_Call[Res]) Run(run func()) *AgentService_IMAMeasurementsClient_Recv_Call[Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AgentService_IMAMeasurementsClient_Recv_Call[Res]) Return(_a0 *agent.IMAMeasurementsResponse, _a1 error) *AgentService_IMAMeasurementsClient_Recv_Call[Res] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AgentService_IMAMeasurementsClient_Recv_Call[Res]) RunAndReturn(run func() (*agent.IMAMeasurementsResponse, error)) *AgentService_IMAMeasurementsClient_Recv_Call[Res] {
	_c.Call.Return(run)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *AgentService_IMAMeasurementsClient[Res]) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for RecvMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AgentService_IMAMeasurementsClient_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type AgentService_IMAMeasurementsClient_RecvMsg_Call[Res interface{}] struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *AgentService_IMAMeasurementsClient_Expecter[Res]) RecvMsg(m interface{}) *AgentService_IMAMeasurementsClient_RecvMsg_Call[Res] {
	return &AgentService_IMAMeasurementsClient_RecvMsg_Call[Res]{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *AgentService_IMAMeasurementsClient_RecvMsg_Call[Res]) Run(run func(m interface{})) *AgentService_IMAMeasurementsClient_RecvMsg_Call[Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *AgentService_IMAMeasurementsClient_RecvMsg_Call[Res]) Return(_a0 error) *AgentService_IMAMeasurementsClient_RecvMsg_Call[Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AgentService_IMAMeasurementsClient_RecvMsg_Call[Res]) RunAndReturn(run func(interface{}) error) *AgentService_IMAMeasurementsClient_RecvMsg_Call[Res] {
	_c.Call.Return(run)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *AgentService_IMAMeasurementsClient[Res]) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for SendMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AgentService_IMAMeasurementsClient_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type AgentService_IMAMeasurementsClient_SendMsg_Call[Res interface{}] struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *AgentService_IMAMeasurementsClient_Expecter[Res]) SendMsg(m interface{}) *AgentService_IMAMeasurementsClient_SendMsg_Call[Res] {
	return &AgentService_IMAMeasurementsClient_SendMsg_Call[Res]{Call: _e.mock.On("SendMsg", m)}
}

func (_c *AgentService_IMAMeasurementsClient_SendMsg_Call[Res]) Run(run func(m interface{})) *AgentService_IMAMeasurementsClient_SendMsg_Call[Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *AgentService_IMAMeasurementsClient_SendMsg_Call[Res]) Return(_a0 error) *AgentService_IMAMeasurementsClient_SendMsg_Call[Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AgentService_IMAMeasurementsClient_SendMsg_Call[Res]) RunAndReturn(run func(interface{}) error) *AgentService_IMAMeasurementsClient_SendMsg_Call[Res] {
	_c.Call.Return(run)
	return _c
}

// Trailer provides a mock function with no fields
func (_m *AgentService_IMAMeasurementsClient[Res]) Trailer() metadata.MD {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Trailer")
	}

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	return r0
}

// AgentService_IMAMeasurementsClient_Trailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Trailer'
type AgentService_IMAMeasurementsClient_Trailer_Call[Res interface{}] struct {
	*mock.Call
}

// Trailer is a helper method to define mock.On call
func (_e *AgentService_IMAMeasurementsClient_Expecter[Res]) Trailer() *AgentService_IMAMeasurementsClient_Trailer_Call[Res] {
	return &AgentService_IMAMeasurementsClient_Trailer_Call[Res]{Call: _e.mock.On("Trailer")}
}

func (_c *AgentService_IMAMeasurementsClient_Trailer_Call[Res]) Run(run func()) *AgentService_IMAMeasurementsClient_Trailer_Call[Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AgentService_IMAMeasurementsClient_Trailer_Call[Res]) Return(_a0 metadata.MD) *AgentService_IMAMeasurementsClient_Trailer_Call[Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AgentService_IMAMeasurementsClient_Trailer_Call[Res]) RunAndReturn(run func() metadata.MD) *AgentService_IMAMeasurementsClient_Trailer_Call[Res] {
	_c.Call.Return(run)
	return _c
}

// NewAgentService_IMAMeasurementsClient creates a new instance of AgentService_IMAMeasurementsClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAgentService_IMAMeasurementsClient[Res interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *AgentService_IMAMeasurementsClient[Res] {
	mock := &AgentService_IMAMeasurementsClient[Res]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

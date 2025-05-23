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

// AgentService_DataClient is an autogenerated mock type for the AgentService_DataClient type
type AgentService_DataClient[Req interface{}, Res interface{}] struct {
	mock.Mock
}

type AgentService_DataClient_Expecter[Req interface{}, Res interface{}] struct {
	mock *mock.Mock
}

func (_m *AgentService_DataClient[Req, Res]) EXPECT() *AgentService_DataClient_Expecter[Req, Res] {
	return &AgentService_DataClient_Expecter[Req, Res]{mock: &_m.Mock}
}

// CloseAndRecv provides a mock function with no fields
func (_m *AgentService_DataClient[Req, Res]) CloseAndRecv() (*agent.DataResponse, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CloseAndRecv")
	}

	var r0 *agent.DataResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() (*agent.DataResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *agent.DataResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*agent.DataResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AgentService_DataClient_CloseAndRecv_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseAndRecv'
type AgentService_DataClient_CloseAndRecv_Call[Req interface{}, Res interface{}] struct {
	*mock.Call
}

// CloseAndRecv is a helper method to define mock.On call
func (_e *AgentService_DataClient_Expecter[Req, Res]) CloseAndRecv() *AgentService_DataClient_CloseAndRecv_Call[Req, Res] {
	return &AgentService_DataClient_CloseAndRecv_Call[Req, Res]{Call: _e.mock.On("CloseAndRecv")}
}

func (_c *AgentService_DataClient_CloseAndRecv_Call[Req, Res]) Run(run func()) *AgentService_DataClient_CloseAndRecv_Call[Req, Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AgentService_DataClient_CloseAndRecv_Call[Req, Res]) Return(_a0 *agent.DataResponse, _a1 error) *AgentService_DataClient_CloseAndRecv_Call[Req, Res] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AgentService_DataClient_CloseAndRecv_Call[Req, Res]) RunAndReturn(run func() (*agent.DataResponse, error)) *AgentService_DataClient_CloseAndRecv_Call[Req, Res] {
	_c.Call.Return(run)
	return _c
}

// CloseSend provides a mock function with no fields
func (_m *AgentService_DataClient[Req, Res]) CloseSend() error {
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

// AgentService_DataClient_CloseSend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseSend'
type AgentService_DataClient_CloseSend_Call[Req interface{}, Res interface{}] struct {
	*mock.Call
}

// CloseSend is a helper method to define mock.On call
func (_e *AgentService_DataClient_Expecter[Req, Res]) CloseSend() *AgentService_DataClient_CloseSend_Call[Req, Res] {
	return &AgentService_DataClient_CloseSend_Call[Req, Res]{Call: _e.mock.On("CloseSend")}
}

func (_c *AgentService_DataClient_CloseSend_Call[Req, Res]) Run(run func()) *AgentService_DataClient_CloseSend_Call[Req, Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AgentService_DataClient_CloseSend_Call[Req, Res]) Return(_a0 error) *AgentService_DataClient_CloseSend_Call[Req, Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AgentService_DataClient_CloseSend_Call[Req, Res]) RunAndReturn(run func() error) *AgentService_DataClient_CloseSend_Call[Req, Res] {
	_c.Call.Return(run)
	return _c
}

// Context provides a mock function with no fields
func (_m *AgentService_DataClient[Req, Res]) Context() context.Context {
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

// AgentService_DataClient_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type AgentService_DataClient_Context_Call[Req interface{}, Res interface{}] struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *AgentService_DataClient_Expecter[Req, Res]) Context() *AgentService_DataClient_Context_Call[Req, Res] {
	return &AgentService_DataClient_Context_Call[Req, Res]{Call: _e.mock.On("Context")}
}

func (_c *AgentService_DataClient_Context_Call[Req, Res]) Run(run func()) *AgentService_DataClient_Context_Call[Req, Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AgentService_DataClient_Context_Call[Req, Res]) Return(_a0 context.Context) *AgentService_DataClient_Context_Call[Req, Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AgentService_DataClient_Context_Call[Req, Res]) RunAndReturn(run func() context.Context) *AgentService_DataClient_Context_Call[Req, Res] {
	_c.Call.Return(run)
	return _c
}

// Header provides a mock function with no fields
func (_m *AgentService_DataClient[Req, Res]) Header() (metadata.MD, error) {
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

// AgentService_DataClient_Header_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Header'
type AgentService_DataClient_Header_Call[Req interface{}, Res interface{}] struct {
	*mock.Call
}

// Header is a helper method to define mock.On call
func (_e *AgentService_DataClient_Expecter[Req, Res]) Header() *AgentService_DataClient_Header_Call[Req, Res] {
	return &AgentService_DataClient_Header_Call[Req, Res]{Call: _e.mock.On("Header")}
}

func (_c *AgentService_DataClient_Header_Call[Req, Res]) Run(run func()) *AgentService_DataClient_Header_Call[Req, Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AgentService_DataClient_Header_Call[Req, Res]) Return(_a0 metadata.MD, _a1 error) *AgentService_DataClient_Header_Call[Req, Res] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AgentService_DataClient_Header_Call[Req, Res]) RunAndReturn(run func() (metadata.MD, error)) *AgentService_DataClient_Header_Call[Req, Res] {
	_c.Call.Return(run)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *AgentService_DataClient[Req, Res]) RecvMsg(m interface{}) error {
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

// AgentService_DataClient_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type AgentService_DataClient_RecvMsg_Call[Req interface{}, Res interface{}] struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *AgentService_DataClient_Expecter[Req, Res]) RecvMsg(m interface{}) *AgentService_DataClient_RecvMsg_Call[Req, Res] {
	return &AgentService_DataClient_RecvMsg_Call[Req, Res]{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *AgentService_DataClient_RecvMsg_Call[Req, Res]) Run(run func(m interface{})) *AgentService_DataClient_RecvMsg_Call[Req, Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *AgentService_DataClient_RecvMsg_Call[Req, Res]) Return(_a0 error) *AgentService_DataClient_RecvMsg_Call[Req, Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AgentService_DataClient_RecvMsg_Call[Req, Res]) RunAndReturn(run func(interface{}) error) *AgentService_DataClient_RecvMsg_Call[Req, Res] {
	_c.Call.Return(run)
	return _c
}

// Send provides a mock function with given fields: _a0
func (_m *AgentService_DataClient[Req, Res]) Send(_a0 *agent.DataRequest) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*agent.DataRequest) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AgentService_DataClient_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type AgentService_DataClient_Send_Call[Req interface{}, Res interface{}] struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - _a0 *agent.DataRequest
func (_e *AgentService_DataClient_Expecter[Req, Res]) Send(_a0 interface{}) *AgentService_DataClient_Send_Call[Req, Res] {
	return &AgentService_DataClient_Send_Call[Req, Res]{Call: _e.mock.On("Send", _a0)}
}

func (_c *AgentService_DataClient_Send_Call[Req, Res]) Run(run func(_a0 *agent.DataRequest)) *AgentService_DataClient_Send_Call[Req, Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*agent.DataRequest))
	})
	return _c
}

func (_c *AgentService_DataClient_Send_Call[Req, Res]) Return(_a0 error) *AgentService_DataClient_Send_Call[Req, Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AgentService_DataClient_Send_Call[Req, Res]) RunAndReturn(run func(*agent.DataRequest) error) *AgentService_DataClient_Send_Call[Req, Res] {
	_c.Call.Return(run)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *AgentService_DataClient[Req, Res]) SendMsg(m interface{}) error {
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

// AgentService_DataClient_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type AgentService_DataClient_SendMsg_Call[Req interface{}, Res interface{}] struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *AgentService_DataClient_Expecter[Req, Res]) SendMsg(m interface{}) *AgentService_DataClient_SendMsg_Call[Req, Res] {
	return &AgentService_DataClient_SendMsg_Call[Req, Res]{Call: _e.mock.On("SendMsg", m)}
}

func (_c *AgentService_DataClient_SendMsg_Call[Req, Res]) Run(run func(m interface{})) *AgentService_DataClient_SendMsg_Call[Req, Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *AgentService_DataClient_SendMsg_Call[Req, Res]) Return(_a0 error) *AgentService_DataClient_SendMsg_Call[Req, Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AgentService_DataClient_SendMsg_Call[Req, Res]) RunAndReturn(run func(interface{}) error) *AgentService_DataClient_SendMsg_Call[Req, Res] {
	_c.Call.Return(run)
	return _c
}

// Trailer provides a mock function with no fields
func (_m *AgentService_DataClient[Req, Res]) Trailer() metadata.MD {
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

// AgentService_DataClient_Trailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Trailer'
type AgentService_DataClient_Trailer_Call[Req interface{}, Res interface{}] struct {
	*mock.Call
}

// Trailer is a helper method to define mock.On call
func (_e *AgentService_DataClient_Expecter[Req, Res]) Trailer() *AgentService_DataClient_Trailer_Call[Req, Res] {
	return &AgentService_DataClient_Trailer_Call[Req, Res]{Call: _e.mock.On("Trailer")}
}

func (_c *AgentService_DataClient_Trailer_Call[Req, Res]) Run(run func()) *AgentService_DataClient_Trailer_Call[Req, Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AgentService_DataClient_Trailer_Call[Req, Res]) Return(_a0 metadata.MD) *AgentService_DataClient_Trailer_Call[Req, Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AgentService_DataClient_Trailer_Call[Req, Res]) RunAndReturn(run func() metadata.MD) *AgentService_DataClient_Trailer_Call[Req, Res] {
	_c.Call.Return(run)
	return _c
}

// NewAgentService_DataClient creates a new instance of AgentService_DataClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAgentService_DataClient[Req interface{}, Res interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *AgentService_DataClient[Req, Res] {
	mock := &AgentService_DataClient[Req, Res]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

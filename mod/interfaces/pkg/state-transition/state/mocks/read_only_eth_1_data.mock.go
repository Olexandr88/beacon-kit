// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ReadOnlyEth1Data is an autogenerated mock type for the ReadOnlyEth1Data type
type ReadOnlyEth1Data[Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}] struct {
	mock.Mock
}

type ReadOnlyEth1Data_Expecter[Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}] struct {
	mock *mock.Mock
}

func (_m *ReadOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT]) EXPECT() *ReadOnlyEth1Data_Expecter[Eth1DataT, ExecutionPayloadHeaderT] {
	return &ReadOnlyEth1Data_Expecter[Eth1DataT, ExecutionPayloadHeaderT]{mock: &_m.Mock}
}

// GetEth1Data provides a mock function with given fields:
func (_m *ReadOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT]) GetEth1Data() (Eth1DataT, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetEth1Data")
	}

	var r0 Eth1DataT
	var r1 error
	if rf, ok := ret.Get(0).(func() (Eth1DataT, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() Eth1DataT); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(Eth1DataT)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadOnlyEth1Data_GetEth1Data_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEth1Data'
type ReadOnlyEth1Data_GetEth1Data_Call[Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}] struct {
	*mock.Call
}

// GetEth1Data is a helper method to define mock.On call
func (_e *ReadOnlyEth1Data_Expecter[Eth1DataT, ExecutionPayloadHeaderT]) GetEth1Data() *ReadOnlyEth1Data_GetEth1Data_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	return &ReadOnlyEth1Data_GetEth1Data_Call[Eth1DataT, ExecutionPayloadHeaderT]{Call: _e.mock.On("GetEth1Data")}
}

func (_c *ReadOnlyEth1Data_GetEth1Data_Call[Eth1DataT, ExecutionPayloadHeaderT]) Run(run func()) *ReadOnlyEth1Data_GetEth1Data_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ReadOnlyEth1Data_GetEth1Data_Call[Eth1DataT, ExecutionPayloadHeaderT]) Return(_a0 Eth1DataT, _a1 error) *ReadOnlyEth1Data_GetEth1Data_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReadOnlyEth1Data_GetEth1Data_Call[Eth1DataT, ExecutionPayloadHeaderT]) RunAndReturn(run func() (Eth1DataT, error)) *ReadOnlyEth1Data_GetEth1Data_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	_c.Call.Return(run)
	return _c
}

// GetEth1DepositIndex provides a mock function with given fields:
func (_m *ReadOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT]) GetEth1DepositIndex() (uint64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetEth1DepositIndex")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadOnlyEth1Data_GetEth1DepositIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEth1DepositIndex'
type ReadOnlyEth1Data_GetEth1DepositIndex_Call[Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}] struct {
	*mock.Call
}

// GetEth1DepositIndex is a helper method to define mock.On call
func (_e *ReadOnlyEth1Data_Expecter[Eth1DataT, ExecutionPayloadHeaderT]) GetEth1DepositIndex() *ReadOnlyEth1Data_GetEth1DepositIndex_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	return &ReadOnlyEth1Data_GetEth1DepositIndex_Call[Eth1DataT, ExecutionPayloadHeaderT]{Call: _e.mock.On("GetEth1DepositIndex")}
}

func (_c *ReadOnlyEth1Data_GetEth1DepositIndex_Call[Eth1DataT, ExecutionPayloadHeaderT]) Run(run func()) *ReadOnlyEth1Data_GetEth1DepositIndex_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ReadOnlyEth1Data_GetEth1DepositIndex_Call[Eth1DataT, ExecutionPayloadHeaderT]) Return(_a0 uint64, _a1 error) *ReadOnlyEth1Data_GetEth1DepositIndex_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReadOnlyEth1Data_GetEth1DepositIndex_Call[Eth1DataT, ExecutionPayloadHeaderT]) RunAndReturn(run func() (uint64, error)) *ReadOnlyEth1Data_GetEth1DepositIndex_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	_c.Call.Return(run)
	return _c
}

// GetLatestExecutionPayloadHeader provides a mock function with given fields:
func (_m *ReadOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT]) GetLatestExecutionPayloadHeader() (ExecutionPayloadHeaderT, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLatestExecutionPayloadHeader")
	}

	var r0 ExecutionPayloadHeaderT
	var r1 error
	if rf, ok := ret.Get(0).(func() (ExecutionPayloadHeaderT, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() ExecutionPayloadHeaderT); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(ExecutionPayloadHeaderT)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadOnlyEth1Data_GetLatestExecutionPayloadHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLatestExecutionPayloadHeader'
type ReadOnlyEth1Data_GetLatestExecutionPayloadHeader_Call[Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}] struct {
	*mock.Call
}

// GetLatestExecutionPayloadHeader is a helper method to define mock.On call
func (_e *ReadOnlyEth1Data_Expecter[Eth1DataT, ExecutionPayloadHeaderT]) GetLatestExecutionPayloadHeader() *ReadOnlyEth1Data_GetLatestExecutionPayloadHeader_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	return &ReadOnlyEth1Data_GetLatestExecutionPayloadHeader_Call[Eth1DataT, ExecutionPayloadHeaderT]{Call: _e.mock.On("GetLatestExecutionPayloadHeader")}
}

func (_c *ReadOnlyEth1Data_GetLatestExecutionPayloadHeader_Call[Eth1DataT, ExecutionPayloadHeaderT]) Run(run func()) *ReadOnlyEth1Data_GetLatestExecutionPayloadHeader_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ReadOnlyEth1Data_GetLatestExecutionPayloadHeader_Call[Eth1DataT, ExecutionPayloadHeaderT]) Return(_a0 ExecutionPayloadHeaderT, _a1 error) *ReadOnlyEth1Data_GetLatestExecutionPayloadHeader_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReadOnlyEth1Data_GetLatestExecutionPayloadHeader_Call[Eth1DataT, ExecutionPayloadHeaderT]) RunAndReturn(run func() (ExecutionPayloadHeaderT, error)) *ReadOnlyEth1Data_GetLatestExecutionPayloadHeader_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	_c.Call.Return(run)
	return _c
}

// NewReadOnlyEth1Data creates a new instance of ReadOnlyEth1Data. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReadOnlyEth1Data[Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *ReadOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT] {
	mock := &ReadOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

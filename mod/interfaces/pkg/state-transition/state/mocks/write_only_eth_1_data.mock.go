// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// WriteOnlyEth1Data is an autogenerated mock type for the WriteOnlyEth1Data type
type WriteOnlyEth1Data[Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}] struct {
	mock.Mock
}

type WriteOnlyEth1Data_Expecter[Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}] struct {
	mock *mock.Mock
}

func (_m *WriteOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT]) EXPECT() *WriteOnlyEth1Data_Expecter[Eth1DataT, ExecutionPayloadHeaderT] {
	return &WriteOnlyEth1Data_Expecter[Eth1DataT, ExecutionPayloadHeaderT]{mock: &_m.Mock}
}

// SetEth1Data provides a mock function with given fields: _a0
func (_m *WriteOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT]) SetEth1Data(_a0 Eth1DataT) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for SetEth1Data")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Eth1DataT) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteOnlyEth1Data_SetEth1Data_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetEth1Data'
type WriteOnlyEth1Data_SetEth1Data_Call[Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}] struct {
	*mock.Call
}

// SetEth1Data is a helper method to define mock.On call
//   - _a0 Eth1DataT
func (_e *WriteOnlyEth1Data_Expecter[Eth1DataT, ExecutionPayloadHeaderT]) SetEth1Data(_a0 interface{}) *WriteOnlyEth1Data_SetEth1Data_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	return &WriteOnlyEth1Data_SetEth1Data_Call[Eth1DataT, ExecutionPayloadHeaderT]{Call: _e.mock.On("SetEth1Data", _a0)}
}

func (_c *WriteOnlyEth1Data_SetEth1Data_Call[Eth1DataT, ExecutionPayloadHeaderT]) Run(run func(_a0 Eth1DataT)) *WriteOnlyEth1Data_SetEth1Data_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Eth1DataT))
	})
	return _c
}

func (_c *WriteOnlyEth1Data_SetEth1Data_Call[Eth1DataT, ExecutionPayloadHeaderT]) Return(_a0 error) *WriteOnlyEth1Data_SetEth1Data_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WriteOnlyEth1Data_SetEth1Data_Call[Eth1DataT, ExecutionPayloadHeaderT]) RunAndReturn(run func(Eth1DataT) error) *WriteOnlyEth1Data_SetEth1Data_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	_c.Call.Return(run)
	return _c
}

// SetEth1DepositIndex provides a mock function with given fields: _a0
func (_m *WriteOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT]) SetEth1DepositIndex(_a0 uint64) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for SetEth1DepositIndex")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteOnlyEth1Data_SetEth1DepositIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetEth1DepositIndex'
type WriteOnlyEth1Data_SetEth1DepositIndex_Call[Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}] struct {
	*mock.Call
}

// SetEth1DepositIndex is a helper method to define mock.On call
//   - _a0 uint64
func (_e *WriteOnlyEth1Data_Expecter[Eth1DataT, ExecutionPayloadHeaderT]) SetEth1DepositIndex(_a0 interface{}) *WriteOnlyEth1Data_SetEth1DepositIndex_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	return &WriteOnlyEth1Data_SetEth1DepositIndex_Call[Eth1DataT, ExecutionPayloadHeaderT]{Call: _e.mock.On("SetEth1DepositIndex", _a0)}
}

func (_c *WriteOnlyEth1Data_SetEth1DepositIndex_Call[Eth1DataT, ExecutionPayloadHeaderT]) Run(run func(_a0 uint64)) *WriteOnlyEth1Data_SetEth1DepositIndex_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *WriteOnlyEth1Data_SetEth1DepositIndex_Call[Eth1DataT, ExecutionPayloadHeaderT]) Return(_a0 error) *WriteOnlyEth1Data_SetEth1DepositIndex_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WriteOnlyEth1Data_SetEth1DepositIndex_Call[Eth1DataT, ExecutionPayloadHeaderT]) RunAndReturn(run func(uint64) error) *WriteOnlyEth1Data_SetEth1DepositIndex_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	_c.Call.Return(run)
	return _c
}

// SetLatestExecutionPayloadHeader provides a mock function with given fields: _a0
func (_m *WriteOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT]) SetLatestExecutionPayloadHeader(_a0 ExecutionPayloadHeaderT) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for SetLatestExecutionPayloadHeader")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(ExecutionPayloadHeaderT) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteOnlyEth1Data_SetLatestExecutionPayloadHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetLatestExecutionPayloadHeader'
type WriteOnlyEth1Data_SetLatestExecutionPayloadHeader_Call[Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}] struct {
	*mock.Call
}

// SetLatestExecutionPayloadHeader is a helper method to define mock.On call
//   - _a0 ExecutionPayloadHeaderT
func (_e *WriteOnlyEth1Data_Expecter[Eth1DataT, ExecutionPayloadHeaderT]) SetLatestExecutionPayloadHeader(_a0 interface{}) *WriteOnlyEth1Data_SetLatestExecutionPayloadHeader_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	return &WriteOnlyEth1Data_SetLatestExecutionPayloadHeader_Call[Eth1DataT, ExecutionPayloadHeaderT]{Call: _e.mock.On("SetLatestExecutionPayloadHeader", _a0)}
}

func (_c *WriteOnlyEth1Data_SetLatestExecutionPayloadHeader_Call[Eth1DataT, ExecutionPayloadHeaderT]) Run(run func(_a0 ExecutionPayloadHeaderT)) *WriteOnlyEth1Data_SetLatestExecutionPayloadHeader_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(ExecutionPayloadHeaderT))
	})
	return _c
}

func (_c *WriteOnlyEth1Data_SetLatestExecutionPayloadHeader_Call[Eth1DataT, ExecutionPayloadHeaderT]) Return(_a0 error) *WriteOnlyEth1Data_SetLatestExecutionPayloadHeader_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WriteOnlyEth1Data_SetLatestExecutionPayloadHeader_Call[Eth1DataT, ExecutionPayloadHeaderT]) RunAndReturn(run func(ExecutionPayloadHeaderT) error) *WriteOnlyEth1Data_SetLatestExecutionPayloadHeader_Call[Eth1DataT, ExecutionPayloadHeaderT] {
	_c.Call.Return(run)
	return _c
}

// NewWriteOnlyEth1Data creates a new instance of WriteOnlyEth1Data. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWriteOnlyEth1Data[Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *WriteOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT] {
	mock := &WriteOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

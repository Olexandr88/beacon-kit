// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	bytes "github.com/berachain/beacon-kit/mod/primitives/pkg/bytes"
	mock "github.com/stretchr/testify/mock"
)

// WriteOnlyRandaoMixes is an autogenerated mock type for the WriteOnlyRandaoMixes type
type WriteOnlyRandaoMixes struct {
	mock.Mock
}

type WriteOnlyRandaoMixes_Expecter struct {
	mock *mock.Mock
}

func (_m *WriteOnlyRandaoMixes) EXPECT() *WriteOnlyRandaoMixes_Expecter {
	return &WriteOnlyRandaoMixes_Expecter{mock: &_m.Mock}
}

// UpdateRandaoMixAtIndex provides a mock function with given fields: _a0, _a1
func (_m *WriteOnlyRandaoMixes) UpdateRandaoMixAtIndex(_a0 uint64, _a1 bytes.B32) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRandaoMixAtIndex")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, bytes.B32) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteOnlyRandaoMixes_UpdateRandaoMixAtIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateRandaoMixAtIndex'
type WriteOnlyRandaoMixes_UpdateRandaoMixAtIndex_Call struct {
	*mock.Call
}

// UpdateRandaoMixAtIndex is a helper method to define mock.On call
//   - _a0 uint64
//   - _a1 bytes.B32
func (_e *WriteOnlyRandaoMixes_Expecter) UpdateRandaoMixAtIndex(_a0 interface{}, _a1 interface{}) *WriteOnlyRandaoMixes_UpdateRandaoMixAtIndex_Call {
	return &WriteOnlyRandaoMixes_UpdateRandaoMixAtIndex_Call{Call: _e.mock.On("UpdateRandaoMixAtIndex", _a0, _a1)}
}

func (_c *WriteOnlyRandaoMixes_UpdateRandaoMixAtIndex_Call) Run(run func(_a0 uint64, _a1 bytes.B32)) *WriteOnlyRandaoMixes_UpdateRandaoMixAtIndex_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64), args[1].(bytes.B32))
	})
	return _c
}

func (_c *WriteOnlyRandaoMixes_UpdateRandaoMixAtIndex_Call) Return(_a0 error) *WriteOnlyRandaoMixes_UpdateRandaoMixAtIndex_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WriteOnlyRandaoMixes_UpdateRandaoMixAtIndex_Call) RunAndReturn(run func(uint64, bytes.B32) error) *WriteOnlyRandaoMixes_UpdateRandaoMixAtIndex_Call {
	_c.Call.Return(run)
	return _c
}

// NewWriteOnlyRandaoMixes creates a new instance of WriteOnlyRandaoMixes. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWriteOnlyRandaoMixes(t interface {
	mock.TestingT
	Cleanup(func())
}) *WriteOnlyRandaoMixes {
	mock := &WriteOnlyRandaoMixes{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

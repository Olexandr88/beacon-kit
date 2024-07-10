// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// BlobSidecars is an autogenerated mock type for the BlobSidecars type
type BlobSidecars[T interface{}, BlobSidecarT interface{}] struct {
	mock.Mock
}

type BlobSidecars_Expecter[T interface{}, BlobSidecarT interface{}] struct {
	mock *mock.Mock
}

func (_m *BlobSidecars[T, BlobSidecarT]) EXPECT() *BlobSidecars_Expecter[T, BlobSidecarT] {
	return &BlobSidecars_Expecter[T, BlobSidecarT]{mock: &_m.Mock}
}

// Get provides a mock function with given fields: idx
func (_m *BlobSidecars[T, BlobSidecarT]) Get(idx uint32) (BlobSidecarT, error) {
	ret := _m.Called(idx)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 BlobSidecarT
	var r1 error
	if rf, ok := ret.Get(0).(func(uint32) (BlobSidecarT, error)); ok {
		return rf(idx)
	}
	if rf, ok := ret.Get(0).(func(uint32) BlobSidecarT); ok {
		r0 = rf(idx)
	} else {
		r0 = ret.Get(0).(BlobSidecarT)
	}

	if rf, ok := ret.Get(1).(func(uint32) error); ok {
		r1 = rf(idx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlobSidecars_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type BlobSidecars_Get_Call[T interface{}, BlobSidecarT interface{}] struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - idx uint32
func (_e *BlobSidecars_Expecter[T, BlobSidecarT]) Get(idx interface{}) *BlobSidecars_Get_Call[T, BlobSidecarT] {
	return &BlobSidecars_Get_Call[T, BlobSidecarT]{Call: _e.mock.On("Get", idx)}
}

func (_c *BlobSidecars_Get_Call[T, BlobSidecarT]) Run(run func(idx uint32)) *BlobSidecars_Get_Call[T, BlobSidecarT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint32))
	})
	return _c
}

func (_c *BlobSidecars_Get_Call[T, BlobSidecarT]) Return(_a0 BlobSidecarT, _a1 error) *BlobSidecars_Get_Call[T, BlobSidecarT] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BlobSidecars_Get_Call[T, BlobSidecarT]) RunAndReturn(run func(uint32) (BlobSidecarT, error)) *BlobSidecars_Get_Call[T, BlobSidecarT] {
	_c.Call.Return(run)
	return _c
}

// GetSidecars provides a mock function with given fields:
func (_m *BlobSidecars[T, BlobSidecarT]) GetSidecars() []BlobSidecarT {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSidecars")
	}

	var r0 []BlobSidecarT
	if rf, ok := ret.Get(0).(func() []BlobSidecarT); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]BlobSidecarT)
		}
	}

	return r0
}

// BlobSidecars_GetSidecars_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSidecars'
type BlobSidecars_GetSidecars_Call[T interface{}, BlobSidecarT interface{}] struct {
	*mock.Call
}

// GetSidecars is a helper method to define mock.On call
func (_e *BlobSidecars_Expecter[T, BlobSidecarT]) GetSidecars() *BlobSidecars_GetSidecars_Call[T, BlobSidecarT] {
	return &BlobSidecars_GetSidecars_Call[T, BlobSidecarT]{Call: _e.mock.On("GetSidecars")}
}

func (_c *BlobSidecars_GetSidecars_Call[T, BlobSidecarT]) Run(run func()) *BlobSidecars_GetSidecars_Call[T, BlobSidecarT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BlobSidecars_GetSidecars_Call[T, BlobSidecarT]) Return(_a0 []BlobSidecarT) *BlobSidecars_GetSidecars_Call[T, BlobSidecarT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlobSidecars_GetSidecars_Call[T, BlobSidecarT]) RunAndReturn(run func() []BlobSidecarT) *BlobSidecars_GetSidecars_Call[T, BlobSidecarT] {
	_c.Call.Return(run)
	return _c
}

// IsNil provides a mock function with given fields:
func (_m *BlobSidecars[T, BlobSidecarT]) IsNil() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsNil")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// BlobSidecars_IsNil_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsNil'
type BlobSidecars_IsNil_Call[T interface{}, BlobSidecarT interface{}] struct {
	*mock.Call
}

// IsNil is a helper method to define mock.On call
func (_e *BlobSidecars_Expecter[T, BlobSidecarT]) IsNil() *BlobSidecars_IsNil_Call[T, BlobSidecarT] {
	return &BlobSidecars_IsNil_Call[T, BlobSidecarT]{Call: _e.mock.On("IsNil")}
}

func (_c *BlobSidecars_IsNil_Call[T, BlobSidecarT]) Run(run func()) *BlobSidecars_IsNil_Call[T, BlobSidecarT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BlobSidecars_IsNil_Call[T, BlobSidecarT]) Return(_a0 bool) *BlobSidecars_IsNil_Call[T, BlobSidecarT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlobSidecars_IsNil_Call[T, BlobSidecarT]) RunAndReturn(run func() bool) *BlobSidecars_IsNil_Call[T, BlobSidecarT] {
	_c.Call.Return(run)
	return _c
}

// Len provides a mock function with given fields:
func (_m *BlobSidecars[T, BlobSidecarT]) Len() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Len")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// BlobSidecars_Len_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Len'
type BlobSidecars_Len_Call[T interface{}, BlobSidecarT interface{}] struct {
	*mock.Call
}

// Len is a helper method to define mock.On call
func (_e *BlobSidecars_Expecter[T, BlobSidecarT]) Len() *BlobSidecars_Len_Call[T, BlobSidecarT] {
	return &BlobSidecars_Len_Call[T, BlobSidecarT]{Call: _e.mock.On("Len")}
}

func (_c *BlobSidecars_Len_Call[T, BlobSidecarT]) Run(run func()) *BlobSidecars_Len_Call[T, BlobSidecarT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BlobSidecars_Len_Call[T, BlobSidecarT]) Return(_a0 int) *BlobSidecars_Len_Call[T, BlobSidecarT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlobSidecars_Len_Call[T, BlobSidecarT]) RunAndReturn(run func() int) *BlobSidecars_Len_Call[T, BlobSidecarT] {
	_c.Call.Return(run)
	return _c
}

// NewFromSidecars provides a mock function with given fields: sidecars
func (_m *BlobSidecars[T, BlobSidecarT]) NewFromSidecars(sidecars []BlobSidecarT) T {
	ret := _m.Called(sidecars)

	if len(ret) == 0 {
		panic("no return value specified for NewFromSidecars")
	}

	var r0 T
	if rf, ok := ret.Get(0).(func([]BlobSidecarT) T); ok {
		r0 = rf(sidecars)
	} else {
		r0 = ret.Get(0).(T)
	}

	return r0
}

// BlobSidecars_NewFromSidecars_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewFromSidecars'
type BlobSidecars_NewFromSidecars_Call[T interface{}, BlobSidecarT interface{}] struct {
	*mock.Call
}

// NewFromSidecars is a helper method to define mock.On call
//   - sidecars []BlobSidecarT
func (_e *BlobSidecars_Expecter[T, BlobSidecarT]) NewFromSidecars(sidecars interface{}) *BlobSidecars_NewFromSidecars_Call[T, BlobSidecarT] {
	return &BlobSidecars_NewFromSidecars_Call[T, BlobSidecarT]{Call: _e.mock.On("NewFromSidecars", sidecars)}
}

func (_c *BlobSidecars_NewFromSidecars_Call[T, BlobSidecarT]) Run(run func(sidecars []BlobSidecarT)) *BlobSidecars_NewFromSidecars_Call[T, BlobSidecarT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]BlobSidecarT))
	})
	return _c
}

func (_c *BlobSidecars_NewFromSidecars_Call[T, BlobSidecarT]) Return(_a0 T) *BlobSidecars_NewFromSidecars_Call[T, BlobSidecarT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlobSidecars_NewFromSidecars_Call[T, BlobSidecarT]) RunAndReturn(run func([]BlobSidecarT) T) *BlobSidecars_NewFromSidecars_Call[T, BlobSidecarT] {
	_c.Call.Return(run)
	return _c
}

// ValidateBlockRoots provides a mock function with given fields:
func (_m *BlobSidecars[T, BlobSidecarT]) ValidateBlockRoots() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ValidateBlockRoots")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BlobSidecars_ValidateBlockRoots_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateBlockRoots'
type BlobSidecars_ValidateBlockRoots_Call[T interface{}, BlobSidecarT interface{}] struct {
	*mock.Call
}

// ValidateBlockRoots is a helper method to define mock.On call
func (_e *BlobSidecars_Expecter[T, BlobSidecarT]) ValidateBlockRoots() *BlobSidecars_ValidateBlockRoots_Call[T, BlobSidecarT] {
	return &BlobSidecars_ValidateBlockRoots_Call[T, BlobSidecarT]{Call: _e.mock.On("ValidateBlockRoots")}
}

func (_c *BlobSidecars_ValidateBlockRoots_Call[T, BlobSidecarT]) Run(run func()) *BlobSidecars_ValidateBlockRoots_Call[T, BlobSidecarT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BlobSidecars_ValidateBlockRoots_Call[T, BlobSidecarT]) Return(_a0 error) *BlobSidecars_ValidateBlockRoots_Call[T, BlobSidecarT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlobSidecars_ValidateBlockRoots_Call[T, BlobSidecarT]) RunAndReturn(run func() error) *BlobSidecars_ValidateBlockRoots_Call[T, BlobSidecarT] {
	_c.Call.Return(run)
	return _c
}

// VerifyInclusionProofs provides a mock function with given fields: kzgOffset
func (_m *BlobSidecars[T, BlobSidecarT]) VerifyInclusionProofs(kzgOffset uint64) error {
	ret := _m.Called(kzgOffset)

	if len(ret) == 0 {
		panic("no return value specified for VerifyInclusionProofs")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(kzgOffset)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BlobSidecars_VerifyInclusionProofs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifyInclusionProofs'
type BlobSidecars_VerifyInclusionProofs_Call[T interface{}, BlobSidecarT interface{}] struct {
	*mock.Call
}

// VerifyInclusionProofs is a helper method to define mock.On call
//   - kzgOffset uint64
func (_e *BlobSidecars_Expecter[T, BlobSidecarT]) VerifyInclusionProofs(kzgOffset interface{}) *BlobSidecars_VerifyInclusionProofs_Call[T, BlobSidecarT] {
	return &BlobSidecars_VerifyInclusionProofs_Call[T, BlobSidecarT]{Call: _e.mock.On("VerifyInclusionProofs", kzgOffset)}
}

func (_c *BlobSidecars_VerifyInclusionProofs_Call[T, BlobSidecarT]) Run(run func(kzgOffset uint64)) *BlobSidecars_VerifyInclusionProofs_Call[T, BlobSidecarT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *BlobSidecars_VerifyInclusionProofs_Call[T, BlobSidecarT]) Return(_a0 error) *BlobSidecars_VerifyInclusionProofs_Call[T, BlobSidecarT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlobSidecars_VerifyInclusionProofs_Call[T, BlobSidecarT]) RunAndReturn(run func(uint64) error) *BlobSidecars_VerifyInclusionProofs_Call[T, BlobSidecarT] {
	_c.Call.Return(run)
	return _c
}

// NewBlobSidecars creates a new instance of BlobSidecars. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlobSidecars[T interface{}, BlobSidecarT interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *BlobSidecars[T, BlobSidecarT] {
	mock := &BlobSidecars[T, BlobSidecarT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	math "github.com/berachain/beacon-kit/mod/primitives/pkg/math"

	mock "github.com/stretchr/testify/mock"
)

// Contract is an autogenerated mock type for the Contract type
type Contract[DepositT interface{}] struct {
	mock.Mock
}

type Contract_Expecter[DepositT interface{}] struct {
	mock *mock.Mock
}

func (_m *Contract[DepositT]) EXPECT() *Contract_Expecter[DepositT] {
	return &Contract_Expecter[DepositT]{mock: &_m.Mock}
}

// ReadDeposits provides a mock function with given fields: ctx, blkNum
func (_m *Contract[DepositT]) ReadDeposits(ctx context.Context, blkNum math.U64) ([]DepositT, error) {
	ret := _m.Called(ctx, blkNum)

	if len(ret) == 0 {
		panic("no return value specified for ReadDeposits")
	}

	var r0 []DepositT
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, math.U64) ([]DepositT, error)); ok {
		return rf(ctx, blkNum)
	}
	if rf, ok := ret.Get(0).(func(context.Context, math.U64) []DepositT); ok {
		r0 = rf(ctx, blkNum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]DepositT)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, math.U64) error); ok {
		r1 = rf(ctx, blkNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Contract_ReadDeposits_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadDeposits'
type Contract_ReadDeposits_Call[DepositT interface{}] struct {
	*mock.Call
}

// ReadDeposits is a helper method to define mock.On call
//   - ctx context.Context
//   - blkNum math.U64
func (_e *Contract_Expecter[DepositT]) ReadDeposits(ctx interface{}, blkNum interface{}) *Contract_ReadDeposits_Call[DepositT] {
	return &Contract_ReadDeposits_Call[DepositT]{Call: _e.mock.On("ReadDeposits", ctx, blkNum)}
}

func (_c *Contract_ReadDeposits_Call[DepositT]) Run(run func(ctx context.Context, blkNum math.U64)) *Contract_ReadDeposits_Call[DepositT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(math.U64))
	})
	return _c
}

func (_c *Contract_ReadDeposits_Call[DepositT]) Return(_a0 []DepositT, _a1 error) *Contract_ReadDeposits_Call[DepositT] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Contract_ReadDeposits_Call[DepositT]) RunAndReturn(run func(context.Context, math.U64) ([]DepositT, error)) *Contract_ReadDeposits_Call[DepositT] {
	_c.Call.Return(run)
	return _c
}

// NewContract creates a new instance of Contract. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewContract[DepositT interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *Contract[DepositT] {
	mock := &Contract[DepositT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

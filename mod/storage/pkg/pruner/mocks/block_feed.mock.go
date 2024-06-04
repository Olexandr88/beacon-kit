// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	pruner "github.com/berachain/beacon-kit/mod/storage/pkg/pruner"
	mock "github.com/stretchr/testify/mock"
)

// BlockFeed is an autogenerated mock type for the BlockFeed type
type BlockFeed[BeaconBlockT pruner.BeaconBlock, BlockEventT pruner.BlockEvent[BeaconBlockT], SubscriptionT pruner.Subscription] struct {
	mock.Mock
}

type BlockFeed_Expecter[BeaconBlockT pruner.BeaconBlock, BlockEventT pruner.BlockEvent[BeaconBlockT], SubscriptionT pruner.Subscription] struct {
	mock *mock.Mock
}

func (_m *BlockFeed[BeaconBlockT, BlockEventT, SubscriptionT]) EXPECT() *BlockFeed_Expecter[BeaconBlockT, BlockEventT, SubscriptionT] {
	return &BlockFeed_Expecter[BeaconBlockT, BlockEventT, SubscriptionT]{mock: &_m.Mock}
}

// Subscribe provides a mock function with given fields: _a0
func (_m *BlockFeed[BeaconBlockT, BlockEventT, SubscriptionT]) Subscribe(_a0 chan<- BlockEventT) SubscriptionT {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Subscribe")
	}

	var r0 SubscriptionT
	if rf, ok := ret.Get(0).(func(chan<- BlockEventT) SubscriptionT); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(SubscriptionT)
	}

	return r0
}

// BlockFeed_Subscribe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Subscribe'
type BlockFeed_Subscribe_Call[BeaconBlockT pruner.BeaconBlock, BlockEventT pruner.BlockEvent[BeaconBlockT], SubscriptionT pruner.Subscription] struct {
	*mock.Call
}

// Subscribe is a helper method to define mock.On call
//   - _a0 chan<- BlockEventT
func (_e *BlockFeed_Expecter[BeaconBlockT, BlockEventT, SubscriptionT]) Subscribe(_a0 interface{}) *BlockFeed_Subscribe_Call[BeaconBlockT, BlockEventT, SubscriptionT] {
	return &BlockFeed_Subscribe_Call[BeaconBlockT, BlockEventT, SubscriptionT]{Call: _e.mock.On("Subscribe", _a0)}
}

func (_c *BlockFeed_Subscribe_Call[BeaconBlockT, BlockEventT, SubscriptionT]) Run(run func(_a0 chan<- BlockEventT)) *BlockFeed_Subscribe_Call[BeaconBlockT, BlockEventT, SubscriptionT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(chan<- BlockEventT))
	})
	return _c
}

func (_c *BlockFeed_Subscribe_Call[BeaconBlockT, BlockEventT, SubscriptionT]) Return(_a0 SubscriptionT) *BlockFeed_Subscribe_Call[BeaconBlockT, BlockEventT, SubscriptionT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlockFeed_Subscribe_Call[BeaconBlockT, BlockEventT, SubscriptionT]) RunAndReturn(run func(chan<- BlockEventT) SubscriptionT) *BlockFeed_Subscribe_Call[BeaconBlockT, BlockEventT, SubscriptionT] {
	_c.Call.Return(run)
	return _c
}

// NewBlockFeed creates a new instance of BlockFeed. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlockFeed[BeaconBlockT pruner.BeaconBlock, BlockEventT pruner.BlockEvent[BeaconBlockT], SubscriptionT pruner.Subscription](t interface {
	mock.TestingT
	Cleanup(func())
}) *BlockFeed[BeaconBlockT, BlockEventT, SubscriptionT] {
	mock := &BlockFeed[BeaconBlockT, BlockEventT, SubscriptionT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	bytes "github.com/berachain/beacon-kit/mod/primitives/pkg/bytes"
	eip4844 "github.com/berachain/beacon-kit/mod/primitives/pkg/eip4844"

	mock "github.com/stretchr/testify/mock"
)

// BlobProofVerifier is an autogenerated mock type for the BlobProofVerifier type
type BlobProofVerifier[BlobProofArgsT interface{}] struct {
	mock.Mock
}

type BlobProofVerifier_Expecter[BlobProofArgsT interface{}] struct {
	mock *mock.Mock
}

func (_m *BlobProofVerifier[BlobProofArgsT]) EXPECT() *BlobProofVerifier_Expecter[BlobProofArgsT] {
	return &BlobProofVerifier_Expecter[BlobProofArgsT]{mock: &_m.Mock}
}

// GetImplementation provides a mock function with given fields:
func (_m *BlobProofVerifier[BlobProofArgsT]) GetImplementation() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetImplementation")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// BlobProofVerifier_GetImplementation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetImplementation'
type BlobProofVerifier_GetImplementation_Call[BlobProofArgsT interface{}] struct {
	*mock.Call
}

// GetImplementation is a helper method to define mock.On call
func (_e *BlobProofVerifier_Expecter[BlobProofArgsT]) GetImplementation() *BlobProofVerifier_GetImplementation_Call[BlobProofArgsT] {
	return &BlobProofVerifier_GetImplementation_Call[BlobProofArgsT]{Call: _e.mock.On("GetImplementation")}
}

func (_c *BlobProofVerifier_GetImplementation_Call[BlobProofArgsT]) Run(run func()) *BlobProofVerifier_GetImplementation_Call[BlobProofArgsT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BlobProofVerifier_GetImplementation_Call[BlobProofArgsT]) Return(_a0 string) *BlobProofVerifier_GetImplementation_Call[BlobProofArgsT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlobProofVerifier_GetImplementation_Call[BlobProofArgsT]) RunAndReturn(run func() string) *BlobProofVerifier_GetImplementation_Call[BlobProofArgsT] {
	_c.Call.Return(run)
	return _c
}

// VerifyBlobProof provides a mock function with given fields: blob, proof, commitment
func (_m *BlobProofVerifier[BlobProofArgsT]) VerifyBlobProof(blob *eip4844.Blob, proof bytes.B48, commitment eip4844.KZGCommitment) error {
	ret := _m.Called(blob, proof, commitment)

	if len(ret) == 0 {
		panic("no return value specified for VerifyBlobProof")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*eip4844.Blob, bytes.B48, eip4844.KZGCommitment) error); ok {
		r0 = rf(blob, proof, commitment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BlobProofVerifier_VerifyBlobProof_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifyBlobProof'
type BlobProofVerifier_VerifyBlobProof_Call[BlobProofArgsT interface{}] struct {
	*mock.Call
}

// VerifyBlobProof is a helper method to define mock.On call
//   - blob *eip4844.Blob
//   - proof bytes.B48
//   - commitment eip4844.KZGCommitment
func (_e *BlobProofVerifier_Expecter[BlobProofArgsT]) VerifyBlobProof(blob interface{}, proof interface{}, commitment interface{}) *BlobProofVerifier_VerifyBlobProof_Call[BlobProofArgsT] {
	return &BlobProofVerifier_VerifyBlobProof_Call[BlobProofArgsT]{Call: _e.mock.On("VerifyBlobProof", blob, proof, commitment)}
}

func (_c *BlobProofVerifier_VerifyBlobProof_Call[BlobProofArgsT]) Run(run func(blob *eip4844.Blob, proof bytes.B48, commitment eip4844.KZGCommitment)) *BlobProofVerifier_VerifyBlobProof_Call[BlobProofArgsT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*eip4844.Blob), args[1].(bytes.B48), args[2].(eip4844.KZGCommitment))
	})
	return _c
}

func (_c *BlobProofVerifier_VerifyBlobProof_Call[BlobProofArgsT]) Return(_a0 error) *BlobProofVerifier_VerifyBlobProof_Call[BlobProofArgsT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlobProofVerifier_VerifyBlobProof_Call[BlobProofArgsT]) RunAndReturn(run func(*eip4844.Blob, bytes.B48, eip4844.KZGCommitment) error) *BlobProofVerifier_VerifyBlobProof_Call[BlobProofArgsT] {
	_c.Call.Return(run)
	return _c
}

// VerifyBlobProofBatch provides a mock function with given fields: _a0
func (_m *BlobProofVerifier[BlobProofArgsT]) VerifyBlobProofBatch(_a0 BlobProofArgsT) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for VerifyBlobProofBatch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(BlobProofArgsT) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BlobProofVerifier_VerifyBlobProofBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifyBlobProofBatch'
type BlobProofVerifier_VerifyBlobProofBatch_Call[BlobProofArgsT interface{}] struct {
	*mock.Call
}

// VerifyBlobProofBatch is a helper method to define mock.On call
//   - _a0 BlobProofArgsT
func (_e *BlobProofVerifier_Expecter[BlobProofArgsT]) VerifyBlobProofBatch(_a0 interface{}) *BlobProofVerifier_VerifyBlobProofBatch_Call[BlobProofArgsT] {
	return &BlobProofVerifier_VerifyBlobProofBatch_Call[BlobProofArgsT]{Call: _e.mock.On("VerifyBlobProofBatch", _a0)}
}

func (_c *BlobProofVerifier_VerifyBlobProofBatch_Call[BlobProofArgsT]) Run(run func(_a0 BlobProofArgsT)) *BlobProofVerifier_VerifyBlobProofBatch_Call[BlobProofArgsT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(BlobProofArgsT))
	})
	return _c
}

func (_c *BlobProofVerifier_VerifyBlobProofBatch_Call[BlobProofArgsT]) Return(_a0 error) *BlobProofVerifier_VerifyBlobProofBatch_Call[BlobProofArgsT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlobProofVerifier_VerifyBlobProofBatch_Call[BlobProofArgsT]) RunAndReturn(run func(BlobProofArgsT) error) *BlobProofVerifier_VerifyBlobProofBatch_Call[BlobProofArgsT] {
	_c.Call.Return(run)
	return _c
}

// NewBlobProofVerifier creates a new instance of BlobProofVerifier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlobProofVerifier[BlobProofArgsT interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *BlobProofVerifier[BlobProofArgsT] {
	mock := &BlobProofVerifier[BlobProofArgsT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

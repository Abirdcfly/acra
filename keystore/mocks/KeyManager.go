// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	context "context"

	kms "github.com/cossacklabs/acra/keystore/kms"
	mock "github.com/stretchr/testify/mock"
)

// KeyManager is an autogenerated mock type for the KeyManager type
type KeyManager struct {
	mock.Mock
}

// CreateKey provides a mock function with given fields: ctx, metaData
func (_m *KeyManager) CreateKey(ctx context.Context, metaData kms.CreateKeyMetadata) (*kms.KeyMetadata, error) {
	ret := _m.Called(ctx, metaData)

	var r0 *kms.KeyMetadata
	if rf, ok := ret.Get(0).(func(context.Context, kms.CreateKeyMetadata) *kms.KeyMetadata); ok {
		r0 = rf(ctx, metaData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kms.KeyMetadata)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, kms.CreateKeyMetadata) error); ok {
		r1 = rf(ctx, metaData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Decrypt provides a mock function with given fields: ctx, keyID, data, _a3
func (_m *KeyManager) Decrypt(ctx context.Context, keyID []byte, data []byte, _a3 []byte) ([]byte, error) {
	ret := _m.Called(ctx, keyID, data, _a3)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, []byte, []byte, []byte) []byte); ok {
		r0 = rf(ctx, keyID, data, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte, []byte, []byte) error); ok {
		r1 = rf(ctx, keyID, data, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Encrypt provides a mock function with given fields: ctx, keyID, data, _a3
func (_m *KeyManager) Encrypt(ctx context.Context, keyID []byte, data []byte, _a3 []byte) ([]byte, error) {
	ret := _m.Called(ctx, keyID, data, _a3)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, []byte, []byte, []byte) []byte); ok {
		r0 = rf(ctx, keyID, data, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte, []byte, []byte) error); ok {
		r1 = rf(ctx, keyID, data, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ID provides a mock function with given fields:
func (_m *KeyManager) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsKeyExist provides a mock function with given fields: ctx, keyID
func (_m *KeyManager) IsKeyExist(ctx context.Context, keyID string) (bool, error) {
	ret := _m.Called(ctx, keyID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, keyID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, keyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewKeyManagerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewKeyManager creates a new instance of KeyManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewKeyManager(t NewKeyManagerT) *KeyManager {
	mock := &KeyManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

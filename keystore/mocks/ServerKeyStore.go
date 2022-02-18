// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	keystore "github.com/cossacklabs/acra/keystore"
	keys "github.com/cossacklabs/themis/gothemis/keys"

	mock "github.com/stretchr/testify/mock"
)

// ServerKeyStore is an autogenerated mock type for the ServerKeyStore type
type ServerKeyStore struct {
	mock.Mock
}

// CacheOnStart provides a mock function with given fields:
func (_m *ServerKeyStore) CacheOnStart() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateClientIDSymmetricKey provides a mock function with given fields: id
func (_m *ServerKeyStore) GenerateClientIDSymmetricKey(id []byte) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateDataEncryptionKeys provides a mock function with given fields: clientID
func (_m *ServerKeyStore) GenerateDataEncryptionKeys(clientID []byte) error {
	ret := _m.Called(clientID)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(clientID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GeneratePoisonRecordSymmetricKey provides a mock function with given fields:
func (_m *ServerKeyStore) GeneratePoisonRecordSymmetricKey() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateZoneIDSymmetricKey provides a mock function with given fields: id
func (_m *ServerKeyStore) GenerateZoneIDSymmetricKey(id []byte) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateZoneKey provides a mock function with given fields:
func (_m *ServerKeyStore) GenerateZoneKey() ([]byte, []byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func() []byte); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetClientIDEncryptionPublicKey provides a mock function with given fields: clientID
func (_m *ServerKeyStore) GetClientIDEncryptionPublicKey(clientID []byte) (*keys.PublicKey, error) {
	ret := _m.Called(clientID)

	var r0 *keys.PublicKey
	if rf, ok := ret.Get(0).(func([]byte) *keys.PublicKey); ok {
		r0 = rf(clientID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keys.PublicKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(clientID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetClientIDSymmetricKey provides a mock function with given fields: id
func (_m *ServerKeyStore) GetClientIDSymmetricKey(id []byte) ([]byte, error) {
	ret := _m.Called(id)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetClientIDSymmetricKeys provides a mock function with given fields: id
func (_m *ServerKeyStore) GetClientIDSymmetricKeys(id []byte) ([][]byte, error) {
	ret := _m.Called(id)

	var r0 [][]byte
	if rf, ok := ret.Get(0).(func([]byte) [][]byte); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHMACSecretKey provides a mock function with given fields: id
func (_m *ServerKeyStore) GetHMACSecretKey(id []byte) ([]byte, error) {
	ret := _m.Called(id)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLogSecretKey provides a mock function with given fields:
func (_m *ServerKeyStore) GetLogSecretKey() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPoisonKeyPair provides a mock function with given fields:
func (_m *ServerKeyStore) GetPoisonKeyPair() (*keys.Keypair, error) {
	ret := _m.Called()

	var r0 *keys.Keypair
	if rf, ok := ret.Get(0).(func() *keys.Keypair); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keys.Keypair)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPoisonPrivateKeys provides a mock function with given fields:
func (_m *ServerKeyStore) GetPoisonPrivateKeys() ([]*keys.PrivateKey, error) {
	ret := _m.Called()

	var r0 []*keys.PrivateKey
	if rf, ok := ret.Get(0).(func() []*keys.PrivateKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*keys.PrivateKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPoisonSymmetricKey provides a mock function with given fields:
func (_m *ServerKeyStore) GetPoisonSymmetricKey() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPoisonSymmetricKeys provides a mock function with given fields:
func (_m *ServerKeyStore) GetPoisonSymmetricKeys() ([][]byte, error) {
	ret := _m.Called()

	var r0 [][]byte
	if rf, ok := ret.Get(0).(func() [][]byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServerDecryptionPrivateKey provides a mock function with given fields: id
func (_m *ServerKeyStore) GetServerDecryptionPrivateKey(id []byte) (*keys.PrivateKey, error) {
	ret := _m.Called(id)

	var r0 *keys.PrivateKey
	if rf, ok := ret.Get(0).(func([]byte) *keys.PrivateKey); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keys.PrivateKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServerDecryptionPrivateKeys provides a mock function with given fields: id
func (_m *ServerKeyStore) GetServerDecryptionPrivateKeys(id []byte) ([]*keys.PrivateKey, error) {
	ret := _m.Called(id)

	var r0 []*keys.PrivateKey
	if rf, ok := ret.Get(0).(func([]byte) []*keys.PrivateKey); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*keys.PrivateKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetZoneIDSymmetricKey provides a mock function with given fields: id
func (_m *ServerKeyStore) GetZoneIDSymmetricKey(id []byte) ([]byte, error) {
	ret := _m.Called(id)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetZoneIDSymmetricKeys provides a mock function with given fields: id
func (_m *ServerKeyStore) GetZoneIDSymmetricKeys(id []byte) ([][]byte, error) {
	ret := _m.Called(id)

	var r0 [][]byte
	if rf, ok := ret.Get(0).(func([]byte) [][]byte); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetZonePrivateKey provides a mock function with given fields: id
func (_m *ServerKeyStore) GetZonePrivateKey(id []byte) (*keys.PrivateKey, error) {
	ret := _m.Called(id)

	var r0 *keys.PrivateKey
	if rf, ok := ret.Get(0).(func([]byte) *keys.PrivateKey); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keys.PrivateKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetZonePrivateKeys provides a mock function with given fields: id
func (_m *ServerKeyStore) GetZonePrivateKeys(id []byte) ([]*keys.PrivateKey, error) {
	ret := _m.Called(id)

	var r0 []*keys.PrivateKey
	if rf, ok := ret.Get(0).(func([]byte) []*keys.PrivateKey); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*keys.PrivateKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetZonePublicKey provides a mock function with given fields: zoneID
func (_m *ServerKeyStore) GetZonePublicKey(zoneID []byte) (*keys.PublicKey, error) {
	ret := _m.Called(zoneID)

	var r0 *keys.PublicKey
	if rf, ok := ret.Get(0).(func([]byte) *keys.PublicKey); ok {
		r0 = rf(zoneID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*keys.PublicKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(zoneID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasZonePrivateKey provides a mock function with given fields: id
func (_m *ServerKeyStore) HasZonePrivateKey(id []byte) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func([]byte) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ListKeys provides a mock function with given fields:
func (_m *ServerKeyStore) ListKeys() ([]keystore.KeyDescription, error) {
	ret := _m.Called()

	var r0 []keystore.KeyDescription
	if rf, ok := ret.Get(0).(func() []keystore.KeyDescription); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]keystore.KeyDescription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reset provides a mock function with given fields:
func (_m *ServerKeyStore) Reset() {
	_m.Called()
}

// RotateSymmetricZoneKey provides a mock function with given fields: zoneID
func (_m *ServerKeyStore) RotateSymmetricZoneKey(zoneID []byte) error {
	ret := _m.Called(zoneID)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(zoneID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RotateZoneKey provides a mock function with given fields: zoneID
func (_m *ServerKeyStore) RotateZoneKey(zoneID []byte) ([]byte, error) {
	ret := _m.Called(zoneID)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(zoneID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(zoneID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveDataEncryptionKeys provides a mock function with given fields: clientID, keypair
func (_m *ServerKeyStore) SaveDataEncryptionKeys(clientID []byte, keypair *keys.Keypair) error {
	ret := _m.Called(clientID, keypair)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, *keys.Keypair) error); ok {
		r0 = rf(clientID, keypair)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveZoneKeypair provides a mock function with given fields: zoneID, keypair
func (_m *ServerKeyStore) SaveZoneKeypair(zoneID []byte, keypair *keys.Keypair) error {
	ret := _m.Called(zoneID, keypair)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, *keys.Keypair) error); ok {
		r0 = rf(zoneID, keypair)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

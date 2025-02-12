// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"

	valsettypes "github.com/palomachain/paloma/x/valset/types"
)

// ValsetKeeper is an autogenerated mock type for the ValsetKeeper type
type ValsetKeeper struct {
	mock.Mock
}

// FindSnapshotByID provides a mock function with given fields: ctx, id
func (_m *ValsetKeeper) FindSnapshotByID(ctx types.Context, id uint64) (*valsettypes.Snapshot, error) {
	ret := _m.Called(ctx, id)

	var r0 *valsettypes.Snapshot
	if rf, ok := ret.Get(0).(func(types.Context, uint64) *valsettypes.Snapshot); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*valsettypes.Snapshot)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrentSnapshot provides a mock function with given fields: ctx
func (_m *ValsetKeeper) GetCurrentSnapshot(ctx types.Context) (*valsettypes.Snapshot, error) {
	ret := _m.Called(ctx)

	var r0 *valsettypes.Snapshot
	if rf, ok := ret.Get(0).(func(types.Context) *valsettypes.Snapshot); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*valsettypes.Snapshot)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsJailed provides a mock function with given fields: ctx, val
func (_m *ValsetKeeper) IsJailed(ctx types.Context, val types.ValAddress) bool {
	ret := _m.Called(ctx, val)

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.Context, types.ValAddress) bool); ok {
		r0 = rf(ctx, val)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Jail provides a mock function with given fields: ctx, valAddr, reason
func (_m *ValsetKeeper) Jail(ctx types.Context, valAddr types.ValAddress, reason string) error {
	ret := _m.Called(ctx, valAddr, reason)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, types.ValAddress, string) error); ok {
		r0 = rf(ctx, valAddr, reason)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// KeepValidatorAlive provides a mock function with given fields: ctx, valAddr
func (_m *ValsetKeeper) KeepValidatorAlive(ctx types.Context, valAddr types.ValAddress) error {
	ret := _m.Called(ctx, valAddr)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, types.ValAddress) error); ok {
		r0 = rf(ctx, valAddr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetValidatorBalance provides a mock function with given fields: ctx, valAddr, chainType, chainReferenceID, externalAddress, balance
func (_m *ValsetKeeper) SetValidatorBalance(ctx types.Context, valAddr types.ValAddress, chainType string, chainReferenceID string, externalAddress string, balance *big.Int) error {
	ret := _m.Called(ctx, valAddr, chainType, chainReferenceID, externalAddress, balance)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, types.ValAddress, string, string, string, *big.Int) error); ok {
		r0 = rf(ctx, valAddr, chainType, chainReferenceID, externalAddress, balance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewValsetKeeper interface {
	mock.TestingT
	Cleanup(func())
}

// NewValsetKeeper creates a new instance of ValsetKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewValsetKeeper(t mockConstructorTestingTNewValsetKeeper) *ValsetKeeper {
	mock := &ValsetKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

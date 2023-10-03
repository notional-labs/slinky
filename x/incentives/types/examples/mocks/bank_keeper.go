// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	math "cosmossdk.io/math"

	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// BankKeeper is an autogenerated mock type for the BankKeeper type
type BankKeeper struct {
	mock.Mock
}

// MintCoins provides a mock function with given fields: ctx, moduleName, amt
func (_m *BankKeeper) MintCoins(ctx context.Context, moduleName string, amt math.Int) error {
	ret := _m.Called(ctx, moduleName, amt)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, math.Int) error); ok {
		r0 = rf(ctx, moduleName, amt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendCoinsFromModuleToAccount provides a mock function with given fields: ctx, senderModule, recipientAddr, amt
func (_m *BankKeeper) SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr types.ValAddress, amt math.Int) error {
	ret := _m.Called(ctx, senderModule, recipientAddr, amt)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ValAddress, math.Int) error); ok {
		r0 = rf(ctx, senderModule, recipientAddr, amt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBankKeeper creates a new instance of BankKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBankKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *BankKeeper {
	mock := &BankKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

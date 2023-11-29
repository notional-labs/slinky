// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	big "math/big"

	mock "github.com/stretchr/testify/mock"

	oracletypes "github.com/skip-mev/slinky/x/oracle/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// CurrencyPairStrategy is an autogenerated mock type for the CurrencyPairStrategy type
type CurrencyPairStrategy struct {
	mock.Mock
}

// FromID provides a mock function with given fields: ctx, id
func (_m *CurrencyPairStrategy) FromID(ctx types.Context, id uint64) (oracletypes.CurrencyPair, error) {
	ret := _m.Called(ctx, id)

	var r0 oracletypes.CurrencyPair
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, uint64) (oracletypes.CurrencyPair, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint64) oracletypes.CurrencyPair); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(oracletypes.CurrencyPair)
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDecodedPrice provides a mock function with given fields: ctx, cp, priceBytes
func (_m *CurrencyPairStrategy) GetDecodedPrice(ctx types.Context, cp oracletypes.CurrencyPair, priceBytes []byte) (*big.Int, error) {
	ret := _m.Called(ctx, cp, priceBytes)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, oracletypes.CurrencyPair, []byte) (*big.Int, error)); ok {
		return rf(ctx, cp, priceBytes)
	}
	if rf, ok := ret.Get(0).(func(types.Context, oracletypes.CurrencyPair, []byte) *big.Int); ok {
		r0 = rf(ctx, cp, priceBytes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, oracletypes.CurrencyPair, []byte) error); ok {
		r1 = rf(ctx, cp, priceBytes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEncodedPrice provides a mock function with given fields: ctx, cp, price
func (_m *CurrencyPairStrategy) GetEncodedPrice(ctx types.Context, cp oracletypes.CurrencyPair, price *big.Int) ([]byte, error) {
	ret := _m.Called(ctx, cp, price)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, oracletypes.CurrencyPair, *big.Int) ([]byte, error)); ok {
		return rf(ctx, cp, price)
	}
	if rf, ok := ret.Get(0).(func(types.Context, oracletypes.CurrencyPair, *big.Int) []byte); ok {
		r0 = rf(ctx, cp, price)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, oracletypes.CurrencyPair, *big.Int) error); ok {
		r1 = rf(ctx, cp, price)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ID provides a mock function with given fields: ctx, cp
func (_m *CurrencyPairStrategy) ID(ctx types.Context, cp oracletypes.CurrencyPair) (uint64, error) {
	ret := _m.Called(ctx, cp)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, oracletypes.CurrencyPair) (uint64, error)); ok {
		return rf(ctx, cp)
	}
	if rf, ok := ret.Get(0).(func(types.Context, oracletypes.CurrencyPair) uint64); ok {
		r0 = rf(ctx, cp)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(types.Context, oracletypes.CurrencyPair) error); ok {
		r1 = rf(ctx, cp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCurrencyPairStrategy creates a new instance of CurrencyPairStrategy. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCurrencyPairStrategy(t interface {
	mock.TestingT
	Cleanup(func())
}) *CurrencyPairStrategy {
	mock := &CurrencyPairStrategy{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

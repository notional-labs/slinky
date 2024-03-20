package types_test

import (
	"testing"

	slinkytypes "github.com/skip-mev/slinky/pkg/types"

	"github.com/stretchr/testify/require"

	"github.com/skip-mev/slinky/x/marketmap/types"
)

var (
	btcusdt = types.Market{
		Ticker: types.Ticker{
			CurrencyPair: slinkytypes.CurrencyPair{
				Base:  "BITCOIN",
				Quote: "USDT",
			},
			Decimals:         8,
			MinProviderCount: 1,
		},
		Paths: types.Paths{
			Paths: []types.Path{
				{
					Operations: []types.Operation{
						{
							CurrencyPair: slinkytypes.CurrencyPair{
								Base:  "BITCOIN",
								Quote: "USDT",
							},
						},
					},
				},
			},
		},
		Providers: types.Providers{
			Providers: []types.ProviderConfig{
				{
					Name:           "kucoin",
					OffChainTicker: "btc-usdt",
				},
			},
		},
	}

	usdtusd = types.Market{
		Ticker: types.Ticker{
			CurrencyPair: slinkytypes.CurrencyPair{
				Base:  "USDT",
				Quote: "USD",
			},
			Decimals:         8,
			MinProviderCount: 1,
		},
		Paths: types.Paths{
			Paths: []types.Path{
				{
					Operations: []types.Operation{
						{
							CurrencyPair: slinkytypes.CurrencyPair{
								Base:  "USDT",
								Quote: "USD",
							},
						},
					},
				},
			},
		},
		Providers: types.Providers{
			Providers: []types.ProviderConfig{
				{
					Name:           "kucoin",
					OffChainTicker: "usdt-usd",
				},
			},
		},
	}

	usdcusd = types.Market{
		Ticker: types.Ticker{
			CurrencyPair: slinkytypes.CurrencyPair{
				Base:  "USDC",
				Quote: "USD",
			},
			Decimals:         8,
			MinProviderCount: 1,
		},
		Paths: types.Paths{
			Paths: []types.Path{
				{
					Operations: []types.Operation{
						{
							CurrencyPair: slinkytypes.CurrencyPair{
								Base:  "USDC",
								Quote: "USD",
							},
						},
					},
				},
			},
		},
		Providers: types.Providers{
			Providers: []types.ProviderConfig{
				{
					Name:           "kucoin",
					OffChainTicker: "usdc-usd",
				},
			},
		},
	}

	ethusdt = types.Market{
		Ticker: types.Ticker{
			CurrencyPair: slinkytypes.CurrencyPair{
				Base:  "ETHEREUM",
				Quote: "USDT",
			},
			Decimals:         8,
			MinProviderCount: 1,
		},
		Paths: types.Paths{
			Paths: []types.Path{
				{
					Operations: []types.Operation{
						{
							CurrencyPair: slinkytypes.CurrencyPair{
								Base:  "ETHEREUM",
								Quote: "USDT",
							},
						},
					},
				},
			},
		},
		Providers: types.Providers{
			Providers: []types.ProviderConfig{
				{
					Name:           "kucoin",
					OffChainTicker: "eth-usdt",
				},
			},
		},
	}

	markets = map[string]types.Market{
		btcusdt.Ticker.String(): btcusdt,
		usdtusd.Ticker.String(): usdtusd,
		usdcusd.Ticker.String(): usdcusd,
		ethusdt.Ticker.String(): ethusdt,
	}
)

func TestPaths(t *testing.T) {
	testCases := []struct {
		name         string
		paths        types.Paths
		currencyPair slinkytypes.CurrencyPair
		expErr       bool
	}{
		{
			name:         "valid",
			paths:        btcusdt.Paths,
			currencyPair: btcusdt.Ticker.CurrencyPair,
			expErr:       false,
		},
		{
			name:         "invalid",
			paths:        types.Paths{},
			currencyPair: btcusdt.Ticker.CurrencyPair,
			expErr:       true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.paths.ValidateBasic(tc.currencyPair)
			if tc.expErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestPathsEqual(t *testing.T) {
	cases := []struct {
		name  string
		paths types.Paths
		other types.Paths
		exp   bool
	}{
		{
			name:  "equal paths",
			paths: btcusdt.Paths,
			other: btcusdt.Paths,
			exp:   true,
		},
		{
			name:  "different length",
			paths: btcusdt.Paths,
			other: types.Paths{},
			exp:   false,
		},
		{
			name:  "different paths",
			paths: btcusdt.Paths,
			other: ethusdt.Paths,
			exp:   false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.paths.Equal(tc.other), tc.exp)
		})
	}
}

func TestPath(t *testing.T) {
	testCases := []struct {
		name   string
		path   types.Path
		target string
		expErr bool
	}{
		{
			name:   "empty path",
			path:   types.Path{},
			target: "",
			expErr: true,
		},
		{
			name: "invalid path with a single operation",
			path: types.Path{
				Operations: []types.Operation{
					{
						CurrencyPair: slinkytypes.CurrencyPair{},
					},
				},
			},
			target: "",
			expErr: true,
		},
		{
			name: "invalid path with multiple operations with a bad ticker in the route",
			path: types.Path{
				Operations: []types.Operation{
					{
						CurrencyPair: btcusdt.Ticker.CurrencyPair,
					},
					{
						CurrencyPair: slinkytypes.CurrencyPair{},
					},
				},
			},
			target: "",
			expErr: true,
		},
		{
			name: "invalid path with multiple operations and mismatching tickers",
			path: types.Path{
				Operations: []types.Operation{
					{
						CurrencyPair: btcusdt.Ticker.CurrencyPair,
					},
					{
						CurrencyPair: ethusdt.Ticker.CurrencyPair,
					},
				},
			},
			target: "",
			expErr: true,
		},
		{
			name: "invalid path with multiple operations and cyclic graph",
			path: types.Path{
				Operations: []types.Operation{
					{
						CurrencyPair: btcusdt.Ticker.CurrencyPair,
					},
					{
						CurrencyPair: usdtusd.Ticker.CurrencyPair,
					},
					{
						CurrencyPair: usdtusd.Ticker.CurrencyPair,
						Invert:       true,
					},
				},
			},
			target: "",
			expErr: true,
		},
		{
			name: "valid path with a single operation",
			path: types.Path{
				Operations: []types.Operation{
					{
						CurrencyPair: btcusdt.Ticker.CurrencyPair,
					},
				},
			},
			target: "BITCOIN/USDT",
			expErr: false,
		},
		{
			name: "valid path with multiple operations",
			path: types.Path{
				Operations: []types.Operation{
					{
						CurrencyPair: btcusdt.Ticker.CurrencyPair,
					},
					{
						CurrencyPair: usdtusd.Ticker.CurrencyPair,
					},
				},
			},
			target: "BITCOIN/USD",
			expErr: false,
		},
		{
			name: "valid path with multiple operations and inverted tickers",
			path: types.Path{
				Operations: []types.Operation{
					{
						CurrencyPair: usdtusd.Ticker.CurrencyPair,
						Invert:       true,
					},
					{
						CurrencyPair: btcusdt.Ticker.CurrencyPair,
						Invert:       true,
					},
				},
			},
			target: "USD/BITCOIN",
			expErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.path.ValidateBasic()
			if tc.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.True(t, tc.path.Match(tc.target))
			}
		})
	}
}

func TestPathEqual(t *testing.T) {
	cases := []struct {
		name  string
		path  types.Path
		other types.Path
		exp   bool
	}{
		{
			name: "equal paths",
			path: types.Path{
				Operations: []types.Operation{
					{
						CurrencyPair: btcusdt.Ticker.CurrencyPair,
					},
					{
						CurrencyPair: usdtusd.Ticker.CurrencyPair,
					},
				},
			},
			other: types.Path{
				Operations: []types.Operation{
					{
						CurrencyPair: btcusdt.Ticker.CurrencyPair,
					},
					{
						CurrencyPair: usdtusd.Ticker.CurrencyPair,
					},
				},
			},
			exp: true,
		},
		{
			name: "different length",
			path: types.Path{
				Operations: []types.Operation{
					{
						CurrencyPair: btcusdt.Ticker.CurrencyPair,
					},
				},
			},
			other: types.Path{
				Operations: []types.Operation{
					{
						CurrencyPair: btcusdt.Ticker.CurrencyPair,
					},
					{
						CurrencyPair: usdtusd.Ticker.CurrencyPair,
					},
				},
			},
			exp: false,
		},
		{
			name: "different operations",
			path: types.Path{
				Operations: []types.Operation{
					{
						CurrencyPair: btcusdt.Ticker.CurrencyPair,
					},
				},
			},
			other: types.Path{
				Operations: []types.Operation{
					{
						CurrencyPair: usdtusd.Ticker.CurrencyPair,
					},
				},
			},
			exp: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.path.Equal(tc.other), tc.exp)
		})
	}
}

func TestOperation(t *testing.T) {
	t.Run("valid operation", func(t *testing.T) {
		cp := slinkytypes.CurrencyPair{
			Base:  "BITCOIN",
			Quote: "USDT",
		}

		_, err := types.NewOperation(cp, false)
		require.NoError(t, err)
	})

	t.Run("invalid operation", func(t *testing.T) {
		cp := slinkytypes.CurrencyPair{}
		_, err := types.NewOperation(cp, false)
		require.Error(t, err)
	})
}

func TestOperationEqual(t *testing.T) {
	cases := []struct {
		name  string
		op    types.Operation
		other types.Operation
		exp   bool
	}{
		{
			name: "equal operations",
			op: types.Operation{
				CurrencyPair: slinkytypes.CurrencyPair{
					Base:  "BITCOIN",
					Quote: "USDT",
				},
				Provider: "kucoin",
				Invert:   false,
			},
			other: types.Operation{
				CurrencyPair: slinkytypes.CurrencyPair{
					Base:  "BITCOIN",
					Quote: "USDT",
				},
				Provider: "kucoin",
				Invert:   false,
			},
			exp: true,
		},
		{
			name: "different base",
			op: types.Operation{
				CurrencyPair: slinkytypes.CurrencyPair{
					Base:  "BITCOIN",
					Quote: "USDT",
				},
				Provider: "kucoin",
				Invert:   false,
			},
			other: types.Operation{
				CurrencyPair: slinkytypes.CurrencyPair{
					Base:  "ETHEREUM",
					Quote: "USDT",
				},
				Provider: "kucoin",
				Invert:   false,
			},
			exp: false,
		},
		{
			name: "different invert strategy",
			op: types.Operation{
				CurrencyPair: slinkytypes.CurrencyPair{
					Base:  "BITCOIN",
					Quote: "USDT",
				},
				Provider: "kucoin",
				Invert:   false,
			},
			other: types.Operation{
				CurrencyPair: slinkytypes.CurrencyPair{
					Base:  "BITCOIN",
					Quote: "USDT",
				},
				Provider: "kucoin",
				Invert:   true,
			},
			exp: false,
		},
		{
			name: "different provider",
			op: types.Operation{
				CurrencyPair: slinkytypes.CurrencyPair{
					Base:  "BITCOIN",
					Quote: "USDT",
				},
				Provider: "kucoin",
				Invert:   false,
			},
			other: types.Operation{
				CurrencyPair: slinkytypes.CurrencyPair{
					Base:  "BITCOIN",
					Quote: "USDT",
				},
				Provider: "binance",
				Invert:   false,
			},
			exp: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.op.Equal(tc.other), tc.exp)
		})
	}
}

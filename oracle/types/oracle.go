package types

import (
	"math/big"

	"github.com/skip-mev/slinky/providers/base"
	apihandlers "github.com/skip-mev/slinky/providers/base/api/handlers"
	wshandlers "github.com/skip-mev/slinky/providers/base/websocket/handlers"
	providertypes "github.com/skip-mev/slinky/providers/types"
	"github.com/skip-mev/slinky/providers/types/factory"
	mmtypes "github.com/skip-mev/slinky/x/marketmap/types"
)

type (
	// PriceProviderFactory is a type alias for the price provider factory.
	PriceProviderFactory = factory.ProviderFactory[mmtypes.Ticker, *big.Int]

	// PriceProvider is a type alias for the price provider.
	PriceProvider = providertypes.Provider[mmtypes.Ticker, *big.Int]

	// PriceAPIDataHandler is a type alias for the price API data handler. This
	// is responsible for parsing http responses and returning the resolved
	// and unresolved prices.
	PriceAPIDataHandler = apihandlers.APIDataHandler[mmtypes.Ticker, *big.Int]

	// PriceAPIQueryHandler is a type alias for the price API query handler. This
	// is responsible for building the API query for the price provider and
	// returning the resolved and unresolved prices.
	PriceAPIQueryHandler = apihandlers.APIQueryHandler[mmtypes.Ticker, *big.Int]

	// PriceWebSocketDataHandler is a type alias for the price web socket data handler.
	// This is responsible for parsing web socket messages and returning the resolved
	// and unresolved prices.
	PriceWebSocketDataHandler = wshandlers.WebSocketDataHandler[mmtypes.Ticker, *big.Int]

	// PriceWebSocketQueryHandler is a type alias for the price web socket query handler.
	// This is responsible for building the web socket query for the price provider and
	// returning the resolved and unresolved prices.
	PriceWebSocketQueryHandler = wshandlers.WebSocketQueryHandler[mmtypes.Ticker, *big.Int]

	// PriceResponse is a type alias for the price response. A price response is
	// composed of a map of resolved prices and a map of unresolved prices. Resolved
	// prices are the prices that were successfully fetched from the API, while
	// unresolved prices are the prices that were not successfully fetched from the API.
	PriceResponse = providertypes.GetResponse[mmtypes.Ticker, *big.Int]

	// ResolvedPrices is a type alias for the resolved prices.
	ResolvedPrices = map[mmtypes.Ticker]providertypes.Result[*big.Int]

	// UnResolvedPrices is a type alias for the unresolved prices.
	UnResolvedPrices = map[mmtypes.Ticker]error

	// TickerPrices is a type alias for the map of prices. This is a map of tickers i.e.
	// BTC/USD, ETH/USD, etc. to their respective prices.
	TickerPrices = map[mmtypes.Ticker]*big.Int
)

var (
	// NewPriceResult is a function alias for the new price result.
	NewPriceResult = providertypes.NewResult[*big.Int]

	// NewPricesResponse is a function alias for the new price response.
	NewPriceResponse = providertypes.NewGetResponse[mmtypes.Ticker, *big.Int]

	// NewPriceResponseWithErr is a function alias for the new price response with errors.
	NewPriceResponseWithErr = providertypes.NewGetResponseWithErr[mmtypes.Ticker, *big.Int]

	// NewPriceProvider is a function alias for the new price provider.
	NewPriceProvider = base.NewProvider[mmtypes.Ticker, *big.Int]

	// NewPriceAPIQueryHandler is a function alias for the new API query handler meant to be
	// used by the price providers.
	NewPriceAPIQueryHandler = apihandlers.NewAPIQueryHandler[mmtypes.Ticker, *big.Int]

	// NewPriceWebSocketQueryHandler is a function alias for the new web socket query handler meant to be
	// used by the price providers.
	NewPriceWebSocketQueryHandler = wshandlers.NewWebSocketQueryHandler[mmtypes.Ticker, *big.Int]
)

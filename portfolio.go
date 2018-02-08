package kiteconnect

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/zerodhatech/go-kite-connect/models"
)

// Holding is the holdings response per user.
type Holding struct {
	Tradingsymbol   string `json:"tradingsymbol"`
	Exchange        string `json:"exchange"`
	InstrumentToken uint32 `json:"instrument_token"`
	ISIN            string `json:"isin"`
	Product         string `json:"product"`

	Price              float64 `json:"price"`
	Quantity           int     `json:"quantity"`
	T1Quantity         int     `json:"t1_quantity"`
	RealisedQuantity   int     `json:"realised_quantity"`
	CollateralQuantity int     `json:"collateral_quantity"`
	CollateralType     string  `json:"collateral_type"`

	AveragePrice        float64 `json:"average_price"`
	LastPrice           float64 `json:"last_price"`
	ClosePrice          float64 `json:"close_price"`
	PnL                 float64 `json:"pnl"`
	DayChange           float64 `json:"day_change"`
	DayChangePercentage float64 `json:"day_change_percentage"`
}

// Holdings is a list of Holding response
type Holdings []Holding

// Position represents a single position entry.
type Position struct {
	Tradingsymbol   string `json:"tradingsymbol"`
	Exchange        string `json:"exchange"`
	InstrumentToken uint32 `json:"instrument_token"`
	Product         string `json:"product"`

	Quantity          int     `json:"quantity"`
	OvernightQuantity int     `json:"overnight_quantity"`
	Multiplier        float64 `json:"multiplier"`

	AveragePrice float64 `json:"average_price"`
	ClosePrice   float64 `json:"close_price"`
	LastPrice    float64 `json:"last_price"`
	Value        float64 `json:"value"`
	PnL          float64 `json:"pnl"`
	M2M          float64 `json:"m2m"`
	Unrealised   float64 `json:"unrealised"`
	Realised     float64 `json:"realised"`

	BuyQuantity int     `json:"buy_quantity"`
	BuyPrice    float64 `json:"buy_price"`
	BuyValue    float64 `json:"buy_value"`
	BuyM2MValue float64 `json:"buy_m2m"`

	SellQuantity int     `json:"sell_quantity"`
	SellPrice    float64 `json:"sell_price"`
	SellValue    float64 `json:"sell_value"`
	SellM2MValue float64 `json:"sell_m2m"`

	DayBuyQuantity int     `json:"day_buy_quantity"`
	DayBuyPrice    float64 `json:"day_buy_price"`
	DayBuyValue    float64 `json:"day_buy_value"`

	DaySellQuantity int     `json:"day_sell_quantity"`
	DaySellPrice    float64 `json:"day_sell_price"`
	DaySellValue    float64 `json:"day_sell_value"`
}

// Positions defines the response structure for positions calls.
type Positions struct {
	Net []Position `json:"net"`
	Day []Position `json:"day"`
}

// PositionModParams represents the input params of a Position Conversion request.
type ConvertPositionParams struct {
	Exchange        string `url:"exchange"`
	TradingSymbol   string `url:"tradingsymbol"`
	OldProduct      string `url:"old_product"`
	NewProduct      string `url:"new_product"`
	PositionType    string `url:"position_type"`
	TransactionType string `url:"transaction_type"`
	Quantity        string `url:"quantity"`
}

// GetHoldings gets list of user holdings.
func (c *Client) GetHoldings() (Holdings, error) {
	var holdings Holdings
	err := c.doEnvelope(http.MethodGet, URIGetHoldings, nil, nil, &holdings)
	return holdings, err
}

// GetPositions gets user positions.
func (c *Client) GetPositions() (Positions, error) {
	var positions Positions
	err := c.doEnvelope(http.MethodGet, URIGetPositions, nil, nil, &positions)
	return positions, err
}

// ConvertPosition converts postion's product type
func (c *Client) ConvertPosition(positionParams ConvertPositionParams) (bool, error) {
	var (
		b      bool
		err    error
		params url.Values
		out    models.PlainResponse
	)

	if params, err = query.Values(positionParams); err != nil {
		return false, NewError(InputError, fmt.Sprintf("Error decoding order params: %v", err), nil)
	}

	if err = c.doEnvelope(http.MethodPut, URIConvertPosition, params, nil, &out); err == nil {
		b = true
	}

	return b, err
}
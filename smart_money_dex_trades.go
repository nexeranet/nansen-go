package nansen

import (
	"context"
	"net/http"
)

type GetSmartMoneyDEXTradesBody struct {
	Chains     []string        `json:"chains"`
	Filters    map[string]any  `json:"filters,omitempty"`
	Pagination *PaginationBody `json:"pagination,omitempty"`
	OrderBy    []SortOrderBody `json:"order_by,omitempty"`
}

type GetSmartMoneyDEXTradesData struct {
	Chain                string  `json:"chain"`
	BlockTimestamp       string  `json:"block_timestamp"`
	TransactionHash      string  `json:"transaction_hash"`
	TraderAddress        string  `json:"trader_address"`
	TraderAddressLabel   string  `json:"trader_address_label"`
	TokenBoughtAddress   string  `json:"token_bought_address"`
	TokenSoldAddress     string  `json:"token_sold_address"`
	TokenBoughtAmount    float64 `json:"token_bought_amount"`
	TokenSoldAmount      float64 `json:"token_sold_amount"`
	TokenBoughtSymbol    string  `json:"token_bought_symbol"`
	TokenSoldSymbol      string  `json:"token_sold_symbol"`
	TokenBoughtAgeDays   float64 `json:"token_bought_age_days"`
	TokenSoldAgeDays     float64 `json:"token_sold_age_days"`
	TokenBoughtMarketCap float64 `json:"token_bought_market_cap"`
	TokenSoldMarketCap   float64 `json:"token_sold_market_cap"`
	TradeValueUsd        float64 `json:"trade_value_usd"`
}

type GetSmartMoneyDEXTradesResponse struct {
	Data       []GetSmartMoneyDEXTradesData `json:"data"`
	Pagination *PaginationResponse          `json:"pagination,omitempty"`
}

func (c *Client) GetSmartMoneyDEXTrades(ctx context.Context, body GetSmartMoneyDEXTradesBody) (result *GetSmartMoneyDEXTradesResponse, err error) {
	var response GetSmartMoneyDEXTradesResponse
	request := NewRequest(c.Url("/smart-money/dex-trades"), body, http.MethodPost)
	_, err = c.doCall(ctx, request, &response)
	if err != nil {
		return result, err
	}
	return &response, nil
}

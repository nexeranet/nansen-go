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
	Chain                string `json:"chain"`
	BlockTimestamp       string `json:"block_timestamp"`
	TransactionHash      string `json:"transaction_hash"`
	TraderAddress        string `json:"trader_address"`
	TraderAddressLabel   string `json:"trader_address_label"`
	TokenBoughtAddress   string `json:"token_bought_address"`
	TokenSoldAddress     string `json:"token_sold_address"`
	TokenBoughtAmount    string `json:"token_bought_amount,string"`
	TokenSoldAmount      string `json:"token_sold_amount,string"`
	TokenBoughtSymbol    string `json:"token_bought_symbol"`
	TokenSoldSymbol      string `json:"token_sold_symbol"`
	TokenBoughtAgeDays   string `json:"token_bought_age_days,string"`
	TokenSoldAgeDays     string `json:"token_sold_age_days,string"`
	TokenBoughtMarketCap string `json:"token_bought_market_cap,string"`
	TokenSoldMarketCap   string `json:"token_sold_market_cap,string"`
	TradeValueUsd        string `json:"trade_value_usd,string"`
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

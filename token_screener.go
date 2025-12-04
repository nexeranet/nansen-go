package nansen

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetTokenScreeningBody struct {
	Chains     []string        `json:"chains"`
	Date       DateBody        `json:"date"`
	Pagination *PaginationBody `json:"pagination,omitempty"`
	Filters    map[string]any  `json:"filters,omitempty"`
	OrderBy    []SortOrderBody `json:"order_by,omitempty"`
}

type TokenScreeningData struct {
	Chain           string      `json:"chain"`
	TokenAddress    string      `json:"token_address"`
	TokenSymbol     string      `json:"token_symbol"`
	TokenAgeDays    json.Number `json:"token_age_days"`
	MarketCapUsd    json.Number `json:"market_cap_usd"`
	Liquidity       json.Number `json:"liquidity"`
	PriceUsd        json.Number `json:"price_usd"`
	PriceChange     json.Number `json:"price_change"`
	Fdv             json.Number `json:"fdv"`
	FdvMcRatio      json.Number `json:"fdv_mc_ratio"`
	BuyVolume       json.Number `json:"buy_volume"`
	InflowFdvRatio  json.Number `json:"inflow_fdv_ratio"`
	OutflowFdvRatio json.Number `json:"outflow_fdv_ratio"`
	SellVolume      json.Number `json:"sell_volume"`
	Volume          json.Number `json:"volume"`
	Netflow         json.Number `json:"netflow"`
}

type GetTokenScreeningResponse struct {
	Data       []TokenScreeningData `json:"data"`
	Pagination *PaginationResponse  `json:"pagination,omitempty"`
}

func (c *Client) GetTokenScreening(ctx context.Context, body GetTokenScreeningBody) (result *GetTokenScreeningResponse, err error) {
	request := NewRequest(c.Url("/token-screener"), body, http.MethodPost)
	_, err = c.doCall(ctx, request, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

package nansen

import (
	"context"
	"net/http"
)

type GetTokenScreeningBody struct {
	Chains []string `json:"chains"`
	//Date       DateBody        `json:"date"`
	TimeFrame  string          `json:"timeframe"`
	Pagination *PaginationBody `json:"pagination,omitempty"`
	Filters    map[string]any  `json:"filters,omitempty"`
	OrderBy    []SortOrderBody `json:"order_by,omitempty"`
}

type TokenScreeningData struct {
	Chain           string  `json:"chain"`
	TokenAddress    string  `json:"token_address"`
	TokenSymbol     string  `json:"token_symbol"`
	TokenAgeDays    float64 `json:"token_age_days"`
	MarketCapUsd    float64 `json:"market_cap_usd"`
	Liquidity       float64 `json:"liquidity"`
	PriceUsd        float64 `json:"price_usd"`
	PriceChange     float64 `json:"price_change"`
	Fdv             float64 `json:"fdv"`
	FdvMcRatio      float64 `json:"fdv_mc_ratio"`
	BuyVolume       float64 `json:"buy_volume"`
	InflowFdvRatio  float64 `json:"inflow_fdv_ratio"`
	OutflowFdvRatio float64 `json:"outflow_fdv_ratio"`
	SellVolume      float64 `json:"sell_volume"`
	Volume          float64 `json:"volume"`
	Netflow         float64 `json:"netflow"`
}

type GetTokenScreeningResponse struct {
	Data       []TokenScreeningData `json:"data"`
	Pagination *PaginationResponse  `json:"pagination,omitempty"`
}

func (c *Client) GetTokenScreening(ctx context.Context, body GetTokenScreeningBody) (result *GetTokenScreeningResponse, err error) {
	var response GetTokenScreeningResponse
	request := NewRequest(c.Url("/token-screener"), body, http.MethodPost)
	_, err = c.doCall(ctx, request, &response)
	if err != nil {
		return result, err
	}
	return &response, nil
}

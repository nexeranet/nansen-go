package nansen

import (
	"context"
	"net/http"
)

type GetAddressPnLSummaryBody struct {
	Address    *string  `json:"address,omitempty"`
	EntityName *string  `json:"entity_name,omitempty"`
	Chain      string   `json:"chain"`
	Date       DateBody `json:"date"`
}
type Top5TokenStruct struct {
	RealizedPnl  float64 `json:"realized_pnl"`
	RealizedRoi  float64 `json:"realized_roi"`
	TokenAddress string  `json:"token_address"`
	TokenSymbol  string  `json:"token_symbol"`
	Chain        string  `json:"chain"`
}

type GetAddressPnLSummaryResponse struct {
	Top5Tokens         []Top5TokenStruct   `json:"top5_tokens"`
	TradedTokenCount   float64             `json:"traded_token_count"`
	TradedTimes        float64             `json:"traded_times"`
	RealizedPnlUsd     float64             `json:"realized_pnl_usd"`
	RealizedPnlPercent float64             `json:"realized_pnl_percent"`
	WinRate            float64             `json:"win_rate"`
	Pagination         *PaginationResponse `json:"pagination,omitempty"`
}

func (c *Client) GetAddressPnLSummary(ctx context.Context, body GetAddressPnLSummaryBody) (result *GetAddressPnLSummaryResponse, err error) {
	var response GetAddressPnLSummaryResponse
	request := NewRequest(c.Url("/profiler/address/pnl-summary"), body, http.MethodPost)
	_, err = c.doCall(ctx, request, &response)
	if err != nil {
		return result, err
	}
	return &response, nil
}

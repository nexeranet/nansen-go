package nansen

import (
	"context"
	"net/http"
	"time"
)

type GetTGMDEXTradesBody struct {
	Chain          string          `json:"chain"`
	TokenAddress   string          `json:"token_address"`
	OnlySmartMoney *bool           `json:"only_smart_money,omitempty"`
	Date           *DateBody       `json:"date,omitempty"`
	Pagination     PaginationBody  `json:"pagination"`
	Filters        map[string]any  `json:"filters,omitempty"`
	OrderBy        []SortOrderBody `json:"order_by,omitempty"`
}

type GetTGMDEXTradesData struct {
	BlockTimestamp        time.Time `json:"block_timestamp"`
	TransactionHash       string    `json:"transaction_hash"`
	TraderAddress         string    `json:"trader_address"`
	TraderAddressLabel    string    `json:"trader_address_label"`
	Action                string    `json:"action"`
	TokenAddress          string    `json:"token_address"`
	TokenName             string    `json:"token_name"`
	TokenAmount           float64   `json:"token_amount"`
	TradedTokenAddress    string    `json:"traded_token_address"`
	TradedTokenName       string    `json:"traded_token_name"`
	TradedTokenAmount     float64   `json:"traded_token_amount"`
	EstimatedSwapPriceUsd float64   `json:"estimated_swap_price_usd"`
	EstimatedValueUsd     float64   `json:"estimated_value_usd"`
}

type GetTGMDEXTradesResponse struct {
	Data       []GetTGMDEXTradesData `json:"data"`
	Pagination *PaginationResponse   `json:"pagination,omitempty"`
}

func (c *Client) GetTGMDEXTrades(ctx context.Context, body GetTGMDEXTradesBody) (result *GetTGMDEXTradesResponse, err error) {
	request := NewRequest(c.Url("/tgm/dex-trades"), body, http.MethodPost)
	_, err = c.doCall(ctx, request, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

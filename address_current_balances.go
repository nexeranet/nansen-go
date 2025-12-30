package nansen

import (
	"context"
	"net/http"
)

type GetAddressCurrentBalanceBody struct {
	Address       *string         `json:"address,omitempty"`
	EntityName    *string         `json:"entity_name,omitempty"`
	Chain         string          `json:"chain"`
	HideSpamToken *bool           `json:"hide_spam_token,omitempty"`
	Filters       map[string]any  `json:"filters,omitempty"`
	Pagination    *PaginationBody `json:"pagination,omitempty"`
	OrderBy       []SortOrderBody `json:"order_by,omitempty"`
}

type GetAddressCurrentBalanceData struct {
	Chain        string  `json:"chain"`
	Address      string  `json:"address"`
	TokenAddress string  `json:"token_address"`
	TokenSymbol  string  `json:"token_symbol"`
	TokenName    string  `json:"token_name"`
	TokenAmount  float64 `json:"token_amount"`
	PriceUsd     float64 `json:"price_usd"`
	ValueUsd     float64 `json:"value_usd"`
}

type GetAddressCurrentBalanceResponse struct {
	Data       []GetAddressCurrentBalanceData `json:"data"`
	Pagination *PaginationResponse            `json:"pagination,omitempty"`
}

func (c *Client) GetAddressCurrentBalance(ctx context.Context, body GetAddressCurrentBalanceBody) (result *GetAddressCurrentBalanceResponse, err error) {
	request := NewRequest(c.Url("/profiler/address/current-balance"), body, http.MethodPost)
	_, err = c.doCall(ctx, request, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

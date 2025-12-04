package nansen

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetAddressCurrentBalanceBody struct {
	Address       string          `json:"address"`
	EntityName    string          `json:"entity_name"`
	Chain         string          `json:"chain"`
	HideSpamToken *bool           `json:"hide_spam_token,omitempty"`
	Filter        map[string]any  `json:"filter,omitempty"`
	Pagination    *PaginationBody `json:"pagination,omitempty"`
	OrderBy       []SortOrderBody `json:"order_by,omitempty"`
}

type GetAddressCurrentBalanceData struct {
	Chain        string      `json:"chain"`
	Address      string      `json:"address"`
	TokenAddress string      `json:"token_address"`
	TokenSymbol  string      `json:"token_symbol"`
	TokenName    string      `json:"token_name"`
	TokenAmount  json.Number `json:"token_amount"`
	PriceUsd     json.Number `json:"price_usd"`
	ValueUsd     json.Number `json:"value_usd"`
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

package nansen

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetAddressHistoricalBalancesBody struct {
	Address    *string         `json:"address,omitempty"`
	EntityName *string         `json:"entity_name,omitempty"`
	Chain      string          `json:"chain"`
	Date       DateBody        `json:"date"`
	Filter     map[string]any  `json:"filter,omitempty"`
	Pagination *PaginationBody `json:"pagination,omitempty"`
	OrderBy    []SortOrderBody `json:"order_by,omitempty"`
}

type GetAddressHistoricalBalancesData struct {
	BlockTimestamp string      `json:"block_timestamp"`
	TokenAddress   string      `json:"token_address"`
	Chain          string      `json:"chain"`
	TokenAmount    json.Number `json:"token_amount"`
	ValueUsd       json.Number `json:"value_usd"`
	TokenSymbol    string      `json:"token_symbol"`
}

type GetAddressHistoricalBalancesResponse struct {
	Data       []GetAddressHistoricalBalancesData `json:"data"`
	Pagination *PaginationResponse                `json:"pagination,omitempty"`
}

func (c *Client) GetAddressHistoricalBalances(ctx context.Context, body GetAddressHistoricalBalancesBody) (result *GetAddressHistoricalBalancesResponse, err error) {
	request := NewRequest(c.Url("/profiler/address/historical-balances"), body, http.MethodPost)
	_, err = c.doCall(ctx, request, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

package nansen

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetTGMHoldersBody struct {
	Chain             string          `json:"chain"`
	TokenAddress      string          `json:"token_address"`
	AggregateByEntity *bool           `json:"aggregate_by_entity,omitempty"`
	LabelType         string          `json:"label_type,omitempty"`
	Pagination        *PaginationBody `json:"pagination,omitempty"`
	Filter            map[string]any  `json:"filter,omitempty"`
	OrderBy           []SortOrderBody `json:"order_by,omitempty"`
}

type GetTGMHoldersData struct {
	Address             string      `json:"address"`
	AddressLabel        string      `json:"address_label"`
	TokenAmount         json.Number `json:"token_amount"`
	TotalOutflow        json.Number `json:"total_outflow"`
	TotalInflow         json.Number `json:"total_inflow"`
	BalanceChange24H    json.Number `json:"balance_change_24h"`
	BalanceChange7D     json.Number `json:"balance_change_7d"`
	BalanceChange30D    json.Number `json:"balance_change_30d"`
	OwnershipPercentage json.Number `json:"ownership_percentage"`
	ValueUsd            json.Number `json:"value_usd"`
}

type GetTGMHoldersResponse struct {
	Data       []GetTGMHoldersData `json:"data"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

func (c *Client) GetTGMHolders(ctx context.Context, body GetTGMTransfersBody) (result *GetTGMTransfersResponse, err error) {
	request := NewRequest(c.Url("/tgm/transfers"), body, http.MethodPost)
	_, err = c.doCall(ctx, request, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

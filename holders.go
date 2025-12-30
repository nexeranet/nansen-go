package nansen

import (
	"context"
	"net/http"
)

type GetTGMHoldersBody struct {
	Chain             string          `json:"chain"`
	TokenAddress      string          `json:"token_address"`
	AggregateByEntity *bool           `json:"aggregate_by_entity,omitempty"`
	LabelType         string          `json:"label_type,omitempty"`
	Pagination        *PaginationBody `json:"pagination,omitempty"`
	Filters           map[string]any  `json:"filters,omitempty"`
	OrderBy           []SortOrderBody `json:"order_by,omitempty"`
}

type GetTGMHoldersData struct {
	Address             string  `json:"address"`
	AddressLabel        string  `json:"address_label"`
	TokenAmount         float64 `json:"token_amount"`
	TotalOutflow        float64 `json:"total_outflow"`
	TotalInflow         float64 `json:"total_inflow"`
	BalanceChange24H    float64 `json:"balance_change_24h"`
	BalanceChange7D     float64 `json:"balance_change_7d"`
	BalanceChange30D    float64 `json:"balance_change_30d"`
	OwnershipPercentage float64 `json:"ownership_percentage"`
	ValueUsd            float64 `json:"value_usd"`
}

type GetTGMHoldersResponse struct {
	Data       []GetTGMHoldersData `json:"data"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

func (c *Client) GetTGMHolders(ctx context.Context, body GetTGMHoldersBody) (result *GetTGMHoldersResponse, err error) {
	result = &GetTGMHoldersResponse{}
	request := NewRequest(c.Url("/tgm/transfers"), body, http.MethodPost)
	_, err = c.doCall(ctx, request, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

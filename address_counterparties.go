package nansen

import (
	"context"
	"net/http"
)

type GetAddressCounterpartiesBody struct {
	Address     string          `json:"address"`
	EntityName  string          `json:"entity_name"`
	Chain       string          `json:"chain"`
	Date        DateBody        `json:"date"`
	SourceInput *string         `json:"source_input,omitempty"`
	GroupBy     *string         `json:"group_by,omitempty"`
	Filters     map[string]any  `json:"filters,omitempty"`
	Pagination  *PaginationBody `json:"pagination,omitempty"`
	OrderBy     []SortOrderBody `json:"order_by,omitempty"`
}

type GetAddressCounterpartiesData struct {
	CounterpartyAddress      string   `json:"counterparty_address"`
	CounterpartyAddressLabel []string `json:"counterparty_address_label"`
	InteractionCount         float64  `json:"interaction_count"`
	TotalVolumeUsd           float64  `json:"total_volume_usd"`
	VolumeInUsd              float64  `json:"volume_in_usd"`
	VolumeOutUsd             float64  `json:"volume_out_usd"`
	TokensInfo               []struct {
		TokenAddress string `json:"token_address"`
		TokenSymbol  string `json:"token_symbol"`
		TokenName    string `json:"token_name"`
		NumTransfer  string `json:"num_transfer"`
	} `json:"tokens_info"`
}

type GetAddressCounterpartiesResponse struct {
	Data       []GetAddressCounterpartiesData `json:"data"`
	Pagination *PaginationResponse            `json:"pagination,omitempty"`
}

func (c *Client) GetAddressCounterparties(ctx context.Context, body GetAddressCounterpartiesBody) (result *GetAddressCounterpartiesResponse, err error) {
	request := NewRequest(c.Url("/profiler/address/counterparties"), body, http.MethodPost)
	_, err = c.doCall(ctx, request, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

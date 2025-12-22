package nansen

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetTGMFlowsBody struct {
	Chain        string          `json:"chain"`
	TokenAddress string          `json:"token_address"`
	TimeFrame    TimeFrame       `json:"time_frame"`
	Date         *DateBody       `json:"date,omitempty"`
	Label        string          `json:"label"`
	Pagination   PaginationBody  `json:"pagination"`
	Filters      map[string]any  `json:"filters,omitempty"`
	OrderBy      []SortOrderBody `json:"order_by,omitempty"`
}

type GetTGMFlowsData struct {
	Date               string      `json:"date"`
	PriceUsd           json.Number `json:"price_usd"`
	TokenAmount        json.Number `json:"token_amount"`
	ValueUsd           json.Number `json:"value_usd"`
	HoldersCount       json.Number `json:"holders_count"`
	TotalInflowsCount  json.Number `json:"total_inflows_count"`
	TotalOutflowsCount json.Number `json:"total_outflows_count"`
}

type GetTGMFlowsResponse struct {
	Data       []GetTGMFlowsData   `json:"data"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

func (c *Client) GetTGMFlows(ctx context.Context, body GetTGMFlowsBody) (result *GetTGMFlowsResponse, err error) {
	request := NewRequest(c.Url("/tgm/flows"), body, http.MethodPost)
	_, err = c.doCall(ctx, request, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

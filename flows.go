package nansen

import (
	"context"
	"net/http"
)

type GetTGMFlowsBody struct {
	Chain        string          `json:"chain"`
	TokenAddress string          `json:"token_address"`
	Date         *DateBody       `json:"date,omitempty"`
	Label        string          `json:"label"`
	Pagination   PaginationBody  `json:"pagination"`
	Filters      map[string]any  `json:"filters,omitempty"`
	OrderBy      []SortOrderBody `json:"order_by,omitempty"`
}

type GetTGMFlowsData struct {
	Date               string  `json:"date"`
	PriceUsd           float64 `json:"price_usd"`
	TokenAmount        float64 `json:"token_amount"`
	ValueUsd           float64 `json:"value_usd"`
	HoldersCount       float64 `json:"holders_count"`
	TotalInflowsCount  float64 `json:"total_inflows_count"`
	TotalOutflowsCount float64 `json:"total_outflows_count"`
}

type GetTGMFlowsResponse struct {
	Data       []GetTGMFlowsData   `json:"data"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

func (c *Client) GetTGMFlows(ctx context.Context, body GetTGMFlowsBody) (result *GetTGMFlowsResponse, err error) {
	var response GetTGMFlowsResponse
	request := NewRequest(c.Url("/tgm/flows"), body, http.MethodPost)
	_, err = c.doCall(ctx, request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

package nansen

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetSmartMoneyNetflowBody struct {
	Chains     []string        `json:"chains"`
	Filters    map[string]any  `json:"filters,omitempty"`
	Pagination *PaginationBody `json:"pagination,omitempty"`
	OrderBy    []SortOrderBody `json:"order_by,omitempty"`
}

type GetSmartMoneyNetflowData struct {
	TokenAddress  string      `json:"token_address"`
	TokenSymbol   string      `json:"token_symbol"`
	NetFlow24HUsd json.Number `json:"net_flow_24h_usd,string"`
	NetFlow7DUsd  json.Number `json:"net_flow_7d_usd,string"`
	NetFlow30DUsd json.Number `json:"net_flow_30d_usd"`
	Chain         string      `json:"chain"`
	TokenSectors  []string    `json:"token_sectors"`
	TraderCount   json.Number `json:"trader_count,string"`
	TokenAgeDays  json.Number `json:"token_age_days,string"`
	MarketCapUsd  json.Number `json:"market_cap_usd,string"`
}

type GetSmartMoneyNetflowResponse struct {
	Data       []GetSmartMoneyNetflowData `json:"data"`
	Pagination *PaginationResponse        `json:"pagination,omitempty"`
}

func (c *Client) GetSmartMoneyNetflow(ctx context.Context, body GetSmartMoneyNetflowBody) (*GetSmartMoneyNetflowResponse, error) {
	var response GetSmartMoneyNetflowResponse
	request := NewRequest(c.Url("/smart-money/netflow"), body, http.MethodPost)
	_, err := c.doCall(ctx, request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

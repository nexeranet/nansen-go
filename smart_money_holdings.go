package nansen

import (
	"context"
	"net/http"
)

type GetSmartMoneyHoldingsBody struct {
	Chains     []string        `json:"chains"`
	Filters    map[string]any  `json:"filters,omitempty"`
	Pagination *PaginationBody `json:"pagination,omitempty"`
	OrderBy    []SortOrderBody `json:"order_by,omitempty"`
}

type GetSmartMoneyHoldingsData struct {
	Chain                   string   `json:"chain"`
	TokenAddress            string   `json:"token_address"`
	TokenSymbol             string   `json:"token_symbol"`
	TokenSectors            []string `json:"token_sectors"`
	ValueUsd                string   `json:"value_usd,string"`
	Balance24HPercentChange string   `json:"balance_24h_percent_change,string"`
	HoldersCount            string   `json:"holders_count,string"`
	ShareOfHoldingsPercent  string   `json:"share_of_holdings_percent,string"`
	TokenAgeDays            string   `json:"token_age_days,string"`
	MarketCapUsd            string   `json:"market_cap_usd,string"`
}

type GetSmartMoneyHoldingsResponse struct {
	Data       []GetSmartMoneyHoldingsData `json:"data"`
	Pagination *PaginationResponse         `json:"pagination,omitempty"`
}

func (c *Client) GetSmartMoneyHoldings(ctx context.Context, body GetSmartMoneyHoldingsBody) (*GetSmartMoneyHoldingsResponse, error) {
	var response GetSmartMoneyHoldingsResponse
	request := NewRequest(c.Url("/smart-money/holdings"), body, http.MethodPost)
	_, err := c.doCall(ctx, request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

package nansen

import (
	"context"
	"net/http"
)

type GetAddressPnLBody struct {
	Chain        string          `json:"chain"`
	TokenAddress string          `json:"token_address"`
	Date         *DateBody       `json:"date,omitempty"`
	Label        string          `json:"label"`
	Pagination   PaginationBody  `json:"pagination"`
	Filters      map[string]any  `json:"filters,omitempty"`
	OrderBy      []SortOrderBody `json:"order_by,omitempty"`
}

type GetAddressPnLData struct {
	TokenAddress         string  `json:"token_address"`
	TokenSymbol          string  `json:"token_symbol"`
	TokenPrice           float64 `json:"token_price"`
	RoiPercentRealised   float64 `json:"roi_percent_realised"`
	PnlUsdRealised       float64 `json:"pnl_usd_realised"`
	PnlUsdUnrealised     float64 `json:"pnl_usd_unrealised"`
	RoiPercentUnrealised float64 `json:"roi_percent_unrealised"`
	BoughtAmount         float64 `json:"bought_amount"`
	BoughtUsd            float64 `json:"bought_usd"`
	CostBasisUsd         float64 `json:"cost_basis_usd"`
	SoldAmount           float64 `json:"sold_amount"`
	SoldUsd              float64 `json:"sold_usd"`
	AvgSoldPriceUsd      float64 `json:"avg_sold_price_usd"`
	HoldingAmount        float64 `json:"holding_amount"`
	HoldingUsd           float64 `json:"holding_usd"`
	NofBuys              string  `json:"nof_buys"`
	NofSells             string  `json:"nof_sells"`
	MaxBalanceHeld       float64 `json:"max_balance_held"`
	MaxBalanceHeldUsd    float64 `json:"max_balance_held_usd"`
}

type GetAddressPnLResponse struct {
	Data       []GetAddressPnLData `json:"data"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

func (c *Client) GetAddressPnL(ctx context.Context, body GetAddressPnLBody) (*GetAddressPnLResponse, error) {
	var response GetAddressPnLResponse
	request := NewRequest(c.Url("/profiler/address/pnl"), body, http.MethodPost)
	_, err := c.doCall(ctx, request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

package nansen

import (
	"context"
	"net/http"
)

type GetTGMFlowIntelligenceBody struct {
	Chain        string         `json:"chain"`
	TokenAddress string         `json:"token_address"`
	TimeFrame    string         `json:"timeframe"`
	Filters      map[string]any `json:"filters"`
}

type GetTGMFlowIntelligenceData struct {
	PublicFigureNetFlowUsd  float64 `json:"public_figure_net_flow_usd"`
	PublicFigureAvgFlowUsd  float64 `json:"public_figure_avg_flow_usd"`
	PublicFigureWalletCount float64 `json:"public_figure_wallet_count"`
	TopPnlNetFlowUsd        float64 `json:"top_pnl_net_flow_usd"`
	TopPnlAvgFlowUsd        float64 `json:"top_pnl_avg_flow_usd"`
	TopPnlWalletCount       float64 `json:"top_pnl_wallet_count"`
	WhaleNetFlowUsd         float64 `json:"whale_net_flow_usd"`
	WhaleAvgFlowUsd         float64 `json:"whale_avg_flow_usd"`
	WhaleWalletCount        float64 `json:"whale_wallet_count"`
	SmartTraderNetFlowUsd   float64 `json:"smart_trader_net_flow_usd"`
	SmartTraderAvgFlowUsd   float64 `json:"smart_trader_avg_flow_usd"`
	SmartTraderWalletCount  float64 `json:"smart_trader_wallet_count"`
	ExchangeNetFlowUsd      float64 `json:"exchange_net_flow_usd"`
	ExchangeAvgFlowUsd      float64 `json:"exchange_avg_flow_usd"`
	ExchangeWalletCount     float64 `json:"exchange_wallet_count"`
	FreshWalletsNetFlowUsd  float64 `json:"fresh_wallets_net_flow_usd"`
	FreshWalletsAvgFlowUsd  float64 `json:"fresh_wallets_avg_flow_usd"`
	FreshWalletsWalletCount float64 `json:"fresh_wallets_wallet_count"`
}

type GetTGMFlowIntelligenceResponse struct {
	Data []GetTGMFlowIntelligenceData `json:"data"`
}

func (c *Client) GetTGMFlowIntelligence(ctx context.Context, body GetTGMFlowIntelligenceBody) (result *[]GetTGMFlowIntelligenceData, err error) {
	var response GetTGMFlowIntelligenceResponse
	request := NewRequest(c.Url("/tgm/flow-intelligence"), body, http.MethodPost)
	_, err = c.doCall(ctx, request, &response)
	if err != nil {
		return result, err
	}
	return &response.Data, nil
}

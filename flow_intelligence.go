package nansen

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetTGMFlowIntelligenceBody struct {
	Chain        string         `json:"chain"`
	TokenAddress string         `json:"token_address"`
	TimeFrame    TimeFrame      `json:"time_frame"`
	Filters      map[string]any `json:"filters"`
}

type GetTGMFlowIntelligenceData struct {
	PublicFigureNetFlowUsd  json.Number `json:"public_figure_net_flow_usd"`
	PublicFigureAvgFlowUsd  json.Number `json:"public_figure_avg_flow_usd"`
	PublicFigureWalletCount json.Number `json:"public_figure_wallet_count"`
	TopPnlNetFlowUsd        json.Number `json:"top_pnl_net_flow_usd"`
	TopPnlAvgFlowUsd        json.Number `json:"top_pnl_avg_flow_usd"`
	TopPnlWalletCount       json.Number `json:"top_pnl_wallet_count"`
	WhaleNetFlowUsd         json.Number `json:"whale_net_flow_usd"`
	WhaleAvgFlowUsd         json.Number `json:"whale_avg_flow_usd"`
	WhaleWalletCount        json.Number `json:"whale_wallet_count"`
	SmartTraderNetFlowUsd   json.Number `json:"smart_trader_net_flow_usd"`
	SmartTraderAvgFlowUsd   json.Number `json:"smart_trader_avg_flow_usd"`
	SmartTraderWalletCount  json.Number `json:"smart_trader_wallet_count"`
	ExchangeNetFlowUsd      json.Number `json:"exchange_net_flow_usd"`
	ExchangeAvgFlowUsd      json.Number `json:"exchange_avg_flow_usd"`
	ExchangeWalletCount     json.Number `json:"exchange_wallet_count"`
	FreshWalletsNetFlowUsd  json.Number `json:"fresh_wallets_net_flow_usd"`
	FreshWalletsAvgFlowUsd  json.Number `json:"fresh_wallets_avg_flow_usd"`
	FreshWalletsWalletCount json.Number `json:"fresh_wallets_wallet_count"`
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

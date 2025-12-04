package nansen

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetPortfolioDeFiHoldingsBody struct {
	WalletAddresses string `json:"wallet_addresses"`
}
type GetPortfolioDeFiHoldingsResponse struct {
	Summary struct {
		TotalValueUsd   json.Number `json:"total_value_usd"`
		TotalAssetsUsd  json.Number `json:"total_assets_usd"`
		TotalDebtsUsd   json.Number `json:"total_debts_usd"`
		TotalRewardsUsd json.Number `json:"total_rewards_usd"`
		TokenCount      json.Number `json:"token_count"`
		ProtocolCount   json.Number `json:"protocol_count"`
	} `json:"summary"`
	Protocols []struct {
		ProtocolName    string      `json:"protocol_name"`
		Chain           string      `json:"chain"`
		TotalValueUsd   json.Number `json:"total_value_usd"`
		TotalAssetsUsd  json.Number `json:"total_assets_usd"`
		TotalDebtsUsd   json.Number `json:"total_debts_usd"`
		TotalRewardsUsd json.Number `json:"total_rewards_usd"`
		Tokens          []struct {
			Address      string      `json:"address"`
			Symbol       string      `json:"symbol"`
			Amount       json.Number `json:"amount"`
			ValueUsd     json.Number `json:"value_usd"`
			PositionType string      `json:"position_type"`
		} `json:"tokens"`
	} `json:"protocols"`
}

func (c *Client) GetPortfolioDeFiHoldings(ctx context.Context, body GetPortfolioDeFiHoldingsBody) (result *GetPortfolioDeFiHoldingsResponse, err error) {
	request := NewRequest(c.Url("/portfolio/defi-holdings"), body, http.MethodPost)
	_, err = c.doCall(ctx, request, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

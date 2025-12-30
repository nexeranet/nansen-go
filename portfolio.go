package nansen

import (
	"context"
	"net/http"
)

type GetPortfolioDeFiHoldingsBody struct {
	WalletAddresses string `json:"wallet_addresses"`
}
type GetPortfolioDeFiHoldingsResponse struct {
	Summary struct {
		TotalValueUsd   float64 `json:"total_value_usd"`
		TotalAssetsUsd  float64 `json:"total_assets_usd"`
		TotalDebtsUsd   float64 `json:"total_debts_usd"`
		TotalRewardsUsd float64 `json:"total_rewards_usd"`
		TokenCount      float64 `json:"token_count"`
		ProtocolCount   float64 `json:"protocol_count"`
	} `json:"summary"`
	Protocols []struct {
		ProtocolName    string  `json:"protocol_name"`
		Chain           string  `json:"chain"`
		TotalValueUsd   float64 `json:"total_value_usd"`
		TotalAssetsUsd  float64 `json:"total_assets_usd"`
		TotalDebtsUsd   float64 `json:"total_debts_usd"`
		TotalRewardsUsd float64 `json:"total_rewards_usd"`
		Tokens          []struct {
			Address      string  `json:"address"`
			Symbol       string  `json:"symbol"`
			Amount       float64 `json:"amount"`
			ValueUsd     float64 `json:"value_usd"`
			PositionType string  `json:"position_type"`
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

package nansen

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type GetTGMTokenInformationBody struct {
	Chain        string    `json:"chain"`
	TokenAddress string    `json:"token_address"`
	TimeFrame    TimeFrame `json:"time_frame"`
}

type GetTGMTokenInformation struct {
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	ContractAddress string `json:"contract_address"`
	Logo            string `json:"logo"`
	TokenDetails    struct {
		TokenDeploymentDate time.Time   `json:"token_deployment_date"`
		Website             string      `json:"website"`
		X                   string      `json:"x"`
		Telegram            string      `json:"telegram"`
		MarketCapUsd        json.Number `json:"market_cap_usd"`
		FdvUsd              json.Number `json:"fdv_usd"`
		CirculatingSupply   json.Number `json:"circulating_supply"`
		TotalSupply         json.Number `json:"total_supply"`
	} `json:"token_details"`
	SpotMetrics struct {
		VolumeTotalUsd json.Number `json:"volume_total_usd"`
		BuyVolumeUsd   json.Number `json:"buy_volume_usd"`
		SellVolumeUsd  json.Number `json:"sell_volume_usd"`
		TotalBuys      json.Number `json:"total_buys"`
		TotalSells     json.Number `json:"total_sells"`
		UniqueBuyers   json.Number `json:"unique_buyers"`
		UniqueSellers  json.Number `json:"unique_sellers"`
		LiquidityUsd   json.Number `json:"liquidity_usd"`
		TotalHolders   json.Number `json:"total_holders"`
	} `json:"spot_metrics"`
}

type GetTGMTokenInformationResponse struct {
	Data GetTGMTokenInformation `json:"data"`
}

func (c *Client) GetTGMTokenInformation(ctx context.Context, body GetTGMTokenInformationBody) (result *GetTGMTokenInformation, err error) {
	var response GetTGMTokenInformationResponse
	request := NewRequest(c.Url("/tgm/token-information"), body, http.MethodPost)
	_, err = c.doCall(ctx, request, &response)
	if err != nil {
		return result, err
	}
	return &response.Data, nil
}

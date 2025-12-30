package nansen

import (
	"context"
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
		TokenDeploymentDate time.Time `json:"token_deployment_date"`
		Website             string    `json:"website"`
		X                   string    `json:"x"`
		Telegram            string    `json:"telegram"`
		MarketCapUsd        float64   `json:"market_cap_usd"`
		FdvUsd              float64   `json:"fdv_usd"`
		CirculatingSupply   float64   `json:"circulating_supply"`
		TotalSupply         float64   `json:"total_supply"`
	} `json:"token_details"`
	SpotMetrics struct {
		VolumeTotalUsd float64 `json:"volume_total_usd"`
		BuyVolumeUsd   float64 `json:"buy_volume_usd"`
		SellVolumeUsd  float64 `json:"sell_volume_usd"`
		TotalBuys      float64 `json:"total_buys"`
		TotalSells     float64 `json:"total_sells"`
		UniqueBuyers   float64 `json:"unique_buyers"`
		UniqueSellers  float64 `json:"unique_sellers"`
		LiquidityUsd   float64 `json:"liquidity_usd"`
		TotalHolders   float64 `json:"total_holders"`
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

package nansen

import (
	"context"
	"encoding/json"
	"net/http"
)

type WhoBoughtSoldBody struct {
	Chain        string          `json:"chain"`
	TokenAddress string          `json:"token_address"`
	Date         DateBody        `json:"date"`
	Filter       map[string]any  `json:"filter,omitempty"`
	OrderBy      []SortOrderBody `json:"order_by,omitempty"`
	Pagination   *PaginationBody `json:"pagination,omitempty"`
	// BUY OR SELL string. Default: BUY
	BuyOrSell *string `json:"buy_or_sell,omitempty"`
}

type WhoBoughtSoldData struct {
	Address           string      `json:"address"`
	AddressLabel      string      `json:"address_label"`
	BoughtTokenVolume json.Number `json:"bought_token_volume"`
	SoldTokenVolume   json.Number `json:"sold_token_volume"`
	TokenTradeVolume  json.Number `json:"token_trade_volume"`
	BoughtVolumeUsd   json.Number `json:"bought_volume_usd"`
	SoldVolumeUsd     json.Number `json:"sold_volume_usd"`
	TradeVolumeUsd    json.Number `json:"trade_volume_usd"`
}

type WhoBoughtSoldResponse struct {
	Data       []WhoBoughtSoldData `json:"data"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

func (c *Client) WhoBoughtSold(ctx context.Context, body WhoBoughtSoldBody) (result *WhoBoughtSoldResponse, err error) {
	request := NewRequest(c.Url("/tgm/who-bought-sold"), body, http.MethodPost)
	_, err = c.doCall(ctx, request, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

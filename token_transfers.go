package nansen

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type GetTGMTransfersBody struct {
	Chain        string          `json:"chain"`
	TokenAddress string          `json:"token_address"`
	Date         DateBody        `json:"date"`
	Pagination   *PaginationBody `json:"pagination,omitempty"`
	Filters      map[string]any  `json:"filters,omitempty"`
	OrderBy      []SortOrderBody `json:"order_by,omitempty"`
}

type GetTGMTransfersData struct {
	BlockTimestamp   time.Time   `json:"block_timestamp"`
	TransactionHash  string      `json:"transaction_hash"`
	FromAddress      string      `json:"from_address"`
	ToAddress        string      `json:"to_address"`
	FromAddressLabel string      `json:"from_address_label"`
	ToAddressLabel   string      `json:"to_address_label"`
	TransactionType  string      `json:"transaction_type"`
	TransferAmount   json.Number `json:"transfer_amount"`
	TransferValueUsd json.Number `json:"transfer_value_usd"`
}

type GetTGMTransfersResponse struct {
	Data       []GetTGMTransfersData `json:"data"`
	Pagination *PaginationResponse   `json:"pagination,omitempty"`
}

func (c *Client) GetTGMTransfers(ctx context.Context, body GetTGMTransfersBody) (result *GetTGMTransfersResponse, err error) {
	request := NewRequest(c.Url("/tgm/transfers"), body, http.MethodPost)
	_, err = c.doCall(ctx, request, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

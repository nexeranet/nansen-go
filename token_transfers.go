package nansen

import (
	"context"
	"net/http"
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
	BlockTimestamp   string  `json:"block_timestamp,omitempty"`
	TransactionHash  string  `json:"transaction_hash,omitempty"`
	FromAddress      string  `json:"from_address,omitempty"`
	ToAddress        string  `json:"to_address,omitempty"`
	FromAddressLabel string  `json:"from_address_label,omitempty"`
	ToAddressLabel   string  `json:"to_address_label,omitempty"`
	TransactionType  string  `json:"transaction_type,omitempty"`
	TransferAmount   float64 `json:"transfer_amount,omitempty"`
	TransferValueUsd float64 `json:"transfer_value_usd,omitempty"`
}

type GetTGMTransfersResponse struct {
	Data       []GetTGMTransfersData `json:"data"`
	Pagination *PaginationResponse   `json:"pagination,omitempty"`
}

func (c *Client) GetTGMTransfers(ctx context.Context, body GetTGMTransfersBody) (result *GetTGMTransfersResponse, err error) {
	var response GetTGMTransfersResponse
	request := NewRequest(c.Url("/tgm/transfers"), body, http.MethodPost)
	_, err = c.doCall(ctx, request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

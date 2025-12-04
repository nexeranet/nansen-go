package nansen

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetSmartMoneyJupiterDCAsBody struct {
	Filters    map[string]any  `json:"filters,omitempty"`
	Pagination *PaginationBody `json:"pagination,omitempty"`
	OrderBy    []SortOrderBody `json:"order_by,omitempty"`
}

type GetSmartMoneyJupiterDCAsData struct {
	DcaCreatedAt              string      `json:"dca_created_at"`
	DcaUpdatedAt              string      `json:"dca_updated_at"`
	TraderAddress             string      `json:"trader_address"`
	TransactionHash           string      `json:"transaction_hash"`
	TraderAddressLabel        string      `json:"trader_address_label"`
	DcaVaultAddress           string      `json:"dca_vault_address"`
	InputTokenAddress         string      `json:"input_token_address"`
	OutputTokenAddress        string      `json:"output_token_address"`
	DepositTokenAmount        json.Number `json:"deposit_token_amount"`
	TokenSpentAmount          json.Number `json:"token_spent_amount"`
	OutputTokenRedeemedAmount json.Number `json:"output_token_redeemed_amount"`
	DcaStatus                 string      `json:"dca_status"`
	InputTokenSymbol          string      `json:"input_token_symbol"`
	OutputTokenSymbol         string      `json:"output_token_symbol"`
	DepositValueUsd           json.Number `json:"deposit_value_usd"`
}

type GetSmartMoneyJupiterDCAsResponse struct {
	Data       []GetSmartMoneyJupiterDCAsData `json:"data"`
	Pagination *PaginationResponse            `json:"pagination,omitempty"`
}

func (c *Client) GetSmartMoneyJupiterDCAs(ctx context.Context, body GetSmartMoneyJupiterDCAsBody) (*GetSmartMoneyJupiterDCAsResponse, error) {
	var response GetSmartMoneyJupiterDCAsResponse
	request := NewRequest(c.Url("/smart-money/dcas"), body, http.MethodPost)
	_, err := c.doCall(ctx, request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

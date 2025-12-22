package nansen

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetTransactionwithTokenTransferLookupBody struct {
	Chain           string `json:"chain"`
	TransactionHash string `json:"transaction_hash"`
	BlockTimestamp  string `json:"block_timestamp"`
}

type GetTransactionwithTokenTransferLookupData struct {
	Chain                 string `json:"chain"`
	TransactionHash       string `json:"transaction_hash"`
	FromAddress           string `json:"from_address"`
	FromAddressLabel      string `json:"from_address_label"`
	ToAddress             string `json:"to_address"`
	ToAddressLabel        string `json:"to_address_label"`
	NativeValue           int    `json:"native_value"`
	DatedNativePrice      int    `json:"dated_native_price"`
	DatedNativeValueUsd   int    `json:"dated_native_value_usd"`
	CurrentNativePrice    int    `json:"current_native_price"`
	CurrentNativeValueUsd int    `json:"current_native_value_usd"`
	ReceiptStatus         int    `json:"receipt_status"`
	BlockTimestamp        string `json:"block_timestamp"`
	TokenTransferArray    []struct {
		FromAddress      string `json:"from_address"`
		FromAddressLabel string `json:"from_address_label"`
		ToAddress        string `json:"to_address"`
		ToAddressLabel   string `json:"to_address_label"`
		TokenAddress     string `json:"token_address"`
		TokenSymbol      string `json:"token_symbol"`
		TokenAmount      int    `json:"token_amount"`
		DatedPriceUsd    int    `json:"dated_price_usd"`
		DatedValueUsd    int    `json:"dated_value_usd"`
		CurrentPriceUsd  int    `json:"current_price_usd"`
		CurrentValueUsd  int    `json:"current_value_usd"`
		TransferID       string `json:"transfer_id"`
	} `json:"token_transfer_array"`
	NftTransferArray []struct {
		FromAddress      string `json:"from_address"`
		FromAddressLabel string `json:"from_address_label"`
		ToAddress        string `json:"to_address"`
		ToAddressLabel   string `json:"to_address_label"`
		ProjectID        string `json:"project_id"`
		CollectionName   string `json:"collection_name"`
		NftID            string `json:"nft_id"`
	} `json:"nft_transfer_array"`
}

type GetTransactionwithTokenTransferLookupResponse struct {
	Data []GetTransactionwithTokenTransferLookupData `json:"data"`
}

func (c *Client) GetTransactionwithTokenTransferLookup(ctx context.Context, body GetTransactionwithTokenTransferLookupBody) (result *GetTGMTokenInformation, err error) {
	var response GetTGMTokenInformationResponse
	request := NewRequest(c.Url("/transaction-with-token-transfer-lookup"), body, http.MethodPost)
	_, err = c.doCall(ctx, request, &response)
	if err != nil {
		return result, err
	}
	return &response.Data, nil
}

type GetAddressTransactionsBody struct {
	Address       string          `json:"address"`
	Chain         string          `json:"chain"`
	Date          DateBody        `json:"date"`
	HideSpamToken *bool           `json:"hide_spam_token,omitempty"`
	Pagination    *PaginationBody `json:"pagination,omitempty"`
	Filters       map[string]any  `json:"filters,omitempty"`
	OrderBy       []SortOrderBody `json:"order_by,omitempty"`
}
type GetAddressTransactionsData struct {
	Chain      string `json:"chain"`
	Method     string `json:"method"`
	TokensSent []struct {
		TokenSymbol      string      `json:"token_symbol"`
		TokenAmount      json.Number `json:"token_amount"`
		PriceUsd         json.Number `json:"price_usd"`
		ValueUsd         json.Number `json:"value_usd"`
		TokenAddress     string      `json:"token_address"`
		Chain            string      `json:"chain"`
		FromAddress      string      `json:"from_address"`
		ToAddress        string      `json:"to_address"`
		FromAddressLabel string      `json:"from_address_label"`
		ToAddressLabel   string      `json:"to_address_label"`
	} `json:"tokens_sent"`
	TokensReceived []struct {
		TokenSymbol      string      `json:"token_symbol"`
		TokenAmount      json.Number `json:"token_amount"`
		PriceUsd         json.Number `json:"price_usd"`
		ValueUsd         json.Number `json:"value_usd"`
		TokenAddress     string      `json:"token_address"`
		Chain            string      `json:"chain"`
		FromAddress      string      `json:"from_address"`
		ToAddress        string      `json:"to_address"`
		FromAddressLabel string      `json:"from_address_label"`
		ToAddressLabel   string      `json:"to_address_label"`
	} `json:"tokens_received"`
	VolumeUsd       json.Number `json:"volume_usd"`
	BlockTimestamp  string      `json:"block_timestamp"`
	TransactionHash string      `json:"transaction_hash"`
	SourceType      string      `json:"source_type"`
}

type GetAddressTransactionsResponse struct {
	Data       []GetAddressTransactionsData `json:"data"`
	Pagination *PaginationResponse          `json:"pagination,omitempty"`
}

func (c *Client) GetAddressTransactions(ctx context.Context, body GetAddressTransactionsBody) (result *GetAddressTransactionsResponse, err error) {
	request := NewRequest(c.Url("/profiler/address/transactions"), body, http.MethodPost)
	_, err = c.doCall(ctx, request, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

package nansen

import (
	"context"
	"net/http"
)

type GetAddressRelatedWalletsBody struct {
	Address    string          `json:"address"`
	Chain      string          `json:"chain"`
	Pagination *PaginationBody `json:"pagination,omitempty"`
	OrderBy    []SortOrderBody `json:"order_by,omitempty"`
}

type GetAddressRelatedWalletsData struct {
	Address         string  `json:"address"`
	AddressLabel    string  `json:"address_label"`
	Relation        string  `json:"relation"`
	TransactionHash string  `json:"transaction_hash"`
	BlockTimestamp  string  `json:"block_timestamp"`
	Order           float32 `json:"order"`
	Chain           string  `json:"chain"`
}

type GetAddressRelatedWalletsResponse struct {
	Data       []GetAddressRelatedWalletsData `json:"data"`
	Pagination *PaginationResponse            `json:"pagination,omitempty"`
}

func (c *Client) GetAddressRelatedWallets(ctx context.Context, body GetAddressRelatedWalletsBody) (result *GetAddressRelatedWalletsResponse, err error) {
	result = &GetAddressRelatedWalletsResponse{}
	request := NewRequest(c.Url("/profiler/address/related-wallets"), body, http.MethodPost)
	_, err = c.doCall(ctx, request, result)
	if err != nil {
		return result, err
	}
	return result, nil
}

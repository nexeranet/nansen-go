package nansen

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestGetAddressTransactions(t *testing.T) {
	// Initialize the test environment
	key := os.Getenv("KEY")
	ctx := context.Background()
	now := time.Now()
	weekBack := now.AddDate(0, 0, -7)
	client := NewClient(DefaultURL, key)
	result, err := client.GetAddressTransactions(ctx, GetAddressTransactionsBody{
		Chain:   "solana",
		Address: "9bD2sWCjipw3LsQruoauupftDTNHdmXSnNyRjrEz4XuS",
		Date:    *NewDateBody(&weekBack, &now),
		Pagination: &PaginationBody{
			Page:    1,
			PerPage: 100,
		},
		OrderBy: []SortOrderBody{
			{
				Field:     "block_timestamp",
				Direction: SortingDirectionDESC,
			},
		},
	})
	if err != nil {
		t.Errorf("Error getting smart money dex trades: %v", err)
	}
	//t.Log(result)
	for _, item := range result.Data {
		t.Log(item.TransactionHash)
	}
}

package nansen

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestGetAddressPnL(t *testing.T) {
	// Initialize the test environment
	key := os.Getenv("KEY")
	ctx := context.Background()
	client := NewClient(DefaultURL, key)
	address := "9bD2sWCjipw3LsQruoauupftDTNHdmXSnNyRjrEz4XuS"
	now := time.Now()
	dayBack := now.AddDate(0, 0, -7)
	result, err := client.GetAddressPnL(ctx, GetAddressPnLBody{
		Chain:   "solana",
		Address: &address,
		Date:    NewDateBody(&dayBack, &now),
		Pagination: PaginationBody{
			Page:    1,
			PerPage: 20,
		},
	})
	if err != nil {
		t.Errorf("Error getting smart money dex trades: %v", err)
	}
	//t.Log(result)
	for _, item := range result.Data {
		t.Log(item.TokenSymbol, item.TokenPrice, item.PnlUsdRealised)
	}
}

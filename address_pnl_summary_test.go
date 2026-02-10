package nansen

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestGetAddressPnLSummary(t *testing.T) {
	// Initialize the test environment
	key := os.Getenv("KEY")
	ctx := context.Background()
	client := NewClient(DefaultURL, key)
	address := "9bD2sWCjipw3LsQruoauupftDTNHdmXSnNyRjrEz4XuS"
	now := time.Now()
	dayBack := now.AddDate(0, 0, -7)
	result, err := client.GetAddressPnLSummary(ctx, GetAddressPnLSummaryBody{
		Chain:   "solana",
		Address: &address,
		Date:    *NewDateBody(&dayBack, &now),
	})
	if err != nil {
		t.Errorf("Error getting smart money dex trades: %v", err)
	}
	t.Log(result.RealizedPnlPercent, result.RealizedPnlUsd)
	for _, item := range result.Top5Tokens {
		t.Log(item.TokenSymbol, item.RealizedPnl)
	}
}

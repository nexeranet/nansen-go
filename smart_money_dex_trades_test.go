package nansen

import (
	"context"
	"os"
	"testing"
)

func TestGetSmartMoneyDexTrades(t *testing.T) {
	// Initialize the test environment
	key := os.Getenv("KEY")
	ctx := context.Background()
	client := NewClient(DefaultURL, key)
	result, err := client.GetSmartMoneyDEXTrades(ctx, GetSmartMoneyDEXTradesBody{
		Chains: []string{"solana"},
		Pagination: &PaginationBody{
			Page:    1,
			PerPage: 100,
		},
		Filters: map[string]any{
			"trade_value_usd": map[string]any{
				"min": 10000,
			},
		},
		OrderBy: []SortOrderBody{
			{
				Field:     "trade_value_usd",
				Direction: SortingDirectionDESC,
			},
		},
	})
	if err != nil {
		t.Errorf("Error getting smart money dex trades: %v", err)
	}
	t.Log(result)
}

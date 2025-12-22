package nansen

import (
	"context"
	"os"
	"testing"
)

func TestGetSmartMoneyHoldings(t *testing.T) {
	key := os.Getenv("KEY")
	ctx := context.Background()
	client := NewClient(DefaultURL, key)
	result, err := client.GetSmartMoneyHoldings(ctx, GetSmartMoneyHoldingsBody{
		Chains: []string{"all"},
		Pagination: &PaginationBody{
			Page:    1,
			PerPage: 100,
		},
		Filters: map[string]any{
			"balance_24h_percent_change": map[string]any{
				"min": 10,
			},
		},
		OrderBy: []SortOrderBody{
			{
				Field:     "balance_24h_percent_change",
				Direction: SortingDirectionDESC,
			},
		},
	})
	if err != nil {
		t.Fatalf("GetSmartMoneyHoldings() error = %v", err)
	}
	if result == nil {
		t.Fatalf("GetSmartMoneyHoldings() result = %v, want non-nil", result)
	}
	if len(result.Data) == 0 {
		t.Fatalf("GetSmartMoneyHoldings() result.Data = %v, want non-empty", result.Data)
	}
	t.Logf("GetSmartMoneyHoldings() result.Data = %v", result.Data)
}

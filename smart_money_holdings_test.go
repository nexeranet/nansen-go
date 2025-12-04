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
			PerPage: 10,
		},
		Filters: map[string]any{
			"only_smart_money": true,
		},
		OrderBy: []SortOrderBody{
			{
				Field:     "chain",
				Direction: SortingDirectionASC,
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

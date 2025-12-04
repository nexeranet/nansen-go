package nansen

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestGetTokenScreening(t *testing.T) {
	key := os.Getenv("KEY")
	ctx := context.Background()
	client := NewClient(DefaultURL, key)
	from := time.Now().Add(-24 * time.Hour)
	to := time.Now()
	result, err := client.GetTokenScreening(ctx, GetTokenScreeningBody{
		Chains: []string{"ethereum", "solana", "base"},
		Date:   *NewDateBody(&from, &to),
		Pagination: &PaginationBody{
			Page:    1,
			PerPage: 10,
		},
		//Filters
		//"filters": {
		//	"only_smart_money": true,
		//	"token_age_days": {
		//		"max": 365,
		//		"min": 1
		//	}
		//},
		Filters: map[string]any{
			"only_smart_money": true,
			"token_age_days": map[string]any{
				"max": 365,
				"min": 1,
			},
		},
		OrderBy: []SortOrderBody{
			{
				Field:     "chain",
				Direction: SortingDirectionASC,
			},
		},
	})
	if err != nil {
		t.Fatalf("GetTokenScreening() error = %v", err)
	}
	if result == nil {
		t.Fatalf("GetTokenScreening() result = %v, want non-nil", result)
	}
	if len(result.Data) == 0 {
		t.Fatalf("GetTokenScreening() result.Data = %v, want non-empty", result.Data)
	}
	t.Logf("GetTokenScreening() result.Data = %v", result.Data)
}

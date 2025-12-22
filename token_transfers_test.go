package nansen

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestGetTGMTransfers(t *testing.T) {
	key := os.Getenv("KEY")
	ctx := context.Background()
	client := NewClient(DefaultURL, key)
	to := time.Now()
	from := to.AddDate(0, 0, -7)
	result, err := client.GetTGMTransfers(ctx, GetTGMTransfersBody{
		Chain:        "tron",
		TokenAddress: "TXL6rJbvmjD46zeN1JssfgxvSo99qC8MRT",
		Date:         *NewDateBody(&from, &to),
		Pagination: &PaginationBody{
			Page:    1,
			PerPage: 10,
		},
		Filters: map[string]any{
			"include_cex":            true,
			"include_dex":            true,
			"non_exchange_transfers": true,
			"only_smart_money":       true,
		},
	})
	if err != nil {
		t.Fatalf("GetTGMTransfers() error = %v", err)
	}
	if result == nil {
		t.Fatalf("GetTGMTransfers() result = %v, want non-nil", result)
	}
	if len(result.Data) == 0 {
		t.Fatalf("GetTGMTransfers() result.Data = %v, want non-empty", result.Data)
	}
	t.Logf("GetTGMTransfers() result.Data = %v", result.Data)
}

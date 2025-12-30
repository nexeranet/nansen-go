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
	address := "AV2TbZDXQyw6h92xKfwGwb6snauMGVPe1aBDWUSiBtrH"
	now := time.Now()
	dayBack := now.AddDate(0, 0, -1)
	result, err := client.GetAddressPnL(ctx, GetAddressPnLBody{
		Chain:   "all",
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
	t.Log(result)
}

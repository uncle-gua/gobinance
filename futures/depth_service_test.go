package futures_test

import (
	"context"
	"testing"

	"github.com/uncle-gua/gobinance/futures"
)

func TestDepth(t *testing.T) {
	client := futures.NewClient("", "", false)
	resp, err := client.NewDepthService().Symbol("BTCUSDT").Limit(100).Do(context.Background())
	if err != nil {
		t.Error(t)
	}

	t.Log(resp)
}

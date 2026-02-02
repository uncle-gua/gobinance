package futures_test

import (
	"context"
	"testing"

	"github.com/uncle-gua/gobinance/futures"
)

func TestMarkPriceKline(t *testing.T) {
	client := futures.NewClient("", "", false)
	res, err := client.NewMarkPriceKlinesService().Symbol("BTCUSDT").Limit(1500).Interval("1m").Do(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Log(len(res))
	t.Log(res[0])
}

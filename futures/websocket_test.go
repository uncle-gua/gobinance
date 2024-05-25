package futures_test

import (
	"testing"
	"time"

	"github.com/uncle-gua/gobinance/futures"
)

func TestWsKline(t *testing.T) {
	if _, _, err := futures.WsKlineServe("BTCUSDT", "1m", func(event *futures.WsKlineEvent) {
		t.Log(event)
	}, func(err error) {
		t.Error(err)
	}); err != nil {
		t.Error(err)
	}
	time.Sleep(15 * time.Second)
}

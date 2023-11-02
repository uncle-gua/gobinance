package futures

import (
	"context"
	"net/http"
)

// KlinesService list klines
type KlinesService struct {
	c         *Client
	symbol    string
	interval  string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (s *KlinesService) Symbol(symbol string) *KlinesService {
	s.symbol = symbol
	return s
}

// Interval set interval
func (s *KlinesService) Interval(interval string) *KlinesService {
	s.interval = interval
	return s
}

// Limit set limit
func (s *KlinesService) Limit(limit int) *KlinesService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *KlinesService) StartTime(startTime int64) *KlinesService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *KlinesService) EndTime(endTime int64) *KlinesService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *KlinesService) Do(ctx context.Context, opts ...RequestOption) (res []*Kline, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/klines",
	}
	r.setParam("symbol", s.symbol)
	r.setParam("interval", s.interval)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(data, &res)
	return res, err
}

// Kline define kline info
type Kline struct {
	OpenTime                 int64   `json:"openTime"`
	Open                     float64 `json:"open,string"`
	High                     float64 `json:"high,string"`
	Low                      float64 `json:"low,string"`
	Close                    float64 `json:"close,string"`
	Volume                   float64 `json:"volume,string"`
	CloseTime                int64   `json:"closeTime"`
	QuoteAssetVolume         float64 `json:"quoteAssetVolume,string"`
	TradeNum                 int64   `json:"tradeNum"`
	TakerBuyBaseAssetVolume  float64 `json:"takerBuyBaseAssetVolume,string"`
	TakerBuyQuoteAssetVolume float64 `json:"takerBuyQuoteAssetVolume,string"`
}

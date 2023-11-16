package futures

import (
	"context"
	"errors"
	"net/http"

	jsoniter "github.com/json-iterator/go"
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

func (kline *Kline) UnmarshalJSON(data []byte) error {
	iter := jsoniter.Get(data)
	if iter.Size() < 11 {
		return errors.New("invalid kline response")
	}

	openTime := iter.Get(0).ToInt64()
	open := iter.Get(1).ToFloat64()
	high := iter.Get(2).ToFloat64()
	low := iter.Get(3).ToFloat64()
	close := iter.Get(4).ToFloat64()
	volume := iter.Get(5).ToFloat64()
	closeTime := iter.Get(6).ToInt64()
	quoteAssetVolume := iter.Get(7).ToFloat64()
	tradeNum := iter.Get(8).ToInt64()
	takerBuyBaseAssetVolume := iter.Get(9).ToFloat64()
	takerBuyQuoteAssetVolume := iter.Get(10).ToFloat64()

	kline.OpenTime = openTime
	kline.Open = open
	kline.High = high
	kline.Low = low
	kline.Close = close
	kline.Volume = volume
	kline.CloseTime = closeTime
	kline.QuoteAssetVolume = quoteAssetVolume
	kline.TradeNum = tradeNum
	kline.TakerBuyBaseAssetVolume = takerBuyBaseAssetVolume
	kline.TakerBuyQuoteAssetVolume = takerBuyQuoteAssetVolume

	return nil
}

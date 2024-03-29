package futures

import (
	"context"
	"net/http"

	"github.com/uncle-gua/gobinance/common"
)

// PremiumIndexService get premium index
type PremiumIndexService struct {
	c      *Client
	symbol *string
}

// Symbol set symbol
func (s *PremiumIndexService) Symbol(symbol string) *PremiumIndexService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *PremiumIndexService) Do(ctx context.Context, opts ...RequestOption) (res []*PremiumIndex, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/premiumIndex",
		secType:  secTypeNone,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	data = common.ToJSONList(data)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(data, &res)
	return res, err
}

// PremiumIndex define premium index of mark price
type PremiumIndex struct {
	Symbol          string  `json:"symbol"`
	MarkPrice       float64 `json:"markPrice,string"`
	LastFundingRate float64 `json:"lastFundingRate,string"`
	NextFundingTime int64   `json:"nextFundingTime"`
	Time            int64   `json:"time"`
}

// FundingRateService get funding rate
type FundingRateService struct {
	c         *Client
	symbol    string
	startTime *int64
	endTime   *int64
	limit     *int
}

// Symbol set symbol
func (s *FundingRateService) Symbol(symbol string) *FundingRateService {
	s.symbol = symbol
	return s
}

// StartTime set startTime
func (s *FundingRateService) StartTime(startTime int64) *FundingRateService {
	s.startTime = &startTime
	return s
}

// EndTime set startTime
func (s *FundingRateService) EndTime(endTime int64) *FundingRateService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *FundingRateService) Limit(limit int) *FundingRateService {
	s.limit = &limit
	return s
}

// Do send request
func (s *FundingRateService) Do(ctx context.Context, opts ...RequestOption) (res []*FundingRate, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/fundingRate",
		secType:  secTypeNone,
	}
	r.setParam("symbol", s.symbol)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(data, &res)
	return res, err
}

// FundingRate define funding rate of mark price
type FundingRate struct {
	Symbol      string  `json:"symbol"`
	FundingRate float64 `json:"fundingRate,string"`
	FundingTime int64   `json:"fundingTime"`
	Time        int64   `json:"time"`
}

// GetLeverageBracketService get funding rate
type GetLeverageBracketService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *GetLeverageBracketService) Symbol(symbol string) *GetLeverageBracketService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetLeverageBracketService) Do(ctx context.Context, opts ...RequestOption) (res []*LeverageBracket, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/leverageBracket",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	if s.symbol != "" {
		data = common.ToJSONList(data)
	}
	err = json.Unmarshal(data, &res)
	return res, err
}

// LeverageBracket define the leverage bracket
type LeverageBracket struct {
	Symbol   string    `json:"symbol"`
	Brackets []Bracket `json:"brackets"`
}

// Bracket define the bracket
type Bracket struct {
	Bracket          int     `json:"bracket"`
	InitialLeverage  int     `json:"initialLeverage"`
	NotionalCap      float64 `json:"notionalCap"`
	NotionalFloor    float64 `json:"notionalFloor"`
	MaintMarginRatio float64 `json:"maintMarginRatio"`
	Cum              float64 `json:"cum"`
}

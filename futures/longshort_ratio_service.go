package futures

import (
	"context"
	"net/http"
)

// LongShortRatioService list open history data of a symbol.
type LongShortRatioService struct {
	c         *Client
	symbol    string
	period    string
	limit     *int
	startTime *int64
	endTime   *int64
}

type LongShortRatio struct {
	Symbol         string  `json:"symbol"`
	LongShortRatio float64 `json:"longShortRatio,string"`
	LongAccount    float64 `json:"longAccount,string"`
	ShortAccount   float64 `json:"shortAccount,string"`
	Timestamp      int64   `json:"timestamp"`
}

// Symbol set symbol
func (s *LongShortRatioService) Symbol(symbol string) *LongShortRatioService {
	s.symbol = symbol
	return s
}

// Period set period interval
func (s *LongShortRatioService) Period(period string) *LongShortRatioService {
	s.period = period
	return s
}

// Limit set limit
func (s *LongShortRatioService) Limit(limit int) *LongShortRatioService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *LongShortRatioService) StartTime(startTime int64) *LongShortRatioService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *LongShortRatioService) EndTime(endTime int64) *LongShortRatioService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *LongShortRatioService) Do(ctx context.Context, opts ...RequestOption) (res []*LongShortRatio, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/futures/data/globalLongShortAccountRatio",
	}

	r.setParam("symbol", s.symbol)
	r.setParam("period", s.period)

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
		return []*LongShortRatio{}, err
	}

	res = make([]*LongShortRatio, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*LongShortRatio{}, err
	}

	return res, nil
}

// TopLongShortAccountRatioService list open history data of a symbol.
type TopLongShortAccountRatioService struct {
	c         *Client
	symbol    string
	period    string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (s *TopLongShortAccountRatioService) Symbol(symbol string) *TopLongShortAccountRatioService {
	s.symbol = symbol
	return s
}

// Period set period interval
func (s *TopLongShortAccountRatioService) Period(period string) *TopLongShortAccountRatioService {
	s.period = period
	return s
}

// Limit set limit
func (s *TopLongShortAccountRatioService) Limit(limit int) *TopLongShortAccountRatioService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *TopLongShortAccountRatioService) StartTime(startTime int64) *TopLongShortAccountRatioService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *TopLongShortAccountRatioService) EndTime(endTime int64) *TopLongShortAccountRatioService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *TopLongShortAccountRatioService) Do(ctx context.Context, opts ...RequestOption) (res []*LongShortRatio, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/futures/data/topLongShortAccountRatio",
	}

	r.setParam("symbol", s.symbol)
	r.setParam("period", s.period)

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
		return []*LongShortRatio{}, err
	}

	res = make([]*LongShortRatio, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*LongShortRatio{}, err
	}

	return res, nil
}

// TopLongShortAccountRatioService list open history data of a symbol.
type TopLongShortPositionRatioService struct {
	c         *Client
	symbol    string
	period    string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (s *TopLongShortPositionRatioService) Symbol(symbol string) *TopLongShortPositionRatioService {
	s.symbol = symbol
	return s
}

// Period set period interval
func (s *TopLongShortPositionRatioService) Period(period string) *TopLongShortPositionRatioService {
	s.period = period
	return s
}

// Limit set limit
func (s *TopLongShortPositionRatioService) Limit(limit int) *TopLongShortPositionRatioService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *TopLongShortPositionRatioService) StartTime(startTime int64) *TopLongShortPositionRatioService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *TopLongShortPositionRatioService) EndTime(endTime int64) *TopLongShortPositionRatioService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *TopLongShortPositionRatioService) Do(ctx context.Context, opts ...RequestOption) (res []*LongShortRatio, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/futures/data/topLongShortPositionRatio",
	}

	r.setParam("symbol", s.symbol)
	r.setParam("period", s.period)

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
		return []*LongShortRatio{}, err
	}

	res = make([]*LongShortRatio, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*LongShortRatio{}, err
	}

	return res, nil
}

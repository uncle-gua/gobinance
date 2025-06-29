package futures

import (
	"context"
	"net/http"
)

// PingService ping server
type PingService struct {
	c *Client
}

// Do send request
func (s *PingService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/ping",
	}
	_, _, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// ServerTimeService get server time
type ServerTimeService struct {
	c *Client
}

// Do send request
func (s *ServerTimeService) Do(ctx context.Context, opts ...RequestOption) (int64, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/time",
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return 0, err
	}
	var res ServerTime
	err = json.Unmarshal(data, &res)
	return res.ServerTime, err
}

type ServerTime struct {
	ServerTime int64 `json:"serverTime"`
}

// SetServerTimeService set server time
type SetServerTimeService struct {
	c *Client
}

// Do send request
func (s *SetServerTimeService) Do(ctx context.Context, opts ...RequestOption) (timeOffset int64, err error) {
	serverTime, err := s.c.NewServerTimeService().Do(ctx)
	if err != nil {
		return 0, err
	}
	timeOffset = currentTimestamp() - serverTime
	s.c.TimeOffset = timeOffset
	return timeOffset, nil
}

type TradingStatusService struct {
	c      *Client
	symbol string
}

type TradingStatus struct {
	Indicators map[string][]Violated `json:"indicators"`
	UpdateTime int64                 `json:"updateTime"`
}

type Violated struct {
	IsLocked     bool    `json:"isLocked"`
	Indicator    string  `json:"indicator"`
	Value        float64 `json:"value"`
	TriggerValue float64 `json:"triggerValue"`
	RecoverTime  int64   `json:"plannedRecoverTime"`
}

func (s *TradingStatusService) Symbol(symbol string) *TradingStatusService {
	s.symbol = symbol
	return s
}

func (s *TradingStatusService) Do(ctx context.Context, opts ...RequestOption) (TradingStatus, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/apiTradingStatus",
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}

	var res TradingStatus
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(data, &res)
	return res, err
}

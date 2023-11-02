package binance

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
		endpoint: "/api/v3/ping",
	}
	_, err = s.c.callAPI(ctx, r, opts...)
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
		endpoint: "/api/v3/time",
	}
	data, err := s.c.callAPI(ctx, r, opts...)
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

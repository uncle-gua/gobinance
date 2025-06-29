package futures

import (
	"context"
	"net/http"
)

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

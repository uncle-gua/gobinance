package futures

import (
	"context"
	"net/http"
)

// GetPositionRiskService get account balance
type GetPositionRiskService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *GetPositionRiskService) Symbol(symbol string) *GetPositionRiskService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetPositionRiskService) Do(ctx context.Context, opts ...RequestOption) (res []*PositionRisk, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v2/positionRisk",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*PositionRisk{}, err
	}
	res = make([]*PositionRisk, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*PositionRisk{}, err
	}
	return res, nil
}

// PositionRisk define position risk info
type PositionRisk struct {
	EntryPrice       float64 `json:"entryPrice,string"`
	MarginType       string  `json:"marginType"`
	IsAutoAddMargin  bool    `json:"isAutoAddMargin"`
	IsolatedMargin   float64 `json:"isolatedMargin,string"`
	Leverage         int     `json:"leverage,string"`
	LiquidationPrice float64 `json:"liquidationPrice,string"`
	MarkPrice        float64 `json:"markPrice,string"`
	MaxNotionalValue float64 `json:"maxNotionalValue,string"`
	PositionAmt      float64 `json:"positionAmt,string"`
	Symbol           string  `json:"symbol"`
	UnRealizedProfit float64 `json:"unRealizedProfit,string"`
	PositionSide     string  `json:"positionSide"`
	Notional         float64 `json:"notional,string"`
	IsolatedWallet   float64 `json:"isolatedWallet,string"`
}

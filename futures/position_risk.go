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
		return res, err
	}
	err = json.Unmarshal(data, &res)
	return res, err
}

// PositionRisk define position risk info
type PositionRisk struct {
	Symbol           string           `json:"symbol"`
	PositionSide     PositionSideType `json:"positionSide"`
	Leverage         int              `json:"leverage,string"`
	MarginType       string           `json:"marginType"`
	IsAutoAddMargin  bool             `json:"isAutoAddMargin,string"`
	PositionAmt      float64          `json:"positionAmt,string"`
	EntryPrice       float64          `json:"entryPrice,string"`
	LiquidationPrice float64          `json:"liquidationPrice,string"`
	MarkPrice        float64          `json:"markPrice,string"`
	BreakEvenPrice   float64          `json:"breakEvenPrice,string"`
	MaxNotionalValue float64          `json:"maxNotionalValue,string"`
	UnRealizedProfit float64          `json:"unRealizedProfit,string"`
	Notional         float64          `json:"notional,string"`
	IsolatedMargin   float64          `json:"isolatedMargin,string"`
	IsolatedWallet   float64          `json:"isolatedWallet,string"`
}

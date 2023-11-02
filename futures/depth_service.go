package futures

import (
	"context"
	"net/http"

	"github.com/uncle-gua/gobinance/common"
)

// DepthService show depth info
type DepthService struct {
	c      *Client
	symbol string
	limit  *int
}

// Symbol set symbol
func (s *DepthService) Symbol(symbol string) *DepthService {
	s.symbol = symbol
	return s
}

// Limit set limit
func (s *DepthService) Limit(limit int) *DepthService {
	s.limit = &limit
	return s
}

// Do send request
func (s *DepthService) Do(ctx context.Context, opts ...RequestOption) (res *DepthResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/depth",
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(DepthResponse)
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DepthResponse define depth info with bids and asks
type DepthResponse struct {
	LastUpdateID int64 `json:"lastUpdateId"`
	Time         int64 `json:"E"`
	TradeTime    int64 `json:"T"`
	Bids         []Bid `json:"bids"`
	Asks         []Ask `json:"asks"`
}

// Ask is a type alias for PriceLevel.
type Ask = common.PriceLevel

// Bid is a type alias for PriceLevel.
type Bid = common.PriceLevel

package futures

import (
	"context"
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// CreateAlgoOrderService create order
type CreateAlgoOrderService struct {
	c                       *Client
	symbol                  string
	side                    SideType
	positionSide            *PositionSideType
	orderType               OrderType
	timeInForce             *TimeInForceType
	quantity                string
	price                   *string
	triggerPrice            *string
	workingType             *WorkingType
	priceMatch              *PriceMatchType
	closePosition           *bool
	priceProtect            *bool
	reduceOnly              *bool
	activatePrice           *string
	callbackRate            *string
	clientAlgoId            *string
	newOrderRespType        NewOrderRespType
	selfTradePreventionMode *SelfTradePreventionModeType
	goodTillDate            int64
}

// Symbol set symbol
func (s *CreateAlgoOrderService) Symbol(symbol string) *CreateAlgoOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CreateAlgoOrderService) Side(side SideType) *CreateAlgoOrderService {
	s.side = side
	return s
}

// PositionSide set side
func (s *CreateAlgoOrderService) PositionSide(positionSide PositionSideType) *CreateAlgoOrderService {
	s.positionSide = &positionSide
	return s
}

// Type set type
func (s *CreateAlgoOrderService) Type(orderType OrderType) *CreateAlgoOrderService {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *CreateAlgoOrderService) TimeInForce(timeInForce TimeInForceType) *CreateAlgoOrderService {
	s.timeInForce = &timeInForce
	return s
}

// GoodTillDate set goodTillDate
func (s *CreateAlgoOrderService) GoodTillDate(goodTillDate int64) *CreateAlgoOrderService {
	s.goodTillDate = goodTillDate
	return s
}

// Quantity set quantity
func (s *CreateAlgoOrderService) Quantity(quantity string) *CreateAlgoOrderService {
	s.quantity = quantity
	return s
}

// ReduceOnly set reduceOnly
func (s *CreateAlgoOrderService) ReduceOnly(reduceOnly bool) *CreateAlgoOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

// Price set price
func (s *CreateAlgoOrderService) Price(price string) *CreateAlgoOrderService {
	s.price = &price
	return s
}

// NewClientOrderID set newClientOrderID
func (s *CreateAlgoOrderService) ClientAlgoId(clientAlgoId string) *CreateAlgoOrderService {
	s.clientAlgoId = &clientAlgoId
	return s
}

// ActivatePrice set activatePrice
func (s *CreateAlgoOrderService) ActivatePrice(activatePrice string) *CreateAlgoOrderService {
	s.activatePrice = &activatePrice
	return s
}

// ActivatePrice set activatePrice
func (s *CreateAlgoOrderService) TriggerPrice(triggerPrice string) *CreateAlgoOrderService {
	s.triggerPrice = &triggerPrice
	return s
}

// WorkingType set workingType
func (s *CreateAlgoOrderService) WorkingType(workingType WorkingType) *CreateAlgoOrderService {
	s.workingType = &workingType
	return s
}

// CallbackRate set callbackRate
func (s *CreateAlgoOrderService) CallbackRate(callbackRate string) *CreateAlgoOrderService {
	s.callbackRate = &callbackRate
	return s
}

// PriceProtect set priceProtect
func (s *CreateAlgoOrderService) PriceProtect(priceProtect bool) *CreateAlgoOrderService {
	s.priceProtect = &priceProtect
	return s
}

// PriceMatch set priceMatch
func (s *CreateAlgoOrderService) PriceMatch(priceMatch PriceMatchType) *CreateAlgoOrderService {
	s.priceMatch = &priceMatch
	return s
}

// SelfTradePreventionMode set selfTradePreventionMode
func (s *CreateAlgoOrderService) SelfTradePreventionMode(selfTradePreventionMode SelfTradePreventionModeType) *CreateAlgoOrderService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

// NewOrderResponseType set newOrderResponseType
func (s *CreateAlgoOrderService) NewOrderResponseType(newOrderResponseType NewOrderRespType) *CreateAlgoOrderService {
	s.newOrderRespType = newOrderResponseType
	return s
}

// ClosePosition set closePosition
func (s *CreateAlgoOrderService) ClosePosition(closePosition bool) *CreateAlgoOrderService {
	s.closePosition = &closePosition
	return s
}

func (s *CreateAlgoOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, header *http.Header, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"algoType":         "CONDITIONAL",
		"symbol":           s.symbol,
		"side":             s.side,
		"type":             s.orderType,
		"quantity":         s.quantity,
		"newOrderRespType": s.newOrderRespType,
	}
	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.goodTillDate > 0 {
		m["goodTillDate"] = s.goodTillDate
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.clientAlgoId != nil {
		m["clientAlgoId"] = *s.clientAlgoId
	} else {
		pre := "x-dNUwr2u2"
		rnd := strings.ReplaceAll(fmt.Sprintf("%8x", rand.Uint32()), " ", "0")
		tim := strconv.FormatInt(time.Now().UTC().UnixNano(), 36)
		m["clientAlgoId"] = fmt.Sprintf("%s%s%s", pre, tim, rnd)
	}
	if s.triggerPrice != nil {
		m["triggerPrice"] = *s.triggerPrice
	}
	if s.activatePrice != nil {
		m["activatePrice"] = *s.activatePrice
	}
	if s.workingType != nil {
		m["workingType"] = *s.workingType
	}
	if s.priceProtect != nil {
		m["priceProtect"] = *s.priceProtect
	}
	if s.priceMatch != nil {
		m["priceMatch"] = *s.priceMatch
	}
	if s.selfTradePreventionMode != nil {
		m["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}
	if s.callbackRate != nil {
		m["callbackRate"] = *s.callbackRate
	}
	if s.closePosition != nil {
		m["closePosition"] = *s.closePosition
	}
	r.setFormParams(m)
	data, header, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	return data, header, nil
}

// Do send request
func (s *CreateAlgoOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateOrderResponse, err error) {
	data, header, err := s.createOrder(ctx, "/fapi/v1/algoOrder", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateOrderResponse)
	err = json.Unmarshal(data, res)
	res.RateLimitOrder10s = header.Get("X-Mbx-Order-Count-10s")
	res.RateLimitOrder1m = header.Get("X-Mbx-Order-Count-1m")

	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateOrderResponse define create order response
type CreateAlgoOrderResponse struct {
	AlgoID                  int64                       `json:"algoId"`
	ClientAlgoId            string                      `json:"clientAlgoId"`
	AlgoType                string                      `json:"algoType"`
	Symbol                  string                      `json:"symbol"`
	Side                    SideType                    `json:"side"`
	PositionSide            PositionSideType            `json:"positionSide"`
	TimeInForce             TimeInForceType             `json:"timeInForce"`
	Quantity                float64                     `json:"quantity,string"`
	AlgoStatus              string                      `json:"algoStatus"`
	TriggerPrice            float64                     `json:"triggerPrice,string"`
	Price                   float64                     `json:"price,string"`
	SelfTradePreventionMode SelfTradePreventionModeType `json:"selfTradePreventionMode"`
	WorkingType             WorkingType                 `json:"workingType"`
	PriceMatch              PriceMatchType              `json:"priceMatch"`
	ClosePosition           bool                        `json:"closePosition"`
	PriceProtect            bool                        `json:"priceProtect"`
	ReduceOnly              bool                        `json:"reduceOnly"`
	ActivatePrice           float64                     `json:"activatePrice,string"`
	CallbackRate            float64                     `json:"callbackRate,string"`
	CreateTime              int64                       `json:"createTime"`
	UpdateTime              int64                       `json:"updateTime"`
	TriggerTime             int64                       `json:"triggerTime"`
	GoodTillDate            int64                       `json:"goodTillDate"`
	RateLimitOrder10s       string                      `json:"rateLimitOrder10s,omitempty"`
	RateLimitOrder1m        string                      `json:"rateLimitOrder1m,omitempty"`
}

// ListOpenOrdersService list opened orders
type ListAlgoOpenOrdersService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *ListAlgoOpenOrdersService) Symbol(symbol string) *ListAlgoOpenOrdersService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *ListAlgoOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*AlgoOrder, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/openAlgoOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*AlgoOrder{}, err
	}
	res = make([]*AlgoOrder, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*AlgoOrder{}, err
	}
	return res, nil
}

type AlgoOrder struct {
	AlgoID                  int64                       `json:"algoId"`
	ClientAlgoId            string                      `json:"clientAlgoId"`
	AlgoType                string                      `json:"algoType"`
	Symbol                  string                      `json:"symbol"`
	Side                    SideType                    `json:"side"`
	PositionSide            PositionSideType            `json:"positionSide"`
	OrderType               OrderType                   `json:"orderType"`
	TimeInForce             TimeInForceType             `json:"timeInForce"`
	Quantity                float64                     `json:"quantity,string"`
	AlgoStatus              string                      `json:"algoStatus"`
	TriggerPrice            float64                     `json:"triggerPrice,string"`
	Price                   float64                     `json:"price,string"`
	SelfTradePreventionMode SelfTradePreventionModeType `json:"selfTradePreventionMode"`
	WorkingType             WorkingType                 `json:"workingType"`
	PriceMatch              PriceMatchType              `json:"priceMatch"`
	ClosePosition           bool                        `json:"closePosition"`
	PriceProtect            bool                        `json:"priceProtect"`
	ReduceOnly              bool                        `json:"reduceOnly"`
	ActivatePrice           float64                     `json:"activatePrice,string"`
	CallbackRate            float64                     `json:"callbackRate,string"`
	CreateTime              int64                       `json:"createTime"`
	UpdateTime              int64                       `json:"updateTime"`
	TriggerTime             int64                       `json:"triggerTime"`
	GoodTillDate            int64                       `json:"goodTillDate"`
	RateLimitOrder10s       string                      `json:"rateLimitOrder10s,omitempty"`
	RateLimitOrder1m        string                      `json:"rateLimitOrder1m,omitempty"`
}

// CancelOrderService cancel an order
type CancelAlgoOrderService struct {
	c            *Client
	symbol       string
	algoId       *int64
	clientAlgoId *string
}

// Symbol set symbol
func (s *CancelAlgoOrderService) Symbol(symbol string) *CancelAlgoOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CancelAlgoOrderService) AlgoId(algoId int64) *CancelAlgoOrderService {
	s.algoId = &algoId
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *CancelAlgoOrderService) ClientAlgoId(clientAlgoId string) *CancelAlgoOrderService {
	s.clientAlgoId = &clientAlgoId
	return s
}

// Do send request
func (s *CancelAlgoOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelAlgoOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/fapi/v1/algoOrder",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	if s.algoId != nil {
		r.setFormParam("algoId", *s.algoId)
	}
	if s.clientAlgoId != nil {
		r.setFormParam("clientAlgoId", *s.clientAlgoId)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelAlgoOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CancelOrderResponse define response of canceling order
type CancelAlgoOrderResponse struct {
	Code         int64  `json:"code,string"`
	Msg          string `json:"msg"`
	AlgoId       int64  `json:"algoId"`
	ClientAlgoId string `json:"clientAlgoId"`
}

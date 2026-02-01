package futures

import (
	"context"
	"errors"
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
	goodTillDate            int64
	quantity                *string
	reduceOnly              *bool
	price                   *string
	newClientOrderID        *string
	stopPrice               *string
	workingType             *WorkingType
	activationPrice         *string
	callbackRate            *string
	priceProtect            *bool
	priceMatch              *PriceMatchType
	selfTradePreventionMode *SelfTradePreventionModeType
	newOrderRespType        NewOrderRespType
	closePosition           *bool
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
	s.quantity = &quantity
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
func (s *CreateAlgoOrderService) NewClientOrderID(newClientOrderID string) *CreateAlgoOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// StopPrice set stopPrice
func (s *CreateAlgoOrderService) StopPrice(stopPrice string) *CreateAlgoOrderService {
	s.stopPrice = &stopPrice
	return s
}

// WorkingType set workingType
func (s *CreateAlgoOrderService) WorkingType(workingType WorkingType) *CreateAlgoOrderService {
	s.workingType = &workingType
	return s
}

// ActivationPrice set activationPrice
func (s *CreateAlgoOrderService) ActivationPrice(activationPrice string) *CreateAlgoOrderService {
	s.activationPrice = &activationPrice
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
	if s.quantity != nil {
		m["quantity"] = *s.quantity
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.newClientOrderID != nil {
		m["newClientOrderId"] = *s.newClientOrderID
	} else {
		pre := "x-dNUwr2u2"
		rnd := strings.ReplaceAll(fmt.Sprintf("%8x", rand.Uint32()), " ", "0")
		tim := strconv.FormatInt(time.Now().UTC().UnixNano(), 36)
		m["newClientOrderId"] = fmt.Sprintf("%s%s%s", pre, tim, rnd)
	}
	if s.stopPrice != nil {
		m["stopPrice"] = *s.stopPrice
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
	if s.activationPrice != nil {
		m["activationPrice"] = *s.activationPrice
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
func (s *CreateAlgoOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateAlgoOrderResponse, err error) {
	data, header, err := s.createOrder(ctx, "/fapi/v1/algoOrder", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateAlgoOrderResponse)
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
	Symbol                  string                      `json:"symbol"`
	OrderID                 int64                       `json:"orderId"`
	ClientOrderID           string                      `json:"clientOrderId"`
	Price                   float64                     `json:"price,string"`
	OrigQuantity            float64                     `json:"origQty,string"`
	ExecutedQuantity        float64                     `json:"executedQty,string"`
	CumQuote                float64                     `json:"cumQuote,string"`
	ReduceOnly              bool                        `json:"reduceOnly"`
	Status                  OrderStatusType             `json:"status"`
	StopPrice               float64                     `json:"stopPrice,string"`
	TimeInForce             TimeInForceType             `json:"timeInForce"`
	Type                    OrderType                   `json:"type"`
	OrigType                OrderType                   `json:"origType"`
	Side                    SideType                    `json:"side"`
	UpdateTime              int64                       `json:"updateTime"`
	WorkingType             WorkingType                 `json:"workingType"`
	ActivatePrice           float64                     `json:"activatePrice,string"`
	PriceRate               float64                     `json:"priceRate,string"`
	AvgPrice                float64                     `json:"avgPrice,string"`
	PositionSide            PositionSideType            `json:"positionSide"`
	ClosePosition           bool                        `json:"closePosition"`
	PriceProtect            bool                        `json:"priceProtect"`
	PriceMatch              PriceMatchType              `json:"priceMatch"`
	SelfTradePreventionMode SelfTradePreventionModeType `json:"selfTradePreventionMode"`
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
func (s *ListAlgoOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*Order, err error) {
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
		return []*Order{}, err
	}
	res = make([]*Order, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Order{}, err
	}
	return res, nil
}

// GetOpenOrderService query current open order
type GetAlgoOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

func (s *GetAlgoOrderService) Symbol(symbol string) *GetAlgoOrderService {
	s.symbol = symbol
	return s
}

func (s *GetAlgoOrderService) OrderID(orderID int64) *GetAlgoOrderService {
	s.orderID = &orderID
	return s
}

func (s *GetAlgoOrderService) OrigClientOrderID(origClientOrderID string) *GetAlgoOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

func (s *GetAlgoOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/openOrder",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID == nil && s.origClientOrderID == nil {
		return nil, errors.New("either orderId or origClientOrderId must be sent")
	}
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(Order)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Order define order info
type AlgoOrder struct {
	Symbol                  string                      `json:"symbol"`
	OrderID                 int64                       `json:"orderId"`
	ClientOrderID           string                      `json:"clientOrderId"`
	Price                   float64                     `json:"price,string"`
	ReduceOnly              bool                        `json:"reduceOnly"`
	OrigQuantity            float64                     `json:"origQty,string"`
	ExecutedQuantity        float64                     `json:"executedQty,string"`
	CumQuantity             float64                     `json:"cumQty,string"`
	CumQuote                float64                     `json:"cumQuote,string"`
	Status                  OrderStatusType             `json:"status"`
	TimeInForce             TimeInForceType             `json:"timeInForce"`
	GoodTillDate            int64                       `json:"goodTillDate"`
	Type                    OrderType                   `json:"type"`
	OrigType                OrderType                   `json:"origType"`
	Side                    SideType                    `json:"side"`
	StopPrice               float64                     `json:"stopPrice,string"`
	Time                    int64                       `json:"time"`
	UpdateTime              int64                       `json:"updateTime"`
	WorkingType             WorkingType                 `json:"workingType"`
	ActivatePrice           float64                     `json:"activatePrice,string"`
	PriceRate               float64                     `json:"priceRate,string"`
	AvgPrice                float64                     `json:"avgPrice,string"`
	PositionSide            PositionSideType            `json:"positionSide"`
	PriceProtect            bool                        `json:"priceProtect"`
	ClosePosition           bool                        `json:"closePosition"`
	PriceMatch              PriceMatchType              `json:"priceMatch"`
	SelfTradePreventionMode SelfTradePreventionModeType `json:"selfTradePreventionMode"`
}

// ListOrdersService all account orders; active, canceled, or filled
type ListAlgoOrdersService struct {
	c         *Client
	symbol    string
	orderID   *int64
	startTime *int64
	endTime   *int64
	limit     *int
}

// Symbol set symbol
func (s *ListAlgoOrdersService) Symbol(symbol string) *ListAlgoOrdersService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *ListAlgoOrdersService) OrderID(orderID int64) *ListAlgoOrdersService {
	s.orderID = &orderID
	return s
}

// StartTime set starttime
func (s *ListAlgoOrdersService) StartTime(startTime int64) *ListAlgoOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *ListAlgoOrdersService) EndTime(endTime int64) *ListAlgoOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListAlgoOrdersService) Limit(limit int) *ListAlgoOrdersService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ListAlgoOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*AlgoOrder, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/allOrders",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
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

// CancelOrderService cancel an order
type CancelAlgoOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

// Symbol set symbol
func (s *CancelAlgoOrderService) Symbol(symbol string) *CancelAlgoOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CancelAlgoOrderService) OrderID(orderID int64) *CancelAlgoOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *CancelAlgoOrderService) OrigClientOrderID(origClientOrderID string) *CancelAlgoOrderService {
	s.origClientOrderID = &origClientOrderID
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
	if s.orderID != nil {
		r.setFormParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setFormParam("origClientOrderId", *s.origClientOrderID)
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

// CancelAlgoOrderResponse define response of canceling order
type CancelAlgoOrderResponse struct {
	ClientOrderID    string           `json:"clientOrderId"`
	CumQuantity      float64          `json:"cumQty,string"`
	CumQuote         float64          `json:"cumQuote,string"`
	ExecutedQuantity float64          `json:"executedQty,string"`
	OrderID          int64            `json:"orderId"`
	OrigQuantity     float64          `json:"origQty,string"`
	Price            float64          `json:"price,string"`
	ReduceOnly       bool             `json:"reduceOnly"`
	Side             SideType         `json:"side"`
	Status           OrderStatusType  `json:"status"`
	StopPrice        float64          `json:"stopPrice,string"`
	Symbol           string           `json:"symbol"`
	TimeInForce      TimeInForceType  `json:"timeInForce"`
	Type             OrderType        `json:"type"`
	OrigType         OrderType        `json:"origType"`
	UpdateTime       int64            `json:"updateTime"`
	WorkingType      WorkingType      `json:"workingType"`
	ActivatePrice    float64          `json:"activatePrice,string"`
	PriceRate        float64          `json:"priceRate,string"`
	PositionSide     PositionSideType `json:"positionSide"`
	PriceProtect     bool             `json:"priceProtect"`
}

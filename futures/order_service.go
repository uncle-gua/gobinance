package futures

import (
	"context"
	stdjson "encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// CreateOrderService create order
type CreateOrderService struct {
	c                       *Client
	symbol                  string
	side                    SideType
	positionSide            *PositionSideType
	orderType               OrderType
	timeInForce             *TimeInForceType
	goodTillDate            int64
	quantity                string
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
func (s *CreateOrderService) Symbol(symbol string) *CreateOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CreateOrderService) Side(side SideType) *CreateOrderService {
	s.side = side
	return s
}

// PositionSide set side
func (s *CreateOrderService) PositionSide(positionSide PositionSideType) *CreateOrderService {
	s.positionSide = &positionSide
	return s
}

// Type set type
func (s *CreateOrderService) Type(orderType OrderType) *CreateOrderService {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *CreateOrderService) TimeInForce(timeInForce TimeInForceType) *CreateOrderService {
	s.timeInForce = &timeInForce
	return s
}

// GoodTillDate set goodTillDate
func (s *CreateOrderService) GoodTillDate(goodTillDate int64) *CreateOrderService {
	s.goodTillDate = goodTillDate
	return s
}

// Quantity set quantity
func (s *CreateOrderService) Quantity(quantity string) *CreateOrderService {
	s.quantity = quantity
	return s
}

// ReduceOnly set reduceOnly
func (s *CreateOrderService) ReduceOnly(reduceOnly bool) *CreateOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

// Price set price
func (s *CreateOrderService) Price(price string) *CreateOrderService {
	s.price = &price
	return s
}

// NewClientOrderID set newClientOrderID
func (s *CreateOrderService) NewClientOrderID(newClientOrderID string) *CreateOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// StopPrice set stopPrice
func (s *CreateOrderService) StopPrice(stopPrice string) *CreateOrderService {
	s.stopPrice = &stopPrice
	return s
}

// WorkingType set workingType
func (s *CreateOrderService) WorkingType(workingType WorkingType) *CreateOrderService {
	s.workingType = &workingType
	return s
}

// ActivationPrice set activationPrice
func (s *CreateOrderService) ActivationPrice(activationPrice string) *CreateOrderService {
	s.activationPrice = &activationPrice
	return s
}

// CallbackRate set callbackRate
func (s *CreateOrderService) CallbackRate(callbackRate string) *CreateOrderService {
	s.callbackRate = &callbackRate
	return s
}

// PriceProtect set priceProtect
func (s *CreateOrderService) PriceProtect(priceProtect bool) *CreateOrderService {
	s.priceProtect = &priceProtect
	return s
}

// PriceMatch set priceMatch
func (s *CreateOrderService) PriceMatch(priceMatch PriceMatchType) *CreateOrderService {
	s.priceMatch = &priceMatch
	return s
}

// SelfTradePreventionMode set selfTradePreventionMode
func (s *CreateOrderService) SelfTradePreventionMode(selfTradePreventionMode SelfTradePreventionModeType) *CreateOrderService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

// NewOrderResponseType set newOrderResponseType
func (s *CreateOrderService) NewOrderResponseType(newOrderResponseType NewOrderRespType) *CreateOrderService {
	s.newOrderRespType = newOrderResponseType
	return s
}

// ClosePosition set closePosition
func (s *CreateOrderService) ClosePosition(closePosition bool) *CreateOrderService {
	s.closePosition = &closePosition
	return s
}

func (s *CreateOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, header *http.Header, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
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
func (s *CreateOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateOrderResponse, err error) {
	data, header, err := s.createOrder(ctx, "/fapi/v1/order", opts...)
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
type CreateOrderResponse struct {
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

// AmendOrderService amend order
type AmendOrderService struct {
	c                 *Client
	origClientOrderID *string
	orderId           int64
	symbol            string
	side              SideType
	quantity          string
	price             *string
	priceMatch        *PriceMatchType
}

// OrigClientOrderID set origClientOrderID
func (s *AmendOrderService) OrigClientOrderID(origClientOrderID string) *AmendOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// OrderId set orderId
func (s *AmendOrderService) OrderId(orderId int64) *AmendOrderService {
	s.orderId = orderId
	return s
}

// Symbol set symbol
func (s *AmendOrderService) Symbol(symbol string) *AmendOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *AmendOrderService) Side(side SideType) *AmendOrderService {
	s.side = side
	return s
}

// Price set price
func (s *AmendOrderService) Price(price string) *AmendOrderService {
	s.price = &price
	return s
}

// PriceMatch set priceMatch mode
func (s *AmendOrderService) PriceMatch(priceMatch PriceMatchType) *AmendOrderService {
	s.priceMatch = &priceMatch
	return s
}

// PriceMatch set priceMatch mode
func (s *AmendOrderService) Quantity(quantity string) *AmendOrderService {
	s.quantity = quantity
	return s
}

func (s *AmendOrderService) amendOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, header *http.Header, err error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":   s.symbol,
		"side":     s.side,
		"quantity": s.quantity,
	}
	if s.origClientOrderID != nil {
		m["origClientOrderId"] = *s.origClientOrderID
	}
	if s.orderId > 0 {
		m["orderId"] = s.orderId
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.priceMatch != nil {
		m["priceMatch"] = *s.priceMatch
	}
	r.setFormParams(m)
	data, header, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	return data, header, nil
}

// Do send request
func (s *AmendOrderService) Do(ctx context.Context, opts ...RequestOption) (res *AmendOrderResponse, err error) {
	data, header, err := s.amendOrder(ctx, "/fapi/v1/order", opts...)
	if err != nil {
		return nil, err
	}
	res = new(AmendOrderResponse)
	err = json.Unmarshal(data, res)
	res.RateLimitOrder10s = header.Get("X-Mbx-Order-Count-10s")
	res.RateLimitOrder1m = header.Get("X-Mbx-Order-Count-1m")

	if err != nil {
		return nil, err
	}
	return res, nil
}

// AmendOrderResponse define amend order response
type AmendOrderResponse struct {
	OrderID                 int64                       `json:"orderId"`
	Symbol                  string                      `json:"symbol"`
	Pair                    string                      `json:"pair"`
	Status                  OrderStatusType             `json:"status"`
	ClientOrderID           string                      `json:"clientOrderId"`
	Price                   float64                     `json:"price,string"`
	AvgPrice                float64                     `json:"avgPrice,string"`
	OrigQuantity            float64                     `json:"origQty,string"`
	ExecutedQuantity        float64                     `json:"executedQty,string"`
	CumQty                  float64                     `json:"cumQty,string"`
	CumBase                 float64                     `json:"cumBase,string"`
	TimeInForce             TimeInForceType             `json:"timeInForce"`
	Type                    OrderType                   `json:"type"`
	ReduceOnly              bool                        `json:"reduceOnly"`
	ClosePosition           bool                        `json:"closePosition"`
	Side                    SideType                    `json:"side"`
	PositionSide            PositionSideType            `json:"positionSide"`
	StopPrice               float64                     `json:"stopPrice,string"`
	WorkingType             WorkingType                 `json:"workingType"`
	PriceProtect            bool                        `json:"priceProtect"`
	OrigType                OrderType                   `json:"origType"`
	PriceMatch              PriceMatchType              `json:"priceMatch"`
	SelfTradePreventionMode SelfTradePreventionModeType `json:"selfTradePreventionMode"`
	GoodTillDate            int64                       `json:"goodTillDate"`
	UpdateTime              int64                       `json:"updateTime"`
	RateLimitOrder10s       string                      `json:"rateLimitOrder10s,omitempty"`
	RateLimitOrder1m        string                      `json:"rateLimitOrder1m,omitempty"`
}

// ListOpenOrdersService list opened orders
type ListOpenOrdersService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *ListOpenOrdersService) Symbol(symbol string) *ListOpenOrdersService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *ListOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/openOrders",
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
type GetOpenOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

func (s *GetOpenOrderService) Symbol(symbol string) *GetOpenOrderService {
	s.symbol = symbol
	return s
}

func (s *GetOpenOrderService) OrderID(orderID int64) *GetOpenOrderService {
	s.orderID = &orderID
	return s
}

func (s *GetOpenOrderService) OrigClientOrderID(origClientOrderID string) *GetOpenOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

func (s *GetOpenOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
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

// GetOrderService get an order
type GetOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

// Symbol set symbol
func (s *GetOrderService) Symbol(symbol string) *GetOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *GetOrderService) OrderID(orderID int64) *GetOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *GetOrderService) OrigClientOrderID(origClientOrderID string) *GetOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// Do send request
func (s *GetOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/order",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
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
type Order struct {
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
	Side                    SideType                    `json:"side"`
	StopPrice               float64                     `json:"stopPrice,string"`
	Time                    int64                       `json:"time"`
	UpdateTime              int64                       `json:"updateTime"`
	WorkingType             WorkingType                 `json:"workingType"`
	ActivatePrice           float64                     `json:"activatePrice,string"`
	PriceRate               float64                     `json:"priceRate,string"`
	AvgPrice                float64                     `json:"avgPrice,string"`
	OrigType                OrderType                   `json:"origType"`
	PositionSide            PositionSideType            `json:"positionSide"`
	PriceProtect            bool                        `json:"priceProtect"`
	ClosePosition           bool                        `json:"closePosition"`
	PriceMatch              PriceMatchType              `json:"priceMatch"`
	SelfTradePreventionMode SelfTradePreventionModeType `json:"selfTradePreventionMode"`
}

// ListOrdersService all account orders; active, canceled, or filled
type ListOrdersService struct {
	c         *Client
	symbol    string
	orderID   *int64
	startTime *int64
	endTime   *int64
	limit     *int
}

// Symbol set symbol
func (s *ListOrdersService) Symbol(symbol string) *ListOrdersService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *ListOrdersService) OrderID(orderID int64) *ListOrdersService {
	s.orderID = &orderID
	return s
}

// StartTime set starttime
func (s *ListOrdersService) StartTime(startTime int64) *ListOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *ListOrdersService) EndTime(endTime int64) *ListOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListOrdersService) Limit(limit int) *ListOrdersService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ListOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*Order, err error) {
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
type CancelOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

// Symbol set symbol
func (s *CancelOrderService) Symbol(symbol string) *CancelOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CancelOrderService) OrderID(orderID int64) *CancelOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *CancelOrderService) OrigClientOrderID(origClientOrderID string) *CancelOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// Do send request
func (s *CancelOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/fapi/v1/order",
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
	res = new(CancelOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CancelOrderResponse define response of canceling order
type CancelOrderResponse struct {
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
	UpdateTime       int64            `json:"updateTime"`
	WorkingType      WorkingType      `json:"workingType"`
	ActivatePrice    float64          `json:"activatePrice,string"`
	PriceRate        float64          `json:"priceRate,string"`
	OrigType         string           `json:"origType"`
	PositionSide     PositionSideType `json:"positionSide"`
	PriceProtect     bool             `json:"priceProtect"`
}

// CancelAllOpenOrdersService cancel all open orders
type CancelAllOpenOrdersService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *CancelAllOpenOrdersService) Symbol(symbol string) *CancelAllOpenOrdersService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *CancelAllOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/fapi/v1/allOpenOrders",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	_, _, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CancelMultiplesOrdersService cancel a list of orders
type CancelMultiplesOrdersService struct {
	c                     *Client
	symbol                string
	orderIDList           []int64
	origClientOrderIDList []string
}

// Symbol set symbol
func (s *CancelMultiplesOrdersService) Symbol(symbol string) *CancelMultiplesOrdersService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CancelMultiplesOrdersService) OrderIDList(orderIDList []int64) *CancelMultiplesOrdersService {
	s.orderIDList = orderIDList
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *CancelMultiplesOrdersService) OrigClientOrderIDList(origClientOrderIDList []string) *CancelMultiplesOrdersService {
	s.origClientOrderIDList = origClientOrderIDList
	return s
}

// Do send request
func (s *CancelMultiplesOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*CancelOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/fapi/v1/batchOrders",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	if s.orderIDList != nil {
		// convert a slice of integers to a string e.g. [1 2 3] => "[1,2,3]"
		orderIDListString := strings.Join(strings.Fields(fmt.Sprint(s.orderIDList)), ",")
		r.setFormParam("orderIdList", orderIDListString)
	}
	if s.origClientOrderIDList != nil {
		r.setFormParam("origClientOrderIdList", s.origClientOrderIDList)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(data, &res)
	return res, err
}

// ListLiquidationOrdersService list liquidation orders
type ListLiquidationOrdersService struct {
	c         *Client
	symbol    *string
	startTime *int64
	endTime   *int64
	limit     *int
}

// Symbol set symbol
func (s *ListLiquidationOrdersService) Symbol(symbol string) *ListLiquidationOrdersService {
	s.symbol = &symbol
	return s
}

// StartTime set startTime
func (s *ListLiquidationOrdersService) StartTime(startTime int64) *ListLiquidationOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set startTime
func (s *ListLiquidationOrdersService) EndTime(endTime int64) *ListLiquidationOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListLiquidationOrdersService) Limit(limit int) *ListLiquidationOrdersService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ListLiquidationOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*LiquidationOrder, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/allForceOrders",
		secType:  secTypeNone,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
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
		return []*LiquidationOrder{}, err
	}
	res = make([]*LiquidationOrder, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*LiquidationOrder{}, err
	}
	return res, nil
}

// LiquidationOrder define liquidation order
type LiquidationOrder struct {
	Symbol           string          `json:"symbol"`
	Price            float64         `json:"price,string"`
	OrigQuantity     float64         `json:"origQty,string"`
	ExecutedQuantity float64         `json:"executedQty,string"`
	AveragePrice     float64         `json:"avragePrice,string"`
	Status           OrderStatusType `json:"status"`
	TimeInForce      TimeInForceType `json:"timeInForce"`
	Type             OrderType       `json:"type"`
	Side             SideType        `json:"side"`
	Time             int64           `json:"time"`
}

// ListUserLiquidationOrdersService lists user's liquidation orders
type ListUserLiquidationOrdersService struct {
	c             *Client
	symbol        *string
	autoCloseType ForceOrderCloseType
	startTime     *int64
	endTime       *int64
	limit         *int
}

// Symbol set symbol
func (s *ListUserLiquidationOrdersService) Symbol(symbol string) *ListUserLiquidationOrdersService {
	s.symbol = &symbol
	return s
}

// AutoCloseType set symbol
func (s *ListUserLiquidationOrdersService) AutoCloseType(autoCloseType ForceOrderCloseType) *ListUserLiquidationOrdersService {
	s.autoCloseType = autoCloseType
	return s
}

// StartTime set startTime
func (s *ListUserLiquidationOrdersService) StartTime(startTime int64) *ListUserLiquidationOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ListUserLiquidationOrdersService) EndTime(endTime int64) *ListUserLiquidationOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListUserLiquidationOrdersService) Limit(limit int) *ListUserLiquidationOrdersService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ListUserLiquidationOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*UserLiquidationOrder, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/forceOrders",
		secType:  secTypeSigned,
	}

	r.setParam("autoCloseType", s.autoCloseType)
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
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
		return []*UserLiquidationOrder{}, err
	}
	res = make([]*UserLiquidationOrder, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*UserLiquidationOrder{}, err
	}
	return res, nil
}

// UserLiquidationOrder defines user's liquidation order
type UserLiquidationOrder struct {
	OrderId          int64            `json:"orderId"`
	Symbol           string           `json:"symbol"`
	Status           OrderStatusType  `json:"status"`
	ClientOrderId    string           `json:"clientOrderId"`
	Price            float64          `json:"price,string"`
	AveragePrice     float64          `json:"avgPrice,string"`
	OrigQuantity     float64          `json:"origQty,string"`
	ExecutedQuantity float64          `json:"executedQty,string"`
	CumQuote         float64          `json:"cumQuote,string"`
	TimeInForce      TimeInForceType  `json:"timeInForce"`
	Type             OrderType        `json:"type"`
	ReduceOnly       bool             `json:"reduceOnly"`
	ClosePosition    bool             `json:"closePosition"`
	Side             SideType         `json:"side"`
	PositionSide     PositionSideType `json:"positionSide"`
	StopPrice        float64          `json:"stopPrice,string"`
	WorkingType      WorkingType      `json:"workingType"`
	OrigType         float64          `json:"origType,string"`
	Time             int64            `json:"time"`
	UpdateTime       int64            `json:"updateTime"`
}

type CreateBatchOrdersService struct {
	c      *Client
	orders []*CreateOrderService
}

type CreateBatchOrdersResponse struct {
	Orders []*Order
}

func (s *CreateBatchOrdersService) OrderList(orders []*CreateOrderService) *CreateBatchOrdersService {
	s.orders = orders
	return s
}

func (s *CreateBatchOrdersService) Do(ctx context.Context, opts ...RequestOption) (res *CreateBatchOrdersResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/fapi/v1/batchOrders",
		secType:  secTypeSigned,
	}

	orders := []params{}
	for _, order := range s.orders {
		m := params{
			"symbol":           order.symbol,
			"side":             order.side,
			"type":             order.orderType,
			"quantity":         order.quantity,
			"newOrderRespType": order.newOrderRespType,
		}

		if order.positionSide != nil {
			m["positionSide"] = *order.positionSide
		}
		if order.timeInForce != nil {
			m["timeInForce"] = *order.timeInForce
		}
		if order.reduceOnly != nil {
			m["reduceOnly"] = *order.reduceOnly
		}
		if order.price != nil {
			m["price"] = *order.price
		}
		if order.newClientOrderID != nil {
			m["newClientOrderId"] = *order.newClientOrderID
		}
		if order.stopPrice != nil {
			m["stopPrice"] = *order.stopPrice
		}
		if order.workingType != nil {
			m["workingType"] = *order.workingType
		}
		if order.priceProtect != nil {
			m["priceProtect"] = *order.priceProtect
		}
		if order.priceMatch != nil {
			m["priceMatch"] = *order.priceMatch
		}
		if order.selfTradePreventionMode != nil {
			m["selfTradePreventionMode"] = *order.selfTradePreventionMode
		}
		if order.activationPrice != nil {
			m["activationPrice"] = *order.activationPrice
		}
		if order.callbackRate != nil {
			m["callbackRate"] = *order.callbackRate
		}
		if order.closePosition != nil {
			m["closePosition"] = *order.closePosition
		}
		orders = append(orders, m)
	}
	b, err := json.Marshal(orders)
	if err != nil {
		return &CreateBatchOrdersResponse{}, err
	}
	m := params{
		"batchOrders": string(b),
	}

	r.setFormParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return &CreateBatchOrdersResponse{}, err
	}

	rawMessages := make([]*stdjson.RawMessage, 0)

	err = json.Unmarshal(data, &rawMessages)

	if err != nil {
		return &CreateBatchOrdersResponse{}, err
	}

	batchCreateOrdersResponse := new(CreateBatchOrdersResponse)

	for _, j := range rawMessages {
		o := new(Order)
		if err := json.Unmarshal(*j, o); err != nil {
			return &CreateBatchOrdersResponse{}, err
		}

		if o.ClientOrderID != "" {
			batchCreateOrdersResponse.Orders = append(batchCreateOrdersResponse.Orders, o)
			continue
		}

	}

	return batchCreateOrdersResponse, nil

}

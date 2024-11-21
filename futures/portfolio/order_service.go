package portfolio

import (
	"context"
	"errors"
	"net/http"
)

// CreateOrderService create order
type CreateOrderService struct {
	c                *Client
	symbol           string
	side             SideType
	positionSide     *PositionSideType
	strategyType     StrategyType
	timeInForce      *TimeInForceType
	quantity         string
	reduceOnly       *bool
	price            *string
	newClientOrderID *string
	stopPrice        *string
	workingType      *WorkingType
	activationPrice  *string
	callbackRate     *string
	priceProtect     *bool
	newOrderRespType NewOrderRespType
	closePosition    *bool
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
func (s *CreateOrderService) Type(strategyType StrategyType) *CreateOrderService {
	s.strategyType = strategyType
	return s
}

// TimeInForce set timeInForce
func (s *CreateOrderService) TimeInForce(timeInForce TimeInForceType) *CreateOrderService {
	s.timeInForce = &timeInForce
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
		"strategyType":     s.strategyType,
		"quantity":         s.quantity,
		"newOrderRespType": s.newOrderRespType,
	}
	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.newClientOrderID != nil {
		m["newClientOrderId"] = *s.newClientOrderID
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
	data, header, err := s.createOrder(ctx, "/papi/v1/um/conditional/order", opts...)
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
	Symbol            string           `json:"symbol"`
	StrategyId        int64            `json:"strategyId"`
	ClientStrategyId  string           `json:"newClientStrategyId"`
	StrategyStatus    string           `json:"strategyStatus"`
	StrategyType      StrategyType     `json:"strategyType"`
	Price             float64          `json:"price,string"`
	OrigQuantity      float64          `json:"origQty,string"`
	ExecutedQuantity  float64          `json:"executedQty,string"`
	CumQuote          float64          `json:"cumQuote,string"`
	ReduceOnly        bool             `json:"reduceOnly"`
	Status            OrderStatusType  `json:"status"`
	StopPrice         float64          `json:"stopPrice,string"`
	TimeInForce       TimeInForceType  `json:"timeInForce"`
	Type              OrderType        `json:"type"`
	Side              SideType         `json:"side"`
	UpdateTime        int64            `json:"updateTime"`
	WorkingType       WorkingType      `json:"workingType"`
	ActivatePrice     float64          `json:"activatePrice,string"`
	PriceRate         float64          `json:"priceRate,string"`
	AvgPrice          float64          `json:"avgPrice,string"`
	PositionSide      PositionSideType `json:"positionSide"`
	ClosePosition     bool             `json:"closePosition"`
	PriceProtect      bool             `json:"priceProtect"`
	RateLimitOrder10s string           `json:"rateLimitOrder10s,omitempty"`
	RateLimitOrder1m  string           `json:"rateLimitOrder1m,omitempty"`
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
		endpoint: "/papi/v1/um/conditional/openOrders",
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
	c                *Client
	symbol           string
	strategyId       *int64
	clientStrategyId *string
}

func (s *GetOpenOrderService) Symbol(symbol string) *GetOpenOrderService {
	s.symbol = symbol
	return s
}

func (s *GetOpenOrderService) StrategyId(strategyId int64) *GetOpenOrderService {
	s.strategyId = &strategyId
	return s
}

func (s *GetOpenOrderService) ClientStrategyId(clientStrategyId string) *GetOpenOrderService {
	s.clientStrategyId = &clientStrategyId
	return s
}

func (s *GetOpenOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/conditional/openOrder",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.strategyId == nil && s.clientStrategyId == nil {
		return nil, errors.New("either strategyId or clientStrategyId must be sent")
	}
	if s.strategyId != nil {
		r.setParam("strategyId", *s.strategyId)
	}
	if s.clientStrategyId != nil {
		r.setParam("newClientStrategyId", *s.clientStrategyId)
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
	c                *Client
	symbol           string
	strategyId       *int64
	clientStrategyId *string
}

// Symbol set symbol
func (s *GetOrderService) Symbol(symbol string) *GetOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *GetOrderService) StrategyId(strategyId int64) *GetOrderService {
	s.strategyId = &strategyId
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *GetOrderService) ClientStrategyId(clientStrategyId string) *GetOrderService {
	s.clientStrategyId = &clientStrategyId
	return s
}

// Do send request
func (s *GetOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/conditional/orderHistory",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.strategyId == nil && s.clientStrategyId == nil {
		return nil, errors.New("either strategyId or clientStrategyId must be sent")
	}
	if s.strategyId != nil {
		r.setParam("strategyId", *s.strategyId)
	}
	if s.clientStrategyId != nil {
		r.setParam("newClientStrategyId", *s.clientStrategyId)
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
	Symbol           string           `json:"symbol"`
	StrategyId       int64            `json:"strategyId"`
	ClientStrategyId string           `json:"newClientStrategyId"`
	StrategyType     StrategyType     `json:"type"`
	Price            float64          `json:"price,string"`
	ReduceOnly       bool             `json:"reduceOnly"`
	OrigQuantity     float64          `json:"origQty,string"`
	ExecutedQuantity float64          `json:"executedQty,string"`
	CumQuantity      float64          `json:"cumQty,string"`
	CumQuote         float64          `json:"cumQuote,string"`
	StrategyStatus   OrderStatusType  `json:"status"`
	TimeInForce      TimeInForceType  `json:"timeInForce"`
	Side             SideType         `json:"side"`
	StopPrice        float64          `json:"stopPrice,string"`
	Time             int64            `json:"time"`
	UpdateTime       int64            `json:"updateTime"`
	WorkingType      WorkingType      `json:"workingType"`
	ActivatePrice    float64          `json:"activatePrice,string"`
	PriceRate        float64          `json:"priceRate,string"`
	AvgPrice         float64          `json:"avgPrice,string"`
	OrigType         string           `json:"origType"`
	PositionSide     PositionSideType `json:"positionSide"`
	PriceProtect     bool             `json:"priceProtect"`
	ClosePosition    bool             `json:"closePosition"`
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
		endpoint: "/papi/v1/um/conditional/allOrders",
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
	c                *Client
	symbol           string
	strategyId       *int64
	clientStrategyId *string
}

// Symbol set symbol
func (s *CancelOrderService) Symbol(symbol string) *CancelOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CancelOrderService) StrategyId(strategyId int64) *CancelOrderService {
	s.strategyId = &strategyId
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *CancelOrderService) ClientStrategyId(clientStrategyId string) *CancelOrderService {
	s.clientStrategyId = &clientStrategyId
	return s
}

// Do send request
func (s *CancelOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/um/conditional/order",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	if s.strategyId != nil {
		r.setFormParam("strategyId", *s.strategyId)
	}
	if s.clientStrategyId != nil {
		r.setFormParam("newClientStrategyId", *s.clientStrategyId)
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
	ClientStrategyId string           `json:"newClientStrategyId"`
	CumQuantity      float64          `json:"cumQty,string"`
	CumQuote         float64          `json:"cumQuote,string"`
	ExecutedQuantity float64          `json:"executedQty,string"`
	StrategyId       int64            `json:"strategyId"`
	OrigQuantity     float64          `json:"origQty,string"`
	Price            float64          `json:"price,string"`
	ReduceOnly       bool             `json:"reduceOnly"`
	Side             SideType         `json:"side"`
	Status           OrderStatusType  `json:"strategyStatus"`
	StopPrice        float64          `json:"stopPrice,string"`
	Symbol           string           `json:"symbol"`
	TimeInForce      TimeInForceType  `json:"timeInForce"`
	StrategyType     StrategyType     `json:"strategyType"`
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
		endpoint: "/papi/v1/um/conditional/allOpenOrders",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	_, _, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}

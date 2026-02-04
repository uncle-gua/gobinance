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

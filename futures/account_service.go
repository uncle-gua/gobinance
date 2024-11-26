package futures

import (
	"context"
	"net/http"
)

// GetBalanceService get account balance
type GetBalanceService struct {
	c *Client
}

// Do send request
func (s *GetBalanceService) Do(ctx context.Context, opts ...RequestOption) (res []*Balance, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v2/balance",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(data, &res)
	return res, err
}

// Balance define user balance of your account
type Balance struct {
	AccountAlias       string  `json:"accountAlias"`
	Asset              string  `json:"asset"`
	Balance            float64 `json:"balance,string"`
	CrossWalletBalance float64 `json:"crossWalletBalance,string"`
	CrossUnPnl         float64 `json:"crossUnPnl,string"`
	AvailableBalance   float64 `json:"availableBalance,string"`
	MaxWithdrawAmount  float64 `json:"maxWithdrawAmount,string"`
}

// GetAccountService get account info
type GetAccountService struct {
	c *Client
}

// Do send request
func (s *GetAccountService) Do(ctx context.Context, opts ...RequestOption) (res *Account, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v2/account",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(Account)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Account define account info
type Account struct {
	Assets                      []*AccountAsset    `json:"assets"`
	FeeTier                     int                `json:"feeTier"`
	CanTrade                    bool               `json:"canTrade"`
	CanDeposit                  bool               `json:"canDeposit"`
	CanWithdraw                 bool               `json:"canWithdraw"`
	UpdateTime                  int64              `json:"updateTime"`
	TotalInitialMargin          float64            `json:"totalInitialMargin,string"`
	TotalMaintMargin            float64            `json:"totalMaintMargin,string"`
	TotalWalletBalance          float64            `json:"totalWalletBalance,string"`
	TotalUnrealizedProfit       float64            `json:"totalUnrealizedProfit,string"`
	TotalMarginBalance          float64            `json:"totalMarginBalance,string"`
	TotalPositionInitialMargin  float64            `json:"totalPositionInitialMargin,string"`
	TotalOpenOrderInitialMargin float64            `json:"totalOpenOrderInitialMargin,string"`
	TotalCrossWalletBalance     float64            `json:"totalCrossWalletBalance,string"`
	TotalCrossUnPnl             float64            `json:"totalCrossUnPnl,string"`
	AvailableBalance            float64            `json:"availableBalance,string"`
	MaxWithdrawAmount           float64            `json:"maxWithdrawAmount,string"`
	Positions                   []*AccountPosition `json:"positions"`
}

// AccountAsset define account asset
type AccountAsset struct {
	Asset                  string  `json:"asset"`
	InitialMargin          float64 `json:"initialMargin,string"`
	MaintMargin            float64 `json:"maintMargin,string"`
	MarginBalance          float64 `json:"marginBalance,string"`
	MaxWithdrawAmount      float64 `json:"maxWithdrawAmount,string"`
	OpenOrderInitialMargin float64 `json:"openOrderInitialMargin,string"`
	PositionInitialMargin  float64 `json:"positionInitialMargin,string"`
	UnrealizedProfit       float64 `json:"unrealizedProfit,string"`
	WalletBalance          float64 `json:"walletBalance,string"`
}

// AccountPosition define account position
type AccountPosition struct {
	Isolated               bool             `json:"isolated"`
	Leverage               string           `json:"leverage"`
	InitialMargin          float64          `json:"initialMargin,string"`
	MaintMargin            float64          `json:"maintMargin,string"`
	OpenOrderInitialMargin float64          `json:"openOrderInitialMargin,string"`
	PositionInitialMargin  float64          `json:"positionInitialMargin,string"`
	Symbol                 string           `json:"symbol"`
	UnrealizedProfit       float64          `json:"unrealizedProfit,string"`
	EntryPrice             float64          `json:"entryPrice,string"`
	MaxNotional            float64          `json:"maxNotional,string"`
	PositionSide           PositionSideType `json:"positionSide"`
	PositionAmt            float64          `json:"positionAmt,string"`
	Notional               float64          `json:"notional,string"`
	IsolatedWallet         float64          `json:"isolatedWallet,string"`
	UpdateTime             int64            `json:"updateTime"`
}

// GetSymbolConfig get account info
type GetSymbolConfig struct {
	c *Client
}

// Do send request
func (s *GetSymbolConfig) Do(ctx context.Context, opts ...RequestOption) (res *SymbolConfig, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/symbolConfig",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SymbolConfig)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SymbolConfig define symbol config
type SymbolConfig struct {
	Symbol           string     `json:"symbol"`
	MarginType       MarginType `json:"marginType"`
	IsAutoAddMargin  bool       `json:"isAutoAddMargin"`
	Leverage         int        `json:"leverage"`
	MaxNotionalValue float64    `json:"maxNotionalValue,string"`
}

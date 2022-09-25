package ws

import (
	sjson "encoding/json"
	"strconv"
	"time"
)

type OrderBookL2 struct {
	ID     int64   `json:"id"`
	Price  float64 `json:"price,string"`
	Side   string  `json:"side"`
	Size   int64   `json:"size"`
	Symbol string  `json:"symbol"`
}

type OrderBookL2Delta struct {
	Delete []*OrderBookL2 `json:"delete"`
	Update []*OrderBookL2 `json:"update"`
	Insert []*OrderBookL2 `json:"insert"`
}

func (o *OrderBookL2) Key() string {
	return strconv.FormatInt(o.ID, 10)
}

type Trade struct {
	Timestamp     time.Time `json:"timestamp"`
	Symbol        string    `json:"symbol"`
	Side          string    `json:"side"`
	Size          int       `json:"size"`
	Price         float64   `json:"price"`
	TickDirection string    `json:"tick_direction"`
	TradeID       string    `json:"trade_id"`
	CrossSeq      int       `json:"cross_seq"`
}

type KLine struct {
	ID       int64   `json:"id"`        // 563
	Symbol   string  `json:"symbol"`    // BTCUSD
	OpenTime int64   `json:"open_time"` // 1539918000
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
	Volume   float64 `json:"volume"`
	Turnover float64 `json:"turnover"` // 0.0013844
	Interval string  `json:"interval"` // 1m
}

type Insurance struct {
	Currency      string    `json:"currency"`
	Timestamp     time.Time `json:"timestamp"`
	WalletBalance int64     `json:"wallet_balance"`
}

type Instrument struct {
	Symbol     string  `json:"symbol"`
	MarkPrice  float64 `json:"mark_price"`
	IndexPrice float64 `json:"index_price"`
}

type Order struct {
	OrderID       string       `json:"order_id"`
	OrderLinkID   string       `json:"order_link_id"`
	Symbol        string       `json:"symbol"`
	Side          string       `json:"side"`
	OrderType     string       `json:"order_type"`
	Price         sjson.Number `json:"price"`
	Qty           float64      `json:"qty"`
	TimeInForce   string       `json:"time_in_force"` // GoodTillCancel/ImmediateOrCancel/FillOrKill/PostOnly
	CreateType    string       `json:"create_type"`
	CancelType    string       `json:"cancel_type"`
	OrderStatus   string       `json:"order_status"`
	LeavesQty     float64      `json:"leaves_qty"`
	CumExecQty    float64      `json:"cum_exec_qty"`
	CumExecValue  sjson.Number `json:"cum_exec_value"`
	CumExecFee    sjson.Number `json:"cum_exec_fee"`
	Timestamp     time.Time    `json:"timestamp"`
	TakeProfit    sjson.Number `json:"take_profit"`
	StopLoss      sjson.Number `json:"stop_loss"`
	TrailingStop  sjson.Number `json:"trailing_stop"`
	LastExecPrice sjson.Number `json:"last_exec_price"`
}

type Execution struct {
	Symbol      string    `json:"symbol"`
	Side        string    `json:"side"`
	OrderID     string    `json:"order_id"`
	ExecID      string    `json:"exec_id"`
	OrderLinkID string    `json:"order_link_id"`
	Price       float64   `json:"price"`
	OrderQty    float64   `json:"order_qty"`
	ExecType    string    `json:"exec_type"`
	ExecQty     float64   `json:"exec_qty"`
	ExecFee     float64   `json:"exec_fee"`
	LeavesQty   float64   `json:"leaves_qty"`
	IsMaker     bool      `json:"is_maker"`
	TradeTime   time.Time `json:"trade_time"`
}

type Position struct {
	UserID           int64   `json:"user_id,string"`
	Symbol           string  `json:"symbol"`
	Size             float64 `json:"size"`
	Side             string  `json:"side"`
	PositionValue    float64 `json:"position_value"`
	EntryPrice       float64 `json:"entry_price"`
	LiqPrice         float64 `json:"liq_price"`
	BustPrice        float64 `json:"bust_price"`
	Leverage         float64 `json:"leverage"`
	OrderMargin      float64 `json:"order_margin"`
	PositionMargin   float64 `json:"position_margin"`
	AvailableBalance float64 `json:"available_balance,string"`
	TakeProfit       float64 `json:"take_profit"`
	StopLoss         float64 `json:"stop_loss"`
	RealisedPnl      float64 `json:"realised_pnl"`
	TrailingStop     float64 `json:"trailing_stop"`
	TrailingActive   float64 `json:"trailing_active"`
	WalletBalance    float64 `json:"wallet_balance,string"`
	RiskID           int     `json:"risk_id,string"`
	OccClosingFee    float64 `json:"occ_closing_fee"`
	OccFundingFee    float64 `json:"occ_funding_fee"`
	AutoAddMargin    int     `json:"auto_add_margin,string"`
	CumRealisedPnl   float64 `json:"cum_realised_pnl"`
	PositionStatus   string  `json:"position_status"`
	PositionSeq      int64   `json:"position_seq,string"`
}

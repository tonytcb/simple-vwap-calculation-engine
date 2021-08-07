package coinbase

import "time"

type Channel struct {
	Name     string   `json:"name"`
	Products []string `json:"product_ids"`
}

type Request struct {
	Type     string    `json:"type"`
	Channels []Channel `json:"channels"`
}

type Response struct {
	Type      string    `json:"type"`
	TradeID   int       `json:"trade_id"`
	ProductID string    `json:"product_id"`
	Size      string    `json:"size"`
	Price     string    `json:"price"`
	Side      string    `json:"side"`
	Time      time.Time `json:"time,string"`
}

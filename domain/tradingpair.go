package domain

import "fmt"

// TradingPair defines the two currencies to create a trading
type TradingPair struct {
	From Currency
	To   Currency
}

// NewTradingPair builds a new TradingPair struct
func NewTradingPair(from Currency, to Currency) TradingPair {
	return TradingPair{From: from, To: to}
}

// String returns the TradingPair struct as a string value
func (t TradingPair) String() string {
	return fmt.Sprintf("%s-%s", t.From.Alias, t.To.Alias)
}

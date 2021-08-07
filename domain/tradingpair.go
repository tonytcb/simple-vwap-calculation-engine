package domain

import "fmt"

type TradingPair struct {
	From Currency
	To   Currency
}

func NewTradingPair(from Currency, to Currency) TradingPair {
	return TradingPair{From: from, To: to}
}

func (t TradingPair) String() string {
	return fmt.Sprintf("%s-%s", t.From.Alias, t.To.Alias)
}

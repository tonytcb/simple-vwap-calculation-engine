package notifier

import (
	"fmt"

	"github.com/tonytcb/simple-vwap-calculation-engine/domain"
)

// Print contains dependencies to execute the print notification
type Print struct {
}

// NewPrint builds a new Print struct
func NewPrint() *Print {
	return &Print{}
}

// Notify sends the notification
func (p Print) Notify(trading domain.Trading, f float64) error {
	fmt.Printf(
		"[%s] trading-id=%d | trading-share=%f | trading-price=%f | vwap=%f\n",
		trading.Pair,
		trading.ID,
		trading.Share,
		trading.Price,
		f,
	)

	return nil
}

package notifier

import (
	"fmt"
	"github.com/tonytcb/simple-vwap-calculation-engine/domain"
)

type Print struct {
}

func NewPrint() *Print {
	return &Print{}
}

func (p Print) Notify(trading domain.Trading, f float64) error {
	fmt.Printf(
		"[%s] trading-id=%d | trading-price=%f | vwap=%f\n",
		trading.ProductID,
		trading.ID,
		trading.Price,
		f,
	)

	return nil
}

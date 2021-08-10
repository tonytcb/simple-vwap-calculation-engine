package usecases

import (
	"log"

	"github.com/tonytcb/simple-vwap-calculation-engine/domain"
)

// Provider defines the behaviour to send and receive data from some data source
type Provider interface {
	Subscribe([]domain.TradingPair) error
	Pull(chan domain.Trading) error
}

// Notifier defines the behaviour to send a notification
type Notifier interface {
	Notify(domain.Trading, float64) error
}

// VWAP holds dependencies to calculate the on-demand VWAP
type VWAP struct {
	provider Provider
	notifier Notifier
}

// NewVWAP builds a new VWAP struct
func NewVWAP(provider Provider, notifier Notifier) *VWAP {
	return &VWAP{provider: provider, notifier: notifier}
}

// CalculateWithMaxTradings calculates the volume-weighted average price
func (v *VWAP) CalculateWithMaxTradings(max int) {
	var ch = make(chan domain.Trading)
	var operations = map[string][]domain.Trading{}

	go func() {
		if err := v.provider.Pull(ch); err != nil {
			log.Fatalln(err)
		}
	}()

	for trading := range ch {
		if _, ok := operations[trading.Pair]; !ok {
			operations[trading.Pair] = make([]domain.Trading, 0)
		}

		operations[trading.Pair] = append(operations[trading.Pair], trading)

		if (max > 0) && (len(operations[trading.Pair]) > max) {
			operations[trading.Pair] = removeTradingByIndex(operations[trading.Pair], 0)
		}

		v.notifier.Notify(trading, vwap(operations[trading.Pair]))
	}
}

func removeTradingByIndex(a []domain.Trading, i int) []domain.Trading {
	copy(a[i:], a[i+1:])
	a[len(a)-1] = domain.Trading{}
	a = a[:len(a)-1]

	return a
}

func vwap(tradings []domain.Trading) float64 {
	var sumPriceTimesVolume float64
	var sumVolume float64

	if (tradings == nil) || (len(tradings) == 0) {
		return 0
	}

	for _, v := range tradings {
		sumPriceTimesVolume += v.Price * v.Share
		sumVolume += v.Share
	}

	if sumVolume == 0 {
		return 0
	}

	return sumPriceTimesVolume / sumVolume
}

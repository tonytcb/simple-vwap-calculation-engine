package services

import (
	"log"

	"github.com/tonytcb/simple-vwap-calculation-engine/domain"
)

// Notifier defines the behaviour to send a notification
type Notifier interface {
	Notify(domain.Trading, float64) error
}

// VWAP holds dependencies to calculate the VWAP on-demand
type VWAP struct {
	provider domain.Provider
	notifier Notifier
}

// NewVWAP builds a new VWAP struct
func NewVWAP(provider domain.Provider, notifier Notifier) *VWAP {
	return &VWAP{provider: provider, notifier: notifier}
}

// WithMaxTradings calculates the volume-weighted average price
func (v *VWAP) WithMaxTradings(max int) {
	var ch = make(chan domain.Trading)
	var operations = map[string][]domain.Trading{}

	go func() {
		if err := v.provider.Pull(ch); err != nil {
			log.Fatalln(err)
		}
	}()

	for trading := range ch {
		if _, ok := operations[trading.ProductID]; !ok {
			operations[trading.ProductID] = make([]domain.Trading, 0)
		}

		operations[trading.ProductID] = append(operations[trading.ProductID], trading)

		if (max > 0) && (len(operations[trading.ProductID]) > max) {
			operations[trading.ProductID] = removeTradingByIndex(operations[trading.ProductID], 0)
		}

		v.notifier.Notify(trading, vwap(operations[trading.ProductID]))
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

	for _, v := range tradings {
		sumPriceTimesVolume += v.Price * v.Share
		sumVolume += v.Share
	}

	return sumPriceTimesVolume / sumVolume
}

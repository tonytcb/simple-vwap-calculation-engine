package services

import (
	"log"

	"github.com/tonytcb/simple-vwap-calculation-engine/domain"
)

type Notifier interface {
	Notify(domain.Trading, float64) error
}

type VWAP struct {
	provider domain.Provider
	notifier Notifier
}

func NewVWAP(provider domain.Provider, notifier Notifier) *VWAP {
	return &VWAP{provider: provider, notifier: notifier}
}

func (v *VWAP) WithMaxTradings(max int) {
	var ch = make(chan domain.Trading)
	var operations = map[string][]domain.Trading{}

	go func() {
		if err := v.provider.Pull(ch); err != nil {
			log.Fatalln(err)
		}
	}()

	for trading := range ch {
		tradings, ok := operations[trading.ProductID]
		if !ok {
			tradings = make([]domain.Trading, 0)
		}

		tradings = append(operations[trading.ProductID], trading)

		operations[trading.ProductID] = tradings

		v.notifier.Notify(trading, vwap(tradings))

		// @todo limit the calc to the max tranding parameters
	}
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

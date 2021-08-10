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
	var (
		ch                  = make(chan domain.Trading)
		operations          = map[string]*domain.TradingQueue{}
		sumPriceTimesVolume = map[string]float64{}
		sumVolume           = map[string]float64{}
	)

	go func() {
		if err := v.provider.Pull(ch); err != nil {
			log.Fatalln(err)
		}
	}()

	for trading := range ch {
		if _, ok := operations[trading.Pair]; !ok {
			operations[trading.Pair] = domain.NewTradingQueue()
		}

		if _, ok := sumPriceTimesVolume[trading.Pair]; !ok {
			sumPriceTimesVolume[trading.Pair] = 0
		}

		if _, ok := sumVolume[trading.Pair]; !ok {
			sumVolume[trading.Pair] = 0
		}

		operations[trading.Pair].Enqueue(trading)

		sumPriceTimesVolume[trading.Pair] += trading.Price * trading.Share
		sumVolume[trading.Pair] += trading.Share

		if (max > 0) && (operations[trading.Pair].Len() > max) {
			var first = operations[trading.Pair].Dequeue()

			sumPriceTimesVolume[trading.Pair] -= first.Price * first.Share
			sumVolume[trading.Pair] -= first.Share
		}

		var result = sumPriceTimesVolume[trading.Pair] / sumVolume[trading.Pair]

		if err := v.notifier.Notify(trading, result); err != nil {
			log.Println("error to send VWAP notification:", err.Error())
		}
	}
}

package usecases

import (
	"fmt"
	"testing"

	"github.com/tonytcb/simple-vwap-calculation-engine/domain"
)

func TestVWAP_WithoutMaxTradings(t *testing.T) {
	var (
		bitcoinToDollarPair  = fmt.Sprintf("%s-%s", domain.Bitcoin.Alias, domain.Dollar.Alias)
		ethereumToDollarPair = fmt.Sprintf("%s-%s", domain.Ethereum.Alias, domain.Dollar.Alias)
		tradingsMock         = []domain.Trading{
			{Share: 0.00217105, Price: 45900.01, Pair: bitcoinToDollarPair},
			{Share: 0.03, Price: 45900, Pair: bitcoinToDollarPair},
			{Share: 0.00208368, Price: 45900.01, Pair: bitcoinToDollarPair},
			{Share: 0.08877488, Price: 45900, Pair: bitcoinToDollarPair},
			{Share: 0.1, Price: 10000, Pair: ethereumToDollarPair},
		}
		expected = []float64{45900.01, 45900.00067484586, 45900.001242085404, 45900.00034582976, 10000}
		results  = make(chan float64)
		i        = 0
	)

	var vwap = &VWAP{
		provider: providerMock{tradings: tradingsMock},
		notifier: notifierMock{ch: results},
	}
	go vwap.CalculateWithMaxTradings(0)

	for got := range results {
		want := expected[i]

		if got != want {
			t.Errorf("vwap() = %v, want %v", got, want)
		}

		i++

		if len(expected) == i {
			// stop test after check all expected vwap
			close(results)
		}
	}
}

func TestVWAP_WithMaxTradings(t *testing.T) {
	var (
		bitcoinToDollarPair  = fmt.Sprintf("%s-%s", domain.Bitcoin.Alias, domain.Dollar.Alias)
		ethereumToDollarPair = fmt.Sprintf("%s-%s", domain.Ethereum.Alias, domain.Dollar.Alias)
		tradingsMock         = []domain.Trading{
			{Share: 1, Price: 45900, Pair: bitcoinToDollarPair},
			{Share: 1, Price: 45900, Pair: bitcoinToDollarPair},
			{Share: 1, Price: 45900, Pair: bitcoinToDollarPair},
			{Share: 1, Price: 10000, Pair: ethereumToDollarPair},
		}
		expected = []float64{45900, 45900, 45900, 10000}
		results  = make(chan float64)
		i        = 0
	)

	var vwap = &VWAP{
		provider: providerMock{tradings: tradingsMock},
		notifier: notifierMock{ch: results},
	}
	go vwap.CalculateWithMaxTradings(1)

	for got := range results {
		want := expected[i]

		if got != want {
			t.Errorf("vwap() = %v, want %v", got, want)
		}

		i++

		if len(expected) == i {
			// stop test after check all expected vwap
			close(results)
		}
	}
}

type providerMock struct {
	tradings []domain.Trading
}

func (p providerMock) Subscribe(_ []domain.TradingPair) error {
	return nil
}

func (p providerMock) Pull(ch chan domain.Trading) error {
	if p.tradings == nil {
		return nil
	}

	for _, v := range p.tradings {
		ch <- v
	}

	return nil
}

type notifierMock struct {
	ch chan float64
}

func (n notifierMock) Notify(_ domain.Trading, f float64) error {
	n.ch <- f

	return nil
}

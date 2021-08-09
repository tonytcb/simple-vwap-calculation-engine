package services

import (
	"testing"

	"github.com/tonytcb/simple-vwap-calculation-engine/domain"
)

func Test_vwap(t *testing.T) {
	type args struct {
		tradings []domain.Trading
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Nil slice should return 0",
			args: args{
				tradings: []domain.Trading{},
			},
			want: 0,
		},
		{
			name: "Empty slice should return 0",
			args: args{
				tradings: []domain.Trading{},
			},
			want: 0,
		},
		{
			name: "Should return 1",
			args: args{
				tradings: []domain.Trading{
					{Share: 1, Price: 1},
					{Share: 1, Price: 1},
					{Share: 1, Price: 1},
					{Share: 1, Price: 1},
					{Share: 1, Price: 1},
				},
			},
			want: 1,
		},
		{
			name: "Should return 45900.01",
			args: args{
				tradings: []domain.Trading{
					{Share: 0.00217105, Price: 45900.01},
				},
			},
			want: 45900.01,
		},
		{
			name: "Should return 45900.00067484586",
			args: args{
				tradings: []domain.Trading{
					{Share: 0.00217105, Price: 45900.01},
					{Share: 0.03, Price: 45900},
				},
			},
			want: 45900.00067484586,
		},
		{
			name: "Should return 45900.001242085404",
			args: args{
				tradings: []domain.Trading{
					{Share: 0.00217105, Price: 45900.01},
					{Share: 0.03, Price: 45900},
					{Share: 0.00208368, Price: 45900.01},
				},
			},
			want: 45900.001242085404,
		},
		{
			name: "Should return 45900.00034582976",
			args: args{
				tradings: []domain.Trading{
					{Share: 0.00217105, Price: 45900.01},
					{Share: 0.03, Price: 45900},
					{Share: 0.00208368, Price: 45900.01},
					{Share: 0.08877488, Price: 45900},
				},
			},
			want: 45900.00034582976,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vwap(tt.args.tradings); got != tt.want {
				t.Errorf("vwap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVWAP_WithMaxTradings(t *testing.T) {
	var (
		tradingsMock = []domain.Trading{
			{Share: 0.00217105, Price: 45900.01},
			{Share: 0.03, Price: 45900},
			{Share: 0.00208368, Price: 45900.01},
			{Share: 0.08877488, Price: 45900},
		}
		expected = []float64{45900.01, 45900.00067484586, 45900.001242085404, 45900.00034582976}
		results  = make(chan float64)
		i        = 0
	)

	var vwap = &VWAP{
		provider: providerMock{tradings: tradingsMock},
		notifier: notifierMock{ch: results},
	}
	go vwap.WithMaxTradings(0)

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

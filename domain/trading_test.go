package domain

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNewTrading(t *testing.T) {
	var (
		now                 = time.Now()
		bitcoinToDollarPair = fmt.Sprintf("%s-%s", Bitcoin.Alias, Dollar.Alias)
	)

	type args struct {
		tradeID   int
		pair      string
		size      string
		price     string
		createdAt time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    Trading
		wantErr bool
	}{
		{
			name: "Should return a valid Trading struct",
			args: args{
				tradeID:   1,
				pair:      bitcoinToDollarPair,
				size:      "0.01",
				price:     "4000",
				createdAt: now,
			},
			want: Trading{
				ID:        1,
				Pair:      bitcoinToDollarPair,
				Share:     0.01,
				Price:     4000,
				CreatedAt: now,
			},
			wantErr: false,
		},
		{
			name: "Should return an error when the size is not a valid float",
			args: args{
				tradeID:   1,
				pair:      bitcoinToDollarPair,
				size:      "x",
				price:     "4000",
				createdAt: now,
			},
			want:    Trading{},
			wantErr: true,
		},
		{
			name: "Should return an error when the price is not a valid float",
			args: args{
				tradeID:   1,
				pair:      bitcoinToDollarPair,
				size:      "0.01",
				price:     "y",
				createdAt: now,
			},
			want:    Trading{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTrading(tt.args.tradeID, tt.args.pair, tt.args.size, tt.args.price, tt.args.createdAt)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTrading() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTrading() got = %v, want %v", got, tt.want)
			}
		})
	}
}

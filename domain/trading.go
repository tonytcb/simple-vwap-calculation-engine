package domain

import (
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// Trading defines the Trading domain
type Trading struct {
	ID        int
	Pair      string
	Share     float64
	Price     float64
	CreatedAt time.Time
}

// NewTrading builds a new Trading structure ensuring its values
func NewTrading(tradeID int, pair string, size string, price string, createdAt time.Time) (Trading, error) {
	sizeNumber, err := strconv.ParseFloat(size, 64)
	if err != nil {
		return Trading{}, errors.Wrap(err, "error to convert size to float64")
	}

	priceNumber, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return Trading{}, errors.Wrap(err, "error to convert price to float64")
	}

	return Trading{ID: tradeID, Pair: pair, Share: sizeNumber, Price: priceNumber, CreatedAt: createdAt}, nil
}

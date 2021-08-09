package domain

import (
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// Trading defines the Trading domain
type Trading struct {
	ID        int
	ProductID string
	Share     float64
	Price     float64
	Time      time.Time
}

// NewTrading builds a new Trading structure ensuring its values
func NewTrading(tradeID int, productID string, size string, price string, time time.Time) (Trading, error) {
	sizeNumber, err := strconv.ParseFloat(size, 64)
	if err != nil {
		return Trading{}, errors.Wrap(err, "error to convert size to float64")
	}

	priceNumber, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return Trading{}, errors.Wrap(err, "error to convert price to float64")
	}

	return Trading{ID: tradeID, ProductID: productID, Share: sizeNumber, Price: priceNumber, Time: time}, nil
}

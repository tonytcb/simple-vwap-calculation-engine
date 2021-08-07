package domain

import (
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type Trading struct {
	TradeID   int
	ProductID string
	Size      float64
	Price     float64
	Time      time.Time
}

func NewTrading(tradeID int, productID string, size string, price string, time time.Time) (Trading, error) {
	sizeNumber, err := strconv.ParseFloat(size, 64)
	if err != nil {
		return Trading{}, errors.Wrap(err, "error to convert size to float64")
	}

	priceNumber, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return Trading{}, errors.Wrap(err, "error to convert size to float64")
	}

	return Trading{TradeID: tradeID, ProductID: productID, Size: sizeNumber, Price: priceNumber, Time: time}, nil
}

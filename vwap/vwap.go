package vwap

import (
	"github.com/tonytcb/simple-vwap-calculation-engine/domain"
)

func VWAP(tradings []domain.Trading) float64 {
	var sumPriceTimesVolume float64
	var sumVolume float64

	for _, v := range tradings {
		sumPriceTimesVolume += v.Price * v.Size
		sumVolume += v.Size
	}

	return sumPriceTimesVolume / sumVolume
}

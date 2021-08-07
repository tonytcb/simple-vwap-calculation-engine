package domain

type Exchange interface {
	Subscribe([]TradingPair) error
	Pull(chan Trading)
}

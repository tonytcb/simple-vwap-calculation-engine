package domain

// Provider defines the behaviour to send and receive data from some data source
type Provider interface {
	Subscribe([]TradingPair) error
	Pull(chan Trading) error
}

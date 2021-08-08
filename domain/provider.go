package domain

// Provider defines the behaviour to send and pull data from some source
type Provider interface {
	Subscribe([]TradingPair) error
	Pull(chan Trading) error
}

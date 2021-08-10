package domain

// TradingQueue holds tradings in a queue structure
type TradingQueue struct {
	tradings []Trading
}

// NewTradingQueue builds an empty TradingQueue struct
func NewTradingQueue() *TradingQueue {
	return &TradingQueue{tradings: make([]Trading, 0)}
}

// Enqueue adds one element
func (t *TradingQueue) Enqueue(trading Trading) {
	t.tradings = append(t.tradings, trading)
}

// Dequeue returns and remove the first element
func (t *TradingQueue) Dequeue() Trading {
	var first = t.tradings[0]
	t.tradings = t.tradings[1:]
	return first
}

// Len returns the length of the queue
func (t *TradingQueue) Len() int {
	return len(t.tradings)
}

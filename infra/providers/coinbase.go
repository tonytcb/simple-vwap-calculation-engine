package providers

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/tonytcb/simple-vwap-calculation-engine/domain"
)

const (
	websocketEndpoint    = "wss://ws-feed.pro.coinbase.com"
	messageSubscribeType = "subscribe"
	matchesChannelName   = "matches"
)

type coinbaseChannel struct {
	Name     string   `json:"name"`
	Products []string `json:"product_ids"`
}

// coinbaseRequestMessage defines the request to be sent to Coinbase Websocket
type coinbaseRequestMessage struct {
	Type     string            `json:"type"`
	Channels []coinbaseChannel `json:"channels"`
}

// coinbaseResponseMessage defines the response coming from Coinbase Websocket
type coinbaseResponseMessage struct {
	Type      string    `json:"type"`
	TradeID   int       `json:"trade_id"`
	ProductID string    `json:"product_id"`
	Size      string    `json:"size"`
	Price     string    `json:"price"`
	Side      string    `json:"side"`
	Time      time.Time `json:"time,string"`
}

// Coinbase defined the Coinbase trading provider
type Coinbase struct {
	conn *websocket.Conn
}

// NewCoinbase creates a new Coinbase client struct
func NewCoinbase() (*Coinbase, error) {
	var wsDialer websocket.Dialer

	conn, _, err := wsDialer.Dial(websocketEndpoint, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error to establish Coinbase Websocket connection")
	}

	return &Coinbase{conn: conn}, nil
}

// Subscribe subscribes to receive data from Matches channel
func (c Coinbase) Subscribe(pairs []domain.TradingPair) error {
	var products []string
	for _, v := range pairs {
		products = append(products, v.String())
	}

	request := coinbaseRequestMessage{
		Type: messageSubscribeType,
		Channels: []coinbaseChannel{
			{
				Name:     matchesChannelName,
				Products: products,
			},
		},
	}

	if err := c.conn.WriteJSON(request); err != nil {
		return errors.Wrap(err, fmt.Sprintf("error to subscribe with %v pairs", products))
	}

	return nil
}

// Pull pulls trading messages, send them to the received channel
func (c Coinbase) Pull(ch chan domain.Trading) error {
	for {
		message := coinbaseResponseMessage{}
		if err := c.conn.ReadJSON(&message); err != nil {
			close(ch)
			return errors.Wrap(err, "error to read message from websocket")
		}

		if message.Type == "error" {
			close(ch)
			return fmt.Errorf("error message: %v", message)
		}

		if message.Type == "subscriptions" {
			// ignore subscription message type
			continue
		}

		trading, err := domain.NewTrading(message.TradeID, message.ProductID, message.Size, message.Price, message.Time)
		if err != nil {
			close(ch)
			return errors.Wrap(err, "error to parse websocket message")
		}

		ch <- trading
	}
}

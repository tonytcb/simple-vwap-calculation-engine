package coinbase

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/tonytcb/simple-vwap-calculation-engine/domain"
)

const (
	websocketEndpoint    = "wss://ws-feed.pro.coinbase.com"
	messageSubscribeType = "subscribe"
	matchesChannelName   = "matches"
)

type Coinbase struct {
	conn *websocket.Conn
}

func NewCoinbase() (*Coinbase, error) {
	var wsDialer websocket.Dialer

	conn, _, err := wsDialer.Dial(websocketEndpoint, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error to establish Coinbase Websocket connection")
	}

	return &Coinbase{conn: conn}, nil
}

func (c Coinbase) Subscribe(pairs []domain.TradingPair) error {
	var products []string
	for _, v := range pairs {
		products = append(products, v.String())
	}

	request := Request{
		Type: messageSubscribeType,
		Channels: []Channel{
			{
				Name:     matchesChannelName,
				Products: products,
			},
		},
	}

	if err := c.conn.WriteJSON(request); err != nil {
		return errors.Wrap(err, fmt.Sprintf("error to subscribe to %v", products))
	}

	return nil
}

func (c Coinbase) Pull(ch chan domain.Trading) error {
	for {
		message := Response{}
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

package main

import (
	"log"

	"github.com/tonytcb/simple-vwap-calculation-engine/domain"
	"github.com/tonytcb/simple-vwap-calculation-engine/exchanges/coinbase"
	"github.com/tonytcb/simple-vwap-calculation-engine/vwap"
)

func main() {
	log.Println("========== starting app")

	calculateVwapWithCoinbase()
}

func calculateVwapWithCoinbase() {
	client, err := coinbase.NewCoinbase()
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Subscribe([]domain.TradingPair{
		domain.NewTradingPair(domain.Bitcoin, domain.Dollar),
		domain.NewTradingPair(domain.Ethereum, domain.Dollar),
		domain.NewTradingPair(domain.Ethereum, domain.Bitcoin),
	})

	if err != nil {
		log.Fatalln(err)
	}

	var ch = make(chan domain.Trading)
	var operations = map[string][]domain.Trading{}

	go func() {
		if err = client.Pull(ch); err != nil {
			log.Fatalln(err)
		}
	}()

	for trading := range ch {
		tradings, ok := operations[trading.ProductID]
		if !ok {
			tradings = make([]domain.Trading, 0)
		}

		tradings = append(operations[trading.ProductID], trading)

		operations[trading.ProductID] = tradings

		result := vwap.VWAP(tradings)
		log.Println(trading.ProductID, result)
	}
}

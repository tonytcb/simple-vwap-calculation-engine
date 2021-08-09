package main

import (
	"log"
	"os"
	"strconv"

	"github.com/tonytcb/simple-vwap-calculation-engine/domain"
	"github.com/tonytcb/simple-vwap-calculation-engine/infra/notifier"
	"github.com/tonytcb/simple-vwap-calculation-engine/infra/providers"
	"github.com/tonytcb/simple-vwap-calculation-engine/services"
)

const defaultMaxTradings = 200

func main() {
	log.Println("========== Starting simple-vwap app")

	coinbaseProvider, err := providers.NewCoinbase()
	if err != nil {
		log.Fatalln(err)
	}

	err = coinbaseProvider.Subscribe([]domain.TradingPair{
		domain.NewTradingPair(domain.Bitcoin, domain.Dollar),
		domain.NewTradingPair(domain.Ethereum, domain.Dollar),
		domain.NewTradingPair(domain.Ethereum, domain.Bitcoin),
	})

	if err != nil {
		log.Fatalln(err)
	}

	var max = maxTradingsParameter()

	services.NewVWAP(coinbaseProvider, notifier.NewPrint()).WithMaxTradings(max)
}

func maxTradingsParameter() int {
	if (len(os.Args) == 1) || (os.Args[1] == "") {
		return defaultMaxTradings
	}

	v, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return defaultMaxTradings
	}

	return v
}

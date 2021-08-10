package main

import (
	"log"
	"os"
	"strconv"

	"github.com/tonytcb/simple-vwap-calculation-engine/domain"
	"github.com/tonytcb/simple-vwap-calculation-engine/infra/notifier"
	"github.com/tonytcb/simple-vwap-calculation-engine/infra/providers"
	"github.com/tonytcb/simple-vwap-calculation-engine/usecases"
)

const defaultMaxTradings = 200

func main() {
	log.Println("========== Starting simple-vwap app")

	var max = maxTradingsParameter()

	ch := make(chan int)

	go startVwap(domain.NewTradingPair(domain.Bitcoin, domain.Dollar), max)
	go startVwap(domain.NewTradingPair(domain.Ethereum, domain.Dollar), max)
	go startVwap(domain.NewTradingPair(domain.Ethereum, domain.Bitcoin), max)

	<-ch
}

func startVwap(pair domain.TradingPair, max int) {
	coinbaseProvider, err := providers.NewCoinbase()
	if err != nil {
		log.Fatalln(err)
	}

	if err = coinbaseProvider.Subscribe([]domain.TradingPair{pair}); err != nil {
		log.Fatalln(err)
	}

	usecases.NewVWAP(coinbaseProvider, notifier.NewPrint()).CalculateWithMaxTradings(max)
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

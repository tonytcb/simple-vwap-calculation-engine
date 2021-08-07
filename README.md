# simple-vwap-calculation-engine

## What is VWAP

In finance, volume-weighted average price (VWAP) is the ratio of the value traded to total volume traded over a particular time horizon (usually one day). It is a measure of the average price at which a stock is traded over the trading horizon. [Wikipedia](https://en.wikipedia.org/wiki/Volume-weighted_average_price)

## Goals

This project was created for **test porposes**.

This project aims to calculate the VWAP given one or more trading pairs pulling the data from Coinbase through its Websocket service.

The VWAP is recalculated after every trading data received, limiting to N tradings, for test purposes.

## Tools

- [Golang 1.16](https://golang.org/)
- [Gorilla WebSocket](https://github.com/gorilla/websocket)
- [Testify](https://github.com/stretchr/testify)
-

## How to run

With docker you can run `make docker-run MAX_TRADINGS=100`.

With go installed in your local environment: `go run . MAX_TRADINGS=100`

_MAX_TRADINGS is equals to 200 as default value._
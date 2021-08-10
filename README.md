# simple-vwap-calculation-engine

## What is VWAP

In finance, volume-weighted average price (VWAP) is the ratio of the value traded to total volume traded over a particular time horizon (usually one day). It is a measure of the average price at which a stock is traded over the trading horizon ([Wikipedia](https://en.wikipedia.org/wiki/Volume-weighted_average_price)).

## Goals

This project aims to calculate the VWAP given one or more trading pairs pulling the data from Coinbase through its Websocket service.

The algorithm runs on-demand, meaning that the VWAP is recalculated after every trading data received, limiting to N tradings. When the limit is reached, the oldest trading will fall off.

## Design Solution

The application architecture follows the principles of the **Clean Architecture**, originally described by Robert C. Martin. The foundation of this kind of architecture if the dependency injection, producing systems that are independent of external agents, testable and easier to maintain.
You can read more about Clean Architecture [here](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

To start the on-demand VWAP calculation, it is required to inform the tradings provider and the structure to stream the VWAP result.

For test purposes, it's only available [Coinbase](https://docs.pro.coinbase.com/#the-matches-channel) as a provider and the notification is printing the result in stdout. But, it's easy to add more providers implementing its interface and propagating the Trading to the proper channel, as well is easy to implement a real notification structure using some real stream data technology, then injecting both dependencies in the VWAP calculation use case.

## Tools

- [Docker](https://www.docker.com/)
- [Golang 1.16](https://golang.org/)
- [Gorilla WebSocket](https://github.com/gorilla/websocket)

## How to run

With docker you can run `make docker-run MAX_TRADINGS=100`.

With go installed in your local environment: `go run . 100`

_MAX_TRADINGS is equals to 200 as default value._

To run all test you can run `make docker-tests` or `make tests`.
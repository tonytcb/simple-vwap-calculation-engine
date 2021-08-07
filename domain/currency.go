package domain

type Currency struct {
	Name  string
	Alias string
}

var (
	Ethereum = Currency{
		Name:  "Ethereum",
		Alias: "ETH",
	}

	Bitcoin = Currency{
		Name:  "Bitcoin",
		Alias: "BTC",
	}

	Dollar = Currency{
		Name:  "Dollar",
		Alias: "USD",
	}
)

package domain

// Currency defines the currency behaviour
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

package domain

// Currency defines the currency behaviour
type Currency struct {
	Name  string
	Alias string
}

var (
	// Ethereum represents the ethereum currency
	Ethereum = Currency{
		Name:  "Ethereum",
		Alias: "ETH",
	}

	// Bitcoin represents the bitcoin currency
	Bitcoin = Currency{
		Name:  "Bitcoin",
		Alias: "BTC",
	}

	// Dollar represents the dollar currency
	Dollar = Currency{
		Name:  "Dollar",
		Alias: "USD",
	}
)

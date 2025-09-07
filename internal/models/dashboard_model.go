package models

// CountryReport holds one country's aggregated sales data
type CountryReport struct {
	Country      string   `json:"country"`
	Products     []string `json:"products"`
	TotalRevenue float64  `json:"totalRevenue"`
	Transactions int      `json:"transactions"`
}

// Product represents a single product with sales and stock information
type Product struct {
	ProductName   string `json:"productName"`
	ItemsSold     int    `json:"itemsSold"`
	StockQuantity int    `json:"stockQuantity"`
}

// MonthlySale represents sales data for a single month
type MonthlySale struct {
	Month     string `json:"month"`
	UnitsSold int    `json:"unitsSold"`
}

// Region represents sales data for a single region
type Region struct {
	Region     string  `json:"region"`
	RevenueUSD float64 `json:"revenueUSD"`
	ItemsSold  int     `json:"itemsSold"`
}

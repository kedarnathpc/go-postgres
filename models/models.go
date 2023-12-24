package models

type Stock struct {
	// StockID is a unique identifier for the stock.
	StockID int64 `json:"stockid"`

	// Name is the name of the stock.
	Name string `json:"name"`

	// Price is the current price of the stock.
	Price int64 `json:"price"`

	// Company is the name of the company associated with the stock.
	Company string `json:"company"`
}

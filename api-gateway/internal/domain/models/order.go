package models

type Order struct {
	ID     string  `json:"id"`
	BookID int64   `json:"book_id"`
	Place  string  `json:"place"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

package dto

type AddOrderInput struct {
	BookId int64   `json:"book_id"`
	Place  string  `json:"place"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

type AddOrderOutput struct {
	ID string `json:"id"`
}

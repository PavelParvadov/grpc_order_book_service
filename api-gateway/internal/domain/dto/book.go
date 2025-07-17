package dto

type AddBookInput struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

type AddBookOutput struct {
	ID int64 `json:"id"`
}

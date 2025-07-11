package models

type Order struct {
	Id     int64
	BookId int64
	Status string
	Price  float64
	Place  string
}

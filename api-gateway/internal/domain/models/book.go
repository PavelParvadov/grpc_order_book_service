package models

type Book struct {
	ID     int64  `json:"id,omitempty"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

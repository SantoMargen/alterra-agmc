package models

type Book struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Page     int    `json:"page"`
	Category string `json:"category"`
	Author   string `json:"author"`
	Price    int    `json:"price"`
}

package models

// Book represents the structure of a book in the system.
type Book struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Author        string  `json:"author"`
	Price         float64 `json:"price"`
	Description   string  `json:"description"`
	PublishedYear int     `json:"published_year"`
}
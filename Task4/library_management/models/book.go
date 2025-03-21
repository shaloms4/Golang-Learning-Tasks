package models

// Book structure to represent a book in the library
type Book struct {
	ID     int
	Title  string
	Author string
	Status string // "Available" or "Reserved"
}

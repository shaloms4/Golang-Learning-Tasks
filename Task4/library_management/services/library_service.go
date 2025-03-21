package services

import (
	"errors"
	"fmt"
	"library_management/models"
	"sync"
	"time"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
	ReserveBook(bookID int, memberID int) error // New method
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]*models.Member
	mu      sync.Mutex // Mutex to prevent race conditions
}

func NewLibrary() *Library {
	library := &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]*models.Member),
	}

	// Load predefined members from models.Members
	for i := range models.Members {
		member := &models.Members[i]
		library.Members[member.ID] = member
	}

	return library
}

func (l *Library) AddBook(book models.Book) {
	l.mu.Lock()
	defer l.mu.Unlock()

	book.Status = "Available"
	l.Books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	delete(l.Books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}

	member, exists := l.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	book.Status = "Borrowed"
	l.Books[bookID] = book

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	member, exists := l.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status == "Available" {
		return errors.New("book is not currently borrowed")
	}

	book.Status = "Available"
	l.Books[bookID] = book

	// Remove the book from the borrowed list
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	l.mu.Lock()
	defer l.mu.Unlock()

	var availableBooks []models.Book
	for _, book := range l.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	l.mu.Lock()
	defer l.mu.Unlock()

	member, exists := l.Members[memberID]
	if !exists {
		return nil
	}
	return member.BorrowedBooks
}

// ReserveBook reserves a book for a member
func (l *Library) ReserveBook(bookID int, memberID int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status == "Reserved" || book.Status == "Borrowed" {
		return errors.New("book is already reserved or borrowed")
	}

	// Reserve the book for the member
	book.Status = "Reserved"
	l.Books[bookID] = book

	// Start a Goroutine to automatically cancel the reservation if not borrowed in 5 seconds
	go l.autoCancelReservation(bookID, memberID)

	// Simulate the reservation is processed asynchronously
	fmt.Printf("Book %d reserved for member %d\n", bookID, memberID)
	return nil
}

// autoCancelReservation cancels the reservation if not borrowed within 5 seconds
func (l *Library) autoCancelReservation(bookID int, memberID int) {
	time.Sleep(5 * time.Second)

	l.mu.Lock()
	defer l.mu.Unlock()

	book, exists := l.Books[bookID]
	if !exists || book.Status != "Reserved" {
		return
	}

	// If the book is still reserved after 5 seconds, cancel the reservation
	book.Status = "Available"
	l.Books[bookID] = book

	fmt.Printf("Reservation for book %d by member %d has been automatically cancelled\n", bookID, memberID)
}

package services

import (
	"errors"
	"library_management/concurrency"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ReserveBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Books             map[int]models.Book
	Members           map[int]*models.Member
	ReservationWorker *concurrency.ReservationWorker
}

func NewLibrary() *Library {
	library := &Library{
		Books:             make(map[int]models.Book),
		Members:           make(map[int]*models.Member),
		ReservationWorker: concurrency.NewReservationWorker(),
	}

	for i := range models.Members {
		member := &models.Members[i]
		library.Members[member.ID] = member
	}

	go library.ReservationWorker.Start()
	return library
}

func (l *Library) AddBook(book models.Book) {
	book.Status = "Available"
	l.Books[book.ID] = book

	l.ReservationWorker.SetBookAvailability(book.ID, true)
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.Books, bookID)

	l.ReservationWorker.RemoveBookAvailability(bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
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

	l.ReservationWorker.SetBookAvailability(bookID, false)

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
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

	l.ReservationWorker.SetBookAvailability(bookID, true)

	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}
	return nil
}

func (l *Library) ReserveBook(bookID int, memberID int) error {
	response := make(chan error)
	l.ReservationWorker.Requests <- concurrency.ReservationRequest{
		BookID:   bookID,
		MemberID: memberID,
		Response: response,
	}
	return <-response
}

func (l *Library) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range l.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := l.Members[memberID]
	if !exists {
		return nil
	}
	return member.BorrowedBooks
}

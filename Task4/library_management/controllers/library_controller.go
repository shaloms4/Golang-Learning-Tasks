package controllers

import (
	"fmt"
	"library_management/models"
	"library_management/services"
)

type LibraryController struct {
	library *services.Library
}

func NewLibraryController(library *services.Library) *LibraryController {
	return &LibraryController{library: library}
}

func (lc *LibraryController) Run() {
	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. Reserve Book")
		fmt.Println("6. List Available Books")
		fmt.Println("7. List Borrowed Books")
		fmt.Println("8. Exit")
		fmt.Print("Enter choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id int
			var title, author string
			fmt.Print("Enter Book ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Book Title: ")
			fmt.Scan(&title)
			fmt.Print("Enter Book Author: ")
			fmt.Scan(&author)
			lc.library.AddBook(models.Book{ID: id, Title: title, Author: author})
			fmt.Println("Book added successfully!")

		case 2:
			var bookID int
			fmt.Print("Enter Book ID to remove: ")
			fmt.Scan(&bookID)
			lc.library.RemoveBook(bookID)
			fmt.Println("Book removed successfully!")

		case 3:
			var bookID, memberID int
			fmt.Print("Enter Book ID to borrow: ")
			fmt.Scan(&bookID)
			fmt.Print("Enter Member ID: ")
			fmt.Scan(&memberID)
			err := lc.library.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book borrowed successfully!")
			}

		case 4:
			var bookID, memberID int
			fmt.Print("Enter Book ID to return: ")
			fmt.Scan(&bookID)
			fmt.Print("Enter Member ID: ")
			fmt.Scan(&memberID)
			err := lc.library.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book returned successfully!")
			}

		case 5:
			var bookID, memberID int
			fmt.Print("Enter Book ID to reserve: ")
			fmt.Scan(&bookID)
			fmt.Print("Enter Member ID: ")
			fmt.Scan(&memberID)
			err := lc.library.ReserveBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book reserved successfully!")
			}

		case 6:
			fmt.Println("Available Books:")
			for _, book := range lc.library.ListAvailableBooks() {
				fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
			}

		case 7:
			var memberID int
			fmt.Print("Enter Member ID: ")
			fmt.Scan(&memberID)
			fmt.Println("Borrowed Books:")
			for _, book := range lc.library.ListBorrowedBooks(memberID) {
				fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
			}

		case 8:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice! Please try again.")
		}
	}
}

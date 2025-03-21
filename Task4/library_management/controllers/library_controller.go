package controllers

import (
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

func ShowMenu() {
	fmt.Println("Library Management System")
	fmt.Println("1. Add Book")
	fmt.Println("2. Remove Book")
	fmt.Println("3. Reserve Book")
	fmt.Println("4. Borrow Book")
	fmt.Println("5. Return Book")
	fmt.Println("6. List Available Books")
	fmt.Println("7. List Borrowed Books")
	fmt.Println("8. Exit")
}

func ReadInput(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return strings.TrimSpace(input)
}

func HandleLibraryActions() {
	library := services.NewLibrary()

	for {
		ShowMenu()
		action := ReadInput("Choose an action (1-8): ")

		switch action {
		case "1":
			title := ReadInput("Enter book title: ")
			author := ReadInput("Enter author: ")
			library.AddBook(models.Book{Title: title, Author: author, Status: "Available"})
		case "2":
			bookID, _ := strconv.Atoi(ReadInput("Enter book ID to remove: "))
			library.RemoveBook(bookID)
		case "3":
			bookID, _ := strconv.Atoi(ReadInput("Enter book ID to reserve: "))
			memberID, _ := strconv.Atoi(ReadInput("Enter member ID: "))
			err := library.ReserveBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "4":
			bookID, _ := strconv.Atoi(ReadInput("Enter book ID to borrow: "))
			memberID, _ := strconv.Atoi(ReadInput("Enter member ID: "))
			err := library.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "5":
			bookID, _ := strconv.Atoi(ReadInput("Enter book ID to return: "))
			memberID, _ := strconv.Atoi(ReadInput("Enter member ID: "))
			err := library.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "6":
			books := library.ListAvailableBooks()
			fmt.Println("Available Books:")
			for _, book := range books {
				fmt.Printf("%d. %s by %s\n", book.ID, book.Title, book.Author)
			}
		case "7":
			memberID, _ := strconv.Atoi(ReadInput("Enter member ID to list borrowed books: "))
			books := library.ListBorrowedBooks(memberID)
			fmt.Println("Borrowed Books:")
			for _, book := range books {
				fmt.Printf("%d. %s by %s\n", book.ID, book.Title, book.Author)
			}
		case "8":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

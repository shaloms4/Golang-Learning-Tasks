# Library Management System Documentation

## Overview
The **Console-Based Library Management System** is a simple Go application that allows users to manage a library through a command-line interface. It provides functionalities for adding, removing, borrowing, and returning books, as well as listing available and borrowed books.

## Features
- Add new books to the library
- Remove existing books
- Borrow books (if available)
- Return borrowed books
- List available books
- List borrowed books by a specific member

## Data Structures
### Book
```go
struct Book {
    ID     int
    Title  string
    Author string
    Status string // "Available" or "Borrowed"
}
```
### Member
```go
struct Member {
    ID            int
    Name          string
    BorrowedBooks []Book
}
```

## LibraryManager Interface
```go
interface LibraryManager {
    AddBook(book Book)
    RemoveBook(bookID int)
    BorrowBook(bookID int, memberID int) error
    ReturnBook(bookID int, memberID int) error
    ListAvailableBooks() []Book
    ListBorrowedBooks(memberID int) []Book
}
```

## Console Menu
The system provides an interactive menu:
```
1. Add Book
2. Remove Book
3. Borrow Book
4. Return Book
5. List Available Books
6. List Borrowed Books
7. Exit
```

## Setup and Execution
### 1. Initialize the Go Module
Run the following command inside the project directory:
```sh
go mod init library_management
```

### 2. Build and Run the Program
```sh
go run main.go
```

## Error Handling
- If a book does not exist, an error message is displayed.
- If a member tries to borrow a book that is already borrowed, an error is shown.
- If an invalid choice is made in the console, the system prompts the user to try again.

## Future Enhancements
- Implement database storage instead of in-memory storage.
- Add a user authentication system.
- Introduce book reservation and due dates.

## Contributor
- **Shalom Habtamu**



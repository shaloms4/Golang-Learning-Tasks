# Console-Based Library Management System with Concurrency

## Features:
- **Borrowing and returning books**: Books can be borrowed and returned with an interface to manage book statuses.
- **Add and remove books**: Allows for the addition and removal of books from the system.
- **Concurrent Book Reservation**: Books can be reserved concurrently. If not borrowed in 5 seconds, reservations are automatically cancelled.
- **Lists of Available and Borrowed Books**: Provides functions to list all available books and all books borrowed by a member.

## Concurrency Approach:
- **Goroutines** are used to handle multiple reservation requests simultaneously.
- **Channels** are used to queue reservation requests and manage their completion asynchronously.
- **Mutexes** ensure that updates to book statuses and member records are safe from race conditions, allowing only one process to update the system at a time.

## Console Interface:
1. **Add Book**
2. **Remove Book**
3. **Reserve Book**
4. **Borrow Book**
5. **Return Book**
6. **List Available Books**
7. **List Borrowed Books**
8. **Exit**

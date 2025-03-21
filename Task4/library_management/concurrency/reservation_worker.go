package concurrency

import (
	"fmt"
	"time"
)

var reservationTimeout = 5 * time.Second
var bookReservations = make(map[int]chan bool) // Channel to track reservations

// ProcessReservation handles the book reservation and auto-cancellation if not borrowed
func ProcessReservation(bookID int, memberID int) {
	// Create a channel to handle reservation
	reservationChannel := make(chan bool)
	bookReservations[bookID] = reservationChannel

	// Start a goroutine to handle the auto-cancellation
	go func() {
		time.Sleep(reservationTimeout)
		// Check if reservation is still not borrowed
		select {
		case <-reservationChannel:
			// Reservation successful, member borrowed the book
			fmt.Printf("Book %d borrowed by member %d\n", bookID, memberID)
		default:
			// Timeout reached, cancel reservation
			fmt.Printf("Reservation for book %d by member %d cancelled due to timeout\n", bookID, memberID)
			bookReservations[bookID] = nil
		}
	}()
}

// CancelReservation cancels a reservation if it's not borrowed within the timeout period
func CancelReservation(bookID int) {
	if ch, exists := bookReservations[bookID]; exists && ch != nil {
		ch <- true // Signal that the book was borrowed
	}
}

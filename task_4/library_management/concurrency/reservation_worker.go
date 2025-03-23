package concurrency

import (
	"errors"
	"sync"
	"time"
)

type ReservationRequest struct {
	BookID   int
	MemberID int
	Response chan error
}

type ReservationWorker struct {
	Requests       chan ReservationRequest
	mu             sync.Mutex
	availableBooks map[int]bool
}

func NewReservationWorker() *ReservationWorker {
	return &ReservationWorker{
		Requests:       make(chan ReservationRequest),
		availableBooks: make(map[int]bool),
	}
}

func (rw *ReservationWorker) SetBookAvailability(bookID int, available bool) {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	rw.availableBooks[bookID] = available
}

func (rw *ReservationWorker) RemoveBookAvailability(bookID int) {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	delete(rw.availableBooks, bookID)
}

func (rw *ReservationWorker) IsBookAvailable(bookID int) bool {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	return rw.availableBooks[bookID]
}

func (rw *ReservationWorker) ReserveBook(bookID int) {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	rw.availableBooks[bookID] = false
}

func (rw *ReservationWorker) ReleaseBook(bookID int) {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	rw.availableBooks[bookID] = true
}

func (rw *ReservationWorker) Start() {
	for req := range rw.Requests {
		if rw.IsBookAvailable(req.BookID) {
			rw.ReserveBook(req.BookID)

			// Auto-cancel reservation after a timeout
			time.AfterFunc(5*time.Second, func() {
				rw.ReleaseBook(req.BookID)
			})

			req.Response <- nil
		} else {
			req.Response <- errors.New("book is not available for reservation")
		}
	}
}

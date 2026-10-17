package booking

import "errors"

var (
	ErrorSeatAlreadyBooked = errors.New("seat is already bokoed")
)

type Booking struct {
	ID      string
	MovieID string
	UserID  string
	Status  string
	SeatID  string
}

type BookingStore interface {
	Book(b Booking) error
	ListBookings(movieID string) []Booking
}

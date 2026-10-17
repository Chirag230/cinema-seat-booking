package booking

import (
	"context"
	"errors"
	"time"
)

var (
	ErrorSeatAlreadyBooked = errors.New("seat is already bokoed")
)

type Booking struct {
	ID        string
	MovieID   string
	UserID    string
	Status    string
	SeatID    string
	ExpiresAt time.Time
}

type BookingStore interface {
	Book(b Booking) (Booking, error)
	ListBookings(movieID string) []Booking

	Confirm(ctx context.Context, sessionID string, userID string) (Booking, error)
	Release(ctx context.Context, sessionID string, userID string) error
}

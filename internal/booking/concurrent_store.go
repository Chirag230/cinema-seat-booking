package booking

import "sync"

type ConcurrentStore struct {
	bookings map[string]Booking
	sync.RWMutex
}

func NewConcurrentStore() *ConcurrentStore {
	return &ConcurrentStore{
		bookings: map[string]Booking{},
	}
}

func (c *ConcurrentStore) Book(b Booking) error {
	c.Lock()
	defer c.Unlock()
	if _, exists := c.bookings[b.SeatID]; exists {
		return ErrorSeatAlreadyBooked
	}
	c.bookings[b.SeatID] = b
	return nil
}

func (c *ConcurrentStore) ListBookings(movieID string) []Booking {
	// bookings for particular movie
	c.RLock()
	defer c.RUnlock()
	var result []Booking
	for _, b := range c.bookings {
		if b.MovieID == movieID {
			result = append(result, b)
		}
	}
	return result
}

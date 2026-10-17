package booking

type MemoryStore struct {
	// seat---->booking
	bookings map[string]Booking //"C2"-->booking
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: map[string]Booking{},
	}
}

func (s *MemoryStore) Book(b Booking) error {
	//first check if seat is available / empty
	// if available .. select it
	if _, exists := s.bookings[b.SeatID]; exists {
		return ErrorSeatAlreadyBooked
	}
	s.bookings[b.SeatID] = b
	return nil
}

func (s *MemoryStore) ListBookings(movieID string) []Booking {
	var result []Booking
	for _, b := range s.bookings {
		if b.MovieID == movieID {
			result = append(result, b)
		}
	}
	return result
}

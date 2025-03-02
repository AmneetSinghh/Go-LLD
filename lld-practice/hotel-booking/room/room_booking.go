package room

import (
	"GO-LLD/lld-practice/hotel-booking/payment"
	"os/user"
	"time"
)

type RoomBookingService interface {
	CreateBooking(req *BookingRequest) (*BookingEntry, error)
	DeleteBooking(bookingId int) error
	GetBookingsList() []*BookingEntry
	CancelBooking(bookingId int) (*BookingEntry, error)
	GetBookingById(bookingId int) *BookingEntry
}

type BookingRequest struct {
	bookedRooms    []Room
	bookedBy       user.User
	bookedForDays  int
	checkedInDate  time.Time
	checkedOutDate time.Time
	roomServices   []string // price calculation
}

type BookingEntry struct {
	bookingId      int
	bookedRooms    []Room
	bookedBy       user.User
	bookedForDays  int
	bookingStatus  BookingStatus
	checkedInDate  time.Time
	checkedOutDate time.Time
	totalPrice     int
	PaymentDetails payment.PaymentDetails
}

type BookingCancellation struct {
	bookingEntry  BookingEntry
	refundDetails payment.Refund
}

type RoomBooking struct {
	RoomManagerService RoomManagerService
	RoomBookings       []*BookingEntry
	BookingEntryMap    map[int]*BookingEntry
}

func NewRoomBookingService(rm RoomManagerService) RoomBookingService {
	return &RoomBooking{
		RoomBookings:       []*BookingEntry{},
		BookingEntryMap:    map[int]*BookingEntry{},
		RoomManagerService: rm, // for asking room details
	}
}

func (rb *RoomBooking) CreateBooking(req *BookingRequest) (*BookingEntry, error) {
	return nil, nil
}
func (rb *RoomBooking) DeleteBooking(bookingId int) error {
	return nil
}
func (rb *RoomBooking) GetBookingsList() []*BookingEntry {
	return nil
}
func (rb *RoomBooking) CancelBooking(bookingid int) (*BookingEntry, error) {
	return nil, nil
}
func (rb *RoomBooking) GetBookingById(bookingId int) *BookingEntry {
	return nil
}
func (rb *RoomBooking) GetBookingByUser(bookedBy user.User) []*BookingEntry {
	return nil
}

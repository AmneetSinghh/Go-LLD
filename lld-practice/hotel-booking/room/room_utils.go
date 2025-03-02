package room

type RoomType string

const (
	DELUX      RoomType = "delux"
	DOUBLE_BED RoomType = "double_bed"
	LUXURY     RoomType = "luxury"
	STANDARD   RoomType = "standard"
	SUITE      RoomType = "suite"
)

type BookingStatus string

const (
	PENDING     BookingStatus = "pending"
	CONFIRMED   BookingStatus = "confirmed"
	CHECKED_IN  BookingStatus = "checked_in"
	CHECKED_OUT BookingStatus = "checked_out"
	CANCELLED   BookingStatus = "cancelled"
)

type RoomStatus string

const (
	BOOKED        RoomStatus = "booked"
	NOT_AVAILABLE RoomStatus = "not_available"
	AVAILABLE     RoomStatus = "available"
)

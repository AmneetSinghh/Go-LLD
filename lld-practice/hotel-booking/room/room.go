package room

import "time"

type RoomService interface {
	GetPricing() int // based on room services.
}

type Room struct {
	roomId        int
	roomStatus    RoomStatus
	roomType      RoomType
	RoomNo        int
	Amenities     []string
	BasePrice     int
	disCountPrice int
	maxGuests     int
	isAvailable   int
	lastCleaned   time.Time
}

func NewRoom(rt RoomType) RoomService {
	var room RoomService
	switch rt {
	case DELUX:
		room = &DeluxRoom{}
	case SUITE:
		room = &SuiteRoom{}
	case STANDARD:
		room = &StandardRoom{}
	}
	return room
}

type DeluxRoom struct {
	Room
}

func (d *DeluxRoom) GetPricing() int {
	// calculate price based on aminities.
	return 0
}

type StandardRoom struct {
	Room
}

func (d *StandardRoom) GetPricing() int {
	return 0
}

type SuiteRoom struct {
	Room
}

func (d *SuiteRoom) GetPricing() int {
	return 0
}

package service

import (
	"GO-LLD/lld-practice/hotel-booking/model"
	"GO-LLD/lld-practice/hotel-booking/room"
	"fmt"
	"math/rand"
	"time"
)

// responsibility hotel level operations.....

type StayService interface {
	CheckIn()
	CheckOut()
	GetRoomManager() room.RoomManagerService
	GetRoomBooking() room.RoomBookingService
	GetRoomAdmin() room.RoomAdminService
	GetStayDetails() *Stay
}

type Stay struct {
	stayid    string
	Name      string
	Address   string
	Location  model.Location
	Amenities []string // wifi, goal,

	// only passed outside through interfaces....
	roomManagerService room.RoomManagerService
	roomAdminService   room.RoomAdminService // only StayHotel can add rooms.... so Room-admin service.
	roomBookingService room.RoomBookingService
}

func NewStay(st StayType) StayService {

	rmWrap := room.NewRoomManager()
	var roomManager room.RoomManagerService = rmWrap
	var roomAdmin room.RoomAdminService = rmWrap
	roomBooking := room.NewRoomBookingService(roomManager)

	switch st {
	case HOTEL:
		return &Hotel{
			Stay: Stay{
				stayid:             fmt.Sprintf("hotel_%v", generateRandomStayID()),
				roomManagerService: roomManager,
				roomAdminService:   roomAdmin,
				roomBookingService: roomBooking,
			},
		}
	case HOME_STAY:
		return &HomeStay{
			Stay: Stay{
				stayid:             fmt.Sprintf("homestay_%v", generateRandomStayID()),
				roomManagerService: roomManager,
				roomAdminService:   roomAdmin,
				roomBookingService: roomBooking,
			},
		}
	case VILLA:
		return &Villa{
			Stay: Stay{
				stayid:             fmt.Sprintf("villa_%v", generateRandomStayID()),
				roomManagerService: roomManager,
				roomAdminService:   roomAdmin,
				roomBookingService: roomBooking,
			},
		}
	default:
		return nil
	}
}

func (s *Stay) CheckIn() {

}

func (s *Stay) CheckOut() {

}

func (s *Stay) GetRoomManager() room.RoomManagerService {
	return s.roomManagerService
}

func (s *Stay) GetRoomBooking() room.RoomBookingService {
	return s.roomBookingService
}

func (s *Stay) GetRoomAdmin() room.RoomAdminService {
	return s.roomAdminService
}

func (s *Stay) GetStayDetails() *Stay {
	return s
}

type HomeStay struct {
	Stay
}

func (s *HomeStay) GetPriceDetails() string {
	return ""
}

type Villa struct {
	Stay
}

func (s *Villa) GetPriceDetails() string {
	return ""
}

type Hotel struct {
	Stay
}

func (s *Hotel) GetPriceDetails() string {
	return ""
}

func generateRandomStayID() int {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	return rand.Intn(100) + 1        // Generate a random number between 1 and 100
}

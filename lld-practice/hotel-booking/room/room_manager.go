package room

import "fmt"

// Room manager is inventory of room and source of truth of room data.
// if anyone wants to get room data .. ask RoomManagerService interface.

type RoomManagerService interface {
	GetRoom(roomNumber int) (*Room, error)
	IsRoomAvailable(roomNumber int) bool
	// room details..
}

type RoomAdminService interface {
	RoomManagerService
	UpdateRoom(room *Room) error
	AddRoom(room *Room) error
	RemoveRoom(roomNo int)
}

type RoomManager struct {
	roomList  []*Room
	roomMap   map[int]*Room // roomNo, Room
	RoomTypes []RoomType    // different room types
}

type RoomManagerWrapper struct {
	RoomAdminService   RoomAdminService
	RoomManagerService RoomManagerService
}

func NewRoomManager() *RoomManagerWrapper {
	rm := &RoomManager{}
	return &RoomManagerWrapper{
		RoomAdminService:   rm,
		RoomManagerService: rm,
	}
}

func (rm *RoomManager) AddRoom(room *Room) error {
	if room == nil {
		return fmt.Errorf("Room is nil, insert failed")
	}
	if _, ok := rm.roomMap[room.RoomNo]; ok {
		return fmt.Errorf("Room with number %d already exist", room.RoomNo)
	}
	rm.roomList = append(rm.roomList, room)
	rm.roomMap[room.RoomNo] = room
	return nil
}

func (rm *RoomManager) RemoveRoom(roomNo int) {
	delete(rm.roomMap, roomNo)
	filteredList := []*Room{}
	for _, room := range rm.roomList {
		if room.RoomNo != roomNo {
			filteredList = append(filteredList, room)
		}
	}
	rm.roomList = filteredList
}

func (rm *RoomManager) UpdateRoom(updatedRoom *Room) error {
	if updatedRoom == nil {
		return fmt.Errorf("Room is nil, update failed")
	}
	rm.roomMap[updatedRoom.RoomNo] = updatedRoom
	for idx, room := range rm.roomList {
		if room.RoomNo == updatedRoom.RoomNo {
			rm.roomList[idx] = room
			break
		}
	}
	return nil
}

func (rm *RoomManager) GetRoom(roomNumber int) (*Room, error) {
	room, ok := rm.roomMap[roomNumber]
	if !ok || room == nil {
		return nil, fmt.Errorf("room does't exist")
	}
	return room, nil
}

func (rm *RoomManager) IsRoomAvailable(roomNumber int) bool {
	return true
}

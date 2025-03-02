package service

// only get interface for listing.
type StayManagerQueryService interface {
	GetStayList() []*Stay
	GetStayById(stayid string) *Stay
}

// this interface means modification and get both
type StaymanagerService interface {
	StayManagerQueryService
	AddStay(stay *Stay)
	RemoveStay(stayid string)
	UpdateStay(stay *Stay)
}

type StayManager struct {
	stays []*Stay
}

func NewStayManager() *StayManager {
	return &StayManager{}
}

func (sm *StayManager) AddStay(stay *Stay) {
	sm.stays = append(sm.stays, stay)
}

func (sm *StayManager) RemoveStay(stayId string) {
	filteredStays := []*Stay{}
	for _, stay := range sm.stays {
		if stayId != stay.stayid {
			filteredStays = append(filteredStays, stay)
		}
	}
	sm.stays = filteredStays
}

func (sm *StayManager) UpdateStay(updatedStay *Stay) {
	for idx, stay := range sm.stays {
		if updatedStay.stayid == stay.stayid {
			sm.stays[idx] = updatedStay
			break
		}
	}
}

func (sm *StayManager) GetStayList() []*Stay {
	return sm.stays
}

func (sm *StayManager) GetStayById(stayId string) *Stay {
	for _, stay := range sm.stays {
		if stay.stayid == stayId {
			return stay
		}
	}
	return nil
}

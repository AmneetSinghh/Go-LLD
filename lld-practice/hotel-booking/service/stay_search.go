package service

import (
	"GO-LLD/lld-practice/hotel-booking/model"
)

type StaySearchReq struct {
	Location  model.Location
	Duration  StayDuration
	Occupancy StayOccupancy
	StayName  string
	stayType  StayType
}

type StayDuration struct {
	CheckinDate  string
	CheckOutDate string
}

type StayOccupancy struct {
	Guests int
	Rooms  int
}

type StaySearchService interface {
	SearchStays(query StaySearchReq) (StayListing, error)
}

func NewStaySearch(stayListingService StayListingService) StaySearchService {
	return &StaySearch{
		stayListingService: stayListingService,
	}
}

type StaySearch struct {
	stayListingService StayListingService
}

func (ss *StaySearch) SearchStays(query *StaySearchReq) (StayListingService, error) {

}

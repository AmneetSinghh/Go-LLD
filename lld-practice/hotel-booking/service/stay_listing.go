package service

type StayFilter struct {
	PriceFilter  PriceFilterStrategy  `json:"price_filter"`
	RatingFilter RatingFilterStrategy `json:"rating_filter"`
}

type StaySort struct {
	SortBy    SortByStrategy    `json:"sort_by"`
	SortOrder SortOrderStrategy `json:"sort_order"`
}

type StayListingService interface {
	GetStays() []Stay
	GetPaginatedStays() []Stay
	GetCurrentAppliedFilters() StayFilter
	GetCurrentSortOrder() SortOrderStrategy
	FilterStays(stayFilter StayFilter, query StaySearch) (StayListing, error)
}

func NewStayListing() StayListingService {
	return &StayListing{}
}

type StayListing struct {
	StayManager StayManagerQueryService
}

// filtering , sorting, pagination part of this.

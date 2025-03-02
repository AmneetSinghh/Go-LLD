package service

type SortByStrategy string

const (
	SORT_BY_PRICE    SortByStrategy = "sort_by_price"
	SORT_BY_RATING   SortByStrategy = "sort_by_rating"
	SORT_BY_DISTANCE SortByStrategy = "sort_by_distance"
)

type SortOrderStrategy string

const (
	ASCENDING  SortOrderStrategy = "asc"
	DESCENDING SortOrderStrategy = "desc"
)

type PriceFilterStrategy string

const (
	LESS_THAN_1000 PriceFilterStrategy = "less_than_1000"
	LESS_THAN_2000 PriceFilterStrategy = "less_than_2000"
	LESS_THAN_3000 PriceFilterStrategy = "less_than_3000"
	ABOVE_3000     PriceFilterStrategy = "above_3000"
)

type RatingFilterStrategy string

const (
	RATING_BELOW_3 RatingFilterStrategy = "rating_below_3"
	RATING_3_TO_4  RatingFilterStrategy = "rating_3_to_4"
	RATING_ABOVE_4 RatingFilterStrategy = "rating_above_4"
)

type StayType string

const (
	HOTEL     StayType = "hotel"
	HOME_STAY StayType = "home_stay"
	VILLA     StayType = "villa"
)

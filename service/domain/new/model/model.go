package model

type DataRequest struct {
	Search     string   `json:"search"`
	LowPrice   *float64 `json:"low_price"`
	HighPrice  *float64 `json:"high_price"`
	Restaurant string   `json:"restaurant"`
}

type DataResponse struct {
	Id                int32   `json:"id"`
	Name              string  `json:"name"`
	RestaurantId      int32   `json:"restaurant_id"`
	RestaurantName    string  `json:"restaurant_name"`
	RestaurantAddress string  `json:"restaurant_address"`
	Price             float64 `json:"price"`
}

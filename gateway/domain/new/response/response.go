package response

type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Status     string      `json:"status"`
	Timestamp  string      `json:"timestamp"`
	Data       interface{} `json:"data"`
}

type Restaurant struct {
	Id      int32  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type GetDataResponse struct {
	Id         int32      `json:"id"`
	Name       string     `json:"name"`
	Restaurant Restaurant `json:"restaurant"`
	Price      float64    `json:"price"`
}

type GetListDataResponse struct {
	Lists []GetDataResponse `json:"lists"`
}

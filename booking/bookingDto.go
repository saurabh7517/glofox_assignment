package booking

//BookingDto - date transfer object for booking api
type bookingdto struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

//BookingDtoResponse ..
type bookingDtoResponse struct {
	ID  string `json:"id"`
	Msg string `json:"msg"`
}

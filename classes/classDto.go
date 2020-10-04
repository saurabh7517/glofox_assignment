package classes

//ClassDto - data transfer object for classes api.
type classdto struct {
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

//responseDtp to be sent back to the client upon successful request
type responseClassDto struct {
	ID  string `json: "id"`
	Msg string `json: "msg"`
}

package booking

import (
	"encoding/json"
	"fmt"
	"github/saurabh7517/glofox_assignment/utils"
	"log"
	"net/http"
	"strconv"
)

//Controller ..Booking Controller
type Controller struct {
	URL string
}

const urlpathvar string = "bookings"

func (bc Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var path string = utils.ConfirmStringPathVariable(r.URL.Path, urlpathvar)

	if urlpathvar == path {
		switch r.Method {
		case http.MethodGet:

			bookingList := getBookings()
			if bytedata, err := json.Marshal(bookingList); err != nil {
				log.Print("Error in marshalling list to json")
				w.WriteHeader(http.StatusBadRequest)
				return
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write(bytedata)
			}

		case http.MethodPost:
			var bookingDto *bookingdto = &bookingdto{}
			var response *bookingDtoResponse = &bookingDtoResponse{}
			if err := json.NewDecoder(r.Body).Decode(bookingDto); err != nil {
				log.Print("Improper formatting of JSON")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if id, err := addbooking(bookingDto); err != nil {
				response.ID = strconv.Itoa(id)
				response.Msg = err.Error()
			} else {
				response.ID = strconv.Itoa(id)
				response.Msg = "New Booking created"
			}

			if bytedata, err := json.Marshal(response); err != nil {
				log.Print("Error in marshalling response object to json")
			} else {
				w.WriteHeader(http.StatusCreated)
				w.Write(bytedata)
			}
		case http.MethodPut:
			var bookingDto *bookingdto = &bookingdto{}
			var response *bookingDtoResponse = &bookingDtoResponse{}
			var id int
			var flag bool = false
			if flag, id = utils.ConfirmIntPathVariable(r.URL.Path); flag == false {
				log.Print("Improper URL Path")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("{\"Msg\":\"Wrong URL path\"}"))
				return
			}

			if err := json.NewDecoder(r.Body).Decode(bookingDto); err != nil {
				log.Print("Improper formatting of JSON")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if id, err := updatebooking(id, bookingDto); err != nil {
				response.Msg = err.Error()
			} else {
				response.ID = strconv.Itoa(id)
				response.Msg = fmt.Sprintf("Booking with %v is updated", id)

			}
			if bytedata, err := json.Marshal(response); err != nil {
				log.Print("Error in marshalling response object to json")
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusCreated)
				w.Write(bytedata)
			}
		case http.MethodDelete:
			// var bookingDto *BookingDto = &BookingDto{}
			var response *bookingDtoResponse = &bookingDtoResponse{}
			var id int
			var flag bool = false
			if flag, id = utils.ConfirmIntPathVariable(r.URL.Path); flag == false {
				log.Print("Improper URL Path")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("{\"Msg\":\"Wrong URL path\"}"))
				return
			}

			if id, err := removeBooking(id); err != nil {
				response.Msg = err.Error()
				w.WriteHeader(http.StatusNotFound)
			} else {
				response.ID = strconv.Itoa(id)
				response.Msg = fmt.Sprintf("Id %v is deleted", id)
				w.WriteHeader(http.StatusCreated)
			}
			if bytedata, err := json.Marshal(response); err != nil {
				log.Print("Error in marshalling response object to json")
				w.WriteHeader(http.StatusInternalServerError)
			} else {

				w.Write(bytedata)
			}

		case http.MethodOptions:
			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

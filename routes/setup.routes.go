package routes

import (
	"fmt"
	"github/saurabh7517/glofox_assignment/booking"
	"github/saurabh7517/glofox_assignment/classes"

	"net/http"
)

const classAPIPath = "classes"
const bookingAPIPath = "bookings"

//SetupRoutes ...
func SetupRoutes(baseAPIPath string) {
	//Default Serve Mux uses the url pattern and mathces to the closest string.
	var programList string = fmt.Sprintf("%s/%s", baseAPIPath, classAPIPath)
	var program string = fmt.Sprintf("%s/%s/", baseAPIPath, classAPIPath)

	var bookingList string = fmt.Sprintf("%s/%s", baseAPIPath, bookingAPIPath)
	var bookingString string = fmt.Sprintf("%s/%s/", baseAPIPath, bookingAPIPath)

	// b := booking.BookingController{url : bookingList}
	// b := classes.ClassController{url: "kjfdas"}

	http.Handle(programList, addFilters(classes.ClassController{URL: programList}))
	http.Handle(program, addFilters(classes.ClassController{URL: program}))
	http.Handle(bookingList, addFilters(booking.Controller{URL: bookingList}))
	http.Handle(bookingString, addFilters(booking.Controller{URL: bookingString}))

}

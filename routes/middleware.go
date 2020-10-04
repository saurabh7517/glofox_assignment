package routes

import "net/http"

//AddFilters filters before and after request
func addFilters(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Additions made before processing the request goes here
		// ...
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		handler.ServeHTTP(w, r)
		//Additions made to the header after the response

	})
}

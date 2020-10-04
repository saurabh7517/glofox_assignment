package classes

import (
	"encoding/json"
	"fmt"
	"github/saurabh7517/glofox_assignment/utils"
	"log"
	"net/http"
	"strconv"
)

//ClassController ..
type ClassController struct {
	URL string
}

const urlpathvar string = "classes"

func (cc ClassController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var path string = utils.ConfirmStringPathVariable(r.URL.Path, urlpathvar)

	if urlpathvar == path {
		switch r.Method {
		case http.MethodGet:
			classList := getClasses()
			if bytedata, err := json.Marshal(classList); err != nil {
				log.Print("Error in marshalling list to json")
				w.WriteHeader(http.StatusBadRequest)
				return
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write(bytedata)
			}
		case http.MethodPost:
			var classDto *classdto = &classdto{}
			var response *responseClassDto = &responseClassDto{}
			if err := json.NewDecoder(r.Body).Decode(classDto); err != nil {
				log.Print("Improper formatting of JSON")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if id, err := addclass(classDto); err != nil {
				response.ID = strconv.Itoa(id)
				response.Msg = err.Error()
			} else {
				response.ID = strconv.Itoa(id)
				response.Msg = "New Class created"
			}

			if bytedata, err := json.Marshal(response); err != nil {
				log.Print("Error in marshalling response object to json")
			} else {
				w.WriteHeader(http.StatusCreated)
				w.Write(bytedata)
			}
		case http.MethodPut:
			var classDto *classdto = &classdto{}
			var response *responseClassDto = &responseClassDto{}
			var id int
			var flag bool = false
			if flag, id = utils.ConfirmIntPathVariable(r.URL.Path); flag == false {
				log.Print("Improper URL Path")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("{\"Msg\":\"Wrong URL path\"}"))
				return
			}

			if err := json.NewDecoder(r.Body).Decode(classDto); err != nil {
				log.Print("Improper formatting of JSON")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if id, err := updateclass(id, classDto); err != nil {
				response.Msg = err.Error()
			} else {
				response.ID = strconv.Itoa(id)
				response.Msg = fmt.Sprintf("Class with id %v is updated", id)
			}

			if bytedata, err := json.Marshal(response); err != nil {
				log.Print("Error in marshalling response object to json")
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusCreated)
				w.Write(bytedata)
			}

		case http.MethodDelete:
			var response *responseClassDto = &responseClassDto{}
			var id int
			var flag bool = false
			if flag, id = utils.ConfirmIntPathVariable(r.URL.Path); flag == false {
				log.Print("Improper URL Path")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("{\"Msg\":\"Wrong URL path\"}"))
				return
			}

			if id, err := removeClass(id); err != nil {
				w.WriteHeader(http.StatusNotFound)
				response.Msg = err.Error()
			} else {
				response.ID = strconv.Itoa(id)
				response.Msg = fmt.Sprintf("Id %v is deleted", id)
				w.WriteHeader(http.StatusCreated)
			}
			if bytedata, err := json.Marshal(response); err != nil {
				log.Print("Error in marshalling response object to json")
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				// w.WriteHeader(http.StatusCreated)
				w.Write(bytedata)
			}
		case http.MethodOptions:
			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}

	}
}

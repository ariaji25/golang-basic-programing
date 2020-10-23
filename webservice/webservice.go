package webservice

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"example.com/ariaji25/learninggolang/lostnumber"
)

/* Handle Request.
Function to handle request from client to web service, with parameters :
- w : is http response writer to return server response to client.
- r : is http request to get what client request to server */
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// show the request method
	log.Println("Request Method : ", r.Method)
	// chekc the request method using switch case
	switch r.Method {
	// if post method
	case http.MethodPost:
		// post method is allowed Form Url Encoded as request Body
		// try parse the Form
		if err := r.ParseForm(); err != nil {
			// if error return htt error with status Bad Request
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// get value of form as input by key "numbers"
		numbers := r.FormValue("numbers")
		// show input
		fmt.Println(numbers)
		// find the lost number by call LostNumber() from lostnumber package
		lostnumber, err := lostnumber.LostNumber(string(numbers))
		// chekc if find an error
		if err != nil {
			// return http error with error mesaage and status badrequest
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// craete new response to client type a struct BaseBody
		response := BaseResponse{
			Message: "success get Lost Number, Input : " + numbers + ", Output : " + strconv.FormatFloat(lostnumber, 'f', 0, 64),
			Data:    lostnumber,
			Status:  201,
			IsError: false,
		}
		// add Content-Type of response body
		w.Header().Add("Content-Type", "application/json")
		// set teh status code
		w.WriteHeader(http.StatusCreated)
		// return json body by encode the response that created before
		json.NewEncoder(w).Encode(response)

		// if method is Get
	case http.MethodGet:
		// return Get message
		w.Write([]byte("Web Service to find a Lost Number of a given string that contains list of sorted numbers. "))
		break
	default:
		// return not allowed method
		w.WriteHeader(403)
		w.Write([]byte("Method Not Allowed"))
		break
	}
}

/* Web Service.
Web service to serve request to find Lost Number from given list of sorted list numbers.
The service running on port 8000 */
func WebService() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Start The Web Service")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

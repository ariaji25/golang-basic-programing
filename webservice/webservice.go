package webservice

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"example.com/ariaji25/learninggolang/lostnumber"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Request Method : ", r.Method)
	switch r.Method {
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		numbers := r.FormValue("numbers")
		fmt.Println(numbers)
		lostnumber, err := lostnumber.LostNumber(string(numbers))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response := BaseBody{
			Message: "success get Lost Number, Input : " + numbers + ", Output : " + strconv.FormatFloat(lostnumber, 'f', 0, 64),
			Data:    lostnumber,
			Status:  201,
			IsError: false,
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	case http.MethodGet:
		w.Write([]byte("Web Service to find a Lost Number of a given string that contains list of sorted numbers. "))
		break
	default:
		w.WriteHeader(403)
		w.Write([]byte("Method Not Allowed"))
		break
	}
}

func WebService() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Start The Web Service")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

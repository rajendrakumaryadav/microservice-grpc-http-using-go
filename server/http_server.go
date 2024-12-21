package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"myservice/db"
	"myservice/models"

	"github.com/gorilla/mux"
)

func StartHTTPServer() {
	r := mux.NewRouter()
	r.HandleFunc("/data", getData).Methods("GET")
	r.HandleFunc("/", getIndex).Methods("GET")
	log.Println(http.ListenAndServe(":8080", r))
}

func getData(w http.ResponseWriter, r *http.Request) {
	data, err := db.GetData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	data := `{"message": "Hello, world!"}`
	var resp models.DataResponse
	err := json.Unmarshal([]byte(data), &resp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Token", "developed by Rajendra~")
	json.NewEncoder(w).Encode(resp)

}

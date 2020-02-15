package main

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
)



type parking struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	NumberPatent string `json:"numberPatent"`
	Ingreso string `json:"timeIn"`
	Salida string `json:"timeOut"`
	Permanencia string `json:"timeStay"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/parking", postValues).Methods("GET")
	log.Fatal(http.ListenAndServe(":2000", router))	
}

func postValues(w http.ResponseWriter, r *http.Request){
	pParking := parking{"Diego", "Mondini", "KJFV68", "16h30m","20h00m", "3h30m"}
	json.NewEncoder(w).Encode(pParking)
	w.Header().Set("Content-Type", "application/json")
}

func getValues(){}

func updateValues(){}
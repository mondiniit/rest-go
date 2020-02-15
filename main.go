package main

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
)

type profile struct{
	DocumentType string `json:"documentType"`
	DocumentNumber string `json:"documentNumber"`
	FirstName string `json:"firstName"`
	MiddleName string `json:"middleName"`
	AditionalName  string `json:"additionalName"`
  	LastSurname  string `json:"lastSurname"`
  	AditionalSurname  string `json:"ditionalSurname"`
  	FullName  string `json:"fullName"`
  	Names  string `json:"names"`
  	Gender string `json:"gender"`
  	BirthDate string `json:"birthDate"`
  	Nationality  string `json:"nationality"`
   	MaritalStatus  string `json:"maritalStatus"`
}

type parking struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	NumberPatent string `json:"numberPatent"`
	timeIN string `json:"timeIN"`
	timeOUT string `json:"timeOUT"`
	timeSTAY string `json:"timeSTAY"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloName).Methods("GET")
	log.Fatal(http.ListenAndServe(":2000", router))	
}

func helloName(w http.ResponseWriter, r *http.Request) {
	pProfile := profile{"RUT", "23910105-4", "Diego", "","", "Mondini", "", "Diego Mondini", "Diego", "Male", "17-02-1987", "BR", "Married"}
	json.NewEncoder(w).Encode(pProfile)
	w.Header().Set("Content-Type", "application/json")
	
}

func postValues(){}

func getValues(){}

func updateValues(){}
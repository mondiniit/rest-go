package main

import (
   
	"fmt"
	"github.com/sendgrid/rest"
	// "github.com/gorilla/mux"
)

func main() {
	// router := mux.NewRouter()

	baseURL := "https://virtserver.swaggerhub.com/restafarians/Customer/1.0/customer/23910105-4"
	//router.HandleFunc("baseURL", getPosts).Methods("GET")
	method := rest.Get
	request := rest.Request{
		Method:  method,
		BaseURL: baseURL,
	}
	response, err := rest.Send(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

}

// func getPosts(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(posts)
//   }


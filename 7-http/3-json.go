package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	Password  string `json:"-"` // - means ignore the field in json output
	Email     string `json:"email"`
}

func main() {
	http.HandleFunc("/json", sendJson)

	err := http.ListenAndServe(":8080", nil)
	fmt.Println("server started on port 8080")
	if err != nil {
		panic(err)
	}
}

func sendJson(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	u := User{
		FirstName: "abc",
		Password:  "xyz",
		Email:     "@gmail.com",
	}

	//jsonData, err := json.Marshal(u)
	//if err != nil {
	//	//sending the error response
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return // make sure to return from the program, http.Error doesnt stop the program execution
	//}
	//
	//w.WriteHeader(http.StatusOK)
	//w.Write(jsonData)

	w.WriteHeader(http.StatusOK)
	// below line would convert struct to json and write the response over responseWriter
	err := json.NewEncoder(w).Encode(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

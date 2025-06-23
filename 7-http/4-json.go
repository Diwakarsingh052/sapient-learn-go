package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type student struct {
	FirstName string `json:"first_name"`
	Password  string `json:"-"` // - means ignore the field in json output
	Email     string `json:"email"`
}

func main() {
	http.HandleFunc("/json", recvJson)
	http.ListenAndServe(":8080", nil)
}

func recvJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "only post method is allowed", http.StatusMethodNotAllowed)
		return
	}
	//r.URL.Query().Get("name") // reading the query param
	// r.Header.Get("Content-Type") // request header
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var s student
	//fmt.Println(string(body))
	// unmarshal converts json into struct or map
	err = json.Unmarshal(body, &s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("%+v\n", s)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("read the body"))

}

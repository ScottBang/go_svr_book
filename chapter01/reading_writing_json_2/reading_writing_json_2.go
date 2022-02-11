package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldResponse2 struct {
	Message string `json:"message"`
	Author  string `json:"-"` // 출력안됨.
	Date    string `json:",omitempty"`
	ID      string `json:"id,string"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler2)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello World 2"}
	data, err := json.Marshal(response)

	if err != nil {
		panic("Ooops")
	}

	fmt.Fprint(w, string(data))
}

func helloWorldHandler2(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse2{
		Message: "Hello World 2",
		Author:  "scott",
		Date:    "2022/02/11",
		ID:      "scottbang",
	}
	data, err := json.Marshal(response)

	if err != nil {
		panic("Ooops")
	}

	fmt.Fprint(w, string(data))
}

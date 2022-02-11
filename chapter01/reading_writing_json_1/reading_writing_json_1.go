package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// consol: {"Message":"Hello World"}
	response := helloWorldResponse{Message: "Hello World"}

	// consol: ["message","hello world"]
	// response := [...]string{"message", "hello world"}
	//
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}
	// w 라는 파일에 writing 처리.
	fmt.Fprint(w, string(data))
}

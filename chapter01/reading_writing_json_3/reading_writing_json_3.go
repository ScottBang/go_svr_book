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

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "HelloWorld"}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func helloWorldHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\nProto: \n%v\n", r.Proto)
	fmt.Printf("\nProtoMajor: \n%v\n", r.ProtoMajor)
	fmt.Printf("\nProtoMinor: \n%v\n", r.ProtoMinor)
	fmt.Printf("\nTLS: \n%v\n", r.TLS)
	fmt.Printf("\nHeader: \n%v\n", r.Header)
	fmt.Printf("\nTrailer: \n%v\n", r.Trailer)
	fmt.Printf("\nTransferEncoding: \n%v\n", r.TransferEncoding)

	response := helloWorldResponse{Message: "Hello World!"}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

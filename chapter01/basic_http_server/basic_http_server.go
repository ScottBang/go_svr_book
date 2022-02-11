/*
Test:
	url에 이하 주소 입력함. http://localhost:8080/helloworld
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080
	// url 경로에 해당하는 Handler 함수를 지정.
	http.HandleFunc("/helloworld", helloWorldHandler)
	// log 출력
	log.Printf("Server starting on port %v\n", port)
	// Sprintf는 Formating 된 문자열을 리턴함.
	msg := fmt.Sprintf("listen tcp:%v\n", port)
	fmt.Println(msg)
	// HTTP 서버 시작
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}

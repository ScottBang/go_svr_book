package main

import (
	"fmt"
	"log"

	"rpc/client"
	"rpc/server"
)

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)
	log.Println()
}

package main

import (
	"fmt"
	"rpc_http/client"
	"rpc_http/server"
)

func main() {
	server.StartServer()
	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)

}

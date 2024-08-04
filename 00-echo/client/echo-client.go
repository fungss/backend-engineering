package main

import (
	"fmt"
	"net"
)

const (
	SERVER_PORT = "80"
	SERVER_TYPE = "tcp"
)

func main() {
	// Establish connection
	conn, err := net.Dial(SERVER_TYPE, ":"+SERVER_PORT)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	// Write to server
	_, errWrite := conn.Write([]byte("Hello, World!"))
	if errWrite != nil {
		fmt.Println(errWrite)
	}
}

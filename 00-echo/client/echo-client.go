package main

import (
	"fmt"
	"net"
	"os"
)

const (
	SERVER_PORT = "80"
	SERVER_TYPE = "tcp"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println(("Usage: echo-client 'Your message comes here!'"))
		os.Exit(1)
	}
	echo_msg := os.Args[1]

	// Establish connection
	conn, err := net.Dial(SERVER_TYPE, ":"+SERVER_PORT)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	// Write to server
	_, errWrite := conn.Write([]byte(echo_msg))
	if errWrite != nil {
		fmt.Println(errWrite)
	}
}

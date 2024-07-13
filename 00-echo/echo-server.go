package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

const (
	SERVER_PORT = "80"
	SERVER_TYPE = "tcp"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		if err != io.EOF {
			log.Println("Error reading from connection:", err)
		}
	}
	fmt.Println(string(buf[:n]))
}

func main() {
	fmt.Println("Server Running...")
	// Create and bind socket to port 80 and start listening
	server, err := net.Listen(SERVER_TYPE, ":"+SERVER_PORT)
	if err != nil {
		fmt.Printf("Error starting %s server: %s\n", SERVER_TYPE, err.Error())
	}
	fmt.Printf("Listening on port %s...\n", SERVER_PORT)
	defer server.Close() // To be executed at last

	// Wait and accept connection
	conn, err := server.Accept()
	if err != nil {
		log.Fatal(err)
	}

	handleConnection(conn)
}

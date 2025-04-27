package main

import (
	"fmt"
	"net"
)

const (
	PORT = "6379"
)

func main() {
	fmt.Printf("Listening on %s\n", PORT)

	// Create new server
	l, err := net.Listen("tcp", ":"+PORT)

	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}

	// Listen for connections
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Connection could not accepted:", err)
		return
	}

	defer conn.Close()

	for {
		resp := NewResp(conn)
		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(value)

		// ignore request and send back a PONG
		conn.Write([]byte("+OK\r\n"))
	}

}

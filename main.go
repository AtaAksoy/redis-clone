package main

import (
	"fmt"
	"io"
	"net"
	"os"
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
		buf := make([]byte, 1024)

		// read message from client
		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1)
		}

		// ignore request and send back a PONG
		conn.Write([]byte("+OK\r\n"))
	}

}

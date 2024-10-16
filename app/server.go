package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	buff := make([]byte, 128)

	n, err := conn.Read(buff)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		return
	}

	if !strings.Contains(string(buff[:n]), "PING") {
		fmt.Println("Invalid command")
		return
	}

	_, err = conn.Write([]byte("+PONG\r\n"))
	if err != nil {
		fmt.Println("Error writing: ", err.Error())
		return
	}
}

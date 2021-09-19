package main

import (
	"fmt"
	"net"
)

func main() {
	listenPort := "0.0.0.0:34000"
	listener, err := net.Listen("tcp4", listenPort)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Listening on %s", listenPort)
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	fmt.Printf("Serving %s\n", connection.RemoteAddr().String())
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	listenPort := "0.0.0.0:4000"
	listener, err := net.Listen("tcp", listenPort)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Listening on %s\n", listenPort)
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

	for {
		netData, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		message := strings.TrimSpace(string(netData))
		fmt.Println(message)
		if message == "STOP" {
			return
		}

		result := "Thanks!\n"
		connection.Write([]byte(string(result)))
	}
}

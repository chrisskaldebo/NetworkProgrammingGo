package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Listen
	// Accept
	// Handle connection -> thread
	dstream, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println(err)

		return
	}

	defer dstream.Close()

	for {
		con, err := dstream.Accept()
		if err != nil {
			fmt.Println(err)

			return
		}

		go handle(con)
	}
}

func handle(con net.Conn) {
	for {
		data, err := bufio.NewReader(con).ReadString('\n')
		if err != nil {
			fmt.Println(err)

			return
		}
		fmt.Println(data)
	}
}

package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	fmt.Println("Listening on port :3169")

	l, err := net.Listen("tcp", ":3169")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Listen for connections
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		// read message from client
		_, err = conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				continue
			}
			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1)
		}
		fmt.Println(string(buf))
		// ignore request and send back a PONG
		conn.Write([]byte("OK\r\n"))
	}
}

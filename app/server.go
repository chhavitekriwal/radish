package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	input := make([]byte,1024)
	for {
		n,err := conn.Read(input)
		if(err != nil) {
			fmt.Println("Failed to read input")
			os.Exit(1)
		}
		fmt.Printf("Read %d bytes\n",n)
		conn.Write([]byte("+PONG\r\n"))
	}
}

package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		return
	}
	defer l.Close()

	for {
		conn,err := l.Accept()
		if(err != nil) {
			fmt.Println("Error connecting to client..")
			fmt.Println(err)
			return
		}
		go handleReq(conn)
	}
}

func handleReq(conn net.Conn) {
	input := make([]byte,1024)
	
	for {
		n,readErr := conn.Read(input)
		if(readErr !=nil) {
			fmt.Println("Error reading from client..")
			fmt.Println(readErr)
			return
		}
		resp := string(input[:n])	
		parts := strings.Split(resp,"\r\n")
		fmt.Println("Read ",parts)
		switch strings.ToUpper(parts[2]) {
		case "PING":
			conn.Write([]byte("+PONG\r\n"))
		case "ECHO":
			msg := "+"+parts[4]+"\r\n"	
			conn.Write([]byte(msg))
		default:
			conn.Write([]byte("-NOMATCH\r\n"))
		}
	}

}
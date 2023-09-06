package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn,err := l.Accept()
		if(err != nil) {
			fmt.Println(err)
			fmt.Println("Exiting..")
			os.Exit(1)
		}
		go handleReq(conn)
	}
}

func handleReq(conn net.Conn) {
	input := make([]byte,1024)
	
	for {
		n,readErr := conn.Read(input)
		if(readErr !=nil) {
			fmt.Println(readErr)
			fmt.Println("Exiting..")
			os.Exit(1)
		}
		//fmt.Printf("Read %d bytes\n",rsize)
		
		resp := string(input[:n])
		fmt.Println(resp)
		
		parts := strings.Split(resp,"\\r\\n")
		fmt.Println(parts)
		switch parts[2] {
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
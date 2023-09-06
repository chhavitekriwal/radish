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
		rsize,readErr := conn.Read(input)
		if(readErr !=nil) {
			fmt.Println(readErr)
			fmt.Println("Exiting..")
			os.Exit(1)
		}
		fmt.Printf("Read %d bytes\n",rsize)
		_,writerr := conn.Write([]byte("+PONG\r\n"))
		if(writerr != nil) {
			fmt.Println(writerr)
			fmt.Println("Exiting..")
			os.Exit(1)
		}
	}

}
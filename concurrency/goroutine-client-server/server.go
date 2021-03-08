package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil{
		fmt.Println(err)
	}

	for{
		conn, err := listener.Accept()
		if err != nil{
			continue
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn){
	for{
		_, err := io.WriteString(conn, "response from server")
		if err != nil{
			return
		}
		time.Sleep(time.Second)
	}

}

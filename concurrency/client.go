package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader){
	_, err := io.Copy(dst, src)
	if err != nil{
       log.Fatalln(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil{
		log.Fatalln(err)
	}
	mustCopy(os.Stdout, conn)
}


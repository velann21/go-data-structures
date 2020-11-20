package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 2)

	go func() {
		for i:=0 ;i<10; i++{
			ch1 <- fmt.Sprintf("Number %s", i)
		}
	}()

	go func(){
		for i:=0 ;i<10; i++{
			ch2 <- fmt.Sprintf("Number %s", i)
		}
	}()

	for{
		select {
		case ch1Rcv := <-ch1:
			fmt.Println(ch1Rcv)
		case ch2Rcv	:= <- ch2:
			fmt.Println(ch2Rcv)

		case <- time.After(10 * time.Second):
			fmt.Println("Block timedout")
			break
		}

	}
}

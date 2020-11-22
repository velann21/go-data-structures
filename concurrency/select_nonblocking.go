package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string, 1)
	go func(){
		for i:= 0; i<=2; i++{
			ch1 <- "Message"
		}
	}()


	go func(){
		for{
			select{
			case ch1Rcv := <- ch1:
				fmt.Println(ch1Rcv)
			default:
				fmt.Println("It is default block no message")
				time.Sleep(2 * time.Second)
			}
			fmt.Println("Processing .....")
			time.Sleep(1500* time.Millisecond)
		}
	}()
}

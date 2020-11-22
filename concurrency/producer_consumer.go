package main

import (
	"fmt"
	"time"
)

func producer(income chan string){
	for i:=0; i<=10; i++{
		income <- fmt.Sprintf("Message %d", i)
	}
}

func consumer(income chan string, done chan bool){
	for{
		select {
		case <- time.After(time.Second*10):
			fmt.Println("All done")
			done <- true
		case data := <- income:
			fmt.Println(data)
		}
	}

}

func main() {
	income := make(chan string, 10)
	done := make(chan bool, 1)
	go producer(income)
	go consumer(income, done)
	<- done
}
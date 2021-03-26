package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go func(){
		for i:=0; i<=10; i++{
			time.Sleep(2*time.Second)
			ch1 <- i
		}
		close(ch1)
	}()

	for{
		select {
		case data, ok := <- ch1:
			if !ok{
				fmt.Println("Closed ch1")
				ch1 = nil
			}
			fmt.Println(data)
		case <- time.After(1*time.Second):
			fmt.Println("timed out")
			ch1 = nil
		}
		if ch1 == nil{
			break
		}
	}
}

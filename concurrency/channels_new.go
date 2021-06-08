package main

import (
	"fmt"
	"time"
)

func main() {

	go func(){
		fmt.Println("Inside go func")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Inside main")

	owner :=  func()<-chan int{
		ch := make(chan int)
		defer close(ch)
		go func(){
			for i:=0 ;i<=10; i++{
				ch <- i
			}
		}()
		return ch
	}

	consumer := func(inputs <-chan int) {
		for chn := range inputs{
			fmt.Println(chn)
		}
	}

	ch := owner()
	consumer(ch)
}

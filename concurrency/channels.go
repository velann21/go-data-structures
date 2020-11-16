package main

import "fmt"

func main() {
	owner := func() <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i:=0 ; i<=10; i++{
				ch <- 10
			}
			fmt.Println("Done the owner")
		}()

		return ch
	}

	consumer := func(ch <-chan int, done chan bool){
		fmt.Println("Inside the consumer")
		for v := range ch {
			fmt.Println("Channel,", v)
		}
		done <- true
		fmt.Println("Done recieve")
	}


	done := make(chan bool)
	ch := owner()
	consumer(ch, done)
}

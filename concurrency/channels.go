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

	consumer := func(ch <-chan int){
		fmt.Println("Inside the consumer")
		for v := range ch {
			fmt.Println("Channel,", v)
		}
		fmt.Println("Done recieve")
	}
	ch := owner()
	consumer(ch)
}

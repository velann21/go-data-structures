package main

import "fmt"

func main() {
	resultChan := make(chan int)
	go func (result chan int){
		a := 10
		b := 50
		c := a + b
		result <- c
	}(resultChan)

	output := <- resultChan
	fmt.Println(output)
}

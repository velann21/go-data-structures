package main

import "fmt"

func main() {
	resultsChan := make(chan int)
	go func (results chan int){
		for i:=0; i<=10; i++ {
			results <- i*2
		}
		close(results)
	}(resultsChan)

	for v := range resultsChan{
		fmt.Println(v)
	}
}

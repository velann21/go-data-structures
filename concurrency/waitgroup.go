package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var data int

	wg.Add(1)
	go func(){
		fmt.Println("Increasing value")
		data ++
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("checking value first")
	if data == 0{
		fmt.Println("Yes I am zero")
	}
}

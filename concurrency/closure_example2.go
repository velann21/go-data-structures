package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	for i:=0; i<=3; i++{
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("I is  :",i)
		}(i)
	}

	wg.Wait()
	fmt.Println("Done with execution")


}

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i:=0; i<=3; i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("I is  :",i)
		}()
	}
	wg.Wait()
	fmt.Println("Done with execution")
}

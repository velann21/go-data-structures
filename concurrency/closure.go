package main

import (
	"fmt"
	"sync"
)

func main() {

	tempFunc := func (group *sync.WaitGroup){
		group.Add(1)
		var i int

		go func(){
			defer group.Done()
			i++
			fmt.Println("I is ",i)
		}()

		fmt.Println("Returning the function")

	}
	wg := sync.WaitGroup{}
	tempFunc(&wg)

	wg.Wait()

	fmt.Println("I am done with execution")

}

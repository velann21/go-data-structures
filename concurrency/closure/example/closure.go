package main

import (
	"fmt"
	"sync"
)

func main() {
	tempFunc := func (group *sync.WaitGroup){
		for i:=0; i<=3 ;i++{
			group.Add(1)
			go func(){
				defer group.Done()
				fmt.Println("I is ",i)
			}()
		}

		fmt.Println("Returning the function")
	}
	wg := sync.WaitGroup{}
	tempFunc(&wg)
	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
)

var sharedSrc = make(map[string]interface{})
func main() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)

	wg.Add(1)

	go func(){
		defer wg.Done()

		cond.L.Lock()
		for len(sharedSrc) == 0{
			cond.Wait()
		}
		fmt.Println(sharedSrc["rsc1"])

		cond.L.Unlock()

	}()

	cond.L.Lock()
	sharedSrc["rsc1"] = "foo"
	cond.Signal()
	cond.L.Unlock()

	wg.Wait()



}

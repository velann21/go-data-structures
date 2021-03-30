package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int32 = 0
func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i:=0 ;i<100; i++{
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

func Credit(){

}

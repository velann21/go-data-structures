package main

import (
	"fmt"
	"sync"
)

func main() {
	balance := 0
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	creditBalance := func(amount int){
		mutex.Lock()
		defer mutex.Unlock()
		balance += amount
	}

	withdrawBalace := func(amount int){
		mutex.Lock()
		defer mutex.Unlock()
		balance -= amount
	}

	wg.Add(100)
	for i:=0; i<100; i++{
		go func(){
			defer wg.Done()
			creditBalance(1)
		}()

	}

	wg.Add(100)
	for i:=0; i<100; i++{
		go func(){
			defer wg.Done()
			withdrawBalace(1)
		}()
	}

	wg.Wait()

	fmt.Println(balance)
}


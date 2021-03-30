package main

import (
	"fmt"
	"sync"
)

var balance = 0

func main() {
	mutex := &sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(50)
	for i:=0 ; i<50; i++{
		go func (){
			defer wg.Done()
			Deposite(mutex)
		}()
	}

	wg.Add(50)
	for i:=0 ; i<50; i++{
		go func (){
			defer wg.Done()
			withDraw(mutex)
		}()
	}

	wg.Wait()
	fmt.Println(balance)

}

func Deposite(mutex *sync.Mutex){
	mutex.Lock()
	balance = balance+1
	mutex.Unlock()
}

func withDraw(mutex *sync.Mutex){
	mutex.Lock()
	defer mutex.Unlock()
	balance = balance-1
}

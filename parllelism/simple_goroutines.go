package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	go func (){
		wg.Add(1)
		Printer("Go Routine ")
		wg.Done()
	}()

	Printer("Main Routine ")
	wg.Wait()
}

func Printer(printerType string){
	for i:=0; i<=100; i++{
		fmt.Println(printerType,i)
	}
}
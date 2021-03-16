package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	capsChn := make(chan string, 3)
	trimChn := make(chan string, 3)
	resultChan := make(chan string, 10)
	termChan := make(chan bool)
	dataSet := []string{"Indumathi Selvam", "Velan Nandhakumar", "Nandhini Nandhakumar"}

	go func(dataSet []string, capsChn chan string){
		for _, v := range dataSet{
            uppercaseName := strings.ToUpper(v)
			capsChn <- uppercaseName
		}
		close(capsChn)
	}(dataSet, capsChn)

	go func(dataSet []string, trimChn chan string){
		for _, v := range dataSet{
			trimmedName := strings.TrimSpace(v)
			trimChn <- trimmedName
		}
		close(trimChn)
	}(dataSet, trimChn)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(){
		isWait := true
		for isWait{
			select {
			case trimmedValue, ok :=<-trimChn:
				if ok{
					fmt.Println("Trimmed result")
					resultChan <- trimmedValue
				}
			case capsValue, ok :=<-capsChn:
				if ok {
					fmt.Println("Caps result")
					resultChan <- capsValue
				}
			case term := <- termChan:
				if term{
					fmt.Println("Inside term chan block")
					isWait = false
					close(resultChan)
				}
			}

		}
	}()
    go func (){
		for  v := range resultChan{
		  fmt.Println("Result",v)
		}
		fmt.Println("Done result gather")
		close(resultChan)
		fmt.Println("Collected the result")
	}()
	wg.Wait()

	//for  v := range resultChan{
    //   fmt.Println(v)
	//}
	termChan <- true


}

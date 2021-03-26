package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	capsChn := make(chan string, 3)
	trimChn := make(chan string, 3)
	resultChan1 := make(chan string, 10)
	resultChan2 := make(chan string, 10)
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
		for {
			select {
			case trimmedValue, ok :=<-trimChn:
				fmt.Println("Trimmed result")
				resultChan1 <- trimmedValue
				if !ok{
					trimChn = nil
				}
			case capsValue, ok :=<-capsChn:
				fmt.Println("Caps result")
				resultChan2 <- capsValue
				if !ok {
					capsChn = nil
				}
			}
			if trimChn == nil && capsChn == nil {
				close(resultChan1)
				close(resultChan2)
				wg.Done()
				wg.Done()
				break
			}
		}
	}()
    go func (){
		for  v := range resultChan1{
		  fmt.Println("Result",v)
		}
		for v := range resultChan2 {
			fmt.Println("Result",v)
		}
		fmt.Println("Done result gather")
		fmt.Println("Collected the result")
	}()
	wg.Wait()

	//for  v := range resultChan{
    //   fmt.Println(v)
	//}


}

package main
//
//import (
//	"fmt"
//	"strings"
//	"sync"
//)
//
//func main() {
//	//producer := func()chan int{
//	//	dataPipeline := make(chan int)
//	//	go func(){
//	//		defer close(dataPipeline)
//	//		for i := 0; i<=10; i++{
//	//			dataPipeline <- i
//	//		}
//	//	}()
//	//	return dataPipeline
//	//}
//	//
//	//consumer := func(datum chan int) {
//	//	 for data := range datum{
//	//	 	fmt.Println(data)
//	//	 }
//	//}
//
//	pipeline := Producer()
//	termHandler := Consumer(pipeline)
//	<-termHandler
//	close(termHandler)
//
//	streamCaps := []string{"IS", "WAS", "BUT"}
//	streamSmall := []string{"velan", "bloom", "ikea"}
//	result := make(chan string)
//	small := ConvertCapsToSmaller(streamCaps)
//	caps := ConvertSmallToCaps(streamSmall)
//
//	go func() {
//		for {
//			select {
//			case data, ok := <-small:
//				if !ok {
//					small = nil
//				}
//				result <- data
//			case data, ok := <-caps:
//				if !ok {
//					caps = nil
//				}
//				result <- data
//			}
//			if caps == nil && small == nil {
//				close(result)
//				break
//			}
//		}
//	}()
//
//	wg := &sync.WaitGroup{}
//	wg.Add(1)
//	go func(){
//		defer wg.Done()
//		for res := range result {
//			fmt.Println(res)
//		}
//	}()
//
//	wg.Wait()
//}
//
//func ConvertCapsToSmaller(streamCaps []string) chan string {
//	smallData := make(chan string)
//	go func() {
//		defer close(smallData)
//		for _, value := range streamCaps {
//			smallData <- strings.ToLower(value)
//		}
//	}()
//	return smallData
//}
//
//func ConvertSmallToCaps(streamSmall []string) chan string {
//	capsData := make(chan string)
//	go func() {
//		defer close(capsData)
//		for _, value := range streamSmall {
//			capsData <- strings.ToUpper(value)
//		}
//	}()
//	return capsData
//}
//
//func Producer() chan int {
//	dataPipeline := make(chan int)
//	go func() {
//		defer close(dataPipeline)
//		for i := 0; i <= 10; i++ {
//			dataPipeline <- i
//		}
//	}()
//	return dataPipeline
//}
//
//func Consumer(pipeline <-chan int) chan bool {
//	termination := make(chan bool)
//	go func(<-chan int) {
//		for data := range pipeline {
//			fmt.Println(data)
//		}
//		termination <- true
//	}(pipeline)
//	return termination
//}

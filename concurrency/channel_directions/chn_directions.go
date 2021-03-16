package main

import "fmt"

func producer(dataSet []int)chan int{
	processorChn := make(chan int, 10)
	go func(processorChn chan<- int){
		defer close(processorChn)
		for _, v := range dataSet{
			processorChn <- v
		}
	}(processorChn)
	return processorChn
}

func processor(out <-chan int)chan int{
	resultChn := make(chan int, 10)
	go func(result chan<- int){
		defer close(result)
	    output := 0
		for v :=range out{
			output += v
			result <- output
		}
	}(resultChn)
	return resultChn
}

func resultConsumer(out <-chan int)chan bool{
	blockerChn := make(chan bool)
	go func(blockerChn chan<- bool){
		defer close(blockerChn)
		for v := range out{
			fmt.Println(v)
		}
		blockerChn <- true
	}(blockerChn)
	return blockerChn
}

func main() {
	procChn := producer([]int{1,2,3,4})
	resultChn := processor(procChn)
	blockerChn := resultConsumer(resultChn)
	<- blockerChn
}

package main

import (
	"fmt"
	"time"
)

func fun(str string){
	for i:=0; i<3; i++{
		fmt.Println(str)
		time.Sleep(1*time.Millisecond)
	}

}

func main() {
	fun("Direct call")

	go fun("Go routine call")

	time.Sleep(10*time.Second)
}

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

	//Normal Goroutine call
	go fun("goroutine-1")

	//Anonymus func call
	go func (){
		fun("goroutine-2")
	}()

	funcExe := fun
	funcExe("goroutine-3")

	time.Sleep(10*time.Second)
}

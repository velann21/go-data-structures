package main

import "fmt"

func main() {
	arr := []int{1,2,3,4,5}
	fmt.Println(AddEvenPos(arr, len(arr)-1))
}

func AddEvenPos(arr []int, n int)int{
	if n < 0{
		return 0
	}
	if n == 0 {
		return arr[0]
	}

	return arr[n] + AddEvenPos(arr, n-2)

}

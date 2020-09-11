package main

import "fmt"

func main() {
	arr := []int{1,2,3,4,5}
	fmt.Println(SumArray(arr, len(arr)-1))
}

func SumArray(arr []int, n int)int{
	if n == 0{
		return arr[0]
	}
	return arr[n] + SumArray(arr, n-1)
}

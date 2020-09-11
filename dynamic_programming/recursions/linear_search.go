package main

import "fmt"

func main() {
	arr := []int{10,20,30,50,60}
	fmt.Println(LinearSearch(arr, len(arr)-1, 35))
}

func LinearSearch(arr []int, n int, tofind int)int{
	if arr[n] == tofind{
		return n
	}
	if n == 0{
		return -1
	}
	return LinearSearch(arr, n-1, tofind)
}

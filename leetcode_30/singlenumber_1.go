package main

import "fmt"

func singleNumber(nums []int) int {
	a := 0
	for _, v := range nums{
		fmt.Println("Before", a)
		a^=v

		fmt.Println("After", a)
	}
	return a
}

func main() {
	arr := []int{4, 4}
	num := singleNumber(arr)
	fmt.Println(num)
}

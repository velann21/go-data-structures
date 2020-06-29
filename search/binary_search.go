package main

import "fmt"

func main() {
	sr := Search{}
	op := sr.BinarySearch([]int{1,3,5,7,9,10}, 9)
	fmt.Println(op)
}

type Search struct {

}

func (src *Search) BinarySearch(input []int, element int)int{
	left := 0
	right := len(input)-1
	for left <= right{
		fmt.Println(left)
		fmt.Println(right-left)
		middle := left + (right-left)/2
		if input[middle] == element{
			return middle
		}

		if input[middle] < element{
			left = middle+1
		}else {
			right = middle -1;
		}
	}

	return -1
}

package main

import (
	"fmt"
	"sort"
)

func SolutionWithSpace(A []int) int {
	// write your code in Go 1.4
	binset := make([]int, 1000000)
	binset = SetBinSet(A, binset)
	return GetMissingElement(binset)
}

func SetBinSet(A []int, binset []int)[]int{
	for  _, v := range A{
		if v > 0{
			binset[v] = v
		}
	}
	return binset
}

func GetMissingElement(binset []int)int{
	for j, v:= range binset{
		if j != 0 && v == 0{
			return j
		}
	}
	return 0
}

func SolutionWithoutSpace(A []int)int{
	SortArray(A)
	min := FindMin(A)
	return min
}

func FindMin(A []int)int{
	min := 1
	for _, v := range A {
		if v == min {
			min++
		}
	}
	return min
}

func SortArray(A []int){
	sort.Ints(A)
}

func main() {
	fmt.Println(SolutionWithoutSpace([]int{1, 3, 6, 4, 1, 2}))
	fmt.Println(SolutionWithoutSpace([]int{-1, -2}))
	fmt.Println(SolutionWithoutSpace([]int{1, 2, 3}))

	fmt.Println(SolutionWithSpace([]int{1, 3, 6, 4, 1, 2}))
	fmt.Println(SolutionWithSpace([]int{-1, -2}))
	fmt.Println(SolutionWithSpace([]int{1, 2, 3}))
}




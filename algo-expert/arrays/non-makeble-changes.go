package main

import (
	"fmt"
	"sort"
)

func NonConstructibleChange(coins []int) int {
	// Write your code here.
	if len(coins) <= 0{
		return 1
	}
	sortedArray := sortArray(coins)
	if coins[0] != 1{
		return 1
	}
	currentChange := 0
	for _, value := range sortedArray{
		if value > currentChange +1 {
			return currentChange+1
		}else{
			currentChange += value
		}
	}
	return currentChange+1
}

func sortArray(coins []int)[]int{
	sort.Ints(coins)
	return coins
}

func main() {
	c := NonConstructibleChange([]int{1,1,2,5,7,9,22})
	fmt.Println(c)
}
package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{12,5,6,10,3,4,8}
	min, max := findMaxMinInOneGo(arr)
	fmt.Print(min,max)

}

func findMaxMinInOneGo(arr []int)(min int,max int){
	min = math.MaxInt8
	max = math.MinInt8
	for i:=0 ;i<len(arr);i++{
		if (arr[i] < min){
			min = arr[i]
		}
		if(arr[i] > max){
			max = arr[i]
		}
	}
	return min, max
}

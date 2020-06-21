package main

import "fmt"

func main() {
	arr := []int{10,5,6,12,3,4,8}
	min, max := findMaxMinInOneGo(arr)
	fmt.Print(min,max)

}

func findMaxMinInOneGo(arr []int)(min int,max int){
	min = arr[0]
	max = arr[0]
	for i:=0 ;i<len(arr);i++{
		if (arr[i] < min){
			min = arr[i]
		}else if(arr[i] > max){
			max = arr[i]
		}
	}
	return min, max
}

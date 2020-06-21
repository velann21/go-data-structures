package main

import "fmt"

func main() {
	//arr := []int{1,2,4,5,6,8,10,12}
	//pair := 10
	//findPairInSortedArray(arr, pair)
	arr1 := []int{1,12,5,6,8,10,2,4,0}
	pair1 := 10
	findPairInUnsortedArrayUsingMap(arr1, pair1)
}

func findPairInSortedArray(arr []int, toFindPair int){
	i:=0
	j:=len(arr)-1
	for i<j{
		if (arr[i]+arr[j] == toFindPair){
			fmt.Println(arr[i],arr[j])
			i++
			j--
		}else if(arr[i]+arr[j]<toFindPair){
			i++
		}else{
				j--
		}
	}
}

func findPairInUnsortedArrayUsingMap(arr []int, toFindPair int){
	HMap := make([]int, toFindPair+1)
	for i:=0; i<len(arr); i++{
		if arr[i] > toFindPair{
			continue
		}
		diff := toFindPair - arr[i];
		if HMap[diff] == -1{
			fmt.Println(arr[i],diff)
		}else{
			HMap[diff] = -1
			HMap[arr[i]] = -1
		}

	}
}



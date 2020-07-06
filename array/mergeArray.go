package main

import "fmt"

func main() {
	//arr1 := []int{1,3,5,7,9}
	//arr2 := []int{10,13,15,17,19}

	arr1 := []int{1,5,7,9,10,11,12}
	arr2 := []int{3,13,14,15,16,17,18,19}
	mergedArray := mergeSortedArray(arr1, arr2)
	fmt.Println(mergedArray)
}

func mergeSortedArray(arr1 []int, arr2 []int) []int{
	if len(arr1) <= 0 && len(arr2) <= 0 {
		return nil
	}
	i,j := 0,0
	mergedArr := make([]int,0)
	for i<len(arr1) && j<len(arr2){
		fmt.Println(i)
		fmt.Println(len(arr1))

		firstArrPtr := arr1[i]
		secondArrPtr := arr2[j]
		if firstArrPtr < secondArrPtr{
			mergedArr = append(mergedArr, firstArrPtr)
			i++
		}else if secondArrPtr < firstArrPtr{
			mergedArr = append(mergedArr, secondArrPtr)
			j++
		}
	}
	if i < j{
		for i < len(arr1){
			mergedArr = append(mergedArr, arr1[i])
			i++
		}
	}else if j < i{
		for j < len(arr2){
			mergedArr = append(mergedArr, arr2[j])
			j++
		}
	}

		return mergedArr
}

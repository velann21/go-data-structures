package main

import "fmt"

func main(){
	//arr1 := []int{1,10,3,4,5,6}
	//arr2 := []int{6,7,8,9,11,10,5,5}
	arr1 := []int{1,7,10,12,13}
	arr2 := []int{6,7,8,9,10,11}
	intersectedArr := intersectionSortedArr(arr1, arr2)
	fmt.Println(intersectedArr)
}

func intersectUnSortedArr(arr1 []int, arr2 []int)[]int{
	arr3 := make([]int, 0)
	i := 0
	j := 0

	for i < len(arr1){
		arr3 = append(arr3, arr1[i])
		i++
	}

	for j<len(arr2){
		elementToFind := arr2[j]
		currentEle := 0
		for i = 0; i<len(arr3); i++{
			currentEle = arr3[i]
			if elementToFind == currentEle{
				break
			}
		}
		if elementToFind != currentEle {
			arr3 = append(arr3, elementToFind)
		}
		j++;
	}
	return arr3
}

func intersectionSortedArr(arr1 []int, arr2 []int)[]int{
	i := 0
	j := 0
	arr3 := make([]int, 0)
	for i<len(arr1) && j<len(arr2){
		if arr1[i] < arr2[j]{
			arr3 = append(arr3, arr1[i])
			i++
		}else if arr1[i] == arr2[j]{
			arr3 = append(arr3, arr1[i])
			i++
			j++
		}else {
			arr3 = append(arr3, arr2[j])
			j++
		}
	}

	if i < len(arr1){
		for i<len(arr1){
			arr3 = append(arr3, arr1[i])
			i++
		}
	}else if j < len(arr2){
		for j<len(arr2)-1 {
			arr3 = append(arr3, arr1[j])
			j++
		}
	}
	return arr3
}

func union(arr1 []int, arr2 []int){

}

func difference(arr1 []int, arr2 []int){

}
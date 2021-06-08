package main

import "fmt"

func main(){

	x := "hello"
	defer fmt.Println(x)
	x = "world"
	fmt.Println("hello")

	//x := "hello"
	//defer func() { fmt.Println(x) }()
	//x = "world"
	//fmt.Println("hello")


	//arr := []int{1,3,4,5,5,6,7,7,7,7,7,8,8,8,8,8,9}
	//findingCountingDuplicates(arr)
	//arr1 := []int{3,4,1,3,2,4,4,5,6,2, 6, 6}
	//findDuplicatesUnsortedArr(arr1, 6)
	//findDuplicatesUnsortedArrWithoutMap(arr1)
}

func findingCountingDuplicates(arr []int){
	// This outer loop throughs whole arr
	for i:=0; i<len(arr)-1; i++{
		// if current element and next element are same then duplicate is present
		if arr[i] == arr[i+1]{

			j := i+1
			// This loop is to count the number of duplicates on specific dup number
			for j<len(arr) && arr[j] == arr[i]{
				j++
			}

			fmt.Println(arr[i],j-i)
			i = j-1
		}
	}
}

func findDuplicatesUnsortedArr(arr []int, n int){

	H := make([]int,n+1)

	for i:=0; i<len(arr); i++{
		H[arr[i]] ++
	}

	for i:=0; i<len(arr); i++{

		if H[arr[i]] > 1 && H[arr[i]] != -1{
			fmt.Println(arr[i], H[arr[i]])
			H[arr[i]] = -1
		}
	}
}

func findDuplicatesUnsortedArrWithoutMap(arr []int){
	dupNumber := 0
	for i:=0; i<len(arr); i++{
		for j:=i+1; j<len(arr); j++{
			if arr[i] == arr[j] && arr[j] != dupNumber && arr[j] != -1{
				dupNumber = arr[j]
				fmt.Println( "Duplicate Present", arr[j])
				arr[j] = -1
			}
		}
	}
}

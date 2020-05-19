package main

import "fmt"

func main(){
	var i int = 0
	fmt.Println(i)
    arr := make([]int,0 )
    arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)
	arr = append(arr, 4)
	arr = append(arr, 5)
	arr = append(arr, 5)
	arr = append(arr, 6)
	arr = append(arr, 7)

	value := filter(func(value int)bool{
		if value == 2 || value == 4{
			return true
		}
		return false
	}, arr)

	fmt.Println(value)
}

func filter(filter func(value int)bool, array []int)[]int{
	filterValues := make([]int, 0)
	for _,v:= range array{
		if filter(v) {
			filterValues = append(filterValues,v)
		}
	}
	return filterValues
}



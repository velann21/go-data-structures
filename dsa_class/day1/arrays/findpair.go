package main

import "fmt"

//input : [2,4,6,8](1)
//op := 2,4
//input2 := [6,2,4,8]
//op: [2,4]
//input3 := [6,2,-2,8]
//op:=[-2,8]

func main() {
	o1, o2 := findPair([]int{5,6,1,2,3}, 6)
	fmt.Println(o1, ",",o2)
}

func findPair(input []int, target int)(int, int){
	tracker := make(map[int]int)
	pair1 := 0
	pair2 := 0

	//Write the value as the position of array in the map
	for i, v := range input{
		tracker[v] = i
	}
	//Iterate the array and find the diff with target and find whether the diff is present in array
	//If present the return the current index and diffed elements index
	for j, v := range input{
		diff := target-v
		if tracker[diff] != 0 && tracker[diff] != j{
			pair1 = j
			pair2 = tracker[diff]
			break
 		}
	}
	return pair1, pair2
}

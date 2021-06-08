package main

import "fmt"

func SortedSquaredArray(array []int) []int {

	if len(array) <= 0 {
		return nil
	}
	start := 0
	end := len(array)-1
	auxilaryArray := make([]int, len(array))
	auxilaryArrayPtr := len(array)-1
	for start <= end && auxilaryArrayPtr >=0{
		startValue := array[start]
		endValue := array[end]
		startSquare := startValue*startValue
		endSquare := endValue*endValue
		if startSquare > endSquare{
			auxilaryArray[auxilaryArrayPtr] = startSquare
			start++
			auxilaryArrayPtr--
		}else if startSquare == endSquare{
			if start == end{
				auxilaryArray[auxilaryArrayPtr] = startSquare
				break
			}
			auxilaryArray[auxilaryArrayPtr] = startSquare
			auxilaryArrayPtr--
			auxilaryArray[auxilaryArrayPtr] = endSquare
			auxilaryArrayPtr--
			start++
			end--
		}else{
			auxilaryArray[auxilaryArrayPtr] = endSquare
			auxilaryArrayPtr--
			end--
		}
	}
	return auxilaryArray
}

func main() {
	sortedArray := SortedSquaredArray([]int{1, 2, 3, 5, 6, 8, 9})
	fmt.Println(sortedArray)
}
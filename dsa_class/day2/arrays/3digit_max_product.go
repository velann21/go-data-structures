package arrays

import (
	"fmt"
	"math"
)

//input: [4,1,2,5,6]
//op: (4*5*6)
//input2: [-10,-8,-7,1,3,4,5]
//op: (-10*-8*5)
//App1: Sort the array and get the last 3 elements if it has only positive elements
        //If it has negative elements then check the first two negative and last one max and compare it with
        // Eg: 1. max1*max2*max3  2. min1*min2*max1
        //Check which is greater and return that
        // Here the runtime would be nlogn since we should sort the array
//Below approach would solve me in O(n)
func main() {
	fmt.Println(maxProduct([]int{-10,-2,-6,1,6,4}))
}


func maxProduct(input []int)int {

	//Make the min as max int
	min1 := math.MaxInt32
	min2 := math.MaxInt32
	//Make the max as min int
	max1 := math.MinInt32
	max2 := math.MinInt32
	max3 := math.MinInt32

	for _, v := range input {
		if v < min1 {
			min2 = min1
			min1 = v
		} else if v < min2 {
			min2 = v
		} else if v > max1 {
			max3 = max2
			max2 = max1
			max1 = v
		} else if v > max2 {
			max3 = max2
			max2 = v
		} else if v > max3 {
			max3 = v
		}
	}
	prod1 := max1 * max2 * max3
	prod2 := min1 * min2 * max1

	if prod1 > prod2 {
		return prod1
	} else {
		return prod2
	}
}


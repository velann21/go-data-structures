package complexity

import "fmt"

//Time complexity is not equal to amount of time it take it to complete the execution
// In computers alogrithm we should alwasy consider the rate of growth which algo takes in time complexity

//Example as below

//This takes always n
func BigON(array []int){
	for i:=0 ; i<=len(array); i++{ // nTimes
		fmt.Println(i) //1 times n*1 = n
	}
}

func BigON2(array []int){
	for i:=0 ; i<=len(array); i++{ // nTimes
		for i:=0 ; i<=len(array); i++{ // nTimes
			fmt.Println(i) //1 times n*n+1 = n2
		}
	}
}
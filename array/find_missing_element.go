package array

import "fmt"

func main()  {
	//arr := []int{1,2,3,4,5,6,7,9,10,11,12}
	arr := []int{6,7,9,12}
	     //diff= 6,6 ,7,(8,9)
	     //missi     6+2(8), 7+3(10), 8+3(11)
	    //       6,6,6, 8, 8
	//findAllMissingElementUsingBinSet(arr,12)
	findAllMissingElement(arr)
	//fmt.Println(missingElement)
}


func findFirstMissingElementInNNatural(arr []int, n int)int{
	sum := 0
	diff := 0

	// Calculation for finding first n natural numbers
	firstNNum := n*(n+1)/2

	// Looping to sum all the numbers in arr
	for i:=0; i<len(arr);i++{
		sum = sum+arr[i]
	}
	//If sum is lesser than the first n natural then firnd the diff and thats the missing number
	if sum != firstNNum{
		diff = firstNNum - sum
	}
	return diff
}

func findMissingFirstElementInArr(arr []int)int{
	diff := arr[0] - 0
	missingNum := 0
	for i:=0; i<len(arr); i++{
		if  diff != arr[i]-i {
			missingNum = diff+i
			break
		}
	}
	return missingNum
}

func findAllMissingElement(arr []int){
	diff := arr[0] - 0
	for i:=0; i<len(arr); i++{
		if  diff != arr[i]-i {
			for diff < arr[i]-i{
				fmt.Println(diff+i)
				diff ++
			}
		}
	}
}

func findAllMissingElementUsingBinSet(arr []int, n int){
	binArr := make([]int, n+1)
	binArr[0] = 1
	for i:=0; i<len(arr);{
       binArr[arr[i]] = 1
       i++
	}
	fmt.Println(binArr)
	for j:=6; j<len(binArr); j++{
		if binArr[j] == 0{
			fmt.Print(j,",")
		}
	}
}





package array

import (
	"fmt"
)

func main(){

	arr := make([]int, 0)
	arr = append(arr, 10)
	arr = append(arr, 20)
	arr = append(arr, 30)
	arr = append(arr, 40)
	arr = append(arr, 50)
	arr = append(arr, 60)
	arr = append(arr, 70)
	arr = append(arr, 80)

	isSortedBefore := checkArraySorted(&arr)
	fmt.Println(isSortedBefore)

	fmt.Println("Before len :", len(arr))
	fmt.Println("Before cap :", cap(arr))
	updatedArray := insertElementInSortedArray(25, &arr)
	fmt.Println(updatedArray)
	fmt.Println("After len :", len(*updatedArray))
	fmt.Println("After cap :", cap(*updatedArray))

	ptrArr := reverse(&arr)
	fmt.Println(ptrArr)

	ls := leftShift(&arr)
	fmt.Println(ls)

	rs := rightShift(&arr)
	fmt.Println(rs)

	isSortedAfter := checkArraySorted(&arr)
	fmt.Println(isSortedAfter)

	isReverseSorted := checkIfArraySortedReverseWay(&arr)
	fmt.Println(isReverseSorted)
}

func reverse(arr *[]int) *[]int{
	ptrArr := *arr
	if arr == nil || len(ptrArr) <= 0{
		return nil
	}
	i := 0
	j := len(ptrArr)-1
	for i<j{
	   temp := ptrArr[i]
	   ptrArr[i] = ptrArr[j]
	   ptrArr[j] = temp
		i += 1;
		j -= 1;
	}
	ptrArr = nil
	return arr
}

func leftShift(arr *[]int)*[]int{
	ptrArr := *arr
	var firstVar = 0
	for i:=0;i<len(ptrArr);i++ {
       if i == 0{
       	 firstVar = ptrArr[i]
       	 fmt.Println(firstVar)
       	 continue;
	   }
		ptrArr[i-1] = ptrArr[i]
	}
	ptrArr[len(ptrArr)-1] = firstVar
	ptrArr = nil
	return arr
}

func rightShift(arr *[]int)*[]int{
	ptrArr := *arr
	var lastVar = 0
	length := len(ptrArr)-1
	for i:=length; i>=0; i-- {
		if i == length{
			lastVar = ptrArr[length]
			fmt.Println(lastVar)
			continue;
		}
		ptrArr[i+1] = ptrArr[i]
	}
	ptrArr[0] = lastVar
	ptrArr = nil
	return arr
}

func checkArraySorted(arr *[]int) bool {
	ptrArr := *arr
	for i:=0; i<len(ptrArr)-1;i++{
        if ptrArr[i] > ptrArr[i+1]{
        	return false
		}
	}
	ptrArr = nil
	return true
}


func checkIfArraySortedReverseWay(arr *[]int) bool {
	ptrArr := *arr
	for i:=0; i<len(ptrArr)-1;i++{
		if ptrArr[i] < ptrArr[i+1]{
			return false
		}
	}
	ptrArr = nil
	return true
}

func insertElementInSortedArray(element int, arr *[]int) *[]int{
	ptrArr := *arr
	i := 0
	if len(ptrArr) == cap(ptrArr){
		ptrArr = append(ptrArr,0)
	}
	// Here i have take the len-2 otherwise the above appended element is less than last element(Since the appended value
	// is 0 always the last element in prev array is smaller than the element to be inserted) it won't run the loop.
	// it breaks at first condition itself
	//         0       <   80 //true and breaks the loop
	//eg: if ptrArr[i] < element{
	//	break
	// }
	for i=len(ptrArr)-2; i>=0; i--{

		if ptrArr[i] < element{
			break
		}
		ptrArr[i+1] = ptrArr[i]
	}
	ptrArr[i+1] = element
	*arr = ptrArr
	ptrArr = nil
    return arr
}


//Pending questions for the day to solve.
// 1) Move all the positive and negative integers together.
// 2) Merge 2 sorted array
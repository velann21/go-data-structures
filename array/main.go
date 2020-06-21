package main

import (
	"fmt"
	"strconv"
)

func main() {
	//array := []int{1,4,8,9,14,18,20}
	//result := BinarySearch(array, len(array), 20)
	//fmt.Print(result)

	//results := FindFactors(63)
	//fmt.Print(results)
	//FindFactorsRecursively(45, 1)

	//Tail Recursion Call
	//TailRecursionPrint(5)
	//HeadRecursionPrint(5)

	//resp := TailRecursionReturnLocalVaraible(5)
	//fmt.Println(resp)

	//TreeRecursion(5)

	//IndirectRecursion(20)

	fmt.Println(SumNNaturalNumbers(5))

	fmt.Println(Factorial(5))

}

func BinarySearch(array []int, size int, numberTofind int)int{
	startingPoint := 0
	lengthPoint := size-1

	for startingPoint < lengthPoint {
		middlePointer := (startingPoint+lengthPoint)/2

		if array[middlePointer] == numberTofind{
			return middlePointer
		}

		if numberTofind > array[middlePointer]{
			startingPoint = middlePointer+1
		}else{
			lengthPoint = middlePointer-1
		}
	}
	return -1
}

func FindFactors(factorialNumber int) []int{
	divisor := 1;
	factors := make([]int,0)
	for divisor <= factorialNumber{
		if factorialNumber%divisor == 0{
			factors = append(factors, divisor)
			divisor++
		}else {
			divisor++
			continue
		}
	}
	return factors
}

func FindFactorsRecursively(factorialNumber int, divisor int)int{
	if factorialNumber%divisor == 0{
		fmt.Print(divisor, ",")
	}
	if (divisor <= factorialNumber){
		FindFactorsRecursively(factorialNumber, divisor+1)
	}
    return 0
}

func TailRecursionPrint(n int){
	if n>0{
		fmt.Println("TailRecursionPrint n is : "+strconv.Itoa(n))
		TailRecursionPrint(n-1)
	}
}

func HeadRecursionPrint(n int){
	if n>0 {
		HeadRecursionPrint(n-1)
		fmt.Println("HeadRecursionPrint n is :" + strconv.Itoa(n))
	}
}

var temp  = 0
func TailRecursionReturnLocalVaraible(n int)int{
	if n > 0 {
		temp++
		res := TailRecursionReturnLocalVaraible(n-1)+temp
		println(res)
		return res
	}
	return n
}


func TreeRecursion(n int){
	if n > 0 {
		fmt.Print(" , " + strconv.Itoa(n))
		TreeRecursion(n-1)
		TreeRecursion(n-1)
	}
}


func IndirectRecursion(n int){
	A(n)
}

func A(n int){
	if n > 0{
		print(" A call:"+strconv.Itoa(n))
		B(n-1)
	}
}

func B(n int){
	if n > 1{
		print(" B call:"+strconv.Itoa(n))
		A(n/2)
	}
}


func SumNNaturalNumbers(n int)int{
	if n == 0 {
		return 0
	}

	if n > 0{
		return SumNNaturalNumbers(n-1)+1
	}

	return 0
}

func Factorial(n int)int{
	if n == 0{
		return 1
	}
	if n > 0{
		return Factorial(n-1)*n
	}
	return 1
}




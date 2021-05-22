package main

import (
	"fmt"
)


type List interface {
	Add(element int)
	DisplayArray()
	expandArray()
	DeleteElement(index int)
	LinearSearch(element int)int
	BinarySearch(element int)int
}
type ArrayList struct {
	Size int
	Length int
	Array []int
}

func NewArrayList(Size int)List{
	array := make([]int, Size)
	return &ArrayList{Size: Size,  Array: array}
}

func (al *ArrayList) Add(element int){
	if al.Length == al.Size{
		al.expandArray()
	}
	al.Array[al.Length] = element
	al.Length++
}

func (al *ArrayList) DisplayArray(){
	for i:=0; i<al.Length;i++{
		fmt.Println(al.Array[i])
	}
}

func (al *ArrayList) expandArray() {
	newSize := al.Size+10
	newArray := make([]int, newSize)
	for i, v := range al.Array{
		newArray[i] = v
	}
	al.Array = newArray
	al.Size = newSize
}

func (al *ArrayList) DeleteElement(index int){
	if index > al.Length{
		return
	}
	for i:=index; i<=al.Length; i++{
		if i+1 != al.Length+1{
			al.Array[i] = al.Array[i+1]
		}
	}
	al.Length --
}

func (al *ArrayList) LinearSearch(element int)int{
	start :=0
	end := al.Length
	for end >= start {
		if al.Array[start] == element{
			fmt.Println("Yes found")
			return al.Array[start]
		}else if al.Array[end] == element{
			fmt.Println("Yes found")
			return al.Array[end]
		}else{
			end--
			start++
		}
	}
	return 0
}

func (al *ArrayList) BinarySearch(element int)int{
	start := 0
	end := al.Length
	mid := (start + end)/2
	for end >= start{
		fmt.Println(al.Array[mid])
		if al.Array[mid] == element{
			return al.Array[mid]
		}else if al.Array[mid] > element{
			end = mid-1
		}else {
			start = mid+1
		}
		mid = (start + end)/2
	}
	return 0
}

func (al *ArrayList) Get(index int)int{
	if index < 0 || index > al.Length{
		return 0
	}
	return al.Array[index]
}

func (al *ArrayList) Set(index, element int){
	if index < 0 || index > al.Length{
		return
	}
	al.Array[index] = element
}

func (al *ArrayList) FindMaxMin()(int, int){
	if al.Length <=0 {
		return 0, 0
	}
	max := al.Array[0]
	min := al.Array[0]
	for _, value := range al.Array{
		if value > max{
			max = value
		}else if value < min{
			min = value
		}
	}
	return max, min
}

func (al *ArrayList) AddFunc()int{
	if al.Length <=0 {
		return 0
	}
	sum := 0
	for _, value := range al.Array{
		sum += value
	}
	return sum
}

func (al *ArrayList) CalculateAvg()int{
	if al.Length <=0 {
		return 0
	}
	avg := 0
	for _, value := range al.Array{
		avg += value
	}
	avg = avg/al.Length
	return avg
}

func (al *ArrayList) LeftShift()[]int{
	if al.Length <=0 && al.Length == 1 {
		return al.Array
	}
	for i:=0; i<=al.Length; i++{
		if i+1 <= al.Length{
			al.Array[i] = al.Array[i+1]
		}
	}
	return al.Array
}

func (al *ArrayList) ReverseArray()[]int{
	if al.Length <=0 && al.Length == 1 {
		return al.Array
	}
	start := 0
	end := al.Length
	for end > start{
		temp := al.Array[start]
		al.Array[start] = al.Array[end]
		al.Array[end] = temp
	}
	return al.Array
}

func (al *ArrayList) LeftRotate(rotateNumber int)[]int{
	if al.Length <=0 && al.Length == 1 {
		return al.Array
	}
	if rotateNumber == 1{
		tempVar := al.Array[0]
		al.Rotate()
		al.Array[al.Length] = tempVar
	}
	return al.Array
}

func (al *ArrayList) Rotate(){
	for i:=0; i<=al.Length; i++ {
		if i+1 <= al.Length{
			al.Array[i] = al.Array[i+1]
		}
	}
}

func (al *ArrayList) CheckArraySorted()bool{
	if al.Length <=0 && al.Length == 1 {
		return true
	}
	for i:=0; i<al.Length; i++{
		if i+1 <=al.Length && al.Array[i] > al.Array[i+1]{
			return false
		}
	}
	return true
}

func (al *ArrayList) MergeKSortedArrays(array1, array2 []int)[]int{
	if len(array1) <=0 && len(array2) > 0{
		return array2
	}
	if len(array2) <=0 && len(array1) > 0{
		return array1
	}
	if len(array2) <=0 && len(array1)<=0{
		return nil
	}

	arrayOnePointer := 0
	arrayTwoPointer := 0
	auxiliaryArray := make([]int, 0)
	auxiliaryArrayPointer := 0


	for arrayOnePointer<=len(array1) && arrayTwoPointer<=len(array2) {
		if array1[arrayOnePointer] < array2[arrayTwoPointer]{
			auxiliaryArray[auxiliaryArrayPointer] = array1[arrayOnePointer]
			arrayOnePointer ++
			auxiliaryArrayPointer++
		}else if array1[arrayOnePointer] > array2[arrayTwoPointer] {
			auxiliaryArray[auxiliaryArrayPointer] = array2[arrayTwoPointer]
			arrayTwoPointer ++
			auxiliaryArrayPointer++
		}else if array1[arrayOnePointer] == array2[arrayTwoPointer] {
			auxiliaryArray[auxiliaryArrayPointer] = array2[arrayTwoPointer]
			arrayTwoPointer ++
			arrayOnePointer ++
			auxiliaryArrayPointer++
		}
	}

	if arrayOnePointer < len(array1){
		for i := arrayOnePointer; i<= len(array1); i++{
			auxiliaryArray[auxiliaryArrayPointer] = array1[i]
			auxiliaryArrayPointer++
		}
	}else if arrayTwoPointer < len(array2){
		for i := arrayTwoPointer; i<= len(array2); i++{
			auxiliaryArray[auxiliaryArrayPointer] = array2[i]
			auxiliaryArrayPointer++
		}
	}
	return auxiliaryArray
}


func main() {
	list := NewArrayList(2)
	for i:=0; i<=20; i++{
		list.Add(i)
	}
	list.DisplayArray()
	list.DeleteElement(5)

	element := list.LinearSearch(1)
	if element == 0{
		fmt.Println("Not Found")
		return
	}
	fmt.Println("Found")

	IsPresent := list.BinarySearch(4)
	if IsPresent == 0{
		fmt.Println("Not Found")
		return
	}
	fmt.Println("Found")

}




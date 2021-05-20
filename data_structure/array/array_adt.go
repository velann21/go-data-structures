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




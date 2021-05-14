package main

import "fmt"

type List interface {
	Add(element int)
	DisplayArray()
	increaseArray()
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
		al.increaseArray()
	}
	al.Array[al.Length] = element
	al.Length++
}

func (al *ArrayList) DisplayArray(){
	for i:=0; i<al.Length;i++{
		fmt.Println(al.Array[i])
	}
}

func (al *ArrayList) increaseArray() {
	newSize := al.Size+10
	newArray := make([]int, newSize)
	for i, v := range al.Array{
		newArray[i] = v
	}
	al.Array = newArray
	al.Size = newSize
}

func main() {
	list := NewArrayList(2)
	for i:=0; i<=20; i++{
		list.Add(i)
	}
	list.DisplayArray()
}




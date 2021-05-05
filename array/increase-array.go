package main

import "errors"

type ArrayList struct {
	Size int
	Array []int
	LastInsertedPosition int
}
func main() {

	al := NewArrayList()
	al.Set(10)


}

func NewArrayList()*ArrayList{
	arr := make([]int, 5)
	return &ArrayList{Array: arr, LastInsertedPosition: 0}
}

func (al *ArrayList) Set(value int){
	if al.LastInsertedPosition == len(al.Array){
		al.increaseArraySize()
	}
	al.Array[al.LastInsertedPosition] = value
	al.LastInsertedPosition =+1
}

func (al *ArrayList) Get(position int)(int, error){
	if position > al.LastInsertedPosition{
		return 0, errors.New("Invalid ops")
	}
	return al.Array[position], nil

}



func (al *ArrayList) increaseArraySize(){
	newArray := make([]int, al.LastInsertedPosition+10)

	for i:=0; i<=al.LastInsertedPosition; i++{
		newArray[i] = al.Array[i]
	}
	al.Array = newArray

}

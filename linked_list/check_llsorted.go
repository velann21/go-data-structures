package main

import (
	"errors"
	"fmt"
)

func main() {

	node := NodeItem7{}
	node.Append(10)
	node.Append(20)
	node.Append(30)


	node.Append(50)
	node.Append(60)
	node.Append(70)
	node.Append(80)
	node.Append(90)
	node.Append(100)


	isSorted, err := node.IsSorted()
	if err != nil{

	}

	fmt.Println(isSorted)





}

type Node7 struct {
	Data int
	Next *Node7
}

type NodeItem7 struct {
	Head *Node7
	Tail *Node7
}

func (node *NodeItem7) Append(value int){
	if node.Head == nil{
		node.Head = &Node7{Data:value, Next:nil}
		return
	}

	last := node.Head
	for last.Next != nil{
		last = last.Next
	}

	last.Next = &Node7{Data:value, Next:nil}
}

func (node *NodeItem7) IsSorted()(bool,error){
	if node.Head == nil{
		return false, errors.New("Invalid request")
	}
	last := node.Head
	checkPoint :=  node.Head.Data
	for last != nil{
		if last.Data < checkPoint{
			return  false, nil
		}
		checkPoint = last.Data
		last = last.Next
	}
	return true, nil
}

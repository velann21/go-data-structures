package main

import (
	"errors"
	"fmt"
)

func main() {

	node := NodeItem10{}
	node.Append(10)
	node.Append(20)
	node.Append(30)
	node.Append(40)
	node.Append(50)
	node.Append(60)
	node.Append(70)

	node2 := NodeItem10{}
	node2.Append(100)
	node2.Append(120)
	node2.Append(130)
	node2.Append(140)
	node2.Append(150)
	node2.Append(160)
	node2.Append(170)

	err := node.ConcatLL(&node2)
	if err != nil{
		fmt.Println(err)
	}



}

type Node10 struct {
	Data int
	Next *Node10
}

type NodeItem10 struct {
	Head *Node10
	Tail *Node10
}

func (node *NodeItem10) Append(value int){
	if node.Head == nil{
		node.Head = &Node10{Data:value, Next:nil}
		return
	}

	last := node.Head

	for last.Next != nil{
		last = last.Next
	}

	last.Next = &Node10{Data:value, Next:nil}
}


func (node *NodeItem10) ConcatLL(node2 *NodeItem10) error{

	if node.Head == nil || node2.Head == nil{
		return errors.New("Invalid request")
	}

	firstL := node.Head
	for firstL.Next != nil{
		firstL = firstL.Next
	}
	firstL.Next = node2.Head

	return nil
}
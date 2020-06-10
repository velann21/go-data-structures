package main

import (
	"errors"
	"fmt"
)

func main() {
	node := NodeItem14{}
	node.Append(10)
	node.Append(20)
	node.Append(30)
	node.Append(40)
	node.Append(50)
	node.Append(60)
	node.Append(70)

	node.DisplayInBothDir()




}


type Node14 struct {
	Prev *Node14
	Data int
	Next *Node14
}

type NodeItem14 struct {
	Head *Node14
	Tail *Node14
}

func (node *NodeItem14) Append(value int){
	if node.Head == nil{
		node.Head = &Node14{Prev:nil, Data:value, Next:nil}
		node.Tail = node.Head
		return
	}
	last := node.Head
	for last.Next != nil{
		last = last.Next
	}
	last.Next = &Node14{Prev:last, Data:value, Next:nil}
}

func (node *NodeItem14) Display()error{
	if node.Head == nil {
		return errors.New("Invalid Request")
	}
	last := node.Head
	for last != nil{
		fmt.Println(last.Data)
		last = last.Next
	}
	return nil
}

func (node *NodeItem14) DisplayInBothDir()error{
	if node.Head == nil {
		return errors.New("Invalid Request")
	}
	last := node.Head
	var prev *Node14
	for last != nil{
		fmt.Println(last.Data)
		if last.Next == nil{
			prev = last
		}
		last = last.Next
	}

	for prev != nil{
		fmt.Println(prev.Data)
		prev = prev.Prev
	}
	return nil
}
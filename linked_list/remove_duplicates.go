package main

import (
	"errors"
	"fmt"
)

func main() {

	node := NodeItem8{}
	node.Append(10)
	node.Append(20)
	node.Append(30)
	node.Append(30)
	node.Append(50)

	err := node.RemoveDuplicates()
	if err != nil{
		fmt.Println(err)
	}



}

type Node8 struct {
	Data int
	Next *Node8
}

type NodeItem8 struct {
	Head *Node8
	Tail *Node8
}

func (node *NodeItem8) Append(value int){
	if node.Head == nil{
		node.Head = &Node8{Data:value, Next:nil}
		return
	}
	last := node.Head
	for last.Next != nil{
		last = last.Next
	}
	last.Next = &Node8{Data:value, Next:nil}
}

func (node *NodeItem8) RemoveDuplicates()error{
	if node.Head == nil{
		return errors.New("Invalid Request")
	}

	last := node.Head
	fastPtr := node.Head.Next
	for fastPtr != nil{
		if last.Data != fastPtr.Data {
			last = fastPtr
			fastPtr = fastPtr.Next
		}else{
			last.Next = fastPtr.Next
			fastPtr = last.Next
		}
	}

	return nil
}
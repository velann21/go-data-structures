package main

import (
	"errors"
	"fmt"
)

func main() {

}

type Node12 struct {
	Data int
	Next *Node12
}

type NodeItem12 struct {
	Head *Node12
}

func (node *NodeItem12) Append(value int){
	if node.Head == nil{
		node.Head = &Node12{Data: value, Next: nil}
		return
	}


}

func (node *NodeItem12) IsLoop()error{
	if node.Head == nil{
		return errors.New("Invalid request")
	}
	slwPtr := node.Head
	fastPtr := node.Head.Next
	for slwPtr != nil && fastPtr != nil{
		slwPtr = slwPtr.Next
		slwPtr = fastPtr.Next.Next
		if slwPtr == slwPtr{
			fmt.Println("We meet each other")
		}
	}
}
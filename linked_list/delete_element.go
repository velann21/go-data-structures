package main

import (
	"errors"
	"fmt"
)

func main(){
	node := NodeItem1{}
	node.Append(10)
	node.Append(20)
	node.Append(30)
	node.Append(40)
	node.Append(50)
	node.Append(60)
	node.Append(70)
	node.Append(80)
	node.Append(90)

	err := node.Delete(4)
	if err != nil{
		fmt.Println("Error occured")
	}
}

type Node6 struct {
	Data int
	Next *Node6
}

type NodeItem1 struct {
	Head *Node6
	Size int
}

func (node *NodeItem1) Append(value int){
	if node.Head == nil{
		node.Head = &Node6{Data:value, Next:nil}
		node.Size++
		return
	}

	last := node.Head

	for last.Next != nil{
		last = last.Next
	}

	last.Next = &Node6{Data:value, Next:nil}
	node.Size++
}

func (node *NodeItem1) Delete(index int) error{
	if node.Head == nil{
		return errors.New("Invalid Request")
	}

	last := node.Head
	lastIndex := 1
	for last != nil && lastIndex<index-1{
		last = last.Next
		lastIndex++
	}
	last.Next = last.Next.Next
	node.Size--
	return nil
}



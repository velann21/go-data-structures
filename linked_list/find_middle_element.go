package main

import (
	"errors"
	"fmt"
)

type Node15 struct {
	Data int
	Next *Node15
}

type NodeStruct15 struct {
	Head *Node15
	Tail *Node15
}

func (node *NodeStruct15) Insert(element int){
	if node.Head == nil{
		node.Head = &Node15{Data:element, Next:nil}
		return
	}

	last := node.Head
	for last.Next != nil{
		last = last.Next
	}
	last.Next = &Node15{Data:element, Next:nil}
}

func (node *NodeStruct15) Display()error{
	if node.Head == nil{
		return errors.New("Invalid request")
	}
	last := node.Head
	for last != nil{
		fmt.Println(last.Data)
		last = last.Next
	}
	return nil
}

func (node *NodeStruct15) FindMiddlePoint()error{
	if node.Head == nil{
		return errors.New("Invalid request")
	}
	p := node.Head
	q := node.Head
	for q!=nil{
      q = q.Next
      if q!=nil{
      	q = q.Next
	  }
      if q!=nil{
      	p = p.Next
	  }
	}

	fmt.Println(p.Data)
	return nil
}

func main() {
	node := NodeStruct15{}
	node.Insert(10)
	node.Insert(20)
	node.Insert(30)
	node.Insert(40)
	node.Insert(50)
	node.Insert(60)
	node.Insert(70)
	node.Insert(80)

	err := node.Display()
	if err != nil{
		fmt.Println(err)
	}
	_ = node.FindMiddlePoint()
}
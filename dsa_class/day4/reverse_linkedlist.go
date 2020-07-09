package main

import (
	"errors"
	"fmt"
)

func main()  {
	node := Nodestruct{}
	node.Insert(10)
	node.Insert(20)
	node.Insert(30)
	node.Insert(40)
	node.Insert(50)
	node.Insert(60)

	err := node.ReverseLinkedList()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(node)
}


type Node struct {
	Data int
	Next *Node
}

type Nodestruct struct {
	Head *Node
}

func (node *Nodestruct) Insert(data int){
	if node.Head == nil{
		node.Head = &Node{Data:data, Next:nil}
		return
	}

	last := node.Head
	for last.Next != nil{
		last = last.Next
	}
	last.Next = &Node{Data:data, Next:nil}
}

func (node *Nodestruct) ReverseLinkedList()error{
	if node.Head == nil{
		return errors.New("Invalid request")
	}
	main := node.Head
	secFoll := &Node{}
	firFoll := &Node{}
	for main != nil{
		secFoll = firFoll
		firFoll = main
		main = main.Next
		firFoll.Next = secFoll
	}
	node.Head = firFoll
	return nil
}




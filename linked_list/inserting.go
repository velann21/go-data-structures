package main

import "errors"

func main(){

	node := NodeItem{}
	node.Append(10)
	node.Append(20)
	node.Append(30)
	node.Append(40)
	node.Append(50)
	node.Append(60)
	node.Append(70)
	node.Append(80)
	node.Append(90)

	err := node.Insert(35, 4)
	if err != nil{

	}
	err = node.Insert(5, 0)
	if err != nil{

	}

	err = node.Insert(25, 3)
	if err != nil{

	}

	err = node.Insert(100, 13)
	if err != nil{

	}

	err = node.Insert(110, 14)
	if err != nil{

	}

	err = node.Insert(110, 15)
	if err != nil{

	}


}

type Node5 struct {
	Data int
	Next *Node5
}

type NodeItem struct {
	Head *Node5
	Tail *Node5
	LastIndex int
}

func (node *NodeItem) Append(value int){

	if node.Head == nil{
		node.Head = &Node5{Data:value, Next:nil}
		node.LastIndex++
		return
	}

	last := node.Head
	for last.Next != nil{
		last = last.Next
	}

	last.Next = &Node5{Data:value, Next:nil}
	node.LastIndex++
}

func (node *NodeItem) Insert(value int, index int)error{
	if node.Head == nil || index < 0 || index-1 > node.LastIndex{
		return errors.New("Invalid Index")
	}

	if index == 0{
		q := node.Head
		p := Node5{Data:value, Next:q}
		node.Head = &p
		node.LastIndex++
		return nil
	}

	last := node.Head
	lastIndex := 1
	for last != nil && lastIndex < index-1{
		last = last.Next
		lastIndex++
	}

	p := Node5{Data:value, Next:last.Next}
	last.Next = &p
	node.LastIndex++
	return nil
}
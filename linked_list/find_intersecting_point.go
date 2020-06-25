package main

import (
	"errors"
	"fmt"
)

func main()  {

}

type Node16 struct {
	Data int
	Next *Node16
}

type NodeType16 struct {
	Head *Node16
}

func (node *NodeType16) Insert(data int){
	if node.Head == nil{
		node.Head = &Node16{Data:data, Next:nil}
		return
	}

	last := node.Head
	for last.Next != nil{
		last = last.Next
	}

	last.Next = &Node16{Data:data, Next:nil}
}

func (node *NodeType16) FindIntersectingPoint(node2 *NodeType16)error{
	if node.Head == nil{
		return errors.New("Invalid request")
	}
    stk1 := Stack{}
	stk2 := Stack{}
	nodeHead := node.Head
	nodeHead2 := node2.Head
	for nodeHead != nil{
		stk1.Push(nodeHead)
		nodeHead = nodeHead.Next
	}

	for nodeHead2 != nil{
		stk2.Push(nodeHead)
		nodeHead2 = nodeHead2.Next
	}
	p := &Node16{}
	for stk1.Pop() == stk2.Pop(){
		p = stk1.Pop()
	}
	
	fmt.Println(p)

	return nil
}


type Stack struct {
	elements []*Node16
}

func (stk *Stack) Push(data *Node16){
	stk.elements = append(stk.elements, data)
}

func (stk *Stack) Pop()*Node16{
	data := stk.elements[len(stk.elements)-1]
	stk.elements = stk.elements[0:len(stk.elements)-2]
	return data
}

func (stk *Stack) IsEmpty()bool{
	return len(stk.elements) == 0
}
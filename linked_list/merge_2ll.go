package main

import (
	"errors"
	"fmt"
)

func main() {
	node := NodeItem11{}
	node.Append(1)
	node.Append(5)
	node.Append(10)


	node2 := NodeItem11{}
	node2.Append(2)
	node2.Append(3)
	node2.Append(12)

	err := node.MergeLL(&node2)
	if err != nil{
		fmt.Println(err)
	}

}

type Node11 struct {
	Data int
	Next *Node11
}

type NodeItem11 struct {
	Head *Node11
	Tail *Node11
}
//Am repeating this just to increase the muscle memory
func (node *NodeItem11) Append(value int){

	if node.Head == nil{
		node.Head = &Node11{Data:value, Next:nil}
		return
	}

	last := node.Head

	for last.Next != nil{
		last = last.Next
	}

	last.Next = &Node11{Data:value, Next:nil}

}

func (node *NodeItem11) MergeLL(node1 *NodeItem11) error {
	if node.Head == nil{
		return errors.New("Invalid request")
	}
    first := node.Head
    second := node1.Head
	var third *Node11
    var last *Node11
    if first.Data < second.Data{
       third = first
       last = third
       first = first.Next
       last.Next = nil
	}else {
		third = second
		last = third
		second = second.Next
		last.Next = nil
	}

    if last == nil{
    	return errors.New("Something wrong")
	}
    for first != nil && second != nil{
    	if first.Data < second.Data{
			last.Next = first
    		last = first
    		first = first.Next
    		last.Next = nil
		}else {
			last.Next = second
			last = second
			second = second.Next
			last.Next = nil
		}
	}

    if first != nil{
    	last.Next = first
	}else if second != nil{
		last.Next = second
	}
	node.Head = third
	return nil
}

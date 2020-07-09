package main

import "fmt"

type Node15 struct {
	Data int
	Next *Node15
}

type NodeStruct15 struct {
	Head *Node15
}

func (node *NodeStruct15) Insert(data int){
	if node.Head == nil{
		node.Head = &Node15{Data:data, Next:nil}
		return
	}
	last := node.Head
	for last.Next != nil{
		last = last.Next
	}

	last.Next = &Node15{Data:data, Next:nil}
}

func (node *NodeStruct15) MergeLinkedList(ll2 *NodeStruct15){
	f := node.Head
	s := ll2.Head

	third := &Node15{}
	last := &Node15{}
	fmt.Print(third)
	if f.Data < s.Data{
		third = f
		last = f
		if f.Next != nil{
			f = f.Next
		}
		last.Next = nil
	}else{
		third = s
		last = s
		if s.Next != nil{
			s = s.Next
		}
		last.Next = nil
	}
	for f != nil && s!= nil{
		if f.Data < s.Data{
			last = f
			f = f.Next
			last = nil
		}else {
			last = s
			s = s.Next
			last = nil
		}
	}

	if s != nil && last != nil{
		for s != nil {
			last.Next = s
			s =s.Next
		}
	}else if f != nil && last != nil{
		for f != nil {
			last.Next = f
			f =f.Next
		}
	}

}
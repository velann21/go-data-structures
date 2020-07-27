package linked_list

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
			last.Next = f
			last = f
			f = f.Next
			last.Next = nil
		}else {
			last.Next = s
			last = s
			s = s.Next
			last.Next = nil
		}
	}

	if s != nil && last != nil{
		for s != nil {
			last.Next = s
			last = s
			s = s.Next
			last.Next = nil
		}
	}else if f != nil && last != nil{
		for f != nil {
			last.Next = f
			last = f
			f = f.Next
			last.Next = nil
		}
	}
}

func main()  {
	node := &NodeStruct15{}
	node.Insert(10)
	node.Insert(20)
	node.Insert(30)
	node.Insert(40)
	node.Insert(50)
	node.Insert(60)

	node1 := &NodeStruct15{}
	node1.Insert(100)
	node1.Insert(200)
	node1.Insert(300)
	node1.Insert(400)
	node1.Insert(500)
	node1.Insert(600)

	node.MergeLinkedList(node1)

}
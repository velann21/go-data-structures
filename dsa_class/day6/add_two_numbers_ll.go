package main

import "fmt"

func main(){

	node := &NodePtrStruct{}
	node.Insert(10)
	node.Insert(20)
	node.Insert(30)

	node2 := &NodePtrStruct{}
	node2.Insert(1)
	node2.Insert(2)
	node2.Insert(3)

	sumList := node.addTwoNumbers(node2)
	for sumList!= nil{
		fmt.Println(sumList.Data)
		sumList = sumList.Next
	}


}

type NodePtr struct {
	Data int
	Next *NodePtr
}

type NodePtrStruct struct {
	Head *NodePtr
}

func (node *NodePtrStruct) Insert(data int){
	if node.Head == nil{
		node.Head = &NodePtr{Data:data, Next:nil}
		return
	}

	last := node.Head

	for last.Next != nil{
		last = last.Next
	}

	last.Next = &NodePtr{Data:data, Next:nil}
	return
}

func (node *NodePtrStruct) addTwoNumbers(node2 *NodePtrStruct) *NodePtr {
	if node == nil || node2 == nil{
		return nil
	}
	n1 := node.Head
	n2 := node2.Head
	n1Temp := n1
	sum := 0
	for n1 != nil && n2 != nil{
		sum = n1.Data + n2.Data
		first := sum / 10;
		last := sum % 10;
		n1.Data = first
		if n1.Next != nil{
			n1.Next.Data += last
		}else{
			n1.Next = &NodePtr{Data:last, Next:nil}
			n1Temp = n1
		}
		n1 = n1.Next
		n2 = n2.Next
	}

	for n2 != nil {
		n1Temp.Next = n2
		n2 = n2.Next
	}

	return node.Head


}

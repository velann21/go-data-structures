package main

import (
	"errors"
	"fmt"
)

func main() {

	defer recover()

	node := NodeItem9{}
	node.Append(10)
	node.Append(20)
	node.Append(30)
	node.Append(40)
	node.Append(50)
	node.Append(60)
	node.Append(70)
	node.Append(80)
	node.Append(90)

	//err := node.ReverseLinkedListWithAuxilary()
	//if err != nil {
	//	fmt.Println(err)
	//}

	err := node.ReverseLinkedSlidingPointer()
	if err != nil {
		fmt.Println(err)
	}


}

type Node9 struct {
	Data int
	Next *Node9
}

type NodeItem9 struct {
	Head *Node9
}

func (node *NodeItem9) Append(value int){
	if node.Head == nil{
		node.Head = &Node9{Data:value, Next:nil}
		return
	}

	last := node.Head

	for last.Next != nil{
		last = last.Next
	}

	last.Next = &Node9{Data:value, Next:nil}
}


func (node *NodeItem9) ReverseLinkedListWithAuxilary()error{
	if node.Head == nil{
		return errors.New("Invalid request")
	}

	auxilary := [12]int{}
	last := node.Head
	var prev *Node9
	i := 0
	for last != nil{
		auxilary[i] = last.Data
		i++
		last = last.Next
		prev = last
	}

	i--
	last = node.Head

	for i >= 0 && last != nil{
		last.Data = auxilary[i]
		i--
		last = last.Next
	}
    fmt.Println(prev)
	return nil
}

func (node *NodeItem9) ReverseLinkedSlidingPointer()error{
	if node.Head == nil{
		return errors.New("Invalid request")
	}

	last := node.Head
	var firFoll *Node9
	var secFoll *Node9
	
	for last!=nil{

		secFoll = firFoll
		firFoll = last
		last = last.Next
		firFoll.Next = secFoll

	}
    node.Head = firFoll
	return nil
}

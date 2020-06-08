package main

import (
	"errors"
	"fmt"
)

func main() {

	node := NodeItem12{}
	node.AddElement(10, false)
	node.AddElement(20, false)
	node.AddElement(30, false)
	node.AddElement(40, false)
	node.AddElement(50, false)
	node.AddElement(10, true)
	err := node.Insert(05, 0)
	if err != nil{
		fmt.Println(err)
	}

	err = node.Insert(25, 3)
	if err != nil{
		fmt.Println(err)
	}
	err = node.Insert(55, 7)
	if err != nil{
		fmt.Println(err)
	}

	isLoop, err := node.IsLoop()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(isLoop)

}

type Node12 struct {
	Item int
	Next *Node12
}

type NodeItem12 struct {
	Head *Node12
	Tail Node12
}

func (node *NodeItem12) AddElement(value int, circle bool){
	if node.Head == nil {
		node.Head = &Node12{Item:value}
		return
	}
	last := node.Head
	var detectedNode *Node12 = nil
	for last.Next != nil{
		if last.Item == value{
			detectedNode = last
		}
		last = last.Next

	}

	if circle {
       last.Next = detectedNode
        return
	}

	last.Next = &Node12{Item:value, Next:nil}
}

func (node NodeItem12) IsLoop()(bool, error){
	if node.Head == nil{
		return false, errors.New("Invalid request")
	}
	slowPtr := node.Head
	fastPtr := node.Head
	for slowPtr != nil && fastPtr.Next != nil{
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
		if slowPtr == fastPtr{
			fmt.Println("Yes its circluar")
			return true, nil
		}
	}
	return false ,nil
}

func (node NodeItem12) Insert(value int, pos int) error {
	if node.Head == nil {
		return errors.New("Invalid request")
	}
	if pos == 0 {
		newNode := &Node12{Item:value, Next:node.Head}
		last := node.Head
		for last.Next != node.Head && last.Next != nil{
			last = last.Next
		}
		last.Next = newNode
		node.Head = newNode
	}else {
		last :=  node.Head
		position := 1
		for position != pos-1 && last != nil{
			last = last.Next
			position++
		}
		newNode := &Node12{Item:value, Next:last.Next}
		last.Next = newNode
	}
	return nil
}
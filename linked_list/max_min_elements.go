package main

import (
	"errors"
	"fmt"
)

func main() {
	node := Node3{}
	node.Append(10)
	node.Append(20)
	node.Append(30)
	node.Append(140)
	node.Append(110)
	node.Append(50)
	node.Append(60)
	node.Append(70)
	node.Append(80)
	node.Append(90)

	max, _ := node.findMax()
	fmt.Println(max)

	maxRec := node.findMaxUsingRecursive(head3)
	fmt.Println(maxRec)

}

type Node3 struct {
	Value int
	Next *Node3
	Max int
}
var head3 *Node3

func (node3 Node3) Append(value int){
	if head3 == nil{
		head3 = &Node3{Value:value, Next:nil}
	}
	last := head3
	for last.Next != nil{
        last = last.Next
	}
	last.Next = &Node3{Value:value, Next:nil}
}

func (node3 Node3) findMax()(int,error){
	if head3 == nil{
		return  0, errors.New("Invalid request")
	}
	last := head3
	for last != nil{
		lastVal := last.Value
		if lastVal > node3.Max{
			node3.Max = lastVal
		}
		last = last.Next
	}
	return node3.Max,nil
}


func (node3 Node3) findMaxUsingRecursive(node *Node3)int{
	if node == nil{
		return node3.Max
	}
	if node.Value > node3.Max{
		node3.Max = node.Value
	}
	return node3.findMaxUsingRecursive(node.Next)
}

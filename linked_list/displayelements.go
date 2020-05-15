package main

import (
	"errors"
	"fmt"
	"log"
)

type Node1 struct {
	Value int
	Next *Node1
}

var head *Node1


func (node Node1) Append(value int){
	if head == nil{
		head = &Node1{Value:value, Next:nil}
		return
	}
	last := head
	for last.Next != nil {
		last = last.Next
	}
	last.Next = &Node1{Value:value, Next:nil}
}

func (node Node1) Display() error{

	if head == nil{
		return errors.New("No elemet present")
	}
	last := head
	for last != nil{
		fmt.Println(last.Value)
		last = last.Next
	}
	return nil
}

func (node Node1) DisplayRecursively(head *Node1) error{
	fmt.Println()
	if head == nil{
		return errors.New("No element present")
	}
	last := head
	if last != nil{
		fmt.Println(last.Value)
		_ = node.DisplayRecursively(last.Next)

	}
	return nil
}

func main(){

	nod := Node1{}
	nod.Append(10)
	nod.Append(20)
	nod.Append(30)
	nod.Append(40)
	nod.Append(50)
	nod.Append(60)
	nod.Append(70)

	err := nod.Display()
	if err != nil{
		log.Println(err)
	}

	err = nod.DisplayRecursively(head)
	if err != nil{
		log.Println(err)
	}

}

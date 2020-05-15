package main

import (
	"errors"
	"fmt"
)

func main(){
	node := Node2{}
	node.Append(10)
	node.Append(20)
	node.Append(30)
	node.Append(40)
	node.Append(50)
	node.Append(60)
	node.Append(70)
	node.Append(80)

	count, err := node.CountNoOfElements()
	if err!=nil{

	}
	fmt.Println(count)

	sum, err :=node.SumAllElements()
	if err!=nil{

	}

	fmt.Println(sum)

	recCount := node.RecursivelyCountAllElements(head2)
    fmt.Println(recCount)

	recSum := node.RecursivelySumAllElements(head2)
	fmt.Println(recSum)


}

type Node2 struct {
	Data int
	Next *Node2
}

var head2 *Node2
func (node Node2) Append(value int)error{
	if head2 == nil {
		head2 = &Node2{Data:value, Next:nil}
		return nil
	}
	last := head2
	for last.Next != nil{
		last = last.Next
	}
	last.Next = &Node2{Data:value, Next:nil}
	return nil
}

func (node Node2) CountNoOfElements()(int,error){
	count := 0
	if head2 == nil {
		return count, errors.New("Invalid request")
	}
	last := head2
	for last != nil{
		count++
		last = last.Next
	}
	return count, nil
}

func (node Node2) SumAllElements()(int,error){
	sum := 0
	if head2 == nil {
		return sum, errors.New("Invalid request")
	}
	last := head2
	for last != nil{
		sum += last.Data
		last = last.Next
	}
	return sum, nil
}

func (node Node2) RecursivelySumAllElements(cnode *Node2)int{
	if cnode == nil{
		return 0
	}
	return node.RecursivelySumAllElements(cnode.Next) + cnode.Data
}


func (node Node2) RecursivelyCountAllElements(cnode *Node2)int{
	if cnode == nil{
		return 0
	}
	return node.RecursivelyCountAllElements(cnode.Next) + 1
}

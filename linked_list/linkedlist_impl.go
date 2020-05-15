package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	itemNode := ItemNode{}
	itemNode.Insert(10)
	itemNode.Insert(20)
	err := itemNode.SetElement(80, 1)
	if err != nil{
		log.Println(err)
	}
	itemNode.Insert(30)
	itemNode.Insert(40)
	itemNode.Insert(50)
	itemNode.Insert(60)
	_, err = itemNode.GetElement(1)
	if err != nil{
		return
	}
	//fmt.Println(*value)
	value1, err := itemNode.GetElement(5)
	if err != nil{
		return
	}
	if value1 != nil{
		//fmt.Print(*value1)
	}
	err = itemNode.PrintAllElement()
	if err != nil{

	}

	array, err := itemNode.PrintAllElementFunc(
		func() []int{
			i := make([]int, 0)
			last := itemNode.head
			for{
				if last.NextNode == nil{
					i = append(i, last.Value)
					break
				}
				i = append(i, last.Value)
				last = last.NextNode

			}
			return i
		})

	fmt.Println(array)
}

type Node struct {
	NextNode     *Node
	Value    int
}

type ItemNode struct {
	head *Node
	size int
}

func (node *ItemNode) Insert(value int) {

	if node.head == nil{
		node.head = &Node{NextNode:nil, Value:value}
		node.size++
		return
	}
	last := node.head
	for {
		if last.NextNode == nil {
			break
		}
		last = last.NextNode
	}
	last.NextNode = &Node{NextNode:nil, Value:value}
	node.size++

}

func (node *ItemNode) SetElement(value int, index int) error{

	if node.head == nil{
		return errors.New("Invalid request")
	}
	if index > node.size{
		return errors.New("Invalid request")
	}
	last := node.head
	startValue := 0
	for {
		if last.NextNode == nil || startValue == index {
			if startValue == index {
				last.Value = value
			}
			break
		}
		last = last.NextNode
		startValue++
	}
	return nil
}

func (node *ItemNode) GetElement(index int) (*int,error){
	if node.head == nil{
		return nil, errors.New("Invalid request")
	}
	if index > node.size{
		return nil, errors.New("Invalid request")
	}
	last := node.head
	startValue := 0
	for {
		if last.NextNode == nil || startValue == index {
			if startValue == index {
				return &last.Value, nil
			}
			break
		}
		last = last.NextNode
		startValue++
	}
	return nil, nil
}

func (node *ItemNode) PrintAllElement()error{
	if node.head == nil{
		return  errors.New("Invalid request")
	}
	last := node.head
	for {
		if last.NextNode == nil{
			fmt.Println(last.Value)
			return nil
		}
		fmt.Println(last.Value)
		last = last.NextNode
	}
}

func (node *ItemNode) PrintAllElementFunc(tripper func()[]int) ([]int, error){
	if tripper == nil{
		return nil, errors.New("")
	}
	i := tripper()
	return i,nil
}
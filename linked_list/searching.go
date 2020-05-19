package main

import (
	"errors"
	"fmt"
	"log"
)

func main(){

	node := NodeUtil{}
	node.Insert(10)
	node.Insert(20)
	node.Insert(30)
	node.Insert(40)
	node.Insert(50)
	node.Insert(60)
	node.Insert(70)
	node.Insert(80)
	node.Insert(90)

	element, err := node.Search(100)
	if err != nil{
		log.Println(err)
	}
	if element != nil{
		fmt.Println(element.Data)
	}

	p, err := node.MoveToHeadSearch(70)
	if err != nil{
		log.Println(err)
	}
	if p != nil{
		fmt.Println(p.Data)
	}



}

type Node4 struct {
	Data int
	Next *Node4
}

type NodeUtil struct {
	Head *Node4
	Tail *Node4
	Index int
}
func (node *NodeUtil) Insert(value int){
	if node.Head == nil{
		node.Head = &Node4{Data:value}
		node.Index++
		return
	}

	last := node.Head
	for last.Next != nil{
        last = last.Next
	}

	newNode := &Node4{Data:value, Next:nil}
	last.Next = newNode
	node.Tail = newNode
	node.Index++
}


func (node *NodeUtil) Search(key int)(*Node4, error){
	if node.Head == nil{
		return nil, errors.New("Invalid request")
	}
	var finder *Node4 = nil
	last := node.Head
	for last != nil{
		if last.Data == key{
			println(last.Data)
			finder = last
		}
		last = last.Next
	}
	return finder, nil
}


func (node *NodeUtil) MoveToHeadSearch(key int)(*Node4, error){
	if node.Head == nil{
		return nil, errors.New("Invalid request")
	}
	p := node.Head
	var q *Node4 = nil
	for p != nil{
		if p.Data == key {
			if (q != nil) {
				q.Next = p.Next
				p.Next = node.Head
				node.Head = p
			}
          return p, nil
		}
		q = p
		p = p.Next
	}
	return nil, nil
}
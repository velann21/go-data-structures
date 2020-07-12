package main

import "fmt"

func main() {
	node := Node18Struct{}
	node.Insert(10)
	node.Insert(20)
	node.Insert(30)
	node.Insert(40)
	node.Insert(50)
	node.IsEvenOddLength()
	// 10 -> 20 -> 30 -> 40 -> 50
	// 10 -> 20 -> 30 -> 40 -> 50 -> 60


}

type Node18 struct {
	Data int
	Next *Node18
}

type Node18Struct struct {
	Head *Node18
}

func (node *Node18Struct) Insert(data int){
	if node.Head == nil{
		node.Head = &Node18{Data:data, Next:nil}
		return
	}

	last := node.Head
	for last.Next != nil{
		last = last.Next
	}
	last.Next = &Node18{Data:data, Next:nil}
}

func (node *Node18Struct) IsEvenOddLength(){
	if node.Head == nil{
		return
	}

	last := node.Head
	for last.Next != nil {
		last = last.Next.Next
		if last == nil{
			break
		}
	}

	if last == nil{
		fmt.Print("Even")
	}else{
		fmt.Print("Odd")
	}
}

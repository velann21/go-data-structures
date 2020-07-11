package main

import "fmt"

func main() {

	node := Node17Struct{}
	_ = node.MakeCircularLL(10, nil)
	n2 := node.MakeCircularLL(20, nil)
	_ = node.MakeCircularLL(30, nil)
	_ = node.MakeCircularLL(40, nil)
	_ = node.MakeCircularLL(50, n2)



}

type Node17 struct {
	Data int
	Next *Node17
}

type Node17Struct struct {
	Head *Node17
}

func (node *Node17Struct) Insert(data int){
	if node.Head == nil{
		node.Head = &Node17{Data:data, Next:nil}
		return
	}

	last := node.Head
	for last.Next != nil{
		last = last.Next
	}
	last.Next = &Node17{Data:data, Next:nil}
}

func (node *Node17Struct) MakeCircularLL(data int, circular *Node17)*Node17{
	if node.Head == nil{
		node.Head = &Node17{Data:data, Next:nil}
		return node.Head
	}

	last := node.Head
	for last.Next != nil{
		last = last.Next
	}
	last.Next = &Node17{Data:data, Next:circular}
	return last.Next
}

func (node *Node17Struct) DetectLoop(){
	if node.Head == nil{
		return
	}
	slow := node.Head
	fast := node.Head

	for fast != nil{
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast{
			fmt.Print("Loop detected")
			fmt.Println(slow.Data)
			fmt.Println(fast.Data)
			return
		}
	}
}


func (node *Node17Struct) DetectLoopAndFindJunctionPoint(){
	if node.Head == nil{
		return
	}

	slow := node.Head
	fast := node.Head

	for fast != nil{
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast{
			fmt.Print("Loop detected")
			fmt.Println(slow.Data)
			fmt.Println(fast.Data)
			break
		}
	}
	slow = node.Head
	for fast != nil{
		slow = slow.Next
		fast = fast.Next
		if slow == fast{
			fmt.Print("Junction point detected")
			fmt.Println(slow.Data)
		}
	}
}



func (node *Node17Struct) DetectLoopAndFindJunctionAndRemove(){
	if node.Head == nil{
		return
	}

	slow := node.Head
	fast := node.Head

	for fast != nil{
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast{
			fmt.Print("Loop detected")
			fmt.Println(slow.Data)
			fmt.Println(fast.Data)
			break
		}
	}
	slow = node.Head
	fastPrev := &Node17{}
	fmt.Print(fastPrev)
	for fast != nil{
		slow = slow.Next
		fastPrev = fast
		fast = fast.Next
		if slow == fast{
			fmt.Print("Junction point detected")
			fmt.Println(slow.Data)
			break
		}
	}
	fastPrev = nil


}



func (node *Node17Struct) DetectLoopAndFindJunctionAndLength(){
	if node.Head == nil{
		return
	}
	slow := node.Head
	fast := node.Head
	for fast != nil{
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast{
			fmt.Print("Loop detected")
			fmt.Println(slow.Data)
			fmt.Println(fast.Data)
			break
		}
	}
	slow = node.Head
	loopCounter := 1
	for fast != nil{
		slow = slow.Next
		loopCounter += 1
		fast = fast.Next
		if slow == fast{
			fmt.Print("Junction point detected")
			fmt.Println(slow.Data)
			break
		}
	}
	fmt.Print(loopCounter)
}




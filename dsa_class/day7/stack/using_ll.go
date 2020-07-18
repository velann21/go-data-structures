package main

import "fmt"

func main() {
	stk := StackNodeStruct{}
	stk.Push(10)
	stk.Push(20)
	stk.Push(30)
	stk.Push(40)
	stk.Push(50)

	fmt.Println(stk.Pop().Data)
	fmt.Println(stk.Pop().Data)
	fmt.Println(stk.Pop().Data)
	fmt.Println(stk.Pop().Data)
	fmt.Println(stk.Pop().Data)

}

type StackNode struct {
	Data int
	Next *StackNode
}

type StackNodeStruct struct {
	Head *StackNode
}

func (stk *StackNodeStruct) Push(data int){
	if stk.Head == nil{
		stk.Head = &StackNode{Data:data, Next:nil}
		return
	}
	newNode := &StackNode{Data:data, Next:stk.Head}
	stk.Head = newNode
}

func (stk *StackNodeStruct) Pop()*StackNode{
	if stk.Head == nil{
		return nil
	}
	temp := stk.Head
	stk.Head = stk.Head.Next
	return temp
}

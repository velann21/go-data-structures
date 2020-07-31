package main

import "fmt"

func main() {
	arr := []int{1,2,3,4,5}
	resultStack := QueueUsingStack(arr)
	for !resultStack.IsEmpty(){
		fmt.Println(resultStack.Pop().Data)
	}
}

func QueueUsingStack(arr []int)*StackNode1{
	if arr == nil{
		return nil
	}
	mainStk := StackNode1{}
	tempStk := StackNode1{}
	for _, v := range arr{
		if mainStk.IsEmpty(){
			mainStk.Push(v)
		}
		for !mainStk.IsEmpty(){
			tempStk.Push(mainStk.Pop().Data)
		}
		mainStk.Push(v)
		for !tempStk.IsEmpty(){
			mainStk.Push(tempStk.Pop().Data)
		}
	}
	return &mainStk
}

type Stack1 struct {
	Data int
	Next *Stack1
}

type StackNode1 struct {
	Head *Stack1
}

func (stk *StackNode1) Push(data int){
	if stk.Head == nil{
		stk.Head = &Stack1{Data:data, Next:nil}
		return
	}
	newNode := &Stack1{Data:data, Next:stk.Head}
	stk.Head = newNode
	return
}

func (stk *StackNode1) Pop()*Stack1{
	if stk.Head == nil{
		return nil
	}
	temp := stk.Head
	if stk.Head.Next != nil{
		stk.Head = stk.Head.Next
	}else{
		stk.Head = nil
	}
	return temp
}

func (stk *StackNode1) Peek()*Stack1{
	if stk.Head == nil{
		return nil
	}
	return stk.Head
}

func (stk *StackNode1) IsEmpty() bool{
	return stk.Head == nil
}




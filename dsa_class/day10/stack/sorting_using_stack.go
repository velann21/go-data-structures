package main

import "fmt"

func main() {
	arr := []int{3,5,2,1,6,7}
	resArr := Sorting(arr)
	for !resArr.IsEmpty(){
		fmt.Println(resArr.Pop().Data)
	}


}

func Sorting(arr []int)*StackNode{
	if len(arr) <= 0 {
		return nil
	}
	primaryStk :=  StackNode{}
	tempStk := StackNode{}

	if primaryStk.IsEmpty(){
		primaryStk.Push(arr[0])
	}
	for i:=1; i<len(arr); i++{

         for !primaryStk.IsEmpty() && arr[i] >  primaryStk.Peek().Data{
         	data := primaryStk.Pop()
         	tempStk.Push(data.Data)
		 }

		 primaryStk.Push(arr[i])

         for !tempStk.IsEmpty(){
			 tempData := tempStk.Pop()
			 primaryStk.Push(tempData.Data)
		 }
	}

	return &primaryStk
}

type Stack struct {
	Data int
	Next *Stack
}

type StackNode struct {
	Head *Stack
}

func (stk *StackNode) Push(data int){
	if stk.Head == nil{
		stk.Head = &Stack{Data:data, Next:nil}
		return
	}
	newNode := &Stack{Data:data, Next:stk.Head}
	stk.Head = newNode
	return
}

func (stk *StackNode) Pop()*Stack{
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

func (stk *StackNode) Peek()*Stack{
	if stk.Head == nil{
		return nil
	}
	return stk.Head
}

func (stk *StackNode) IsEmpty() bool{
	return stk.Head == nil
}




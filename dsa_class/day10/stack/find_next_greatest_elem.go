package main

import "fmt"

// Problems to solve for this week:
//----

//Gas Station
//Valid Tic-Tac-Toe State
//Maximum Frequency Stack

func main() {
	arr := []int{2,1,5,7,8}
	stk := GreaterEleStack{}
	qp := QueryProcessor{stack:&stk}
	for _, v := range arr{
		qp.FindNextGreaterElement(v)
	}
}

type QueryProcessor struct {
     stack 	*GreaterEleStack
}

func (prc *QueryProcessor)FindNextGreaterElement(data int){
	if prc.stack.IsEmpty(){
		prc.stack.Push(data)
		return
	}

	if !prc.stack.IsEmpty() && data < prc.stack.Peek().Data{
		prc.stack.Push(data)
		return
	}

	for !prc.stack.IsEmpty() && data > prc.stack.Peek().Data{
		elem := prc.stack.Pop()
		fmt.Println("Next Greatest element for ",elem.Data," -> ", data)
	}
	prc.stack.Push(data)
}

type GreaterEleStackNode struct {
	Data int
	Next *GreaterEleStackNode
}

type GreaterEleStack struct {
	Head *GreaterEleStackNode
}

func (ges *GreaterEleStack) Push(data int){
	if ges.Head == nil{
		ges.Head = &GreaterEleStackNode{Data:data, Next:nil}
		return
	}
	newNode := &GreaterEleStackNode{Data:data, Next:ges.Head}
	ges.Head = newNode
	return
}

func (ges *GreaterEleStack) Pop()*GreaterEleStackNode{
	if ges.Head == nil{
		return nil
	}

	temp := ges.Head
	if ges.Head.Next != nil{
		ges.Head = ges.Head.Next
	}else{
		ges.Head = nil
	}
	return temp
}

func (ges *GreaterEleStack) Peek()*GreaterEleStackNode{
	if ges.Head == nil{
		return nil
	}
	return ges.Head
}

func (ges *GreaterEleStack) IsEmpty()bool{
	return ges.Head == nil
}
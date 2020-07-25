package stack

import (
	"fmt"
)

func main() {

	node := stackNodeStruct{}
	node1 := stackNodeStruct{}
	node.Push(10)
	node1.CheckMin(10)
	node.Push(20)
	node1.CheckMin(20)
	node.Push(30)
	node1.CheckMin(30)
	node.Push(05)
	node1.CheckMin(05)
	node.Push(01)
	node1.CheckMin(01)

	fmt.Println("Min ATM: ", node1.GetMin().Data)

	pop1 := node.Pop()
	fmt.Println("Pop1 data",pop1.Data)
	node1.TakeOutIfMatchesMin(pop1)
	fmt.Println("Min ATM: ", node1.GetMin().Data)
	pop2 := node.Pop()
	fmt.Println("Pop2 data",pop2.Data)
	node1.TakeOutIfMatchesMin(pop2)

	fmt.Println("Min ATM: ", node1.GetMin().Data)


}



type stackNode struct {
	Data int
	Next *stackNode
}

type stackNodeStruct struct {
	Head *stackNode
}

func (stk *stackNodeStruct) Push(data int){
	if stk.Head == nil{
		stk.Head = &stackNode{Data:data, Next:nil}
		return
	}
	newNode := &stackNode{Data:data, Next:stk.Head}
	stk.Head = newNode
}

func (stk *stackNodeStruct) Pop()*stackNode{
	if stk.Head == nil{
		return nil
	}
	temp := stk.Head
	if stk.Head.Next != nil{
		stk.Head = stk.Head.Next
	}else {
		stk.Head = nil
	}
	return temp
}

func (stk *stackNodeStruct) CheckMin(data int){
	if stk.Head == nil{
		stk.Push(data)
	}

	if data < stk.Head.Data{
		stk.Push(data)
	}
}

func (stk *stackNodeStruct) GetMin()*stackNode{
	return stk.Head
}

func (stk *stackNodeStruct) TakeOutIfMatchesMin(node *stackNode){
	if stk.Head == nil{
		return
	}
	if stk.Head.Data == node.Data{
		stk.Pop()
	}

}
package main

import "fmt"

func (tree *TreeStruct) PreOrderTraversal1(){
	if tree.Root == nil{
		return
	}
	temp := tree.Root
	stk := StackTraverser{}
	for temp != nil{
		fmt.Println(temp.Data)
		stk.Push(temp)
		temp = temp.LeftChild
	}
	for !stk.IsEmpty(){
		popedData := stk.Pop()
		rightChild := popedData.Data.RightChild
		for rightChild!=nil{
			fmt.Println(rightChild.Data)
			stk.Push(rightChild)
			rightChild = rightChild.LeftChild
		}
	}



}

type Stack1 struct {
	Data *TreeNode1
	Next *Stack1
}

type StackTraverser1 struct {
	Head *Stack1
}

func (stk *StackTraverser1) Push(data *TreeNode1){
	if stk.Head == nil{
		stk.Head = &Stack1{Data:data, Next:nil}
		return
	}
	newNode := &Stack1{Data:data, Next:stk.Head}
	stk.Head = newNode
}

func (stk *StackTraverser1) Pop()*Stack1{
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

func (stk *StackTraverser1) IsEmpty()bool{
	return stk.Head == nil
}

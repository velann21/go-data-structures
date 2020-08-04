package main

import "fmt"

func main() {
	tree := TreeStruct{}
	tree.Insert(20)
	tree.Insert(10)
	tree.Insert(30)
	tree.Insert(8)
	tree.Insert(11)
	tree.Insert(31)
	tree.Insert(28)
	tree.PreOrderTraversal()
}

type TreeNode1 struct {

	Data int
	LeftChild *TreeNode1
	RightChild *TreeNode1
}

type TreeStruct struct {
	Root *TreeNode1
}

func (tree *TreeStruct) Insert(data int){
	if tree.Root == nil{
		tree.Root = &TreeNode1{Data:data, LeftChild:nil, RightChild:nil}
		return
	}
	temp := tree.Root
	parentNode := temp
	for temp != nil{
		parentNode = temp
	    if data < temp.Data{
	   	temp = temp.LeftChild
	    	continue
	    }
	    if data > temp.Data{
	   	 temp = temp.RightChild
	   	 continue
	    }
	}

	if data > parentNode.Data{
		parentNode.RightChild = &TreeNode1{Data:data, LeftChild:nil, RightChild:nil}
	} else{
		parentNode.LeftChild = &TreeNode1{Data:data, LeftChild:nil, RightChild:nil}
	}
}

func (tree *TreeStruct) PreOrderTraversal() {
	if tree.Root == nil {
		return
	}
	temp := tree.Root
	stk := StackTraverser{}
	for temp != nil {
		fmt.Println(temp.Data)
		stk.Push(temp)
		temp = temp.LeftChild
	}
	for !stk.IsEmpty() {
		popedData := stk.Pop()
		rightChild := popedData.Data.RightChild
		for rightChild != nil {
			fmt.Println(rightChild.Data)
			stk.Push(rightChild)
			rightChild = rightChild.LeftChild
		}
	}
}


func (tree *TreeStruct) PreOrderTraversalCleaner(){
	if tree.Root == nil{
		return
	}
	temp := tree.Root
	stk := StackTraverser{}
	for temp != nil || !stk.IsEmpty(){
		if temp != nil{
			fmt.Println(temp.Data)
			stk.Push(temp)
			temp = temp.LeftChild
		}else{
			popedData := stk.Pop()
			temp = popedData.Data.RightChild
		}
	}
}


type Stack struct {
	Data *TreeNode1
	Next *Stack
}

type StackTraverser struct {
	Head *Stack
}

func (stk *StackTraverser) Push(data *TreeNode1){
	if stk.Head == nil{
		stk.Head = &Stack{Data:data, Next:nil}
		return
	}
	newNode := &Stack{Data:data, Next:stk.Head}
	stk.Head = newNode
}

func (stk *StackTraverser) Pop()*Stack{
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

func (stk *StackTraverser) IsEmpty()bool{
	return stk.Head == nil
}



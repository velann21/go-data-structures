package main

import "fmt"

func main() {
	tree := TreeTraverser{}
	tree.Insert(20)
	tree.Insert(10)
	tree.Insert(30)
	tree.Insert(05)
	tree.Insert(11)
	tree.Insert(35)
	tree.Insert(25)
	tree.InorderWithoutRec(tree.Root)
}

type TreeNode3 struct {
	Data  int
	LeftChild *TreeNode3
	RightChild *TreeNode3
}

type TreeTraverser struct {
	Root *TreeNode3
}

func (tree *TreeTraverser) Insert(data int){
	if tree.Root == nil{
		tree.Root = &TreeNode3{Data:data, LeftChild:nil, RightChild:nil}
		return
	}
	temp := tree.Root
	lastParent := temp
	for temp != nil{
		lastParent = temp
	    if data < temp.Data{
	    	temp = temp.LeftChild
		}else if data > temp.Data{
			temp = temp.RightChild
		}
	}
	if data < lastParent.Data{
		lastParent.LeftChild = &TreeNode3{Data:data, LeftChild:nil, RightChild:nil}
	}else if data > lastParent.Data{
		lastParent.RightChild = &TreeNode3{Data:data, LeftChild:nil, RightChild:nil}
	}
}

func (tree *TreeTraverser) InorderWithoutRec(root *TreeNode3){
	if root == nil{
		return
	}
	temp := tree.Root
	stk := StackUtil{}
	for temp != nil || !stk.IsEmpty(){
		if temp != nil{
			stk.Push(temp)
			temp = temp.LeftChild
		}else if !stk.IsEmpty(){
			popedData := stk.Pop().Data
			fmt.Println(popedData.Data)
			temp = popedData.RightChild
		}
	}
}

type Stack3 struct {
	Data *TreeNode3
	Next *Stack3
}

type StackUtil struct {
	Head *Stack3
}

func (stk *StackUtil) Push(data *TreeNode3){
	if stk.Head == nil{
		stk.Head = &Stack3{Data:data, Next:nil}
		return
	}
	newNode := &Stack3{Data:data, Next:stk.Head}
	stk.Head = newNode
}

func (stk *StackUtil) Pop()*Stack3{
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

func (stk *StackUtil) IsEmpty()bool{
	return stk.Head == nil
}



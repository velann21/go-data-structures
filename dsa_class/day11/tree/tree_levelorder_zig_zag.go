package main

import "fmt"

func main() {
	tree := TreeStruct2{}
	tree.Insert(20)
	tree.Insert(15)
	tree.Insert(10)
	tree.Insert(25)
	tree.Insert(30)
	tree.ZigZagTraverser(tree.Root)

}

type TreeNode2 struct {
	Data int
	LeftChild *TreeNode2
	RightChild *TreeNode2
}

type TreeStruct2 struct {
	Root *TreeNode2
}

func (tree *TreeStruct2) Insert(data int){
	if tree.Root == nil{
		tree.Root = &TreeNode2{Data:data, LeftChild:nil, RightChild:nil}
		return
	}
	temp := tree.Root
	parentNode := temp
	for temp != nil{
		parentNode = temp
		if data < temp.Data{
			temp = temp.LeftChild
		}else if data > temp.Data{
			temp = temp.RightChild
		}
	}
	if data < parentNode.Data{
		parentNode.LeftChild = &TreeNode2{Data:data, LeftChild:nil, RightChild:nil}
	}else if data > parentNode.Data{
		parentNode.RightChild = &TreeNode2{Data:data, LeftChild:nil, RightChild:nil}
	}
}


func (tree *TreeStruct2) ZigZagTraverser(root *TreeNode2){
	if root == nil{
		return
	}
	stk1 := StackTraverser2{}
	stk2 := StackTraverser2{}
	stk1.Push(root)
	for !stk1.IsEmpty() || !stk2.IsEmpty(){
         for !stk1.IsEmpty(){
         	data := stk1.Pop()
         	if data.Data.RightChild != nil{
				fmt.Println(data.Data.RightChild.Data)
				stk2.Push(data.Data.RightChild)
			}
         	if data.Data.LeftChild != nil{
				fmt.Println(data.Data.LeftChild.Data)
				stk2.Push(data.Data.LeftChild)
			}
		 }
		for !stk2.IsEmpty(){
			data := stk2.Pop()
			if data.Data.LeftChild != nil{
				fmt.Println(data.Data.LeftChild.Data)
				stk1.Push(data.Data.LeftChild)
			}
			if data.Data.RightChild != nil{
				fmt.Println(data.Data.RightChild.Data)
				stk1.Push(data.Data.RightChild)
			}
		}
	}



}


type Stack2 struct {
	Data *TreeNode2
	Next *Stack2
}

type StackTraverser2 struct {
	Head *Stack2
}

func (stk *StackTraverser2) Push(data *TreeNode2){
	if stk.Head == nil{
		stk.Head = &Stack2{Data:data, Next:nil}
		return
	}
	newNode := &Stack2{Data:data, Next:stk.Head}
	stk.Head = newNode
}

func (stk *StackTraverser2) Pop()*Stack2{
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

func (stk *StackTraverser2) IsEmpty()bool{
	return stk.Head == nil
}






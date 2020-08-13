package main

import "fmt"

func main() {

	tree := TreeUtils{q:&QueueUtil1{}}
	tree.BinaryTree(10, -1)
	tree.BinaryTree(20, 20)
	tree.BinaryTree(30, 40)
	tree.BinaryTree(30, 40)
	tree.BinaryTree(30, 90)
	tree.BinaryTree(30, 90)
	tree.BinaryTree(30, 90)
	tree.BinaryTree(30, 90)
	isLeft := tree.IsLeftSameTree(tree.Root.LeftChild, tree.Root.RightChild)
	isRight := tree.IsRightSameTree(tree.Root.LeftChild, tree.Root.RightChild)
	if isLeft && isRight{
		fmt.Println("Yes same tree")
	}
}

type TreeNode6 struct {
	Data int
	LeftChild *TreeNode6
	RightChild *TreeNode6
}

type TreeUtils struct {
	Root *TreeNode6
	q *QueueUtil1
}

func (tree *TreeUtils) BinaryTree(leftData int, rightData int){
	if tree.Root == nil && leftData != -1{
		tree.Root = &TreeNode6{Data:leftData, LeftChild:nil, RightChild:nil}
		tree.q.Push(tree.Root)
		return
	}

	if !tree.q.IsEmpty(){
		popedData := tree.q.Pop()
		if popedData.LeftChild == nil && leftData != -1{
			popedData.LeftChild = &TreeNode6{Data:leftData, LeftChild:nil, RightChild:nil}
			tree.q.Push(popedData.LeftChild)
		}
		if popedData.RightChild == nil && rightData != -1{
			popedData.RightChild = &TreeNode6{Data:rightData, LeftChild:nil, RightChild:nil}
			tree.q.Push(popedData.RightChild)
		}
	}
}

func (tree *TreeUtils) IsLeftSameTree(left *TreeNode6, right *TreeNode6)bool{
	if left==nil && right==nil{
		return true
	}
	if left != nil && right != nil && left.Data == right.Data{
		return tree.IsLeftSameTree(left.LeftChild, right.LeftChild)
	}else{
		return false
	}
}

func (tree *TreeUtils) IsRightSameTree(left *TreeNode6, right *TreeNode6)bool{
	if left==nil && right==nil{
		return true
	}
	if left != nil && right != nil && left.Data == right.Data{
		return tree.IsRightSameTree(left.RightChild, right.RightChild)
	}else{
		return false
	}
}

type Queue1 struct {
	Data *TreeNode6
	Next *Queue1
}

type QueueUtil1 struct {
	Head *Queue1
}

func (q *QueueUtil1) Push(data *TreeNode6){
	if q.Head == nil{
		q.Head = &Queue1{Data:data, Next:nil}
		return
	}
	temp := q.Head
	for temp.Next != nil{
		temp = temp.Next
	}

	temp.Next = &Queue1{Data:data, Next:nil}
}

func (q *QueueUtil1) Pop()*TreeNode6{
	if q.Head == nil{
		return nil
	}
	temp := q.Head
	if q.Head.Next != nil{
		q.Head = q.Head.Next
	}else{
		q.Head = nil
	}
	return temp.Data
}


func (q *QueueUtil1) IsEmpty()bool{
	return q.Head == nil
}
package main

import "fmt"

func main() {
	tree := TreeTraverser5{q:&QueueUtil{}}
	tree.BSInsert(20, -1)
	tree.BSInsert(10, 10)
	tree.BSInsert(05, 15)
	tree.BSInsert(15, 05)
	isLeftMirror := tree.IsLeftMirror(tree.Root.LeftChild, tree.Root.RightChild)
	isRightMirror := tree.IsRightMirror(tree.Root.LeftChild, tree.Root.RightChild)
	if isLeftMirror && isRightMirror{
		fmt.Println("Yes tree is mirror")
	}

}

type TreeNode5 struct {
	Data int
	LeftChild *TreeNode5
	RightChild *TreeNode5
}

type TreeTraverser5 struct {
	Root *TreeNode5
	q *QueueUtil
}

func (tree *TreeTraverser5) BSTInsert(data int){
	if tree.Root == nil{
		tree.Root = &TreeNode5{Data:data, LeftChild:nil, RightChild:nil}
		return
	}

	temp := tree.Root
	parentNode := temp
	for temp!=nil{
		parentNode = temp
		if data<temp.Data{
			temp = temp.LeftChild
		}else if data > temp.Data{
			temp = temp.RightChild
		}
	}
	if data < parentNode.Data{
		parentNode.LeftChild = &TreeNode5{Data:data, LeftChild:nil, RightChild:nil}
	}else {
		parentNode.RightChild = &TreeNode5{Data:data, LeftChild:nil, RightChild:nil}
	}

}

func (tree *TreeTraverser5) BSInsert(leftData int, rightdata int){
	if tree.Root == nil{
		tree.Root = &TreeNode5{Data:leftData, LeftChild:nil, RightChild:nil}
        tree.q.Push(tree.Root)
		return
	}
	if !tree.q.IsEmpty(){
		poppedData := tree.q.Pop()
		if poppedData.LeftChild == nil && leftData != -1{
			poppedData.LeftChild = &TreeNode5{Data:leftData, LeftChild:nil, RightChild:nil}
			tree.q.Push(poppedData.LeftChild)
		}
		if poppedData.RightChild == nil && rightdata != -1{
			poppedData.RightChild = &TreeNode5{Data:rightdata, LeftChild:nil, RightChild:nil}
			tree.q.Push(poppedData.RightChild)
		}
	}
}

func (tree *TreeTraverser5) IsLeftMirror(left *TreeNode5, right *TreeNode5)bool{
	if left == nil && right == nil{
		return true
	}

	if (left != nil && right!= nil && left.Data == right.Data){
		return tree.IsLeftMirror(left.LeftChild, right.RightChild)
	}else{
		return false
	}
}

func (tree *TreeTraverser5) IsRightMirror(left *TreeNode5, right *TreeNode5)bool{
	if left == nil && right == nil{
		return true
	}

	if (left != nil && right!= nil && left.Data == right.Data){
		return tree.IsRightMirror(left.RightChild, right.LeftChild)

	}else{
		return false
	}
}


type Queue struct {
	Data *TreeNode5
	Next *Queue
}

type QueueUtil struct {
	Head *Queue
}

func (q *QueueUtil) Push(data *TreeNode5){
	if q.Head == nil{
		q.Head = &Queue{Data:data, Next:nil}
		return
	}
	temp := q.Head
	for temp.Next != nil{
		temp = temp.Next
	}

	temp.Next = &Queue{Data:data, Next:nil}
}

func (q *QueueUtil) Pop()*TreeNode5{
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

func (q *QueueUtil) IsEmpty()bool{
	return q.Head == nil
}
package main

import (
	"errors"
	"fmt"
)

func main() {
	tree := TreeNodeStruct5{}
	tree.Insert(8, -1)
	tree.Insert(3, 5)
	tree.Insert(4, 9)
	tree.Insert(7, 2)
	tree.PreOrderTraversal(tree.Root)

}

type TreeNode5 struct {
	Rchild *TreeNode5
	Lchild *TreeNode5
	Data int
}

type TreeNodeStruct5 struct {
	Root *TreeNode5
	Queue Queue3
}

func (node *TreeNodeStruct5) Insert(ldata int, rdata int){
	if node.Root == nil{
		node.Root = &TreeNode5{Data:ldata, Lchild:nil, Rchild:nil}
		node.Queue.Push(node.Root)
		return
	}

	if !node.Queue.IsEmpty(){
		data := node.Queue.Pop()
		if data.Lchild == nil || ldata != -1{
			data.Lchild = &TreeNode5{Data:ldata, Lchild:nil, Rchild:nil}
			node.Queue.Push(data.Lchild)
		}

		if data.Rchild == nil || rdata != -1{
			data.Rchild = &TreeNode5{Data:rdata, Lchild:nil, Rchild:nil}
			node.Queue.Push(data.Rchild)
		}
	}
}


func (node *TreeNodeStruct5) PreOrderTraversal(root *TreeNode5){
	if node.Root == nil || root == nil || root != node.Root{
		return
	}
	q := Stack{}
	for root != nil || !q.IsEmpty(){
		if root != nil{
			fmt.Print(root.Data, " ")
			q.Push(root)
			root = root.Lchild
		}else{
			rSide := q.Pop()
			root = rSide.Rchild
		}
	}
}



type Queue3 struct {
	que []*TreeNode5
}

func (q *Queue3) Push(data *TreeNode5) {
	q.que = append(q.que, data)
}

func (q *Queue3) Pop() *TreeNode5 {
	data := q.que[0]
	q.que = q.que[1:len(q.que)]
	return data
}

func (q *Queue3) IsEmpty() bool {
	return len(q.que) == 0
}

func (q *Queue3) Peek(data int) {

}


type Stack struct {
	array []*TreeNode5
}

func (stk *Stack) Push(data *TreeNode5){
	stk.array = append(stk.array, data)
}

func (stk *Stack) Pop()*TreeNode5{
	data := stk.array[len(stk.array)-1]
	stk.array = stk.array[0:len(stk.array)-1]
	return data
}

func (stk *Stack) Peek(pos int)error{
	if pos > len(stk.array){
		return errors.New("Invalid Request")
	}

	return nil
}

func (stk *Stack) IsEmpty()bool{
	return len(stk.array) == 0
}
package main

import (
	"errors"
	"fmt"
)

func main() {
	tree := NodeStruct3{}
	tree.Insert(8, -1)
	tree.Insert(3, 5)
	tree.Insert(4, 9)
	tree.Insert(7, 2)

    _ = tree.PreorderTraversal(tree.Root)

}

type Node3 struct {
	Lchild *Node3
	Rchild *Node3
	Data int
}

type NodeStruct3 struct {
	Root *Node3
	Queue Queue1
}

func (node *NodeStruct3) Insert(ldata int, rData int){
	if node.Root == nil{
		node.Root = &Node3{Data:ldata, Lchild:nil, Rchild:nil}
		node.Queue.Push(node.Root)
		return
	}
	d := node.Queue.Pop()
	if d.Lchild == nil && ldata != -1{
		d.Lchild = &Node3{Data:ldata, Lchild:nil, Rchild:nil}
		node.Queue.Push(d.Lchild)
	}

	if d.Rchild == nil && rData != -1{
		d.Rchild = &Node3{Data:rData, Lchild:nil, Rchild:nil}
		node.Queue.Push(d.Rchild)
	}
}

func (node *NodeStruct3) PreorderTraversal(rootNode *Node3)error{
	if node.Root == nil{
		return errors.New("Invalid request")
	}
	if rootNode != nil{
		fmt.Print(rootNode.Data, " ")
		_ = node.PreorderTraversal(rootNode.Lchild)
		_ = node.PreorderTraversal(rootNode.Rchild)
	}
	return nil
}



type Queue1 struct {
	que []*Node3
}

func (q *Queue1) Push(data *Node3) {
	q.que = append(q.que, data)
}

func (q *Queue1) Pop() *Node3 {
	data := q.que[0]
	q.que = q.que[1:len(q.que)]
	return data
}

func (q *Queue1) IsEmpty() bool {
	return len(q.que) == 0
}

func (q *Queue1) Peek(data int) {

}

package main

import "fmt"

func main() {
	tree := TreeNodeStruct4{}
	tree.Insert(8, -1)
	tree.Insert(3, 5)
	tree.Insert(4, 9)
	tree.Insert(7, 2)

	tree.InOrderTraversal(tree.Root)
}

type TreeNode4 struct {
	Rchild *TreeNode4
	Lchild *TreeNode4
	Data int
}

type TreeNodeStruct4 struct {
	Root *TreeNode4
	Queue Queue2
}

func (node *TreeNodeStruct4) Insert(ldata int, rdata int){
	if node.Root == nil{
		node.Root = &TreeNode4{Data:ldata, Lchild:nil, Rchild:nil}
		node.Queue.Push(node.Root)
		return
	}

	if !node.Queue.IsEmpty(){
		data := node.Queue.Pop()
		if data.Lchild == nil || ldata != -1{
			data.Lchild = &TreeNode4{Data:ldata, Lchild:nil, Rchild:nil}
			node.Queue.Push(data.Lchild)
		}

		if data.Rchild == nil || rdata != -1{
			data.Rchild = &TreeNode4{Data:rdata, Lchild:nil, Rchild:nil}
			node.Queue.Push(data.Rchild)
		}
	}
}

func (node *TreeNodeStruct4) InOrderTraversal(root *TreeNode4){
	if node.Root == nil{
		return
	}
	if root != nil{
		node.InOrderTraversal(root.Lchild)
		fmt.Print(root.Data, " ")
		node.InOrderTraversal(root.Rchild)
	}
}


type Queue2 struct {
	que []*TreeNode4
}

func (q *Queue2) Push(data *TreeNode4) {
	q.que = append(q.que, data)
}

func (q *Queue2) Pop() *TreeNode4 {
	data := q.que[0]
	q.que = q.que[1:len(q.que)]
	return data
}

func (q *Queue2) IsEmpty() bool {
	return len(q.que) == 0
}

func (q *Queue2) Peek(data int) {

}

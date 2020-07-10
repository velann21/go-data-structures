package main

import "fmt"

func main() {

}

type TreeNode16 struct {
	Lchild *TreeNode16
	Rchild *TreeNode16
	Data int
}

type TreeNodeStruct struct {
	Root *TreeNode16
	Q Queue4

}

func (node *TreeNodeStruct) Insert(ldata int, rdata int){
	if node.Root == nil{
		node.Root = &TreeNode16{Lchild:nil,Rchild:nil, Data:ldata}
        node.Q.Push(node.Root)
		return
	}

	data := node.Q.Pop()
	if data != nil{
		if data.Lchild == nil{
			data.Lchild = &TreeNode16{Lchild:nil,Rchild:nil, Data:ldata}
			node.Q.Push(node.Root)
		}
		if data.Rchild == nil {
			data.Rchild = &TreeNode16{Lchild:nil,Rchild:nil, Data:rdata}
			node.Q.Push(data.Rchild)
		}
	}
}

func (node *TreeNodeStruct) LevelOrder(){
	if node.Root == nil{
		return
	}

	q := Queue4{}
	q.Push(node.Root)

	for !q.IsEmpty(){
		data := q.Pop()
		if data.Lchild != nil{
			fmt.Print(data.Lchild.Data)
			q.Push(data.Lchild)
		}
		if data.Rchild != nil{
			fmt.Print(data.Rchild.Data)
			q.Push(data.Rchild)
		}
	}
}


type Queue4 struct {
	que []*TreeNode16
}

func (q *Queue4) Push(data *TreeNode16) {
	q.que = append(q.que, data)
}

func (q *Queue4) Pop() *TreeNode16 {
	data := q.que[0]
	q.que = q.que[1:len(q.que)]
	return data
}

func (q *Queue4) IsEmpty() bool {
	return len(q.que) == 0
}

func (q *Queue4) Peek(data int) {

}


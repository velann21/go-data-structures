package main

import "fmt"

func main() {

	tree := TreeNodeImpl{}
	tree.Insert(8, -1)
	tree.Insert(3, 5)
	tree.Insert(4, 9)
	tree.Insert(7, 2)
}

type TreeNode struct {
	Data      int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

type TreeNodeImpl struct {
	Root *TreeNode
	q    Queue
}

func (t *TreeNodeImpl) Insert(leftdata int, rightdata int) {
	if t.Root == nil {
		t.Root = &TreeNode{Data: leftdata, LeftNode: nil, RightNode: nil}
		t.q.Push(t.Root)
		return
	}
	p := t.q.Pop()
	fmt.Println(p.Data)
	if p.LeftNode == nil && leftdata != -1 {
		newNode := &TreeNode{Data: leftdata, LeftNode: nil, RightNode: nil}
		p.LeftNode = newNode
		t.q.Push(p.LeftNode)
	}
	if p.RightNode == nil && rightdata != -1 {
		newNode := &TreeNode{Data: rightdata, LeftNode: nil, RightNode: nil}
		p.RightNode = newNode
		t.q.Push(p.RightNode)
	}
}

type Queue struct {
	que []*TreeNode
}

func (q *Queue) Push(data *TreeNode) {
	q.que = append(q.que, data)
}

func (q *Queue) Pop() *TreeNode {
	data := q.que[0]
	q.que = q.que[1:len(q.que)]
	return data
}

func (q *Queue) IsEmpty() bool {
	return len(q.que) == 0
}

func (q *Queue) Peek(data int) {

}

package main

func main() {

}

type TreeNode struct {
	Data int
	LeftNode *TreeNode
	RightNode *TreeNode
}

type TreeNodeImpl struct {
	Root *TreeNode
	q Queue
}

func (t *TreeNodeImpl) Insert(data int){
	if t.Root == nil{
		t.Root = &TreeNode{Data:data, LeftNode:nil, RightNode:nil}
		t.q.Push(t.Root)
		return
	}
	node := t.q.Pop()
	if node.LeftNode == nil{
		node.LeftNode = &TreeNode{Data:data, LeftNode:nil, RightNode:nil}
		t.q.Push(node)
		t.q.Push(node.LeftNode)
	}else {
		node.RightNode = &TreeNode{Data:data, LeftNode:nil, RightNode:nil}
		t.q.Push(node.RightNode)
	}

}


type Queue struct {
	que []*TreeNode
}

func (q *Queue) Push(data *TreeNode){
	q.que = append(q.que, data)
}

func (q *Queue) Pop()(*TreeNode){
	data := q.que[0]
	q.que = q.que[1:len(q.que)-1]
	return data
}

func (q *Queue) IsEmpty()bool{
	return len(q.que) == 0
}

func (q *Queue) Peek(data int){

}

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
}

func (t *TreeNodeImpl) Insert(data int){
	if t.Root == nil{

	}
}


type Queue struct {
	que []int
}

func (q *Queue) Push(data int){
	q.que = append(q.que, data)
}

func (q *Queue) Pop()(int){
	data := q.que[0]
	q.que = q.que[1:len(q.que)-1]
	return data
}

func (q *Queue) IsEmpty()bool{
	return len(q.que) == 0
}

func (q *Queue) Peek(data int){

}

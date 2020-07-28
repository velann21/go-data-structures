package main

func main() {

	q := QueueStruct{}
	treeStruct := TreeStruct{q:&q}
	treeStruct.InsertNode(8, -1)
	treeStruct.InsertNode(10, 12)
	treeStruct.InsertNode(13, 15)
	treeStruct.InsertNode(16, 17)
}

type TreeNode struct {
    Data int
    LeftChild *TreeNode
    RightChild *TreeNode
}

type TreeStruct struct {
	Root *TreeNode
	q *QueueStruct
}

func (tree *TreeStruct) InsertNode(leftData int, RightData int){
	if tree.Root == nil{
		tree.Root = &TreeNode{Data:leftData, LeftChild:nil, RightChild:nil}
        tree.q.Enqueue(tree.Root)
		return
	}

	data := tree.q.Dequeue()
	if data != nil{
		if data.Data.LeftChild == nil && leftData != -1{
			data.Data.LeftChild = &TreeNode{Data:leftData, LeftChild:nil, RightChild:nil}
			tree.q.Enqueue(data.Data.LeftChild)
		}else if data.Data.RightChild == nil && RightData != -1{
			data.Data.RightChild = &TreeNode{Data:RightData, LeftChild:nil, RightChild:nil}
			tree.q.Enqueue(data.Data.RightChild)
		}
	}
}

type Queue struct {
	Data *TreeNode
	Next *Queue
}

type QueueStruct struct {
	Head *Queue
	Tail *Queue
}

func (q *QueueStruct) Enqueue(data *TreeNode){
	if q.Head == nil{
		q.Head = &Queue{Data:data, Next:nil}
		q.Tail = q.Head
		return
	}
	if q.Tail != nil{
		newNode := &Queue{Data:data, Next:nil}
		q.Tail = newNode
		return
	}
}

func (q *QueueStruct) Dequeue()*Queue{
	if q.Head == nil{
		return nil
	}
	temp := q.Head
	if q.Head.Next != nil{
		q.Head = q.Head.Next
	}else{
		q.Head = nil
	}
	return temp
}



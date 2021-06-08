package tree

import (
	"errors"
	"sync"
)

type Node struct {
	LeftNode *Node
	RightNode *Node
	Data int
}

type QueueNode struct {
	Data *Node
	Next *QueueNode
}

type Queue struct {
	Front *QueueNode
	Rear *QueueNode
	Size int
}

func (q *Queue) Enqueue(data *Node){
	if q.Front == nil && q.Rear == nil{
		newNode := &QueueNode{Data: data, Next: nil}
		q.Front = newNode
		q.Rear = newNode
		q.Size++
		return
	}
	newNode := &QueueNode{Data: data, Next: nil}
	q.Rear.Next = newNode
	q.Rear = newNode
	q.Size++
	return
}

func (q *Queue) Dequeue()*Node{
	if q.Front == nil{
		return nil
	}
	if q.Front.Next == nil{
		data := q.Front
		q.Front = nil
		q.Size--
		return data.Data
	}
	data := q.Front
	q.Front = q.Front.Next
	q.Size--
	return data.Data
}

func (q *Queue) IsEmpty()bool{
	return q.Front == nil
}

type Tree struct {
	Root *Node
	lock sync.Mutex
	q Queue
}

func (tree *Tree) Insert(left int, right int)error{
	tree.lock.Lock()
	defer tree.lock.Unlock()

	if tree.Root == nil {
		tree.Root = &Node{LeftNode: nil, Data: left, RightNode: nil}
		tree.q.Enqueue(tree.Root)
		return nil
	}

	if tree.q.IsEmpty(){
		return errors.New("SomethingWrong")
	}
	cNode := tree.q.Dequeue()
	if cNode.LeftNode == nil && left != -1{
		cNode.LeftNode = &Node{LeftNode: nil, Data: left, RightNode: nil}
		tree.q.Enqueue(cNode.LeftNode)
	}

	if cNode.RightNode == nil && right != -1{
		cNode.RightNode = &Node{LeftNode: nil, Data: right, RightNode: nil}
		tree.q.Enqueue(cNode.RightNode)
	}
	return nil
}

package tree

import (
	"errors"
	"fmt"
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
		q.Rear = nil
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
	lock *sync.Mutex
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

func (tree *Tree) PreOrder(node *Node){
	if node == nil{
		return
	}
	fmt.Println(node.Data)
	tree.PreOrder(node.LeftNode)
	tree.PreOrder(node.RightNode)
}


func (tree *Tree) CountNodes(node *Node)int{
	if node == nil{
		return 0
	}
	x := tree.CountNodes(node.LeftNode)
	y := tree.CountNodes(node.RightNode)
	sum :=x+y+1
	return sum
}
func (tree *Tree) SumValues(node *Node)int{
	if node == nil{
		return 0
	}
	x := tree.SumValues(node.LeftNode)
	y := tree.SumValues(node.RightNode)
	sum :=x+y+node.Data
	return sum
}

func (tree *Tree) CountLeafNodes(node *Node)int{
	if node == nil{
		return 0
	}
	x := tree.CountLeafNodes(node.LeftNode)
	y := tree.CountLeafNodes(node.RightNode)
	if node.LeftNode == nil && node.RightNode == nil{
		return x+y+1
	}
	return x+y
}

func (tree *Tree) CountTwoChildedNodes(node *Node)int{
	if node == nil{
		return 0
	}
	x := tree.CountLeafNodes(node.LeftNode)
	y := tree.CountLeafNodes(node.RightNode)
	if node.LeftNode != nil && node.RightNode != nil{
		return x+y+1
	}
	return x+y
}

var sumArray []int
func (tree *Tree) BranchSum(node *Node, brnchSum int)[]int{
	if node == nil{
		return nil
	}
	brnchSum += node.Data
	tree.BranchSum(node.LeftNode, brnchSum)
	tree.BranchSum(node.RightNode, brnchSum)
	fmt.Println("Branhsum", brnchSum)
	sumArray = append(sumArray, brnchSum)
	return sumArray
}
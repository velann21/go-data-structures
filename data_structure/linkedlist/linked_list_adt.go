package linkedlist

import (
	"errors"
	"fmt"
)


var ERRNILHEAD = errors.New("")
type List interface {
	Insert(data int)
	GetHead()*Node
	GetTail()*Node
	GetSize()int
	ReturnNodesData(Head *Node)([]int, error)
	DisplayNodesData(Head *Node)error
	Sum(Head *Node)(int,error)
	Set(data int,  index int)error
	RemoveAt(index int)error
	ReverseLL()error
}

func NewLinkedList()List{
	return &LinkedList{}
}

type Node struct {
	Next *Node
	Data int
}

func NewNode(data int, next *Node)*Node{
	return &Node{Next: next, Data: data}
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (ll *LinkedList) GetHead()*Node{
	return ll.Head
}

func (ll *LinkedList) GetTail()*Node{
	return ll.Tail
}

func (ll *LinkedList) GetSize()int{
	return ll.Size
}

func (ll *LinkedList) Insert(data int){
	if ll.Head == nil{
		newNode := NewNode(data, nil)
		ll.Head = newNode
		ll.Tail = newNode
		ll.Size++
		return
	}
	last := ll.Head
	for last.Next != nil{
		last = last.Next
	}
	newNode := NewNode(data, nil)
	last.Next = newNode
	ll.Tail = newNode
	ll.Size++
}

func (ll *LinkedList) ReturnNodesData(Head *Node)([]int, error){
	if Head == nil{
		return nil, ERRNILHEAD
	}
	auxiliaryArray := make([]int, 0)
	if Head.Next == nil{
		auxiliaryArray = append(auxiliaryArray, Head.Data)
		return auxiliaryArray, nil
	}
	start := Head
	for start.Next != nil {
		auxiliaryArray = append(auxiliaryArray, start.Data)
		start = start.Next
	}
	auxiliaryArray = append(auxiliaryArray, start.Data)
	return auxiliaryArray, nil
}

func (ll *LinkedList) DisplayNodesData(Head *Node)error{
	if Head == nil{
		return ERRNILHEAD
	}
	start := Head
	for start.Next != nil{
		fmt.Println(start.Data)
		start = start.Next
	}
	fmt.Println(start.Data)
	return nil
}

var InvalidReq = errors.New("InvalidRequest")
func (ll *LinkedList) Sum(Head *Node)(int,error){
	if Head == nil{
		return 0, InvalidReq
	}
	sum := 0
	for Head != nil{
		sum += Head.Data
		Head = Head.Next
	}
	return sum, nil
}


func (ll *LinkedList) Set(data int,  index int)error{
	if ll.Head == nil && index > 1{
		return InvalidReq
	}
	if ll.Head == nil && index == 1{
		ll.Head = NewNode(data, nil)
		return nil
	}
	prevnode := ll.Head
	startIndex := 1
	for startIndex < index{
		startIndex ++
		prevnode = prevnode.Next
	}
	newNode := NewNode(data, prevnode.Next)
	prevnode.Next = newNode
	ll.Size++
	return nil
}

func (ll *LinkedList) RemoveAt(index int)error{
	if ll.Head == nil{
		return InvalidReq
	}
	if ll.Head.Next == nil && index>1{
		return InvalidReq
	}
	if index == 1 {
		ll.Head = nil
		return nil
	}
	startIndex := 1
	preNode := ll.Head
	for startIndex < index-1{
		startIndex++
		preNode = preNode.Next
	}
	temp := preNode.Next
	preNode.Next = temp.Next
	return nil
}

func (ll *LinkedList) ReverseLL()error{
	if ll.Head == nil{
		return InvalidReq
	}
	if ll.Head.Next == nil{
		return nil
	}
	var firstFollower *Node
	var secondFollower *Node
	var targetPtr = ll.Head

	for targetPtr != nil{
		firstFollower = secondFollower
		secondFollower = targetPtr
		targetPtr = targetPtr.Next
		secondFollower.Next = firstFollower
	}
	ll.Head = secondFollower
	return nil
}
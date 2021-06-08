package linkedlist

import (
	"errors"
	"fmt"
	"reflect"
)


var ERRNILHEAD = errors.New("")
type List interface {
	Insert(data int)*Node
	GetHead()*Node
	GetTail()*Node
	GetSize()int
	ReturnNodesData(Head *Node)([]int, error)
	DisplayNodesData(Head *Node)error
	Sum(Head *Node)(int,error)
	Set(data int,  index int)error
	RemoveAt(index int)error
	ReverseLL()error
	DetectLoop()(bool, error)
	SlidingWindow(main string, substr string)(bool, error)
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

func (ll *LinkedList) Insert(data int)*Node{
	if ll.Head == nil{
		newNode := NewNode(data, nil)
		ll.Head = newNode
		ll.Tail = newNode
		ll.Size++
		return newNode
	}
	last := ll.Head
	for last.Next != nil{
		last = last.Next
	}
	newNode := NewNode(data, nil)
	last.Next = newNode
	ll.Tail = newNode
	ll.Size++
	return newNode
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
	for start != nil{
		fmt.Println(start.Data)
		start = start.Next
	}
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

func (ll *LinkedList) DetectLoop()(bool, error){
	if ll.Head == nil{
		return false, InvalidReq
	}
	if ll.Head.Next == nil{
		return false, nil
	}
	slowPtr := ll.Head
	fastPtr := ll.Head
	for slowPtr != nil && fastPtr.Next.Next != nil{
		fastPtr = fastPtr.Next.Next
		slowPtr = slowPtr.Next
		if  fastPtr==slowPtr{
			return true, nil
		}
	}
	return fastPtr==slowPtr, nil
}

func (ll *LinkedList) SlidingWindow(main string, substr string)(bool, error){
	if len(main) <= 0 || len(substr) <= 0 {
		return false, InvalidReq
	}

	if  main == substr{
		return true, nil
	}
	windowSize := len(substr)
	for slider:=0; slider<=len(main)-windowSize; slider++{
		//temp:= 0
		firstStr := ""
		secondStr := ""
		for slideCounter := slider; slideCounter<slider+windowSize; slideCounter++{
			firstStr += string(main[slideCounter])
			secondStr += string(substr[slideCounter-slider])
		}
		fmt.Println(firstStr)
		fmt.Println(secondStr)
		if firstStr == secondStr{
			return true, nil
		}
	}
	return false, nil
}

type DDNode struct{
	prev *DDNode
	data int
	next *DDNode
}

func NewDDNode(prev *DDNode, data int, next *DDNode) *DDNode {
	return &DDNode{prev: prev, data: data, next: next}
}

var InvalidDataType = errors.New("InvalidDataType")
var NoData = errors.New("NoData")

type DoublyLinkedList struct{
	Head *DDNode
	Tail *DDNode
}

func NewDoublyLinkedList()*DoublyLinkedList{
	return &DoublyLinkedList{}
}

func (linkedList *DoublyLinkedList) InsertDDL(data int)error{
	if reflect.TypeOf(data).Kind() == reflect.Int{
		return errors.New("Invalid data")
	}
	if linkedList.Head == nil{
		newNode := NewDDNode(nil, data, nil)
		linkedList.Head = newNode
		linkedList.Tail = newNode
		return nil
	}
	last := linkedList.Head
	for last.next != nil{
		last = last.next
	}
	newNode := NewDDNode(last, data, last.next)
	last.next = newNode
	linkedList.Tail = newNode
	return nil
}


func (linkedList *DoublyLinkedList) Display()error{
	if linkedList.Head == nil{
		return NoData
	}
	last := linkedList.Head
	for last != nil{
		fmt.Println(last.data)
		last = last.next
	}
	return nil
}

func (linkedList *DoublyLinkedList) Set(index, data int)error{
	if linkedList.Head == nil && index>1{
		return InvalidDataType
	}
	if linkedList.Head == nil && index == 1{
		newNode := NewDDNode(nil, data, nil)
		linkedList.Head = newNode
		linkedList.Tail = newNode
		return nil
	}
	if index == 1 {
		newNode := NewDDNode(nil, data, linkedList.Head)
		linkedList.Head.prev = newNode
		linkedList.Head = newNode
		return nil
	}


	last := linkedList.Head
	startIndex := 1
	for startIndex < index-1{
		startIndex++
		last = last.next
	}
	newNode := NewDDNode(nil, data, nil)
	if last.next != nil{
		last.next.prev = newNode
	}
	last.next = newNode
	newNode.prev = last


	return nil
}

type Stack struct{
	Head *Node
	Tail *Node
}

func (stk *Stack) Push(data int){

	if stk.Head == nil{
		newNode := NewNode(data, nil)
		stk.Head = newNode
		stk.Tail = newNode
		return
	}
	newNode := NewNode(data, stk.Head)
	stk.Head = newNode

}

func (stk *Stack) IsEmpty()bool{
	return stk.Head==nil
}

func (stk *Stack) Pop() (int, error) {
	if stk.IsEmpty(){
		return 0, NoData
	}
	data := stk.Head.Data
	if stk.Head.Next != nil{
		stk.Head = stk.Head.Next
	}

	return data, nil
}

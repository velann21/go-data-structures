package linkedlist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_InsertFirst(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.Insert(10)
	if linkedList.GetHead() == nil{
		t.Error("Head should not be nil")
	}

	if linkedList.GetTail() == nil{
		t.Error("Tail should not be nil")
	}

	if linkedList.GetSize() <=0{
		t.Error("Size should not be less than or equal to 0")
	}
}

func TestLinkedList_InsertNothing(t *testing.T) {
	linkedList := NewLinkedList()
	if linkedList.GetHead() != nil{
		t.Error("Head should be nil")
	}

	if linkedList.GetTail() != nil{
		t.Error("Tail should be nil")
	}

	if linkedList.GetSize() > 0{
		t.Error("Size should be less than or equal to 0")
	}
}

func TestLinkedList_InsertTwoElements(t *testing.T){
	linkedList := NewLinkedList()
	linkedList.Insert(10)
	linkedList.Insert(20)
	if linkedList.GetHead() == nil{
		t.Error("Head should not be nil")
	}

	if linkedList.GetTail() == nil{
		t.Error("Tail should not be nil")
	}

	if linkedList.GetSize() <= 0{
		t.Error("Size should not be less than or equal to 0")
	}

	if linkedList.GetHead().Next == nil{
		t.Error("Heads Next should not be nil")
	}

	if linkedList.GetTail().Next != nil{
		t.Error("Tails Next should be nil")
	}
}

func TestLinkedList_InsertThreeElements(t *testing.T){
	linkedList := NewLinkedList()
	linkedList.Insert(10)
	linkedList.Insert(20)
	linkedList.Insert(30)
	if linkedList.GetHead() == nil{
		t.Error("Head should not be nil")
	}

	if linkedList.GetTail() == nil{
		t.Error("Tail should not be nil")
	}

	if linkedList.GetSize() <= 0{
		t.Error("Size should not be less than or equal to 0")
	}

	if linkedList.GetHead().Next == nil{
		t.Error("Heads Next should not be nil")
	}

	if linkedList.GetTail().Next != nil{
		t.Error("Tails Next should be nil")
	}

	if linkedList.GetHead().Next.Next == nil{
		t.Error("Heads Next Next should not be nil")
	}
}

func TestLinkedList_ReturnNodesData_WithoutInsert(t *testing.T) {
	tests := []struct{
		Head *Node
		Result []int
		err    error
	}{
		{Head: nil, Result: nil, err: ERRNILHEAD},
		{Head: NewNode(10, nil), Result: []int{10}, err: nil},
		{Head: insert(), Result: []int{10, 20, 30}, err: nil},
	}

	ll := NewLinkedList()
	for _, test := range tests{
		result, err := ll.ReturnNodesData(test.Head)
		assert.IsType(t,  test.err, err)
		assert.Equal(t, test.Result, result)
	}
}

func TestLinkedList_DisplayNodesData(t *testing.T) {
	tests := []struct{
		Head *Node
		err    error
	}{
		{Head: nil, err: ERRNILHEAD},
		{Head: NewNode(10, nil),  err: nil},
		{Head: insert(), err: nil},
	}

	ll := NewLinkedList()
	for _, test := range tests{
		err := ll.DisplayNodesData(test.Head)
		assert.IsType(t,  test.err, err)
	}
}

func TestLinkedList_Sum(t *testing.T) {
	tests := []struct{
		input *Node
		expectedResult int
		err error
	}{{input: insert(), expectedResult: 100, err: nil}}
	ll := NewLinkedList()
	for _, test := range tests{
		output, err := ll.Sum(test.input)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expectedResult, output)
	}
}



func TestLinkedList_Set(t *testing.T) {
	ll := NewLinkedList()
	ll.Insert(10)
	ll.Insert(20)
	ll.Insert(30)
	ll.Insert(40)

	linkedListTwo := NewLinkedList()
	linkedListTwo.Insert(10)
	linkedListTwo.Insert(20)
	linkedListTwo.Insert(30)
	linkedListTwo.Insert(40)

	tests := []struct{
		object List
		index int
		data int
		expectedResult []int
		err error
	}{
		{object: ll, index:3, data: 35, expectedResult:[]int{10, 20, 30,35, 40},  err: nil},
		{object: linkedListTwo, index:1, data: 37,expectedResult:[]int{10,37, 20, 30, 40},  err: nil},
		{object: NewLinkedList(), index:1, data: 37,expectedResult:[]int{37},  err: nil},
		{object: NewLinkedList(), index:4, data: 37,expectedResult:nil,  err: InvalidReq},
	}

	for _, test := range tests{
		err := test.object.Set(test.data,test.index)
		assert.IsType(t, test.err, err)
		data, err := test.object.ReturnNodesData(test.object.GetHead())
		assert.Equal(t, test.expectedResult, data)
	}
}

func TestLinkedList_RemoveAt(t *testing.T) {
	ll := NewLinkedList()
	ll.Insert(10)
	ll.Insert(20)
	ll.Insert(30)
	ll.Insert(40)

	tests := []struct{
		object List
		index int
		expectedResult []int
		err error
	}{
		{object: ll, index: 3, expectedResult: []int{10,20,40}, err: nil},
		{object: ll, index: 2, expectedResult: []int{10,40}, err: nil},
		{object: ll, index: 2, expectedResult: []int{10}, err: nil},
		{object: ll, index: 2, expectedResult: []int{10}, err: InvalidReq},
		{object: ll, index: 1, expectedResult: nil, err: nil},
		{object: NewLinkedList(), index: 2, expectedResult: nil, err: InvalidReq},
	}

	for _, test := range tests{
		err := test.object.RemoveAt(test.index)
		assert.IsType(t, test.err, err)
		data, err := test.object.ReturnNodesData(test.object.GetHead())
		assert.Equal(t, test.expectedResult, data)
	}
}

func TestLinkedList_ReverseLL(t *testing.T) {
	ll := NewLinkedList()
	ll.Insert(10)
	ll.Insert(20)
	ll.Insert(30)
	ll.Insert(40)

	linkedListTwo := NewLinkedList()
	linkedListTwo.Insert(10)
	linkedListThree := NewLinkedList()

	tests := []struct{
		object List
		expectedResult []int
		err error
	}{
		{object: ll,expectedResult: []int{40,30,20,10},err: nil},
		{object: linkedListTwo,expectedResult: []int{10},err: nil},
		{object: linkedListThree,expectedResult: nil,err: InvalidReq},
	}

	for _, test := range tests{
		err := test.object.ReverseLL()
		assert.IsType(t, test.err, err)
		data, err := test.object.ReturnNodesData(test.object.GetHead())
		assert.Equal(t, test.expectedResult, data)
	}
}

func TestLinkedList_DetectLoop(t *testing.T) {
	ll := NewLinkedList()
	ll.Insert(10)
	ll.Insert(20)
	ll.Insert(30)
	ll.Insert(40)

	linkedList2 := NewLinkedList()
	linkedList2.Insert(10)
	node2 := linkedList2.Insert(20)
	linkedList2.Insert(30)
	node5 := linkedList2.Insert(40)
	node5.Next = node2

	tests := []struct{
		object List
		expectedResult bool
		err error
	}{
		{object: ll, expectedResult: false, err: nil},
		{object: linkedList2, expectedResult: true, err: nil},
	}

	for _, test := range tests{
		isLoopDetected, err := test.object.DetectLoop()
		assert.Equal(t, test.expectedResult, isLoopDetected)
		assert.IsType(t, test.err, err)
	}
}

func TestLinkedList_SlidingWindow(t *testing.T) {
	tests := []struct{
		mainStr string
		subStr string
		expectedResult bool
		err error
	}{
		{mainStr: "", subStr: "", expectedResult: false, err: InvalidReq},
		{mainStr: "", subStr: "win", expectedResult: false, err: InvalidReq},
		{mainStr: "win", subStr: "", expectedResult: false, err: InvalidReq},
		{mainStr: "win", subStr: "win", expectedResult: true, err: nil},
		{mainStr: "dodwin", subStr: "win", expectedResult: true, err: nil},
		{mainStr: "window", subStr: "win", expectedResult: true, err: nil},

	}
	linkedList := NewLinkedList()
	for _, test := range tests {
		linkedList.SlidingWindow(test.mainStr, test.subStr)
	}
}


func TestDoublyLinkedList_InsertDDL(t *testing.T) {
	tests := []struct{
		input int
		index int
		err error
	}{
		{input: 10,index: 1, err: nil},
		{input: 20,index: 2, err: nil},
		{input: 30,index: 3, err: nil},
		{input: 40,index: 4, err: nil},
		{input: 40,index: 5, err: nil},
		{input: 50,index: 1, err: nil},
	}
	ddl := NewDoublyLinkedList()
	for _, test := range tests{
		err := ddl.Set(test.index, test.input)
		assert.IsType(t, err, test.err)
	}
	fmt.Println("ssdsd")
}

func insert()*Node{
	ll := NewLinkedList()
	ll.Insert(10)
	ll.Insert(20)
	ll.Insert(30)
	ll.Insert(40)
	return ll.GetHead()
}


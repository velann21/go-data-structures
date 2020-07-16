package main

import "fmt"

func main() {

	node := ListNodeStruct{}
	node.Insert(10)
	node.Insert(20)
	//node.Insert(30)
	//node.Insert(40)
	//node.Insert(50)
	fmt.Println(node.removeNthFromEnd(node.Head, 2))
}

type ListNode struct {
	Val int
	Next *ListNode
}

type ListNodeStruct struct {
	Head *ListNode
}

func (node *ListNodeStruct) Insert(data int){
	if node.Head == nil{
		node.Head = &ListNode{Val:data, Next:nil}
		return
	}

	last := node.Head
	for last.Next != nil{
		last = last.Next
	}
	last.Next = &ListNode{Val:data, Next:nil}
}


func (node *ListNodeStruct) removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil{
		return nil
	}
	if head.Next == nil{
		return nil
	}
	slowPtr := head
	fastPtr := head
	slowPrev := head
	fastPtr = node.moveNTimes(fastPtr, n)
    counter := 0
	for fastPtr != nil {
		slowPrev = slowPtr
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next
		counter += 1
	}
	if counter == 0{
		head = head.Next
		return head
	}

	fmt.Println("Slow previous", slowPrev.Val)
	slowPrev.Next = slowPtr.Next
	slowPtr = nil
	return head
}

// 1 -> 2 -> Nil
//
func (node *ListNodeStruct) moveNTimes(fast *ListNode, n int)*ListNode{
	if fast == nil{
		return nil
	}
	ctr := 1
	for fast != nil{
		if ctr <= n{
			fast = fast.Next
			ctr += 1
			continue
		}
		break;
	}
	return fast
}
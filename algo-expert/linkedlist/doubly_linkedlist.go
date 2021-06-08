// Feel free to add methods and fields to the struct definitions.
package linkedlist

type Node struct {
	Value      int
	Prev, Next *Node
}

type DoublyLinkedList struct {
	Head, Tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{nil, nil}
}

func (ll *DoublyLinkedList) SetHead(node *Node) {
	if ll.Head == nil{
		ll.Head = node
		ll.Head.Next = nil
		ll.Head.Prev = nil
		return
	}
	temp := ll.Head
	ll.Head = node
	ll.Head.Next = temp
	temp.Prev = ll.Head
}

func (ll *DoublyLinkedList) SetTail(node *Node) {
	// Write your code here.
	if ll.Tail == nil{
		ll.Tail = node
		ll.Tail.Next = nil
		ll.Tail.Prev = nil
		return
	}
	temp := ll.Tail
	temp.Next = node
	ll.Tail = node
	ll.Tail.Prev = temp
	ll.Tail.Next = nil
}

func (ll *DoublyLinkedList) InsertBefore(node, nodeToInsert *Node) {
	// Write your code here.
	if ll.Head == nil && ll.Tail == nil{
		ll.SetHead(nodeToInsert)
		ll.SetTail(nodeToInsert)
		return
	}
	if ll.Head == node{
		ll.SetHead(nodeToInsert)
		return
	}
	var slowPtr *Node
	fastPtr := ll.Head
	for fastPtr != node && fastPtr != nil  {
		slowPtr = fastPtr
		fastPtr = fastPtr.Next
	}
	if fastPtr == node{
		slowPtr.Next = nodeToInsert
		nodeToInsert.Prev = slowPtr
		nodeToInsert.Next = fastPtr
		fastPtr.Prev = nodeToInsert
	}
}

func (ll *DoublyLinkedList) InsertAfter(node, nodeToInsert *Node) {
	// Write your code here.

}

func (ll *DoublyLinkedList) InsertAtPosition(position int, nodeToInsert *Node) {
	// Write your code here.
}

func (ll *DoublyLinkedList) RemoveNodesWithValue(value int) {
	// Write your code here.
}

func (ll *DoublyLinkedList) Remove(node *Node) {
	// Write your code here.
}

func (ll *DoublyLinkedList) ContainsNodeWithValue(value int) bool {
	// Write your code here.
	return false
}

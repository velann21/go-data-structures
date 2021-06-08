package linkedlist

func RemoveKthNodeFromEnd(head *LinkedList, k int) {
	// Write your code here.
	if head == nil{
		return
	}
	slowPtr := head
	fastPtr := head
	prevNode := head
	counter := k
	for counter > 0 && fastPtr != nil{
		fastPtr = fastPtr.Next
		counter--
	}
	if counter > 0{
		head = slowPtr.Next
		return
	}
	for fastPtr != nil {
		prevNode =  slowPtr
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next
	}
	prevNode.Next = slowPtr.Next
}

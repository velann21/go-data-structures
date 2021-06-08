package linkedlist

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	// Write your code here.
	if linkedList == nil{
		return nil
	}
	if linkedList.Next == nil{
		return linkedList
	}
	slowPtr := linkedList
	fastPtr := linkedList.Next
	for fastPtr != nil{
		if slowPtr.Value == fastPtr.Value{
			slowPtr.Next = fastPtr.Next
			fastPtr = fastPtr.Next
		}else{
			slowPtr = fastPtr
			fastPtr = fastPtr.Next
		}

	}
	return linkedList
}


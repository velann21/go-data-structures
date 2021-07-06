package linkedlist

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func MergeLinkedLists(headOne *LinkedList, headTwo *LinkedList) *LinkedList {
	// Write your code here.
	if headOne == nil && headTwo == nil{
		return nil
	}
	if headOne == nil && headTwo != nil{
		return headTwo
	}
	if headTwo == nil && headOne != nil{
		return headOne
	}

	firstOne := headOne
	secondOne := headTwo
	var head *LinkedList
	var tail *LinkedList

	if firstOne.Value < secondOne.Value{
		head = firstOne
		tail = firstOne
		firstOne = firstOne.Next
		head.Next = tail
		tail.Next = nil

	}

	for firstOne != nil && secondOne != nil{
		if firstOne.Value < secondOne.Value{
			head = firstOne
			tail = firstOne
			firstOne = firstOne.Next
			head.Next = tail
			tail.Next = nil

		}else if secondOne.Value < firstOne.Value{
			tail = secondOne
			secondOne = secondOne.Next
			head.Next = tail
			tail.Next = nil
		}else{
			tail = secondOne
			head.Next = tail
			secondOne = secondOne.Next
			firstOne = firstOne.Next
		}
	}
	if firstOne != nil{
		tail = firstOne
		firstOne = firstOne.Next
	}

	if secondOne != nil{
		tail = secondOne
		secondOne = secondOne.Next
	}

	return head
}

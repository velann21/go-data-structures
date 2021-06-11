package main

import "fmt"

type LinkedList struct {
	Value int
	Next  *LinkedList
	Tail *LinkedList
}

func FindLoop(head *LinkedList) *LinkedList {
	// Write your code here.
	if head == nil || head.Next == nil{
		return nil
	}
	tortisePointer := head
	rabbitPointer := head
	for rabbitPointer != nil && rabbitPointer.Next != nil{
		rabbitPointer = rabbitPointer.Next.Next
		tortisePointer = tortisePointer.Next
		if rabbitPointer == tortisePointer{
			tortisePointer = head
			break
		}
	}

	for rabbitPointer != nil{
		tortisePointer = tortisePointer.Next
		rabbitPointer = rabbitPointer.Next
		if rabbitPointer == tortisePointer{
			break
		}
	}
	var loopingLinks  *LinkedList
	for rabbitPointer != nil{
		if loopingLinks == nil{
			newLL := &LinkedList{Value: rabbitPointer.Value, Next: nil}
			loopingLinks = newLL
			loopingLinks.Tail = newLL
		}else {
			newLL := &LinkedList{Value: rabbitPointer.Value, Next: nil}
			loopingLinks.Tail.Next = newLL
			loopingLinks.Tail = newLL
		}
		rabbitPointer = rabbitPointer.Next
		if rabbitPointer == tortisePointer{
			break
		}
	}


	return loopingLinks
}

func main() {
	ll := &LinkedList{Value: 0, Next: nil}

	ll2 := &LinkedList{Value: 1, Next: nil}
	ll.Next = ll2
	ll3 := &LinkedList{Value: 2, Next: nil}
	ll2.Next = ll3
	ll4 := &LinkedList{Value: 3, Next: nil}
	ll3.Next = ll4
	ll5 := &LinkedList{Value: 4, Next: nil}
	ll4.Next = ll5
	ll6 := &LinkedList{Value: 5, Next: nil}
	ll5.Next = ll6
	ll7 := &LinkedList{Value: 6, Next: nil}
	ll6.Next = ll7
	ll8 := &LinkedList{Value: 7, Next: nil}
	ll7.Next = ll8
	ll9 := &LinkedList{Value: 8, Next: nil}
	ll8.Next = ll9
	ll0 := &LinkedList{Value: 9, Next: nil}
	ll9.Next = ll0
	ll0.Next = ll5

	outputs := FindLoop(ll)
	for outputs != nil{
		fmt.Println(outputs.Value)
		outputs = outputs.Next
	}
}




//package main
//
//import "fmt"
//
//// This is an input struct. Do not edit.
//type LinkedList struct {
//	Value int
//	Next  *LinkedList
//	Tail *LinkedList
//}
//
//func SumOfLinkedLists(linkedListOne *LinkedList, linkedListTwo *LinkedList) *LinkedList {
//	// Write your code here.
//	if linkedListOne == nil{
//		return linkedListTwo
//	}
//	if linkedListTwo == nil{
//		return linkedListOne
//	}
//
//	var linkedListThree *LinkedList
//	carry := 0
//	for linkedListOne != nil && linkedListTwo != nil{
//		value1 := linkedListOne.Value
//		value2 := linkedListTwo.Value
//		sum := value1+value2+carry
//
//		if sum >= 10{
//
//			carry = sum/10
//			sum = sum%10
//
//		}else {
//			carry = 0
//		}
//		if linkedListThree == nil{
//			newLL := &LinkedList{Value: sum, Next: nil}
//			linkedListThree = newLL
//			linkedListThree.Tail = newLL
//		}else {
//			newLL := &LinkedList{Value: sum, Next: nil}
//			linkedListThree.Tail.Next = newLL
//			linkedListThree.Tail = newLL
//		}
//
//		linkedListOne = linkedListOne.Next
//		linkedListTwo = linkedListTwo.Next
//
//	}
//	for linkedListOne != nil{
//		value1 := linkedListOne.Value
//		sum := value1+carry
//
//		if linkedListThree == nil{
//			newLL := &LinkedList{Value: sum, Next: nil}
//			linkedListThree = newLL
//			linkedListThree.Tail = newLL
//		}else {
//			newLL := &LinkedList{Value: sum, Next: nil}
//			linkedListThree.Tail.Next = newLL
//			linkedListThree.Tail = newLL
//		}
//
//		linkedListOne = linkedListOne.Next
//
//	}
//
//	for linkedListTwo != nil{
//		value2 := linkedListTwo.Value
//		sum := value2+carry
//		if linkedListThree == nil{
//			newLL := &LinkedList{Value: sum, Next: nil}
//			linkedListThree = newLL
//			linkedListThree.Tail = newLL
//		}else {
//			newLL := &LinkedList{Value: sum, Next: nil}
//			linkedListThree.Tail.Next = newLL
//			linkedListThree.Tail = newLL
//		}
//		linkedListTwo = linkedListTwo.Next
//	}
//	if carry != 0{
//		if linkedListThree == nil{
//			newLL := &LinkedList{Value: carry, Next: nil}
//			linkedListThree = newLL
//			linkedListThree.Tail = newLL
//		}else {
//			newLL := &LinkedList{Value: carry, Next: nil}
//			linkedListThree.Tail.Next = newLL
//			linkedListThree.Tail = newLL
//		}
//	}
//	return linkedListThree
//}
//
//func main() {
//	ll := &LinkedList{Value: 2, Next: nil}//Next:&LinkedList{Value: 4, Next:&LinkedList{Value: 7, Next:&LinkedList{Value: 1}}}}
//	ll2 := &LinkedList{Value: 9, Next:  nil}//&LinkedList{Value: 4, Next: &LinkedList{Value: 5}}}
//	ll3 := SumOfLinkedLists(ll, ll2)
//	for ll3 != nil{
//		fmt.Println(ll3.Value)
//		ll3 = ll3.Next
//	}
//}

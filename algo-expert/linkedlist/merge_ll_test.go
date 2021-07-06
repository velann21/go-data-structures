package linkedlist

import (
	"fmt"
	"testing"
)

func TestMergeLinkedLists(t *testing.T) {
	ll := &LinkedList{Value: 1, Next: &LinkedList{Value:3, Next:&LinkedList{Value:4, Next:&LinkedList{Value:8, Next: nil}}}}
	ll2 := &LinkedList{Value: 2, Next: &LinkedList{Value:5, Next:&LinkedList{Value:6, Next:nil}}}

	ll3 := MergeLinkedLists(ll, ll2)
	for ll3.Next != nil{
		fmt.Println(ll3.Value)
	}
}
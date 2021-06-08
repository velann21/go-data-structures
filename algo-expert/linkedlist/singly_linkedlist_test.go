package linkedlist

import "testing"

func TestRemoveKthNodeFromEnd(t *testing.T) {
	ll := &LinkedList{Value: 0, Next: &LinkedList{Value: 1, Next: &LinkedList{Value: 2, Next: nil}}}
	RemoveKthNodeFromEnd(ll, 4)
}
// Feel free to add methods and fields to the struct definitions.
package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoublyLinkedList_SetHead(t *testing.T) {
	tests := []struct{
		input *Node
		expectedResult int
	}{
		{input: &Node{Value: 10, Prev: nil, Next: nil}, expectedResult: 10},
		{input: &Node{Value: 20, Prev: nil, Next: nil}, expectedResult: 20},
	}
	ddl := NewDoublyLinkedList()
	for _, test := range tests{
		ddl.SetHead(test.input)
		assert.Equal(t, test.expectedResult, ddl.Head.Value)
	}

}

func TestDoublyLinkedList_SetTail(t *testing.T) {
	tests := []struct{
		input *Node
		expectedResult int
	}{
		{input: &Node{Value: 30, Prev: nil, Next: nil}, expectedResult: 30},
		{input: &Node{Value: 40, Prev: nil, Next: nil}, expectedResult: 40},
	}
	ddl := NewDoublyLinkedList()
	for _, test := range tests{
		ddl.SetTail(test.input)
		assert.Equal(t, test.expectedResult, ddl.Tail.Value)
	}

}
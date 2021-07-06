package tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	tests := []struct{
		input int
		err error
	}{
		{
		input: 10,
		err: nil,
		},
		{
			input: 20,
			err: nil,
		},
		{
			input: 30,
			err: nil,
		},
	}

	q := Queue{}
	for _, test := range tests{
		q.Enqueue(&Node{nil, nil, test.input})

	}
	data := q.Dequeue()
	fmt.Println(data)
	data2 := q.Dequeue()
	fmt.Println(data2)
	data3 := q.Dequeue()
	fmt.Println(data3)
}


func TestTree_Insert(t *testing.T) {
	q := Queue{}
	mu := &sync.Mutex{}
	tree := Tree{lock: mu, q: q}
	tests := []struct{
		leftNode int
		rightNode int
		err error
	}{
		{leftNode: 10, rightNode: -1, err: nil},
		{leftNode: 20, rightNode: 30, err: nil},
		{leftNode: 40, rightNode: 50, err: nil},
	}
	for _, test := range tests{
		err := tree.Insert(test.leftNode, test.rightNode)
		assert.IsType(t, test.err, err)
	}
	//tree.PreOrder(tree.Root)
	fmt.Println(tree.SumValues(tree.Root))
	fmt.Println(tree.CountNodes(tree.Root))
	fmt.Println(tree.CountLeafNodes(tree.Root))
	fmt.Println(tree.CountTwoChildedNodes(tree.Root))
	fmt.Println(tree.BranchSum(tree.Root, 0 ))
}

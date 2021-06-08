package stack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push2(t *testing.T) {
	stack := NewStackLL()
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	data, err := stack.Peek(2)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(data)

	data1, err := stack.Peek(3)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(data1)

	data3, err := stack.Peek(4)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(data3)
}

func TestStack_QueueUsingStack(t *testing.T) {
	tests := []struct{
		auxiliaryStk   *Stack
		inputData      int
		expectedResult []int
		err            error
	}{
		{auxiliaryStk: nil, inputData: 10, expectedResult: []int{10}, err: nil},
		{auxiliaryStk: NewStackLL(), inputData: 20, expectedResult: []int{10,20}, err: nil},
		{auxiliaryStk: NewStackLL(), inputData: 30, expectedResult: []int{10,20, 30}, err: nil},
	}
	stk := NewStackLL()
	for _, test := range tests{
		err := stk.QueueUsingStack(test.auxiliaryStk, test.inputData)
		assert.IsType(t, test.err, err)
	}

}
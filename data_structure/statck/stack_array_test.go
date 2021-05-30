package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
	tests := []struct {
		data int
		err error
	}{  {data: 10, err: nil},
		{data: -1, err: InvalidData},
		{data: 20, err: nil},
		{data: 30, err: nil},
		}
	stack := NewStack(-1)

	for _, value := range tests{
		err := stack.Push(value.data)
		assert.IsType(t, value.err, err)
	}
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack(-1)
	tests := []struct{
		push error
		expectedData int
		err error
	}{
		{push: stack.Push(10), expectedData: 20, err: nil},
		{push: stack.Push(20), expectedData: 10, err: nil},
		{push: nil, expectedData: 0, err: NoData},
	}
	for _, test := range tests{
		data, err := stack.Pop()
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expectedData, data)
	}
}

func TestStack_Peek(t *testing.T) {
	stack := NewStack(-1)
	tests := []struct{
		push error
		inputPosition int
		expectedData int
		err error
	}{
		{push: stack.Push(10),inputPosition: 2,expectedData: 10, err: nil},
		{push: stack.Push(20),inputPosition: 1, expectedData: 20, err: nil},
	}

	for _, test := range tests{
		data, err := stack.Peek(test.inputPosition)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expectedData, data)
	}
}

func TestStack_Negative(t *testing.T) {
	stack := NewStack(-1)
	tests := []struct{
		push error
		inputPosition int
		expectedData int
		err error
	}{
		{push: nil,inputPosition: 2,expectedData: 10, err: nil},
	}

	for _, test := range tests{
		data, err := stack.Peek(test.inputPosition)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expectedData, data)
	}
}
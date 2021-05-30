package stack

import (
	"fmt"
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
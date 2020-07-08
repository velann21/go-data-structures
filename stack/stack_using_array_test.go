package main

import (
	"fmt"
	"testing"
)
func TestStack_Push(t *testing.T) {
    stk := Stack{}
	stk.Push(10)
	stk.Push(20)
	stk.Push(30)
	stk.Push(40)
	stk.Push(50)
	stk.Push(60)
}

func TestStack_Pop(t *testing.T) {
	stk := Stack{}
	stk.Push(10)
	stk.Push(20)
	stk.Push(30)
	stk.Push(40)
	stk.Push(50)
	stk.Push(60)

	fmt.Println(stk.Pop())
}

func TestStack_IsEmpty(t *testing.T) {
	stk := Stack{}
	fmt.Println(stk.IsEmpty())
}

func TestStack_IsEmptyFalse(t *testing.T) {
	stk := Stack{}
	stk.Push(60)
	fmt.Println(stk.IsEmpty())
}

func TestStack_Peek(t *testing.T) {
	stk := Stack{}
	stk.Push(10)
	stk.Push(20)
	stk.Push(30)
	stk.Push(40)
	stk.Push(50)
	stk.Push(60)
	err := stk.Peek(1)
	if err != nil{
		fmt.Println(err)
	}
}




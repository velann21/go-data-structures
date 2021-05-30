package stack

import (
	"errors"
)

var InvalidData =  errors.New("InvalidData")
var NoData = errors.New("NoData")
var InvalidIndex = errors.New("InvalidIndex")

type stack struct {
	top int
	storage []int
}

func NewStack(top int)*stack{

	return &stack{top: top, storage: make([]int, 0)}
}

func (stk *stack) Push(data int)error{
	if data <= 0{
		return InvalidData
	}
	stk.top = stk.top+1
	stk.storage = append(stk.storage, data)
	return nil
}

func (stk *stack) Pop()(int, error){
	if stk.IsEmpty(){
		return 0, NoData
	}
	data := stk.storage[stk.top]
	stk.top = stk.top-1
	return data, nil
}

func (stk *stack) Peek(pos int)(int, error){
	if stk.IsEmpty(){
		return 0, NoData
	}
	if pos <= 0 || pos > stk.top{
		return 0, InvalidIndex
	}
	lookUpPosition := stk.top-pos+1
	return stk.storage[lookUpPosition], nil
}

func (stk *stack) IsEmpty()bool{
	return stk.top == -1
}



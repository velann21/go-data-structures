package stack

import "errors"

type Stack struct {
	array []int
}

func (stk *Stack) Push(data int){
	stk.array = append(stk.array, data)
}

func (stk *Stack) Pop()int{
	data := stk.array[0]
	stk.array = stk.array[1:len(stk.array)-1]
	return data
}

func (stk *Stack) Peek(pos int)error{
	if pos > len(stk.array){
		return errors.New("Invalid Request")
	}

	return nil
}

func (stk *Stack) IsEmpty()bool{
	return len(stk.array) == 0
}


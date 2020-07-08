package main

import (
	"errors"
)

type Stack struct {
	array []int
}

func (stk *Stack) Push(data int){
	stk.array = append(stk.array, data)
}

func (stk *Stack) Pop()int{
	data := stk.array[len(stk.array)-1]
	stk.array = stk.array[0:len(stk.array)-1]
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


func main(){
	stk := Stack{}
	stk.Push(10)
	stk.Push(20)
	stk.Push(30)
	stk.Push(40)
	stk.Push(50)
	stk.Push(60)


}
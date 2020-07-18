package main

func main() {

}

type Stack struct {
	Storage []int
}

func (stk *Stack) Push(data int){
	stk.Storage = append(stk.Storage, data)
}

func (stk *Stack) Pop()int{
	first := stk.Storage[len(stk.Storage)]
	stk.Storage = stk.Storage[0:len(stk.Storage)-1]
	return first
}

func (stk *Stack) Peek()int{
	first := stk.Storage[len(stk.Storage)]
	return first
}






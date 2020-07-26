package main

import "fmt"

func main() {
	RainWaterTrapping([]int{0,1,0,2,1,0,1,3,2,1,2,1})
}

func RainWaterTrapping(data []int)int{
	TotalUnit := 0
	currentIndex := 0
	stk := RectStackNode1{}

	for currentIndex < len(data){
		if stk.IsEmpty() || data[stk.Peek().Data] > data[currentIndex]{
			stk.Push(currentIndex)
			currentIndex = currentIndex+1
		}else{
			top := stk.Pop().Data
			fmt.Println(top)
			TotalUnit++
		}
	}
	return 0
}

type RectStack1 struct {
	Data int
	Next *RectStack1
}

type RectStackNode1 struct {
	Head *RectStack1
}

func (stk *RectStackNode1) Push(data int){
	if stk.Head == nil{
		stk.Head = &RectStack1{Data:data, Next:nil}
		return
	}
	newNode := &RectStack1{Data:data, Next:stk.Head}
	stk.Head = newNode
	return
}

func (stk *RectStackNode1) Pop()*RectStack1{
	if stk.Head == nil{
		return nil
	}

	temp := stk.Head
	if stk.Head.Next != nil{
		stk.Head = stk.Head.Next
	}else{
		stk.Head = nil
	}
	return temp
}

func (stk *RectStackNode1) Peek()*RectStack1{
	return stk.Head
}

func (stk *RectStackNode1) IsEmpty()bool{
	return stk.Head == nil
}



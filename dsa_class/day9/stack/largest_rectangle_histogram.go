package main

import "fmt"

func main() {
	maxArea := LargestRectInHistogram([]int{1,2,3,4,2})
	fmt.Println(maxArea)
}

func LargestRectInHistogram(data []int)int{
	Area := 0
	maxArea := 0
	currentIndex := 0
	stkNode := RectStackNode{}
	for currentIndex < len(data){
		if stkNode.IsEmpty() || data[stkNode.Peek().Data] < data[currentIndex]{
			stkNode.Push(currentIndex)
			currentIndex = currentIndex+1
		}else{
			top := stkNode.Pop()
			if stkNode.IsEmpty(){
				Area = data[top.Data]*currentIndex
			}else{
				Area = data[top.Data]*(currentIndex - stkNode.Peek().Data -1)
			}
			if Area > maxArea{
				maxArea = Area
			}

		}
	}

	for !stkNode.IsEmpty(){
		top := stkNode.Pop()
		if stkNode.IsEmpty(){
			Area = data[top.Data]*currentIndex
		}else{
			Area = data[top.Data]*(currentIndex - stkNode.Peek().Data -1)
		}
		if Area > maxArea{
			maxArea = Area
		}
	}
	return maxArea
}

type RectStack struct {
	Data int
	Next *RectStack
}

type RectStackNode struct {
	Head *RectStack
}

func (stk *RectStackNode) Push(data int){
	if stk.Head == nil{
		stk.Head = &RectStack{Data:data, Next:nil}
		return
	}
	newNode := &RectStack{Data:data, Next:stk.Head}
	stk.Head = newNode
	return
}

func (stk *RectStackNode) Pop()*RectStack{
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

func (stk *RectStackNode) Peek()*RectStack{
	return stk.Head
}

func (stk *RectStackNode) IsEmpty()bool{
	return stk.Head == nil
}

package stack
import "fmt"

func main() {
	stk := StackImpl{}
	stk1 := StackImpl{}

	var inputNumber  int
	_, _ = fmt.Scanln(&inputNumber)
	for i:=1 ; i<=inputNumber; i++{
		var query  int
		_, _ = fmt.Scanln(&query)
		if query == 1{
			var newData  int
			_, _ = fmt.Scan(&newData)
			stk.Push(newData)
			stk1.CheckMax(newData)
		}else if query == 2{
			popData := stk.Pop()
			if popData != nil{
				stk1.TakeOutIfMatchesMax(popData.Data)
			}

		}else if query == 3{
			max := stk1.GetMax()
			fmt.Println(max.Data)
		}
	}
}

type Stack struct{
	Data int
	Next *Stack
}

type StackImpl struct{
	Head *Stack
}

func (stk *StackImpl) Push(data int){
	if stk.Head == nil{
		stk.Head = &Stack{Data: data, Next: nil}
		return
	}
	newNode := &Stack{Data: data, Next: stk.Head}
	stk.Head = newNode
	return
}

func (stk *StackImpl) Pop()*Stack{
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

func (stk *StackImpl) GetMax()*Stack{
	if stk.Head == nil{
		return nil
	}
	return stk.Head
}

func (stk *StackImpl) CheckMax(data int){
	if stk.Head == nil{
		stk.Push(data)
		return
	}

	if data > stk.Head.Data {
		stk.Push(data)
	}
}

func (stk *StackImpl) TakeOutIfMatchesMax(data int){
	if stk.Head == nil{
		return
	}
	if stk.Head.Data == data{
		stk.Pop()
	}

}

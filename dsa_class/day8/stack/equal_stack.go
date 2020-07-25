package stack

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the equalStacks function below.
 */
func equalStacks(h1 []int32, h2 []int32, h3 []int32) int32 {
	/*
	 * Write your code here.
	 */
	if h1 == nil && h2 == nil && h3 == nil{
		return 0
	}
	if len(h1)<=0 && len(h2)<=0 && len(h3) <= 0{
		return 0
	}
	stk1 := StackTraverser{}
	stk2 := StackTraverser{}
	stk3 := StackTraverser{}
	for i:=len(h1)-1; i>=0; i--{
		stk1.CalculateStackHeight(h1[i])
	}

	for i:=len(h2)-1; i>=0; i--{
		stk2.CalculateStackHeight(h2[i])
	}

	for i:=len(h3)-1; i>=0; i--{
		stk3.CalculateStackHeight(h3[i])
	}

	r1, r2, r3 := stk1.Peek().Data, stk2.Peek().Data, stk3.Peek().Data
	for (r1!=r2 && r2!=r3){
		if r1 >= r2 && r1 >= r3{
			stk1.Pop()
			if stk1.IsEmpty(){
				r1 = 0
			}else{
				r1 = stk1.Peek().Data
			}
		}
		if r2 >= r1 && r2 >= r3{
			stk2.Pop()
			if stk2.IsEmpty(){
				r2 = 0
			}else{
				r2 = stk2.Peek().Data
			}
		}

		if r3 >= r1 && r3 >= r2{
			stk3.Pop()
			if stk3.IsEmpty(){
				r3 = 0
			}else{
				r3 = stk3.Peek().Data
			}
		}
	}

	return r1
}

type Stack1 struct{
	Data int32
	Next *Stack1
}

type StackTraverser struct{
	Head *Stack1
}
func (stk *StackTraverser) Push(data int32){
	if stk.Head == nil{
		stk.Head = &Stack1{Data:data, Next: nil}
		return
	}
	newNode := &Stack1{Data:data, Next: stk.Head}
	stk.Head = newNode
	return
}

func (stk *StackTraverser) Pop()*Stack1{
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

func (stk *StackTraverser) Peek()*Stack1{
	if stk.Head == nil{
		return nil
	}
	return stk.Head
}

func (stk *StackTraverser) IsEmpty()bool{
	return stk.Head == nil
}

func (stk *StackTraverser) CalculateStackHeight(data int32){
	if stk.Head == nil{
		stk.Push(data)
		return
	}
	temp := stk.Head
	currentHeight := temp.Data+data
	stk.Push(currentHeight)
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create("/Users/singaravelannandakumar/go/src/awesomeProject3/dsa_class/day8/Input")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	n1N2N3 := strings.Split(readLine(reader), " ")

	n1Temp, err := strconv.ParseInt(n1N2N3[0], 10, 64)
	checkError(err)
	n1 := int32(n1Temp)

	n2Temp, err := strconv.ParseInt(n1N2N3[1], 10, 64)
	checkError(err)
	n2 := int32(n2Temp)

	n3Temp, err := strconv.ParseInt(n1N2N3[2], 10, 64)
	checkError(err)
	n3 := int32(n3Temp)

	h1Temp := strings.Split(readLine(reader), " ")

	var h1 []int32

	for h1Itr := 0; h1Itr < int(n1); h1Itr++ {
		h1ItemTemp, err := strconv.ParseInt(h1Temp[h1Itr], 10, 64)
		checkError(err)
		h1Item := int32(h1ItemTemp)
		h1 = append(h1, h1Item)
	}

	h2Temp := strings.Split(readLine(reader), " ")

	var h2 []int32

	for h2Itr := 0; h2Itr < int(n2); h2Itr++ {
		h2ItemTemp, err := strconv.ParseInt(h2Temp[h2Itr], 10, 64)
		checkError(err)
		h2Item := int32(h2ItemTemp)
		h2 = append(h2, h2Item)
	}

	h3Temp := strings.Split(readLine(reader), " ")

	var h3 []int32

	for h3Itr := 0; h3Itr < int(n3); h3Itr++ {
		h3ItemTemp, err := strconv.ParseInt(h3Temp[h3Itr], 10, 64)
		checkError(err)
		h3Item := int32(h3ItemTemp)
		h3 = append(h3, h3Item)
	}

	result := equalStacks(h1, h2, h3)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}


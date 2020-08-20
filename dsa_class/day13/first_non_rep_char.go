package main

import "fmt"

func main() {
	util := StringUtil{}
	util.input = []string{"a", "a", "b", "c", "e", "b", "c"}
	fmt.Println(util.FindNonRepeatedCharApproach2())

}

type StringUtil struct {
	input []string
}

func (util *StringUtil) FindNonRepeatedCharApproach1()string{
	return  MovecharByChar(util.input)
}

func MovecharByChar(input []string)string{
	nonRepChar := false
	for i, charPtr := range input {
		nonRepChar = true
		isNonRep := MoveCursor(input, charPtr, i, nonRepChar)
		if isNonRep{
			return charPtr
		}
	}
	return ""
}

func MoveCursor(input []string, charPtr string, charPosition int, nonRepChar bool)bool{
	for j, recursor := range input{
		if charPtr == recursor && charPosition != j{
			nonRepChar = false
			break
		}else{
			nonRepChar = true
		}
	}
	return nonRepChar
}

func (util *StringUtil) FindNonRepeatedCharApproach2()string{
	charscount := countChars(util.input)
	return findRepeatdOne(charscount)
}

func countChars(input []string)map[string]int{
	hash := map[string]int{}
	for _, v := range input{
		hash[v] += 1
	}
	return hash
}

func findRepeatdOne(input map[string]int)string{
	for key, vale := range input{
		if vale == 1{
			return key
		}
	}
	return ""
}
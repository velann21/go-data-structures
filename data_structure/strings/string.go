package strings

import (
	"errors"
	"fmt"
)

type String struct {

}

func (str *String) Length(input string)int{
	if input == ""{
		return 0
	}
	length := 0
	for i:=0;i<len(input);i++{
		length++
	}
	return length
}

func (str *String) Uppercase(input string)[]byte{
	if input == ""{
		return nil
	}
	temp := make([]byte, 0)
	for i:=0; i<len(input); i++{
		temp = append(temp, input[i] - 32)
	}
	fmt.Println(string(temp))
	return temp
}

var InvalidStr = errors.New("InvalidString")
func (str *String) FindDuplicates(input string)([]byte, error){
	if input == ""{
		return nil, InvalidStr
	}
	alphabetArray := [26]int{}
	temp := make([]byte, 0)
	for i:=0; i<len(input); i++{
		if input[i] < 97 || input[i] > 122{
			return nil, InvalidStr
		}else{
			index := input[i] - 97
			alphabetArray[index] += 1
		}
	}
	for index, value := range alphabetArray{
		if value > 1{
			temp = append(temp, byte(97+index))
		}
	}
	return temp, nil
}

func (str *String) FindDuplicatesUsingBits(input string)([]byte, error){
	if input == ""{
		return nil, InvalidStr
	}
	var x int32 = 0
	var y int32 = 0
	auxilaryArray := make([]byte, 0)
	for i :=0; i<len(input);i++{
		if input[i] < 97 || input[i] > 122{
			return nil, InvalidStr
		}
		y = 1
		val := input[i] - 97
		y = y << val
		fmt.Println(y&x)
		if y & x > 0{
			auxilaryArray = append(auxilaryArray, input[i])
		}else{
			x |= 1 << val
		}
	}
	return auxilaryArray, nil
}



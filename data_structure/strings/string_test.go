package strings

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString_Length(t *testing.T) {
	stringObject := String{}
	len := stringObject.Length("velan")
	fmt.Println(len)
}

func TestString_Uppercase(t *testing.T) {
	stringObject := String{}
	upperCase := stringObject.Uppercase("velan")
	fmt.Println(upperCase)
}

func TestString_FindDuplicates(t *testing.T) {
	tests := []struct{
		input string
		expectedResult string
		err error
	}{
		{input: "abc12232", expectedResult: "", err: InvalidStr},
		{input: "findings", expectedResult: "in", err: nil},
		{input: "aaaa", expectedResult: "a", err: nil},
		{input: "", expectedResult: "", err: InvalidStr},
		{input: "abcd", expectedResult: "", err: nil},

	}
	stringObject := String{}
	for _, test := range tests{
		output, err := stringObject.FindDuplicates(test.input)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expectedResult, string(output))
	}
}

func TestString_FindDuplicatesUsingBits(t *testing.T) {
	tests := []struct{
		input string
		expectedResult string
		err error
	}{
		{input: "findn", expectedResult: "n"},
		{input: "aaa", expectedResult: "aa"},
		{input: "findindssasdasas", expectedResult: "indsa"},
	}
	stringObject := String{}
	for _, test := range tests{
		output, err := stringObject.FindDuplicatesUsingBits(test.input)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expectedResult, string(output))
		fmt.Println(string(output))
	}
}

//func (str *String) Uppercase(input string)string{
//	if input == ""{
//		return ""
//	}
//	a := []rune(input)
//	for i:=0; i<len(a); i++{
//		a[i] = rune(input[i] - 32)
//	}
//	return string(a)
//}

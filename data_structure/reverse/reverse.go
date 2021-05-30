package main

import "fmt"

func main() {
	str := "abcdef"
	newStr := ""
	for i :=len(str)-1; i>=0; i--{
		newStr += string(str[i])
	}
	fmt.Println(newStr)
}

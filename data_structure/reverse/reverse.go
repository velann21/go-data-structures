package main

import (
	"fmt"
)

func main() {
	str := "abcdef"
	for i :=len(str)-1; i>=0; i--{
		fmt.Print(string(str[i]))
	}
}

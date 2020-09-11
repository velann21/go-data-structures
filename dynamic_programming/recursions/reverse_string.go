package main

import "fmt"

func main() {
	str := "velan"
	fmt.Println(ReverseString(str, len(str)-1, 0, ""))
}

func ReverseString(str string, bakptr int, frntptr int, newstr string)string{
	if bakptr == 0{
		newstr += string(str[bakptr])
		return newstr
	}
	newstr += string(str[bakptr])
	return ReverseString(str, bakptr-1, frntptr-1, newstr)
}

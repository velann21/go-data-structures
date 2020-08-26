package main

import "fmt"

func main() {
	str := "madam"
	result := ""
	for _, v := range str{
		result = string(v) + result
		fmt.Println(result)
	}
	fmt.Println(result)
	if str == result{
		fmt.Println("Yes I am palindrom")
	}

}

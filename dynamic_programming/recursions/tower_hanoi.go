package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(GetSubseqStr("abc"))

}

//func tower(from []int, to []int, aux []int, n int, start int){
//	if n == 0{
//		to = append(to ,from[len(from)-1])
//	}
//
//	aux = append(aux, from[start])
//	tower(from, to, aux, n-1, start+1)
//	fmt.Println("sdsdsdsd")
//	tower(aux, to, from, n-1, start+1)
//
//}


func TowerOfHan(n int, from string, to string, aux string){
	if n > 0{
		TowerOfHan(n-1, from, aux, to)
		fmt.Println("Moving the disk ",n, " from ", from, " to the ", to, " rod.")
		TowerOfHan(n-1, aux, to, from)
	}

}

func GetSubseqStr(word string)string{
	if word == ""{
		return ""
	}
	first := string(word[0])
	fmt.Println(first)

	rest := GetSubseqStr(word[1:])

	result := ""
	for _, v := range strings.Split(rest, ","){
		result += ","+first+v
		result += ","+v
	}

	return result[1:]
}
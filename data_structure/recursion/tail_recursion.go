package recursion

import "fmt"

//Simple Tail recursion
func TailRecursion(n int){
	if n >= 0{
		TailRecursion(n-1)
		fmt.Println(n)
	}

}

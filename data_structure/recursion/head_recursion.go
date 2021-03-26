package recursion

import "fmt"

// Simple Head recursion
func HeadRecursion(n int){
	if n >= 0{
		fmt.Println(n)
		HeadRecursion(n-1)
	}
}

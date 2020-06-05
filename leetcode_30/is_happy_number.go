package main

import "fmt"

func main() {
	isHap := isHappy(129)
	fmt.Println(isHap)
}

func isHappy(n int) bool {
	result := -1
	Ishappy := false
	if n == 1{
		return true
	}
	for result!=1 && result != 0 {
		rem := 0
		mod := 0

		if findNumberLength(n) == 2{
			rem = n/10
			mod = n%10
		}else if findNumberLength(n) == 3{
			rem = n/100
			mod = n%100

		}else if findNumberLength(n) == 1{
			rem = n/1
			mod = n%1
		}else if findNumberLength(n) == 4{
			rem = n/1000
			mod = n%1000
		}else if findNumberLength(n) == 5{
			rem = n/10000
			mod = n%10000
		}
		result = (rem*rem) + (mod*mod)
		n=result
		if result == 1{
			Ishappy = true
		}
	}
	return Ishappy
}


func findNumberLength(number int)int{
	length := 1;
	if (number >= 100000000) {
		length += 8;
		number /= 100000000;
	}
	if (number >= 10000) {
		length += 4;
		number /= 10000;
	}
	if (number >= 100) {
		length += 2;
		number /= 100;
	}
	if (number >= 10) {
		length += 1;
	}
	return length;
}
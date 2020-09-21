package main

import "fmt"

func main() {
    gas := []int{5,4,3,2,1}

	cost := []int{5,1,5,4,3}
	result := CanCircleBackToStart(gas, cost)
	fmt.Println(result)
}

func CanCircleBackToStart(gas []int, cost []int)int{
	gasRem := 0
	costRem := 0
	start := 0
	tankRem := 0
	for k, v:=range gas{
		gasRem += v
		costRem += cost[k]
		tankRem = gasRem - costRem
		if tankRem < 0{
			tankRem = 0
			start = k+1
		}
	}

	if gasRem < costRem{
		return -1
	}
	return start
}

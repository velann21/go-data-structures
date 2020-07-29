package main

func main() {

}

func CanCircleBack(gas []int, cost []int)int{
	gasRem := 0
	costRem := 0
	start := 0
	tankRem := 0
	for k, v:=range gas{
		gasRem =+ v
		costRem =+ cost[k]
		tankRem = gasRem - costRem
		if tankRem < 0{
			tankRem = 0
			start = k
		}
	}

	if gasRem < costRem{
		return -1
	}
	return start
}

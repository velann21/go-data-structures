package main

import (
	"math/big"
	"fmt"
)

func main() {
	fmt.Println(BinomialCoefficent(20, 30))
	//BigIntFactorial()
}

//B(5, 3) = factorial(5) / (factorial(3) * factorial(2))
//= (1*2*3*4*5) / ((1*2*3) * (1*2))
//= 120 / (6 * 2)
//= 120 / 12
//= 10

func BinomialCoefficent(N int, K int)int{
	isTrue := CheckIsNLessthanK(N, K)
	if isTrue{
		return -1
	}
	NFactResult := factResult(N)
	KFactResult := factResult(K)
	NKFactResult := factResult(N-K)
	mulRes := KFactResult.Mul(KFactResult, NKFactResult)
	finalRes := mulRes.Div(NFactResult, mulRes)
	isGreater := IsNumberGreater(finalRes.Int64())
	if isGreater{
		return -1
	}
	return int(finalRes.Int64())
}

func IsNumberGreater(result int64)bool{
	if result > 1000000000{
		return true
	}
	return false
}

func CheckIsNLessthanK(N int, K int)bool{
	if N < K{
		return true
	}
	return false
}

func factResult(n int)*big.Int{
	var fact = new(big.Int)
	fact.MulRange(1, int64(n))
	return fact
}
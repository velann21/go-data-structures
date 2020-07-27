package arrays

import "fmt"

//problemArry := [0,1,3,4,0]
//ans: [1,3,4,0,0]

func main()  {
	fmt.Println(moveZeros([]int{0,1,3,4,0}))
}

func moveZeros(input []int)[]int{
	count := 0
	zerosCount := 0
	nonZeroCount := 0
	if input == nil{
		return nil
	}
	//Iterate the input array
    for _, v := range input{
    	//If the element is not 0 then move the element to first index and increase the nonZeroCount++ && count++
    	if v != 0{
    		input[count] = v
			count = count+1
			nonZeroCount += 1
			//Increase the zerosCount++ so that atlast we could move the number of zeros to one side
		}  else{
			zerosCount = zerosCount+1
		}
	}
    for nonZeroCount < len(input)-1{
		input[nonZeroCount] = 0
		nonZeroCount = nonZeroCount+1
	}

    return input
}
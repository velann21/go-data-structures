package arrays

import "fmt"

//input : [1,3,3,1,3]
//output : [6]

func main()  {
	fmt.Println(duplicates_triplets([]int{1,5,2,1,5}))
}

func duplicates_triplets(input []int)int{
	temp := 0
	for k:=0; k<len(input); k++{
	  temp ^= input[k]
	}
	return temp
}



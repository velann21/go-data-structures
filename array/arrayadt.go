package array

import (
	"fmt"
)

func main(){
	arr := [3]string{"velan", "indu", "Nandhini"}
	arr1 := arr
	arr[0] = "Nandhakumar"
	fmt.Println(arr1)
	fmt.Println(arr, arr1)

	sli := []string{"velan", "indu", "Nandhini"}
	sli2 := sli
	sli[0] = "Nandhakumar"
	fmt.Println(sli2)
	fmt.Println(sli, sli2)
	goFrom1To10()
}

func Length(array *[10]string)int{
	length := 0
	for _, v := range array{
		if v != ""{
			length ++
		}
	}
	return length
}

func Append(arr [10]string, elem  string){
	len := Length1(arr)
    fmt.Println(len)

}

func Length1(arr [10]string) int{
	for i := len(arr); i >= 0 ; i++{
		if arr[i] == ""{
			continue
		}else {
			return i
		}
	}
	return cap(arr)
}


func goFrom1To10() {
	for i := 0; i <= 10; i++ {
		task(i)
	}
}

func task(i int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic occured: ", r)
			Printst()
			fmt.Println("Am defering done")
		}
	}()

	if i == 2 {
		panic("got 2")
	}
	fmt.Println(i)
}

func Printst(){
	for i:=0; i<=10; i++ {
		fmt.Println(i)
	}
}



package twod_array

import "fmt"

type TwoDimensionalArray struct {

}
func NewTwoDArray()*TwoDimensionalArray{
	return &TwoDimensionalArray{}
}
func (two *TwoDimensionalArray) SpiralTraverse(array [][]int) []int {

	if len(array) == 0{
		return nil
	}

	if len(array[0]) > 0 && len(array) ==1{
		return array[0]
	}
	auxilaryArray := make([]int, 0)
	sr := 0
	sc := 0
	ec := len(array[0])-1
	er := len(array)-1
	rt := 0
	ct := 0

	for sr<er && sc < ec{
		for ct <= ec{
			fmt.Println(array[sr][ct])
			auxilaryArray = append(auxilaryArray, array[sr][ct])
			ct++
		}
		rt = sr+1
		for rt <= er{
			fmt.Println(array[rt][ec])
			auxilaryArray = append(auxilaryArray, array[rt][ec])
			rt++
		}
		ct = ec -1
		for ct >= sc{
			fmt.Println(array[er][ct])
			auxilaryArray = append(auxilaryArray, array[er][ct])
			ct --
		}
		rt = er -1
		for rt > sr{
			fmt.Println(array[rt][sc])
			auxilaryArray = append(auxilaryArray, array[rt][sc])
			rt--
		}
		sc ++
		ec--
		sr++
		er--
		ct = sc
		rt = sr
	}
	return auxilaryArray
}

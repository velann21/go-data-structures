package twod_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpiralTraverse(t *testing.T) {
	tests := []struct{
		inputArray [][]int
		expectedResult []int
	}{
		{inputArray: [][]int{{1, 2, 3},{4, 5, 6}}, expectedResult: []int{1, 2, 3, 6, 5, 4}},
		{inputArray: [][]int{
			{1, 2, 3, 4},
			{12,13,14,5},
			{11,16,15,6},
			{10,9, 8, 7},
		}, expectedResult: []int{1, 2, 3, 4, 5,6,7,8,9,10,11,12,13,14,15,16}},
	}
	tdArray := NewTwoDArray()
	for _, test := range tests{
		output := tdArray.SpiralTraverse(test.inputArray)
		assert.Equal(t, test.expectedResult, output)
	}
}
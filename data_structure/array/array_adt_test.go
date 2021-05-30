package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayList_FindFirstMissingElement(t *testing.T) {
	tests := []struct {
		input []int
		expectedResult int
	}{
		{input: []int{1,2,3,4,6}, expectedResult: 5},
		{input: []int{1}, expectedResult: -1},
		{input: []int{}, expectedResult: -1},
	}
	al := NewArrayList(10)
	for _, test := range tests{
		output := al.FindFirstMissingElement(test.input)
		assert.Equal(t, test.expectedResult, output)
	}
}

func TestArrayList_FindAllMissingElement(t *testing.T) {
	tests := []struct {
		input []int
		expectedResult []int
	}{
		{input: []int{1,2,4,6}, expectedResult: []int{3,5}},
		{input: []int{1,2,3,4,8}, expectedResult: []int{5,6,7}},
		{input: []int{1,2,3,4,5,6}, expectedResult: []int{}},
		{input: []int{1}, expectedResult: nil},
		{input: []int{}, expectedResult: nil},
	}
	al := NewArrayList(10)
	for _, test := range tests{
		output := al.FindAllMissingElement(test.input)
		assert.Equal(t, test.expectedResult, output)
	}
}

func TestArrayList_FindDuplicates(t *testing.T) {
	tests := []struct {
		input []int
		expectedResult []int
	}{
		{input: []int{1,2,3,4,4,5,5,5,6}, expectedResult: []int{4,5}},
		{input: []int{1,2,4,4,6}, expectedResult: []int{4}},
		{input: []int{1,2,3,4,8}, expectedResult: []int{}},
		{input: []int{1}, expectedResult: nil},
		{input: []int{}, expectedResult: nil},
	}
	al := NewArrayList(10)
	for _, test := range tests{
		output := al.FindDuplicates(test.input)
		assert.Equal(t, test.expectedResult, output)
	}
}

func TestArrayList_FindMissingElementUnSortedList(t *testing.T) {
	tests := []struct {
		input []int
		maxNumber int
		expectedResult []int
	}{
		{input: []int{6,1,3,5,7},maxNumber: 7, expectedResult: []int{0,2,4}},
		{input: []int{1},maxNumber: 1, expectedResult: []int{0}},
		{input: []int{},maxNumber:0, expectedResult: nil},
	}
	al := NewArrayList(10)
	for _, test := range tests{
		output := al.FindMissingElementUnSortedList(test.input, test.maxNumber)
		assert.Equal(t, test.expectedResult, output)
	}
}

func TestArrayList_IsValidSeq(t *testing.T) {
	tests := []struct {
		array1 []int
		seqArray []int
		expectedResult bool
	}{
		{array1: []int{5, 1, 22, 25, 6, -1, 8, 10}, seqArray: []int{1, 6, -1, 10}, expectedResult: true},
		{array1: []int{5, 1, 22, 25, 6, -1, 8, 10}, seqArray: []int{6, 1, -1, 10}, expectedResult: false},
		{array1: []int{1,1,1,1,1,1}, seqArray: []int{1,1,1,1,1,1}, expectedResult: true},
		{array1: []int{1,1,1,1,1,1}, seqArray: []int{25}, expectedResult: false},
		{array1: []int{1}, seqArray: []int{25,2,4}, expectedResult: false},
		{array1: []int{1,1,1,1,1,1}, seqArray: []int{1,1,1,}, expectedResult: true},
		{array1: []int{5, 1, 22, 25, 6, -1, 8, 10}, seqArray: []int{5, 1, 22, 25, 6, -1, 8, 10}, expectedResult: true},
	}
	al := NewArrayList(10)
	for _, test := range tests{
		output := al.IsValidSeq(test.array1, test.seqArray)
		assert.Equal(t, test.expectedResult, output)
	}
}
package main

import "fmt"

func main() {
	A := [][]int{
		{5, 4, 4},
		{4, 3, 4},
		{3, 2, 4},
		{2, 2, 2},
		{3, 3, 4},
		{1, 4, 4},
		{4, 1, 1},
	}
	FormMatrixByColor(A)

}

func FormMatrixByColor(areas [][]int) int {
	i := 0
	countryCount := 0
	colorTracker := map[int]int{}
	for row := 0; row < len(areas[i]); row++ {
		for col := 0; col < len(areas[row]); col++ {
			if colorTracker[areas[row][col]] == 0 {
				colorTracker[areas[row][col]] = +1
				FindSameCountriesOrthogonally(areas, row, col)
				countryCount += 1
			}
		}
	}
	return countryCount
}

func FindSameCountriesOrthogonally(areas [][]int, row, col int) {
	MoveColumnsOrthogonally(areas, row, col)
}

func MoveColumnsOrthogonally(areas [][]int, row, col int) {
	fmt.Println( len(areas[row]))
	if  col < len(areas[row])-1 {
		if areas[row][col] != areas[row][col+1] {
			return
		}
		MoveRowsOrthogonally(areas, row+1, col)
		MoveColumnsOrthogonally(areas, row, col+1)
		areas[row][col] = -1
	}
}

func MoveRowsOrthogonally(areas [][]int, row, col int) {
	if areas[row][col] == areas[row-1][col] {
		//tempRow := row
		//for tempRow < len(areas[row]){
		//	tempRow = tempRow
		//	if areas[tempRow][col] == areas[tempRow][col]{
		//		areas[tempRow][col] = -1
		//		areas[tempRow-1][col] = -1
		//	}
		//	tempRow = tempRow+1
		//}
	}
}

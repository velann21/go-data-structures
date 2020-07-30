package main

func main() {
	board := [3][3]string{
		{"X", "O", "X"},
		{"X", "O", "X"},
		{"X", "O", "X"},
	}
	IsValidBoard(board)

}

func IsValidBoard(board [3][3]string) bool {
	XCounts := 0
	ZeroCounts := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == "X" {
				XCounts++
			} else if board[i][j] == "O" {
				ZeroCounts++
			}
		}
	}
	if XCounts == ZeroCounts || XCounts == ZeroCounts+1 {

		return true
	}
	return false
}

func IsWin(board [3][3]string, player string) {


}

func WinPossiblityBoard() [8][3]int {
	winBoard := [8][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 5, 8},
	}
	return winBoard
}

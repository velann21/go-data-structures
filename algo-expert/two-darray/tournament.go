package main

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	// Write your code here.
	if len(competitions) <= 0 || len(results) <= 0 {
		return ""
	}
	if len(results) > len(competitions){
		return ""
	}
	auxilaryMap := make(map[string]int, 0)
	tounamentSize := len(competitions)
	for ts:=0; ts<tounamentSize; ts++{
		if results[ts] == 0{
			winnerName := competitions[ts][results[ts]+1]
			if auxilaryMap[winnerName] == 0{
				auxilaryMap[winnerName] = 3
			}else{
				auxilaryMap[winnerName] += 3
			}
		}else{
			winnerName := competitions[ts][results[ts]-1]
			auxilaryMap[winnerName] += 3
		}
	}

	maxScore := 0
	winnerName := ""
	for key, value := range auxilaryMap{
		if value > maxScore{
			maxScore = value
			winnerName = key
		}
	}
	return winnerName
}

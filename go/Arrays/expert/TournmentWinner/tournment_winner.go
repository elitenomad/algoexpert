package main

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	h := map[string]int{}

	for i := 0; i < len(competitions); i++ {
		teamWon := 1
		if results[i] == 1 {
			teamWon = 0
		}

		h[competitions[i][teamWon]] += 3
	}

	// Now we have hash of winners and the points
	max := 0
	result := ""
	for k, v := range h {
		if max < v {
			max = v
			result = k
		}
	}
	return result
}

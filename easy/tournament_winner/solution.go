package main

import "fmt"

const HOME_TEAM_WON = 1

func main() {
	competitions := [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	}
	results := []int{0, 0, 1}
	fmt.Println(TournamentWinner(competitions, results))
}

func TournamentWinner(competitions [][]string, results []int) string {
	m := make(map[string]int)

	for i := 0; i < len(competitions); i++ {
		if results[i] == HOME_TEAM_WON {
			m[competitions[i][0]] += 3
		} else {
			m[competitions[i][1]] += 3
		}
	}

	winTeam := ""
	highestScore := 0
	for team, score := range m {
		if score > highestScore {
			highestScore = score
			winTeam = team
		}
	}

	return winTeam
}

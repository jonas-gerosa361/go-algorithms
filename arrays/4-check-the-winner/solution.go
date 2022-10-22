package main

import "fmt"

const HOME_TEAM_WON = 1

// O(n) time | O(k) space - n number of competitions and k number of teams
func TournamentWinner(competitions [][]string, results []int) string {
	currentBestTeam := ""
	scores := map[string]int{currentBestTeam: 0}

	for index, competition := range competitions {
		result := results[index]
		homeTeam, awayTeam := competition[0], competition[1]

		winningTeam := awayTeam
		if result == HOME_TEAM_WON {
			winningTeam = homeTeam
		}

		updateSocres(winningTeam, 3, scores)

		if scores[winningTeam] > scores[currentBestTeam] {
			currentBestTeam = winningTeam
		}
	}

	return currentBestTeam
}

func updateSocres(team string, points int, scores map[string]int) {
	scores[team] += points
}

func main() {
	competitions := [][]string{
		{"HTML", "Java"},
		{"Java", "Python"},
		{"Python", "HTML"},
		{"C#", "Python"},
		{"Java", "C#"},
		{"C#", "HTML"},
		{"SQL", "C#"},
		{"HTML", "SQL"},
		{"SQL", "Python"},
		{"SQL", "Java"},
	}

	results := []int{0, 1, 1, 1, 0, 1, 0, 1, 1, 0}

	winner := TournamentWinner(competitions, results)
	fmt.Println(winner)
}

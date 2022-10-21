package main

import "fmt"

func TournamentWinner(competitions [][]string, results []int) string {
	var winner string
	teams := make(map[string]int)

	for index, result := range results {
		if result == 0 {
			result = 1
		} else if result == 1 {
			result = 0
		}
		teams[competitions[index][result]] += 3
	}

	var maxPoints int
	fmt.Println(teams)
	for team, points := range teams {
		if points > maxPoints {
			winner = team
			maxPoints = points
		}
	}

	return winner
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

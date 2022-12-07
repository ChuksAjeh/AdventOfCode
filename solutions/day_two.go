package solutions

import (
	"strings"

	"advent.com/problem/utils"
)

var path_daytwo = "/workspaces/AdventOfCode2022/problem_set/problem_day_two.txt"

func Day_two() int {
	var puzzle = utils.OpenFileForReadAndParse(path_daytwo)
	// puzzle := []string{"A Y", "B X", "C Z"}

	//Running score
	score := 0
	// A/X = Rock B/Y = Paper C/Z = Scissors
	for i := 0; i < len(puzzle); i++ {
		var round = puzzle[i]
		var round_split = strings.Split(round, " ")
		var player_one = round_split[0]
		var player_two = round_split[1]

		if player_two == "X" {
			score++
		} else if player_two == "Y" {
			score += 2
		} else {
			score += 3
		}

		//calculate win:
		score += calculateWinner(player_one, player_two)
	}

	return score
}

func Day_two_partTwo() int {
	var puzzle = utils.OpenFileForReadAndParse(path_daytwo)
	// puzzle := []string{"A Y", "B X", "C Z"}

	//Running score
	score := 0
	// A/X = Rock B/Y = Paper C/Z = Scissors
	for i := 0; i < len(puzzle); i++ {
		var round = puzzle[i]
		var round_split = strings.Split(round, " ")
		var player_one = round_split[0]
		var player_two = round_split[1]

		//calculate Round Ending:
		score += calculateRoundEnd(player_one, player_two)
	}

	return score
}

func calculateWinner(player_one string, player_two string) int {
	//Handle draws:
	if (player_one == "A" && player_two == "X") || (player_one == "B" && player_two == "Y") || (player_one == "C" && player_two == "Z") {
		return 3
	}

	//handle wins
	if (player_one == "C" && player_two == "X") || (player_one == "A" && player_two == "Y") || (player_one == "B" && player_two == "Z") {
		return 6
	}

	//handle losses
	if (player_one == "A" && player_two == "Z") || (player_one == "B" && player_two == "X") || (player_one == "C" && player_two == "Y") {
		return 0
	}

	return -1 //error
}

func calculateRoundEnd(player_one string, player_two string) int {
	//handle loss
	if player_one == "A" && player_two == "X" {
		return 3
	}
	if player_one == "B" && player_two == "X" {
		return 1
	}
	if player_one == "C" && player_two == "X" {
		return 2
	}

	//handle Draw
	if player_one == "A" && player_two == "Y" {
		return 4
	}
	if player_one == "B" && player_two == "Y" {
		return 5
	}
	if player_one == "C" && player_two == "Y" {
		return 6
	}

	//handle wins
	if player_one == "A" && player_two == "Z" {
		return 8
	}
	if player_one == "B" && player_two == "Z" {
		return 9
	}
	if player_one == "C" && player_two == "Z" {
		return 7
	}
	return -1 //error
}

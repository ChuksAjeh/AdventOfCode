package solutions

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"advent.com/problem/utils"
)

// var path_day_four = "/workspaces/AdventOfCode2022/problem_set/problem_day_four.txt"

var path_day_four = "/workspaces/AdventOfCode2022/problem_set/sample_puzzles/problem_day_four_sample.txt"

func Day_four() int {
	var puzzleData = utils.OpenFileForReadAndParse(path_day_four)
	var pairs = solve(puzzleData)
	// var pairs = solve_part_two(puzzleData)
	return pairs
}

func solve(puzzleData []string) int {

	counter := 0
	for i := 0; i < len(puzzleData); i++ {
		var parts = strings.Split(puzzleData[i], ",")
		var firstPair = strings.Split(parts[0], "-")
		var secondPair = strings.Split(parts[1], "-")
		firstPairNumOne, err := strconv.Atoi(firstPair[0])
		firstPairNumTwo, err := strconv.Atoi(firstPair[1])
		secondPairNumOne, err := strconv.Atoi(secondPair[0])
		secondPairNumTwo, err := strconv.Atoi(secondPair[1])

		if err != nil {
			fmt.Println("----------------INTEGER CONVERSION ERROR ------------")
		}

		if math.Abs(float64(firstPairNumOne)-float64(firstPairNumTwo)) > math.Abs(float64(secondPairNumOne)-float64(secondPairNumTwo)) {
			if firstPairNumOne <= secondPairNumOne && firstPairNumTwo >= secondPairNumTwo {
				counter++
			}
		} else {
			if firstPairNumOne >= secondPairNumOne && firstPairNumTwo <= secondPairNumTwo {
				counter++
			}
		}
	}
	return counter
}

// func solve_part_two(puzzleData []string) int {
// 	counter := 0

// 	pattern,_ := regexp.Compile("([0-9]+)-([0-9]+),([0-9]+)-([0-9]+)")
// 	for i:=0; i<len(puzzleData); i++{
// 		var flag = pattern.MatchString(puzzleData[i])
// 		if(flag){
// 			var elf_left_one = strconv.Atoi(patt)
// 		}
// 	}
// 	return counter
// }

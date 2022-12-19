package main

import (
	// "math"
	// "os"
	// "strconv"
	// "strings"

	"advent.com/problem/solutions"
)

// import (
// 	"fmt"
// 	"os"
// 	"time"

// 	"advent.com/problem/solutions"
// )

// day three
// func main() {

// 	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	readFile, err := os.Open("/workspaces/AdventOfCode2022/problem_set/problem_day_three.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer readFile.Close()
// 	fileScan := bufio.NewScanner(readFile)
// 	fileScan.Split(bufio.ScanLines)

// 	sum := 0
// 	var fileLines []string
// 	for fileScan.Scan() {
// 		fileLines = append(fileLines, fileScan.Text())
// 	}

// 	for i := 0; i < len(fileLines); i += 3 {
// 		for y := 0; y < len(fileLines[i]); y++ {
// 			if strings.Contains(fileLines[i+1], string(fileLines[i][y])) && strings.Contains(fileLines[i+2], string(fileLines[i][y])) {
// 				sum += strings.Index(letters, string(fileLines[i][y])) + 1
// 				break
// 			}
// 		}
// 	}

// 	fmt.Println(sum)
// }

//day six

// func main() {
// 	input, err := os.ReadFile("/workspaces/AdventOfCode2022/problem_set/problem_day_six.txt")

// 	if err != nil {
// 		panic(err)
// 	}

// 	result1, result2 := 0, 0

// 	start := time.Now()
// 	for i := 0; i < 10000; i++ {
// 		result1 = solutions.Day_six(input, 4)
// 	}
// 	elapsed1 := time.Since(start) / 10000

// 	start = time.Now()
// 	for i := 0; i < 10000; i++ {
// 		result2 = solutions.Day_six(input, 14)
// 	}
// 	elapsed2 := time.Since(start) / 10000

// 	fmt.Println(result1)
// 	fmt.Println(elapsed1)
// 	fmt.Println(result2)
// 	fmt.Println(elapsed2)
// }

// func main() {
// 	fmt.Println(solutions.Day_four())
// }

// func main() {
// 	input, _ := os.ReadFile("/workspaces/AdventOfCode2022/problem_set/problem_day_five.txt")
// 	split := strings.Split(string(input), "\n\n")
// 	layout := strings.Split(split[0], "\n")
// 	stacks := make(map[int][]rune)
// 	stackKeys := []int{}
// 	for i := len(layout) - 1; i >= 0; i-- {
// 		if i == len(layout)-1 {
// 			for _, k := range strings.Split(strings.TrimSpace(layout[i]), "   ") {
// 				key, _ := strconv.Atoi(k)
// 				stacks[key] = []rune{}
// 				stackKeys = append(stackKeys, key)
// 			}
// 		} else {
// 			for i, c := range layout[i] {
// 				if !strings.ContainsAny(string(c), " []") {
// 					key := int(math.Ceil(float64(i) / 4))
// 					stacks[key] = append(stacks[key], c)
// 				}
// 			}
// 		}

// 	}
// 	moves := strings.Split(split[1], "\n")
// 	// partOne(moves, stacks, stackKeys)
// 	solutions.PartTwo(moves, stacks, stackKeys)
// }

func main() {
	solutions.Day_eight()
}

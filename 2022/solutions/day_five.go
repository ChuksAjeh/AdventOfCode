package solutions

import "fmt"

func PartOne(moves []string, stacks map[int][]rune, stackKeys []int) {
	for _, move := range moves {
		var amount int
		var from int
		var to int
		fmt.Sscanf(move, "move %d from %d to %d", &amount, &from, &to)
		source := stacks[from]
		destination := stacks[to]
		for i := 0; i < amount; i++ {
			n := len(source) - 1 // Top element
			destination = append(destination, source[n])
			source = source[:n] // Pop
		}
		stacks[from] = source
		stacks[to] = destination
	}
	partOne := ""
	for _, k := range stackKeys {
		partOne += string(stacks[k][len(stacks[k])-1])
	}
	fmt.Println(partOne)
}

func PartTwo(moves []string, stacks map[int][]rune, stackKeys []int) {
	for _, move := range moves {
		var amount int
		var from int
		var to int
		fmt.Sscanf(move, "move %d from %d to %d", &amount, &from, &to)
		source := stacks[from]
		destination := stacks[to]
		n := len(source) - amount
		for _, r := range source[n:] {
			destination = append(destination, r)
		}
		source = source[:n]
		stacks[from] = source
		stacks[to] = destination
	}
	partTwo := ""
	for _, k := range stackKeys {
		partTwo += string(stacks[k][len(stacks[k])-1])
	}
	fmt.Println(partTwo)
}

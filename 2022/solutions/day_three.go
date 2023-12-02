package solutions

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
	"advent.com/problem/utils"
)



func Day_three() int {
	var path = "/workspaces/AdventOfCode2022/problem_set/problem_day_three.txt"
	var rucksacks = utils.OpenFileForReadAndParse(path)
	// var rucksacks = []string{
	// 	"vJrwpWtwJgWrhcsFMMfFFhFp",
	// 	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	// 	"PmmdzqPrVvPwwTWBwg",
	// 	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	// 	"ttgJtRGJQctTZtZT",
	// 	"CrZsJsPPZsGzwwsLwLmpwMDw",
	// }

	// var commonChars = goThroughRucksack(rucksacks)
	var commonChars = goThroughRucksackPartTwo(rucksacks)
	fmt.Println(commonChars)

	var total = calcTotal(commonChars)

	return total
}

type void struct{}

var member void

func goThroughRucksack(rucksacks []string) []string {

	var commonChars = []string{}

	for i := 0; i < len(rucksacks); i++ {
		set := make(map[string]void)
		var commonChar = ""
		var rucksack = rucksacks[i]
		var rucksackSize = len(rucksack)
		var compartment_one = rucksack[0 : rucksackSize/2]
		var compartment_two = rucksack[rucksackSize/2:]

		//convert strings to string slices:
		var compartment_one_slice = strings.Split(compartment_one, "")
		var compartment_two_slice = strings.Split(compartment_two, "")

		//sort both:
		sort.Strings(compartment_one_slice)
		sort.Strings(compartment_two_slice)

		//add compartment_one to set:
		for j := 0; j < len(compartment_one_slice); j++ {
			//add compartment one elements:
			set[compartment_one_slice[j]] = member
		}

		for k := 0; k < len(compartment_one_slice); k++ {
			if _, ok := set[compartment_two_slice[k]]; ok {
				commonChar = compartment_two_slice[k]
				commonChars = append(commonChars, commonChar)
				break
			}

		}
	}
	return commonChars
}


//ans 2864
func goThroughRucksackPartTwo(rucksacks []string) []string {

	var commonChars = []string{}
	var group = []string{}
	for i := 0; i < len(rucksacks)+1; i++ {
		var rucksack = ""
		if i < len(rucksacks) {
			rucksack = rucksacks[i]
		}

		//we have reached 3 rucksacks aka a group. Process pool:
		if len(group) == 3 {
			commonChars = append(commonChars, processGroup(group))
			group = nil                     //empty slice
			group = append(group, rucksack) //add in non processed rucksack
		} else {
			group = append(group, rucksack)
		}

		//add compartment_one to set:
		// for j := 0; j < len(compartment_one_slice); j++ {
		// 	//add compartment one elements:
		// 	set[compartment_one_slice[j]] = member
		// }

		// for k := 0; k < len(compartment_one_slice); k++ {
		// 	if _, ok := set[compartment_two_slice[k]]; ok {
		// 		commonChar = compartment_two_slice[k]
		// 		commonChars = append(commonChars, commonChar)
		// 		break
		// 	}

		// }
	}
	return commonChars
}

func processGroup(group []string) string {

	//convert to slices

	// set := make(map[string]void)
	var commonChar = ""

	prim := []bool{}
	PrimCap := []bool{}
	//populate prim
	prim = populateBoolArrayTrue(prim, 53)
	PrimCap = populateBoolArrayTrue(prim, 53)

	for i := 0; i < len(group); i++ {

		//populate secondary:
		sec := []bool{}
		secCap := []bool{}
		sec = populateBoolArrayFalse(sec, 53)
		secCap = populateBoolArrayFalse(sec, 53)

		// for every character of ith string
		for j := 0; j < len(group[i]); j++ {
			var rucksack = []rune(group[i])

			var char = rucksack[j]
			if unicode.IsUpper(char) {
				if PrimCap[int(char)-int('A')] {
					secCap[int(char)-int('A')] = true
				}
			} else {
				if prim[int(char)-int('a')] {
					sec[int(char)-int('a')] = true
				}
			}

		}

		prim = sec       //copy values in into prim
		PrimCap = secCap //copy values in into prim
		sec = sec[:0]    //empty sec
		secCap = sec[:0] //empty sec

	}

	//convert ascii back to char:
	for k := 0; k < 53; k++ {
		if prim[k] {
			commonChar = string(k + 97)
		}
		if PrimCap[k] {
			commonChar = string(k + 65)
		}
	}

	return commonChar

}

func populateBoolArrayTrue(arr []bool, n int) []bool {
	for i := 0; i < n; i++ {
		arr = append(arr, true)
	}
	return arr
}

func populateBoolArrayFalse(arr []bool, n int) []bool {
	for i := 0; i < n; i++ {
		arr = append(arr, false)
	}
	return arr
}

func calcTotal(commonChars []string) int {
	total := 0

	for i := 0; i < len(commonChars); i++ {
		total += charMap(commonChars[i])
	}

	return total
}

func charMap(commonChar string) int {
	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	m["d"] = 4
	m["e"] = 5
	m["f"] = 6
	m["g"] = 7
	m["h"] = 8
	m["i"] = 9
	m["j"] = 10
	m["k"] = 11
	m["l"] = 12
	m["n"] = 13
	m["m"] = 14
	m["o"] = 15
	m["p"] = 16
	m["q"] = 17
	m["r"] = 18
	m["s"] = 19
	m["t"] = 20
	m["u"] = 21
	m["v"] = 22
	m["w"] = 23
	m["x"] = 24
	m["y"] = 25
	m["z"] = 26

	m["A"] = 27
	m["B"] = 28
	m["C"] = 29
	m["D"] = 30
	m["E"] = 31
	m["F"] = 32
	m["G"] = 33
	m["H"] = 34
	m["I"] = 35
	m["J"] = 36
	m["K"] = 37
	m["L"] = 38
	m["N"] = 39
	m["M"] = 40
	m["O"] = 41
	m["P"] = 42
	m["Q"] = 43
	m["R"] = 44
	m["S"] = 45
	m["T"] = 46
	m["U"] = 47
	m["V"] = 48
	m["W"] = 49
	m["X"] = 50
	m["Y"] = 51
	m["Z"] = 52

	value := m[commonChar]
	return value
}

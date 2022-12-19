package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func OpenFileForReadAndParse(path string) []string {
	puzzleData := []string{}

	file, err := os.Open(path)

	if err != nil {
		fmt.Printf("Could not open the file due to this  %s error \n", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		puzzleData = append(puzzleData, fileScanner.Text())
	}

	if err = file.Close(); err != nil {
		fmt.Printf("Could not open the file due to this  %s error \n", err)
	}
	return puzzleData
}

func Day_eight_file_parser(path string) [][]int {
	//populate matrix:
	var matrix [][]int
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("------------------- FILE NOT FOUND ----------")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileRow := strings.Split(scanner.Text(), "")
		// convert to int and append to list:
		intRow := []int{}
		for i := 0; i < len(fileRow); i++ {
			val, err := strconv.Atoi(fileRow[i])
			if err != nil {
				fmt.Println("------------------- Conversion Error ----------")
			}
			intRow = append(intRow, val)
		}
		matrix = append(matrix, intRow)
	}

	return matrix
}

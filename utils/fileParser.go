package utils

import (
	"bufio"
	"fmt"
	"os"
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


func Day_five_file_parser(path string){
	 
}
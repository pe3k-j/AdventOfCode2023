package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createEmptyLine(length int) string {
	var output string
	for i := 0; i < length; i++ {
		output += "."
	}
	return output
}

func findNumbers(input []string) []int {
	var output []int
	for i, line := range input {
		var tempNumber string
		for j, rune := range line {
			_, err := strconv.Atoi(string(rune))
			if err == nil {
				tempNumber += string(rune)
			} else if len(tempNumber) != 0 {
				val, err := strconv.Atoi(tempNumber)
				if err != nil {
					fmt.Println(err)
				} else {
					// checks surrounding runes for characters thet are not "."
					surroundingRunes := input[i-1][j-len(tempNumber)-1 : j+1]
					surroundingRunes += string(line[j-len(tempNumber)-1])
					surroundingRunes += string(line[j])
					surroundingRunes += input[i+1][j-len(tempNumber)-1 : j+1]
					surroundingRunes = strings.ReplaceAll(surroundingRunes, ".", "")
					if len(surroundingRunes) > 0 {
						output = append(output, val)
					}
				}
				tempNumber = ""
			}
		}
	}
	return output

}

func main() {

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	input := []string{createEmptyLine(142)}
	for fileScanner.Scan() {
		input = append(input, "."+fileScanner.Text()+".")
	}
	input = append(input, createEmptyLine(142))
	numbers := findNumbers(input)
	var sum int
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum)

}

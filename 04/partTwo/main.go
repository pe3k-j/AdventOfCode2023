package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func findNums(input string) []int {
	var output []int

	var tempNumber string
	for _, rune := range input + " " {
		_, err := strconv.Atoi(string(rune))
		if err != nil {
			if len(tempNumber) > 0 {
				val, _ := strconv.Atoi(tempNumber)
				output = append(output, val)
				tempNumber = ""
			}
		} else {
			tempNumber += string(rune)
		}
	}
	return output
}

func countMatches(line string) int {
	var output int
	parts := strings.FieldsFunc(line, func(r rune) bool {
		return strings.ContainsRune(":|", r)
	})
	fmt.Print(parts[0], ":\n")
	fmt.Print(parts[1], "\n")
	fmt.Print(parts[2], "\n")
	winningNumbers := findNums(parts[1])
	myNumbers := findNums(parts[2])
	for _, number := range myNumbers {
		if slices.Contains(winningNumbers, number) {
			output++
			fmt.Print("Number: ", number, " Output: ", output, ":\n")
		}
	}
	fmt.Println("==============================")

	return output
}

func main() {

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	totalSum := 0
	var matches []int
	for fileScanner.Scan() {
		matches = append(matches, countMatches(fileScanner.Text()))
	}
	numberOfCards := make([]int, len(matches))
	for i := range numberOfCards {
		numberOfCards[i] = 1
	}
	for i, number := range matches {
		j := number
		for j > 0 {

			numberOfCards[i+j] += numberOfCards[i]
			j--
		}
		fmt.Print(i, ": ", numberOfCards, "\n")
	}
	for i := range numberOfCards {
		totalSum += numberOfCards[i]
	}
	fmt.Println("==============================")
	fmt.Println(matches)
	fmt.Println(numberOfCards)
	fmt.Print("Total: ", totalSum, "\n")
}

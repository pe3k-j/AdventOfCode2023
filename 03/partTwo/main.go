package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func createEmptyLine(length int) string {
	var output string
	for i := 0; i < length; i++ {
		output += "."
	}
	return output
}

func findNums(input [3]string) []int {
	var output []int
	for _, line := range input {
		var tempNumber string
		for j, rune := range line {

			_, err := strconv.Atoi(string(rune))
			if err == nil {
				tempNumber += string(rune)
			} else if len(tempNumber) != 0 {
				x := (j - 1) - (len(tempNumber) - 1)
				y := j - 1
				if (x == 2 || x == 3 || x == 4) || (y == 2 || y == 3 || y == 4) {
					val, _ := strconv.Atoi(tempNumber)
					output = append(output, val)
					tempNumber = ""
				} else {
					tempNumber = ""
				}

			}
		}
	}
	return output
}

func findGears(input []string) []int {
	var output []int
	for i, line := range input {
		for j, rune := range line {
			if rune == 42 {
				area := [3]string{
					input[i-1][j-3:j+4] + ".",
					input[i][j-3:j+4] + ".",
					input[i+1][j-3:j+4] + ".",
				}
				numbers := findNums(area)
				if len(numbers) == 2 {
					fmt.Println("===================")
					fmt.Print(area[0], "\n")
					fmt.Print(area[1], "\n")
					fmt.Print(area[2], "\n")
					output = append(output, numbers[0]*numbers[1])

				}
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
	numbers := findGears(input)
	var sum int
	for _, number := range numbers {
		sum += number
	}
	fmt.Print(sum, "\n")

}

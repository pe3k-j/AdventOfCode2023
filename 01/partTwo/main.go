package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func addFirstAndLastNumber(line string) int {
	var numbersInLine []int
	potential_num := ""

	for _, char := range line {
		potential_num += string(char)

		// add normal number
		val, err := strconv.Atoi(string(char))
		if err == nil {
			numbersInLine = append(numbersInLine, val)
		}

		// add text number
		for textFormat, intFormat := range numbers {

			// checks if current sybol is not an end of a text number
			if strings.HasSuffix(potential_num, textFormat) {
				numbersInLine = append(numbersInLine, intFormat)
			}
		}
	}

	return numbersInLine[0]*10 + numbersInLine[len(numbersInLine)-1]
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var finalSum int
	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		currentSum := addFirstAndLastNumber(currentLine)
		finalSum += currentSum
	}
	fmt.Print(finalSum, "\n")
}

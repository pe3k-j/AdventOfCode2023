package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func addFirstAndLastNumber(line string) int {
	var firstNum, lastNum, numCount int
	for i := 0; i < len(line); i++ {
		val, err := strconv.Atoi(string(line[i]))
		if err == nil {
			if numCount == 0 {
				firstNum = val
			}
			lastNum = val
			numCount++
		}
	}
	val, err := strconv.Atoi(strconv.Itoa(firstNum) + strconv.Itoa(lastNum))
	if err != nil {
		fmt.Println(err)
		return 0
	} else {
		return val
	}
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

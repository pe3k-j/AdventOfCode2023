package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func main() {

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	times := findNums(fileScanner.Text())
	fileScanner.Scan()
	distances := findNums(fileScanner.Text())

	var winsCount []int
	for i := 0; i < len(times); i++ {
		fmt.Println("=============")
		fmt.Println(times[i], distances[i])
		var winCount int
		for j := 0; j < times[i]; j++ {
			val := j * (times[i] - j)
			if val > distances[i] {
				winCount++
			}
		}
		fmt.Println(winCount)
		winsCount = append(winsCount, winCount)
	}

	fmt.Println(winsCount)
	finalValue := 1
	for _, num := range winsCount {
		finalValue *= num
	}

	fmt.Print("FinalValue: ", finalValue, "\n")
}

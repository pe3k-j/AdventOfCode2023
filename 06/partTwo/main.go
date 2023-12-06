package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func findNum(input string) int {
	var output int
	var tempOutput string

	for _, rune := range input + " " {
		_, err := strconv.Atoi(string(rune))
		if err != nil {
		} else {
			tempOutput += string(rune)
		}
	}
	val, err := strconv.Atoi(tempOutput)
	if err != nil {
		fmt.Println(err)
	} else {
		output = val
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
	time := findNum(fileScanner.Text())
	fileScanner.Scan()
	distance := findNum(fileScanner.Text())

	fmt.Println("=============")
	fmt.Println(time, distance)
	var winCount int
	for i := 0; i < time; i++ {
		val := i * (time - i)
		if val > distance {
			winCount++
		}
	}
	fmt.Println(winCount)

	fmt.Print("FinalValue: ", winCount, "\n")
}

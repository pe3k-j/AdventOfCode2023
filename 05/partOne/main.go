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

func findLocaions(input []string) []int {
	output := findNums(input[0])
	fmt.Println(output)
	var mapNums [][]int

	input = append(input, input[1])
	i := 3
	for i < len(input) {
		if input[i] == input[1] {

			fmt.Println(mapNums)

			var tempOutput []int
			for _, val := range output {
				convertedVal := val
				for _, mapNum := range mapNums {
					if mapNum[1] <= val && val < (mapNum[1]+mapNum[2]) {
						convertedVal = mapNum[0] + (val - mapNum[1])
					}
				}
				fmt.Print(val, " ==> ", convertedVal, "\n")
				tempOutput = append(tempOutput, convertedVal)
			}
			output = tempOutput
			mapNums = [][]int{}
			i += 2
		} else {
			mapNums = append(mapNums, findNums(input[i]))
			i++
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

	var input []string
	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}
	locations := findLocaions(input)
	fmt.Print("Locations: ", locations, "\n")
	minLocation := locations[0]
	for _, val := range locations {
		if val < minLocation {
			minLocation = val
		}
	}
	fmt.Print("Min: ", minLocation, "\n")
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cubePower(line string) int {
	r, g, b := 0, 0, 0
	words := strings.FieldsFunc(line, func(r rune) bool {
		return strings.ContainsRune(":,;", r)
	})
	for _, word := range words[1:] {
		i := 1
		for {
			_, err := strconv.Atoi(string(word[i]))
			if err != nil {
				break
			} else {
				i++
			}
		}
		val, err := strconv.Atoi(word[1:i])
		if err != nil {
			fmt.Println(err)
		} else {
			switch string(word[i+1]) {
			case "r":
				if val > r {
					r = val
				}
			case "g":
				if val > g {
					g = val
				}
			case "b":
				if val > b {
					b = val
				}
			}
		}

	}
	return r*g*b
}

func main() {

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	totalSum := 0
	for fileScanner.Scan() {
		totalSum += cubePower(fileScanner.Text())
	}
	fmt.Print("Total: ", totalSum, "\n")
}

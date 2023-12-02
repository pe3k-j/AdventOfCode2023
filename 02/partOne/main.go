package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isPossible(line string) bool {
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
				if val > 12 {
					return false
				}
			case "g":
				if val > 13 {
					return false
				}
			case "b":
				if val > 14 {
					return false
				}
			}
		}

	}
	return true
}

func main() {

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	i := 0
	totalSum := 0
	for fileScanner.Scan() {
		i++
		if isPossible(fileScanner.Text()) {
			totalSum += i
			fmt.Print("Total: ", totalSum, "\n")
			fmt.Println(fileScanner.Text())
		}
	}
	fmt.Print("Total: ", totalSum, "\n")
}

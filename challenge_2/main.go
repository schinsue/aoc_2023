package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func convertToNumber(number string) string {
	switch number {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return number
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var number string
	var sum = 0
	numberWords := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		var lowIndex = 1000
		var highIndex = 0
		var lastWord string
		var firstWord string

		for _, word := range numberWords {
			firstIndex := strings.Index(line, word)
			lastIndex := strings.LastIndex(line, word)

			if firstIndex != -1 {
				if firstIndex <= lowIndex {
					lowIndex = firstIndex
					firstWord = word
				}
			}

			if lastIndex != -1 {
				if lastIndex >= highIndex {
					highIndex = lastIndex
					lastWord = word
				}
			}
		}

		firstWord = convertToNumber(firstWord)
		lastWord = convertToNumber(lastWord)

		number = firstWord + lastWord

		n, err := strconv.Atoi(number)

		//check if error occured
		if err != nil {
			log.Fatal(err)
		} else {
			sum += n
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

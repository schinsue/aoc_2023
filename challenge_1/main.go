package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const SUBSET = "0123456789"

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var number string
	var sum = 0

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		firstIndex := strings.IndexAny(line, SUBSET)
		lastIndex := strings.LastIndexAny(line, SUBSET)

		if firstIndex != -1 {
			number = line[firstIndex : firstIndex+1]
		} else {
			log.Fatal("Could not parse: " + line)
		}

		if lastIndex != -1 {
			number = number + line[lastIndex:lastIndex+1]
		} else {
			log.Fatal("Could not parse: " + line)
		}

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

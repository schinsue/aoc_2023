package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func LineParser(line string) (string, []string) {
	gameSplit := strings.Split(line, ": ")
	gameNumber := strings.Replace(gameSplit[0], "Game ", "", -1)
	games := strings.Split(gameSplit[1], "; ")

	return gameNumber, games
}

func isOverColorLimit(colorString string, colorLimit int, color string) bool {
	stringValue := strings.Split(colorString, " "+color)
	n, err := strconv.Atoi(stringValue[0])

	if err != nil {
		log.Fatal(err)
	}

	if n > colorLimit {
		return true
	}

	return false
}

func isValidGame(games []string, limits map[string]int) bool {

	for _, game := range games {
		splitByColor := strings.Split(game, ", ")
		for _, str := range splitByColor {
			switch {
			case strings.Contains(str, "red"):
				if isOverColorLimit(str, limits["red"], "red") {
					return false
				}
			case strings.Contains(str, "blue"):
				if isOverColorLimit(str, limits["blue"], "blue") {
					return false
				}
			case strings.Contains(str, "green"):
				if isOverColorLimit(str, limits["green"], "green") {
					return false
				}
			}
		}
	}

	return true
}

func main() {
	limits := map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var sum = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		n, games := LineParser(line)

		if isValidGame(games, limits) {
			nr, err := strconv.Atoi(n)

			if err != nil {
				log.Fatal(err)
			} else {
				sum += nr
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

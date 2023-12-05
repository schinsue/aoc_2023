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

func isValidGame(games []string, limits map[string]int) bool {

	for _, game := range games {
		splitByColor := strings.Split(game, ", ")
		for _, y := range splitByColor {
			if strings.Contains(y, "red") {
				stringValue := strings.Split(y, " red")
				n, err := strconv.Atoi(stringValue[0])

				if err != nil {
					log.Fatal(err)
				} else {
					if n > limits["red"] {
						return false
					}
				}
			} else if strings.Contains(y, "blue") {
				if strings.Contains(y, "blue") {
					stringValue := strings.Split(y, " blue")
					n, err := strconv.Atoi(stringValue[0])

					if err != nil {
						log.Fatal(err)
					} else {
						if n > limits["blue"] {
							return false
						}
					}
				}
			} else if strings.Contains(y, "green") {
				if strings.Contains(y, "green") {
					stringValue := strings.Split(y, " green")
					n, err := strconv.Atoi(stringValue[0])

					if err != nil {
						log.Fatal(err)
					} else {
						if n > limits["green"] {
							return false
						}
					}
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

package main

import (
	"reflect"
	"strings"
	"testing"
)

const TEXT_INPUT = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func TestLineParser(t *testing.T) {
	lines := strings.Split(TEXT_INPUT, "\n")

	nr, games := LineParser(lines[0])

	if nr != "1" || !reflect.DeepEqual(games, []string{"3 blue, 4 red", "1 red, 2 green, 6 blue", "2 green"}) {
		t.Errorf("LineParser failed on line 1")
	}

	nr, games = LineParser(lines[1])

	if nr != "2" || !reflect.DeepEqual(games, []string{"1 blue, 2 green", "3 green, 4 blue, 1 red", "1 green, 1 blue"}) {
		t.Errorf("LineParser failed on line 2")
	}

	nr, games = LineParser(lines[2])

	if nr != "3" || !reflect.DeepEqual(games, []string{"8 green, 6 blue, 20 red", "5 blue, 4 red, 13 green", "5 green, 1 red"}) {
		t.Errorf("LineParser failed on line 3")
	}

	nr, games = LineParser(lines[3])

	if nr != "4" || !reflect.DeepEqual(games, []string{"1 green, 3 red, 6 blue", "3 green, 6 red", "3 green, 15 blue, 14 red"}) {
		t.Errorf("LineParser failed on line 4")
	}

	nr, games = LineParser(lines[4])

	if nr != "5" || !reflect.DeepEqual(games, []string{"6 red, 1 blue, 3 green", "2 blue, 1 red, 2 green"}) {
		t.Errorf("LineParser failed on line 5")
	}
}

func TestIsValidGame(t *testing.T) {
	limits := map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}
	lines := strings.Split(TEXT_INPUT, "\n")

	for _, line := range lines {
		nr, games := LineParser(line)
		isValid := isValidGame(games, limits)
		if nr == "1" && !isValid {
			t.Errorf("Game 1 should be valid")
		} else if nr == "2" && !isValid {
			t.Errorf("Game 2 should be valid")
		} else if nr == "3" && isValid {
			t.Errorf("Game 3 should be invalid")
		} else if nr == "4" && isValid {
			t.Errorf("Game 4 should be invalid")
		} else if nr == "5" && !isValid {
			t.Errorf("Game 5 should be valid")
		}
	}
}

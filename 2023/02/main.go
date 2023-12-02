package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input
var input string

func init() {
	// Remove trailing newline
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("No input file")
	}
}

func main() {
	pre, parsedData, post := time.Now(), parseInput(input), time.Now()
	fmt.Printf("Input parsed in %s\n", post.Sub(pre))

	pre, ans, post := time.Now(), part1(parsedData), time.Now()
	expectedAns := 2512
	fmt.Printf("Part1 answer: %d, in %s, pass: %t\n", ans, post.Sub(pre), ans == expectedAns)

	pre, ans, post = time.Now(), part2(parsedData), time.Now()
	expectedAns = 0
	fmt.Printf("Part2 answer: %d, in %s, pass: %t\n", ans, post.Sub(pre), ans == expectedAns)
}

type set = map[string]int

type game struct {
	id   int
	sets []set
}

type games = []game

var colorsMaxVal = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func parseInput(input string) games {
	lines := strings.Split(input, "\n")

	allGames := games{}
	for _, line := range lines {
		fields := strings.Fields(line)[1:]

		gameId := 0
		sets := []set{}
		s := set{}
		i := 0
		for i < len(fields) {
			// It's the game ID
			if i == 0 {
				gId, err := strconv.Atoi(strings.TrimSuffix(fields[i], ":"))
				if err != nil {
					panic(fmt.Errorf("cannot parse input game id: %s", err.Error()))
				}
				gameId = gId
				i++
				continue
			}

			// fields[i] is the number and fields[i+1] is the color
			val, err := strconv.Atoi(fields[i])
			if err != nil {
				panic(fmt.Errorf("cannot parse input set value: %s", err.Error()))
			}
			color := strings.TrimSuffix(strings.TrimSuffix(fields[i+1], ","), ";")
			s[color] = val

			isLastColorInSet := strings.Contains(fields[i+1], ";")
			if isLastColorInSet {
				sets = append(sets, s)
				s = set{}
			}

			i += 2
		}

		if len(s) > 0 {
			sets = append(sets, s)
		}

		g := game{
			id:   gameId,
			sets: sets,
		}

		allGames = append(allGames, g)
	}

	return allGames
}

func part1(parsedData games) int {
	gameIDsSum := 0
	for _, g := range parsedData {
		isGameValid := true
		for _, s := range g.sets {
			for k, v := range s {
				if v > colorsMaxVal[k] {
					isGameValid = false
					break
				}
			}
			if !isGameValid {
				break
			}
		}

		if isGameValid {
			gameIDsSum += g.id
		}
	}

	return gameIDsSum
}

func part2(parsedData games) int {
	return 0
}

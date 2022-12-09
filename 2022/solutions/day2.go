package solutions

import (
	"fmt"
	"strings"

	"github.com/whcass/Advent-of-Code/2022/helpers"
)

type Day2 struct {
	Title string
}

func NewDay2() *Day2 {
	return &Day2{Title: "Day 2"}
}

func (d Day2) GetTitle() string {
	return d.Title
}

func (d Day2) Run() {
	input := helpers.GetInput("2")
	sliceData := strings.Split(input, "\n")
	d.Part1(sliceData)
	d.Part2(sliceData)
}

func (d Day2) Part1(sliceData []string) {
	total := 0
	for i := 0; i < len(sliceData); i++ {
		round := strings.Split(sliceData[i], " ")
		result := parseRound(round[0], round[1])
		total += result
	}
	fmt.Printf("[*] Part 1: %d\n", total)
}

func (d Day2) Part2(sliceData []string) {
	total := 0
	for i := 0; i < len(sliceData); i++ {
		round := strings.Split(sliceData[i], " ")
		result := parseRoundPart2(round[0], round[1])
		total += result
	}
	fmt.Printf("[*] Part 2: %d\n", total)
}

func getWin(move string) string {
	if move == "A" {
		return "B"
	} else if move == "B" {
		return "C"
	} else {
		return "A"
	}
}

func getLose(move string) string {
	if move == "A" {
		return "C"
	} else if move == "C" {
		return "B"
	} else {
		return "A"
	}
}

func parseMove(move string) int {
	switch move {
	case "A", "X":
		return 1
	case "B", "Y":
		return 2
	case "C", "Z":
		return 3
	default:
		return 0
	}
}

func parseRound(aMove string, bMove string) int {
	A := parseMove(aMove)
	B := parseMove(bMove)
	if A == B { // A draw
		return B + 3
	} else if (A == 3 && B == 1) || (A == 2 && B == 3) || (A == 1 && B == 2) {
		return B + 6
	} else {
		return B
	}
}

func parseRoundPart2(aMove string, bMove string) int {
	if bMove == "X" { //lose
		return parseMove(getLose(aMove))
	} else if bMove == "Y" { // draw
		return parseMove(aMove) + 3
	} else { // win
		return parseMove(getWin(aMove)) + 6
	}
}

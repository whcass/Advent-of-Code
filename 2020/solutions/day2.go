package solutions

import (
	"fmt"
	"github.com/whcass/Advent-of-Code/2020/helpers"
	"strconv"
	"strings"
)

type Day2 struct {
	Title string
}

func NewDay2() *Day2 {
	return &Day2{Title: "Day 2"}
}

func (d Day2) Run() {
	input := helpers.GetInput("day2.txt")
	inputSplit := strings.Split(input, "\r\n")
	d.Part1(inputSplit)
	d.Part2(inputSplit)
}

func (d Day2) GetTitle() string {
	return d.Title
}

func (d Day2) Part1(input []string) {
	valid := 0
	for _, rule := range input {
		if rule == "" {
			break
		}
		lower, upper, letter, pass := ruleSplit(rule)
		count := strings.Count(pass, letter)
		if count <= upper && count >= lower {
			valid++
		}
	}
	fmt.Printf("[*] Valid Passwords: %d\n", valid)
}

func (d Day2) Part2(input []string) {
	valid := 0
	for _, rule := range input {
		if rule == "" {
			break
		}
		pos1, pos2, letter, pass := ruleSplit(rule)
		pos1Found := false
		pos2Found := false

		if string(pass[pos1-1]) == letter {
			pos1Found = true
		}
		if string(pass[pos2-1]) == letter {
			pos2Found = true
		}
		if pos1Found && pos2Found {

		} else if pos1Found || pos2Found {
			valid++
		}
	}

	fmt.Printf("[*] Valid Passwords: %d\n", valid)
}

func ruleSplit(rule string) (int, int, string, string) {
	ruleSplit := strings.Split(rule, " ")
	num := ruleSplit[0]
	numSplit := strings.Split(num, "-")
	lower, _ := strconv.Atoi(numSplit[0])
	upper, _ := strconv.Atoi(numSplit[1])
	letter := string(ruleSplit[1][0])
	pass := ruleSplit[2]
	return lower, upper, letter, pass
}

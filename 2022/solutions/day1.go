package solutions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/whcass/Advent-of-Code/2022/helpers"
)

type Day1 struct {
	Title string
}

func NewDay1() *Day1 {
	return &Day1{Title: "Day 1"}
}

func (d Day1) GetTitle() string {
	return d.Title
}

func (d Day1) Run() {
	input := helpers.GetInput("day1.txt")
	sliceData := strings.Split(input, "\n")
	d.Part1(sliceData)
}

func (d Day1) Part1(sliceData []string) {
	runningTotal := 0
	highest := 0
	for i := 0; i < len(sliceData); i++ {
		if sliceData[i] == "" {
			if runningTotal > highest {
				highest = runningTotal
			}
			runningTotal = 0
		} else {
			val, _ := strconv.Atoi(sliceData[i])
			runningTotal += val
		}
	}
	fmt.Printf("[*] Part 1: %d\n", highest)
}

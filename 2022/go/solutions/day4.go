package solutions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/whcass/Advent-of-Code/2022/helpers"
)

type Day4 struct {
	Title string
}

func NewDay4() *Day4 {
	return &Day4{Title: "Day 4"}
}

func (d Day4) GetTitle() string {
	return d.Title
}

func (d Day4) Run() {
	input := helpers.GetInput("4")
	sliceData := helpers.SplitInput(input)
	d.Part1(sliceData)
}

func (d Day4) Part1(data []string) {
	containedCount := 0
	overlappedCount := 0
	for i := 0; i < len(data); i++ {
		ranges := strings.Split(data[i], ",")
		rangeA := strings.Split(ranges[0], "-")
		rangeB := strings.Split(ranges[1], "-")
		if isContained(rangeA, rangeB) {
			containedCount++
		}

		if isOverlapped(rangeA, rangeB) {
			overlappedCount++
		}
	}

	helpers.PrintResult(fmt.Sprintf("Part 1: %d", containedCount))
	helpers.PrintResult(fmt.Sprintf("Part 2: %d", overlappedCount))
}

func getRange(rangeA []string, rangeB []string) (int, int, int, int) {
	aLow, _ := strconv.Atoi(rangeA[0])
	aHigh, _ := strconv.Atoi(rangeA[1])
	bLow, _ := strconv.Atoi(rangeB[0])
	bHigh, _ := strconv.Atoi(rangeB[1])

	return aLow, aHigh, bLow, bHigh
}

func isContained(rangeA []string, rangeB []string) bool {
	aLow, aHigh, bLow, bHigh := getRange(rangeA, rangeB)
	if (aLow <= bLow) && (aHigh >= bHigh) {
		return true
	} else if (bLow <= aLow) && (bHigh >= aHigh) {
		return true
	} else {
		return false
	}
}

func isOverlapped(rangeA []string, rangeB []string) bool {
	aLow, aHigh, bLow, bHigh := getRange(rangeA, rangeB)
	if aLow <= bHigh && bLow <= aHigh {
		return true
	} else {
		return false
	}
}

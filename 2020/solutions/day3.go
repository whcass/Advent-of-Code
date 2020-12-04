package solutions

import (
	"fmt"
	"github.com/whcass/Advent-of-Code/2020/helpers"
	"strings"
)

type Day3 struct {
	Title string
}

func NewDay3() *Day3 {
	return &Day3{Title: "Day 3"}
}

func (d Day3) Run() {
	input := helpers.GetInput("day3.txt")
	slopeMap := generateSlopeMap(input)
	d.Part1(slopeMap)
	d.Part2(slopeMap)
}

func (d Day3) GetTitle() string {
	return d.Title
}

func (d Day3) Part1(slopeMap [][]string) {
	treeCount := checkSlope(slopeMap, 3, 1)

	fmt.Printf("[*] Trees Encountered: %d\n", treeCount)
}

func (d Day3) Part2(slopeMap [][]string) {
	slopes := [5][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	slopeMult := 1
	for _, slope := range slopes {
		treeCount := checkSlope(slopeMap, slope[0], slope[1])
		//fmt.Printf("[+] Trees encountered: %d\n", treeCount)
		slopeMult *= treeCount
	}

	fmt.Printf("[*] Total Tree Mult: %d\n", slopeMult)
}

func checkSlope(slopeMap [][]string, xInc int, yInc int) int {
	x := 0
	y := 0
	treeCount := 0
	for true {
		if y == len(slopeMap)-1 {
			break
		}
		x = (x + xInc) % len(slopeMap[0])
		y += yInc
		if slopeMap[y][x] == "#" {
			treeCount++
		}
	}
	return treeCount
}

func generateSlopeMap(input string) [][]string {
	inputSplit := strings.Split(input, "\r\n")
	slopeMap := make([][]string, len(inputSplit)-1)

	for i := 0; i < len(inputSplit)-1; i++ {
		slopeSplit := strings.Split(strings.TrimSpace(inputSplit[i]), "")
		slopeMap[i] = make([]string, len(slopeSplit))
		slopeMap[i] = slopeSplit
	}

	return slopeMap
}

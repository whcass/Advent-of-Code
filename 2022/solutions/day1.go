package solutions

import (
	"fmt"
	"sort"
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
	input := helpers.GetInput("1")
	sliceData := strings.Split(input, "\n")
	d.Part1(sliceData)
}

func (d Day1) Part1(sliceData []string) {
	runningTotal := 0
	var calories []int
	for i := 0; i < len(sliceData); i++ {
		if sliceData[i] == "" {
			calories = append(calories, runningTotal)
			runningTotal = 0
		} else {
			val, _ := strconv.Atoi(sliceData[i])
			runningTotal += val
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	helpers.PrintResult(fmt.Sprintf("Part 1: %d", calories[0]))
	helpers.PrintResult(fmt.Sprintf("Part 2: %d", calories[0]+calories[1]+calories[2]))
}

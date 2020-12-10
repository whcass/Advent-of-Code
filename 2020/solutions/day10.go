package solutions

import (
	"fmt"
	"github.com/whcass/Advent-of-Code/2020/helpers"
	"sort"
	"strings"
)

type Day10 struct {
	Title string
}

func NewDay10() *Day10 {
	return &Day10{Title: "Day 10"}
}

func (d Day10) Run() {
	input := helpers.GetInput("day10.txt")
	inputData := helpers.PrepareInput(strings.Split(input, "\r\n"))

	d.part1(inputData)
}

func (d Day10) GetTitle() string {
	return d.Title
}

func (d Day10) part1(data []int) {
	sort.Ints(data)
	diffOfOne := 1
	diffOfThree := 1
	for i, adapter := range data {
		if i >= len(data)-1 {
			break
		}
		diff := data[i+1] - adapter
		if diff == 1 {
			diffOfOne++
		} else if diff == 3 {
			diffOfThree++
		} else {
			fmt.Println(diff)
		}
	}
	ans := diffOfThree * diffOfOne
	fmt.Printf("[*] Jolt Diff: %d\n", ans)
}

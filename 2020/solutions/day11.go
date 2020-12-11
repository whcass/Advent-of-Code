package solutions

import (
	"fmt"
	"github.com/whcass/Advent-of-Code/2020/helpers"
	"strings"
)

type Day11 struct {
	Title string
	grid  [][]string
}

func NewDay11() *Day11 {
	return &Day11{Title: "Day 11"}
}

func (d Day11) Run() {
	input := helpers.GetInput("day11.txt")
	inputData := strings.Split(input, "\r\n")
	d.grid = d.makeGrid(inputData)
	d.part1()
}

func (d Day11) part1() {
	fmt.Println(d.grid)
}

func (d Day11) GetTitle() string {
	return d.Title
}

func (d Day11) makeGrid(input []string) [][]string {
	grid := make([][]string, len(input)-1)

	for i := 0; i < len(input)-1; i++ {
		split := strings.Split(strings.TrimSpace(input[i]), "")
		grid[i] = make([]string, len(split))
		grid[i] = split
	}

	return grid
}

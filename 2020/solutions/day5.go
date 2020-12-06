package solutions

import (
	"fmt"
	"github.com/whcass/Advent-of-Code/2020/helpers"
	"math"
	"strings"
)

type Day5 struct {
	Title string
}

func NewDay5() *Day5 {
	return &Day5{Title: "Day 5"}
}

func (d Day5) Run() {
	input := helpers.GetInput("day5.txt")
	inputData := strings.Split(input, "\r\n")
	d.part1(inputData)
	d.part2(inputData)

}

func (d Day5) part2(inputData []string) {
	ids := make(map[int]int)
	for i := 127; i < 900; i++ {
		ids[i] = i
	}
	for _, boardPass := range inputData {
		seat := d.getSeat(boardPass)
		delete(ids, seat)
	}
	fmt.Printf("[*] Boarding Pass ID: %d\n", ids)
}

func (d Day5) part1(inputData []string) {
	highest := 0
	for _, boardPass := range inputData {
		seat := d.getSeat(boardPass)
		if seat > highest {
			highest = seat
		}
	}

	fmt.Printf("[*] Highest Seat: %d\n", highest)
}

func (d Day5) getSeat(boardPass string) int {
	rowEnd := 127
	rowStart := 0
	colEnd := 7
	colStart := 0
	boardPassSplit := strings.Split(boardPass, "")

	for i, location := range boardPassSplit {
		if i < 7 {
			if location == "F" {
				rowEnd = (rowStart + rowEnd) / 2
			} else if location == "B" {
				rowStart = int(math.Ceil(float64((rowStart+rowEnd)/2.0))) + 1
			}
		} else {
			if location == "L" {
				colEnd = (colStart + colEnd) / 2
			} else if location == "R" {
				colStart = int(math.Ceil(float64((colStart+colEnd)/2.0))) + 1
			}
		}
	}

	seat := rowEnd*8 + colEnd
	return seat
}

func (d Day5) GetTitle() string {
	return d.Title
}

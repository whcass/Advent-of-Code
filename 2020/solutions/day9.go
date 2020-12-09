package solutions

import (
	"fmt"
	"github.com/whcass/Advent-of-Code/2020/helpers"
	"strings"
)

type Day9 struct {
	Title string
}

func (d Day9) Run() {
	input := helpers.GetInput("day9.txt")
	inputData := helpers.PrepareInput(strings.Split(input, "\r\n"))

	d.part1(inputData)
	d.part2(inputData)
}

func (d Day9) countSeries(in []int) int {
	count := 0
	for _, i := range in {
		count += i
	}

	return count
}

func (d Day9) minMax(in []int) (int, int) {
	min := 1000000000000
	max := 0
	for _, i := range in {
		if i > max {
			max = i
		}

		if i < min {
			min = i
		}
	}

	return min, max
}

func (d Day9) part2(inputData []int) {
	goal := 167829540
	notFound := true
	length := 2
	for notFound {
		run := true
		i := 0
		for run {
			series := inputData[i : i+length]
			total := d.countSeries(series)
			if total > goal {
				run = false
			} else if total == goal {
				min, max := d.minMax(series)
				total := min + max
				fmt.Printf("[*] Weakness: %d\n", total)
				return
			}
			i++
		}
		length++
	}
}

func (d Day9) part1(inputData []int) {
	for i := 25; i < len(inputData); i++ {
		match := false
		for j := i - 25; j < i; j++ {
			for k := j; k < i; k++ {
				a := inputData[k]
				b := inputData[j]
				if a != b {
					test := a + b
					c := inputData[i]
					if test == c {
						match = true
					}
				}
			}
		}
		if !match {
			fmt.Printf("[*] Number: %d\n", inputData[i])
			break
		}
	}
}

func (d Day9) GetTitle() string {
	return d.Title
}

func NewDay9() *Day9 {
	return &Day9{Title: "Day 9"}
}

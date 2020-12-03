package solutions

import (
	"fmt"
	"github.com/whcass/Advent-of-Code/2020/helpers"
	"strconv"
	"strings"
)

type Day1 struct {
	Title string
}

func NewDay1() *Day1 {
	return &Day1{Title: "Day 1"}
}

func (d Day1) Run() {
	input := helpers.GetInput("day1.txt")
	helpers.PrintInfo(d.Title)
	sliceData := strings.Split(input, "\r\n")
	d.Part1(sliceData)
	d.Part2(sliceData)
}

func (d Day1) Part2(sliceData []string) {
	for i := 0; i < len(sliceData); i++ {
		for j := 0; j < len(sliceData); j++ {
			for k := 0; k < len(sliceData); k++ {
				a, _ := strconv.Atoi(sliceData[i])
				b, _ := strconv.Atoi(sliceData[j])
				c, _ := strconv.Atoi(sliceData[k])
				d := a + b + c

				if d == 2020 {
					fmt.Printf("[*] Part 2: %d\n", a*b*c)
					return
				}
			}
		}
	}
}

func (d Day1) Part1(sliceData []string) {
	for i := 0; i < len(sliceData); i++ {
		for j := i; j < len(sliceData); j++ {
			a, _ := strconv.Atoi(sliceData[i])
			b, _ := strconv.Atoi(sliceData[j])
			c := a + b
			if c == 2020 {
				fmt.Printf("[*] Part 1: %d\n", a*b)
				return
			}
		}
	}
}

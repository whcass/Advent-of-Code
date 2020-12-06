package solutions

import (
	"fmt"
	"github.com/whcass/Advent-of-Code/2020/helpers"
	"strings"
)

type Day6 struct {
	Title string
}

func (d Day6) Run() {
	input := helpers.GetInput("day6.txt")
	inputData := strings.Split(input, "\r\n")
	d.part1(inputData)
	d.part2(inputData)
}

func (d Day6) part1(inputData []string) {
	groupMap := make(map[string]int)
	count := 0
	for _, group := range inputData {
		if group == "" {
			count += len(groupMap)
			groupMap = make(map[string]int)
		} else {
			answers := strings.Split(group, "")
			for _, answer := range answers {
				groupMap[answer]++
			}
		}
	}
	fmt.Printf("[*] Answer Count: %d\n", count)
}

func (d Day6) part2(inputData []string) {
	groupMap := make(map[string]int)
	questionCount := 0
	groupSize := 0
	for _, group := range inputData {
		if group == "" {
			for _, count := range groupMap {
				if count == groupSize {
					questionCount++
				}
			}
			groupSize = 0
			groupMap = make(map[string]int)
		} else {
			groupSize++
			answers := strings.Split(group, "")
			for _, answer := range answers {
				groupMap[answer]++
			}
		}
	}
	fmt.Printf("[*] Answer Count: %d\n", questionCount)
}

func (d Day6) GetTitle() string {
	return d.Title
}

func NewDay6() *Day6 {
	return &Day6{Title: "Day 6"}
}

package solutions

import (
	"fmt"
	"strings"

	"github.com/whcass/Advent-of-Code/2022/helpers"
)

type Day6 struct {
	Title string
}

func NewDay6() *Day6 {
	return &Day6{Title: "Day 6"}
}

func (d Day6) GetTitle() string {
	return d.Title
}

func (d Day6) Run() {
	input := helpers.GetInput("6")
	sliceData := strings.Split(input, "\n")
	stream := strings.Split(sliceData[0], "")
	answer := d.Part(stream, 4)
	fmt.Printf("[*] Part 1: %d\n", answer)
	answer = d.Part(stream, 14)
	fmt.Printf("[*] Part 2: %d\n", answer)
}

func (d Day6) Part(data []string, markerLength int) int {
	for i := 0; i < len(data)-markerLength; i++ {
		marker := data[i : i+markerLength]
		uniques := helpers.UniqueStrings(marker)
		if len(uniques) == markerLength {
			return i + markerLength
		}
	}
	return -1
}

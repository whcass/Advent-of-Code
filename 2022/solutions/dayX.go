package solutions

import (
	"strings"

	"github.com/whcass/Advent-of-Code/2022/helpers"
)

type DayX struct {
	Title string
}

func NewDayX() *DayX {
	return &DayX{Title: "Day X"}
}

func (d DayX) GetTitle() string {
	return d.Title
}

func (d DayX) Run() {
	input := helpers.GetInput("X")
	sliceData := strings.Split(input, "\n")
	d.Part1(sliceData)
}

func (d DayX) Part1(data []string) {

}

package solutions

import (
	"github.com/whcass/Advent-of-Code/2022/entities"
	"github.com/whcass/Advent-of-Code/2022/helpers"
)

type Day10 struct {
	Title string
	cpu   entities.CPU
}

func NewDay10() *Day10 {
	return &Day10{Title: "Day 10", cpu: *entities.NewCPU()}
}

func (d Day10) GetTitle() string {
	return d.Title
}

func (d Day10) Run() {
	input := helpers.GetInput("10")
	sliceData := helpers.SplitInput(input)
	d.Part1(sliceData)
}

func (d *Day10) Part1(data []string) {
	for _, command := range data {
		d.cpu.RunCommand(command)
	}
	d.cpu.PrintScreen()
}

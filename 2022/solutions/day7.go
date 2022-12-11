package solutions

import (
	"strings"

	"github.com/whcass/Advent-of-Code/2022/entities"
	"github.com/whcass/Advent-of-Code/2022/helpers"
)

type Day7 struct {
	Title      string
	fileSystem []*entities.Directory
}

func NewDay7() *Day7 {
	return &Day7{Title: "Day 7"}
}

func (d Day7) GetTitle() string {
	return d.Title
}

func (d Day7) Run() {
	input := helpers.GetInput("7")
	sliceData := strings.Split(input, "\n")
	d.Part1(sliceData)
}

func (d Day7) Part1(data []string) {
	currentDirectory := entities.NewDirectory("/", nil)
	for i := 0; i < len(data); i++ {
		command := strings.Split(data[i], " ")

		if command[0] == "$" {

		}
	}
}

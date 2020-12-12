package solutions

import (
	"fmt"
	"github.com/whcass/Advent-of-Code/2020/helpers"
	"math"
	"strconv"
	"strings"
)

type Day12 struct {
	Title string
}

func NewDay12() *Day12 {
	return &Day12{Title: "Day 12"}
}

func (d Day12) Run() {
	input := helpers.GetInput("day12.txt")
	inputData := strings.Split(input, "\r\n")
	d.part1(inputData)
}

func (d Day12) part1(in []string) {
	x := 0
	y := 0
	headings := []string{
		"N",
		"E",
		"S",
		"W",
	}
	facing := 1
	for _, move := range in {
		moveSplit := strings.SplitN(move, "", 2)
		inst := moveSplit[0]
		val, _ := strconv.Atoi(moveSplit[1])
		switch inst {
		case "N":
			y += val
			break
		case "S":
			y -= val
			break
		case "E":
			x += val
			break
		case "W":
			x -= val
			break
		case "L":
			rotate := val / 90
			facing = ((facing - rotate) + 4) % 4
			break
		case "R":
			rotate := val / 90
			facing = (facing + rotate) % 4
			break
		case "F":
			switch headings[facing] {
			case "N":
				y += val
				break
			case "S":
				y -= val
				break
			case "E":
				x += val
				break
			case "W":
				x -= val
				break
			}

		}
	}
	dist := math.Abs(float64(x)) + math.Abs(float64(y))
	fmt.Printf("[*] Distance: %d\n", int(dist))

}

func (d Day12) GetTitle() string {
	return d.Title
}

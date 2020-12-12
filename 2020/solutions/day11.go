package solutions

import (
	"fmt"
	"github.com/whcass/Advent-of-Code/2020/helpers"
	"strings"
)

type Day11 struct {
	Title   string
	grid    [][]string
	newGrid [][]string
}

func NewDay11() *Day11 {
	return &Day11{Title: "Day 11"}
}

func (d *Day11) Run() {
	input := helpers.GetInput("day11.txt")
	inputData := strings.Split(input, "\r\n")
	d.grid = d.makeGrid(inputData)
	d.part1()
}

func (d *Day11) printGrid() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	for i := 0; i < len(d.grid); i++ {
		fmt.Printf("\r")
		for j := 0; j < len(d.grid[i]); j++ {
			seat := d.grid[i][j]
			if seat == "L" {
				helpers.PrintColour(seat, helpers.Green)
			} else if seat == "#" {
				helpers.PrintColour(seat, helpers.Purple)
			} else {
				fmt.Print("")
				helpers.PrintColour(seat, helpers.White)
				fmt.Print("")
			}
		}
		fmt.Println()
	}
}

func (d *Day11) part1() {
	d.newGrid = d.makeEmptyGrid()
	changed := true
	count := 0
	for changed {
		changed = d.runAutomata()
		count++
	}
	d.printGrid()
	seatCount := 0
	for i := 0; i < len(d.grid); i++ {
		for j := 0; j < len(d.grid[i]); j++ {
			if d.grid[i][j] == "#" {
				seatCount++
			}
		}
	}
	fmt.Printf("[*] Seats: %d\n", seatCount)
}

func (d *Day11) runAutomata() bool {
	d.newGrid = d.makeEmptyGrid()
	changed := false
	for row := 0; row < len(d.grid); row++ {
		for col := 0; col < len(d.grid[row]); col++ {
			switch d.grid[row][col] {
			case "L":
				occupiedSeats := d.countSeats(row, col, "#")
				if occupiedSeats == 0 {
					changed = true
					d.newGrid[row][col] = "#"
				} else {
					d.newGrid[row][col] = "L"
				}
				break
			case "#":
				occupiedSeats := d.countSeats(row, col, "#")
				if occupiedSeats >= 4 {
					changed = true
					d.newGrid[row][col] = "L"
				} else {
					d.newGrid[row][col] = "#"
				}
				break
			case ".":
				d.newGrid[row][col] = "."
				break
			}
		}
	}
	copy(d.grid, d.newGrid)
	return changed
}

func (d *Day11) countSeats(row int, col int, check string) int {
	occupiedSeats := 0
	visited := 0
	for rowOffset := -1; rowOffset <= 1; rowOffset++ {
		for colOffset := -1; colOffset <= 1; colOffset++ {
			a := row + rowOffset
			if a < 0 || a >= len(d.grid) {
				continue
			}
			b := col + colOffset
			if b < 0 || b >= len(d.grid[0]) {
				continue
			}
			if colOffset == 0 && rowOffset == 0 {
				continue
			}
			visited++
			if d.grid[a][b] == check {
				occupiedSeats++
			}
		}
	}
	return occupiedSeats
}

func (d *Day11) GetTitle() string {
	return d.Title
}

func (d *Day11) makeEmptyGrid() [][]string {

	grid := make([][]string, len(d.grid))

	for i := 0; i < len(d.grid); i++ {
		grid[i] = make([]string, len(d.grid[i]))
	}

	return grid
}

func (d *Day11) makeGrid(input []string) [][]string {
	grid := make([][]string, len(input))

	for i := 0; i < len(input); i++ {
		split := strings.Split(strings.TrimSpace(input[i]), "")
		grid[i] = make([]string, len(split))
		grid[i] = split
	}

	return grid
}

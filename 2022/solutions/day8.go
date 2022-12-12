package solutions

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/whcass/Advent-of-Code/2022/helpers"
)

type Day8 struct {
	Title string
	grid  [][]int
}

func NewDay8() *Day8 {
	return &Day8{Title: "Day 8"}
}

func (d Day8) GetTitle() string {
	return d.Title
}

func (d Day8) Run() {
	input := helpers.GetInput("8")
	sliceData := helpers.SplitInput(input)
	d.PrepareGrid(sliceData)
	d.Part1()
}

func (d *Day8) PrepareGrid(data []string) {
	d.grid = make([][]int, len(data))
	for row := 0; row < len(data); row++ {
		colData := strings.Split(data[row], "")
		d.grid[row] = make([]int, len(colData))
		for col := 0; col < len(colData); col++ {
			val, err := strconv.Atoi(colData[col])
			if err != nil {
				log.Fatal(err)
			}
			d.grid[row][col] = val
		}
	}
}

func (d *Day8) Part1() {
	count := 0
	for row := 0; row < len(d.grid); row++ {
		for col := 0; col < len(d.grid[row]); col++ {
			if col == 0 || col == len(d.grid[row])-1 || row == 0 || row == len(d.grid)-1 {
				count++
				continue
			}

			treeHeight := d.grid[row][col]

			if d.visibleCol(row, col, treeHeight) {
				count++
				// fmt.Printf("[*] %d is visible vertically at %d,%d\n", treeHeight, row, col)
				continue
			}

			if d.visibleRow(row, col, treeHeight) {
				count++
				// fmt.Printf("[*] %d is visible horizontally at %d,%d\n", treeHeight, row, col)
				continue
			}
		}
	}

	helpers.PrintResult(fmt.Sprintf("Part 1: %d", count))
}

func (d Day8) visibleCol(rowIndex int, colIndex, treeHeight int) bool {
	column := make([]int, 0)
	for _, row := range d.grid {
		column = append(column, row[colIndex])
	}
	visibleUp := true
	visibleDown := true
	for i := rowIndex + 1; i < len(column); i++ {
		val := column[i]
		if val >= treeHeight {
			visibleDown = false
			break
		}
	}

	for i := rowIndex - 1; i >= 0; i-- {
		val := column[i]
		if val >= treeHeight {
			visibleUp = false
			break
		}
	}

	return visibleDown || visibleUp
}

func (d Day8) visibleRow(rowIndex int, colIndex int, treeHeight int) bool {

	visibleRight := true
	visibleLeft := true
	for i := colIndex + 1; i < len(d.grid[rowIndex]); i++ {
		val := d.grid[rowIndex][i]
		if val >= treeHeight {
			visibleRight = false
			break
		}
	}

	for i := colIndex - 1; i >= 0; i-- {
		val := d.grid[rowIndex][i]
		if val >= treeHeight {
			visibleLeft = false
			break
		}
	}

	return visibleLeft || visibleRight
}

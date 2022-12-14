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
	bestScenicScore := 0
	for row := 0; row < len(d.grid); row++ {
		for col := 0; col < len(d.grid[row]); col++ {
			if col == 0 || col == len(d.grid[row])-1 || row == 0 || row == len(d.grid)-1 {
				count++
				continue
			}

			treeHeight := d.grid[row][col]
			visibleCol, colScore := d.visibleCol(row, col, treeHeight)
			visibleRow, rowScore := d.visibleRow(row, col, treeHeight)

			scenicScore := colScore * rowScore
			if scenicScore > bestScenicScore {
				bestScenicScore = scenicScore
			}

			if visibleCol {
				count++
				continue
			}

			if visibleRow {
				count++
				continue
			}
		}
	}

	helpers.PrintResult(fmt.Sprintf("Part 1: %d", count))
	helpers.PrintResult(fmt.Sprintf("Part 2: %d", bestScenicScore))
}

func (d Day8) visibleCol(rowIndex int, colIndex, treeHeight int) (bool, int) {
	column := make([]int, 0)
	for _, row := range d.grid {
		column = append(column, row[colIndex])
	}
	visibleUp := true
	visibleDown := true
	downScore := 0
	upScore := 0
	for i := rowIndex + 1; i < len(column); i++ {
		downScore++
		val := column[i]
		if val >= treeHeight {
			visibleDown = false
			break
		}
	}

	for i := rowIndex - 1; i >= 0; i-- {
		val := column[i]
		upScore++
		if val >= treeHeight {
			visibleUp = false
			break
		}
	}

	return visibleDown || visibleUp, upScore * downScore
}

func (d Day8) visibleRow(rowIndex int, colIndex int, treeHeight int) (bool, int) {

	visibleRight := true
	visibleLeft := true
	leftScore := 0
	rightScore := 0
	for i := colIndex + 1; i < len(d.grid[rowIndex]); i++ {
		rightScore++
		val := d.grid[rowIndex][i]
		if val >= treeHeight {
			visibleRight = false
			break
		}
	}

	for i := colIndex - 1; i >= 0; i-- {
		leftScore++
		val := d.grid[rowIndex][i]
		if val >= treeHeight {
			visibleLeft = false
			break
		}
	}

	return visibleLeft || visibleRight, rightScore * leftScore
}

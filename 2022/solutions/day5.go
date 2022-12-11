package solutions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
	"github.com/whcass/Advent-of-Code/2022/helpers"
)

type Day5 struct {
	Title string
	stack [][]string
}

func NewDay5() *Day5 {
	return &Day5{Title: "Day 5", stack: makeStack()}
}

func (d Day5) GetTitle() string {
	return d.Title
}

func makeStack() [][]string {
	return [][]string{
		{},                                       //0
		{"D", "B", "J", "V"},                     //1
		{"P", "V", "B", "W", "R", "D", "F"},      //2
		{"R", "G", "F", "L", "D", "C", "W", "Q"}, //3
		{"W", "J", "P", "M", "L", "N", "D", "B"}, //4
		{"H", "N", "B", "P", "C", "S", "Q"},      //5
		{"R", "D", "B", "S", "N", "G"},           //6
		{"Z", "B", "P", "M", "Q", "F", "S", "H"}, //7
		{"W", "L", "F"},                          //8
		{"S", "V", "F", "M", "R"},                //9
	}
}

func (d Day5) Run() {
	input := helpers.GetInput("5")
	sliceData := strings.Split(input, "\n")
	d.Part(sliceData, 1)
	d.stack = makeStack()
	d.Part(sliceData, 2)
}

func (d Day5) Part(data []string, part int) {
	for i := 0; i < len(data); i++ {
		move, from, to := parseCommand(data[i])
		// fmt.Printf("[*] stack: %v\n", d.stack)
		// fmt.Printf("[*] %s\n", data[i])
		fromLen := len(d.stack[from])
		var toVal []string
		if part == 1 {
			toVal = append(d.stack[to], funk.ReverseStrings(d.stack[from][fromLen-move:fromLen])...)
		} else {
			toVal = append(d.stack[to], d.stack[from][fromLen-move:fromLen]...)
		}
		d.setStackVal(to, toVal)
		d.setStackVal(from, d.stack[from][0:fromLen-move])
		// fmt.Printf("[*] stack: %v\n", d.stack)
	}
	fmt.Printf("[*] Part %d: ", part)
	for j := 1; j < len(d.stack); j++ {
		stack := d.stack[j]
		fmt.Printf("%s", stack[len(stack)-1])
	}
	fmt.Println()
}

func (d *Day5) setStackVal(stack int, val []string) {
	d.stack[stack] = val
}

func parseCommand(line string) (int, int, int) {
	command := strings.Split(line, " ")
	move, _ := strconv.Atoi(command[1])
	from, _ := strconv.Atoi(command[3])
	to, _ := strconv.Atoi(command[5])

	return move, from, to
}

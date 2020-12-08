package solutions

import (
	"fmt"
	"github.com/whcass/Advent-of-Code/2020/helpers"
	"strconv"
	"strings"
)

type Day8 struct {
	Title     string
	acc       int
	loc       int
	program   []string
	instCount map[int]int
	try       map[int]int
}

func NewDay8() *Day8 {
	return &Day8{Title: "Day 8", acc: 0, loc: 0, instCount: make(map[int]int), try: make(map[int]int)}
}

func (d *Day8) Run() {
	input := helpers.GetInput("day8.txt")
	d.program = strings.Split(input, "\r\n")
	d.part1()
	d.reset()
	d.part2()
}

func (d *Day8) reset() {
	d.acc = 0
	d.loc = 0
	d.instCount = make(map[int]int)
}

func (d *Day8) GetTitle() string {
	return d.Title
}

func (d *Day8) part2() {
	trying := false
	for true {
		if d.loc > len(d.program)-1 {
			fmt.Printf("[*] Final Value : %d\n", d.acc)
			return
		}
		if _, ok := d.instCount[d.loc]; !ok {
			d.instCount[d.loc] = 1
		} else {
			d.reset()
			d.part2()
			return
		}
		inst := d.program[d.loc]
		instSplit := strings.Split(inst, " ")
		value, _ := strconv.Atoi(instSplit[1])
		switch instSplit[0] {
		case "acc":
			increment := value
			d.acc += increment
			d.loc += 1
			break
		case "jmp":
			if _, ok := d.try[d.loc]; !ok && !trying {
				trying = true
				d.try[d.loc] = 1
				d.loc += 1
				break
			}
			jmpValue := value
			d.loc += jmpValue
			break
		case "nop":
			if _, ok := d.try[d.loc]; !ok && !trying {
				d.try[d.loc] = 1
				trying = true
				jmpValue := value
				d.loc += jmpValue
				break
			}
			d.loc += 1
			break
		}
	}
}

func (d *Day8) part1() {
	for true {
		if _, ok := d.instCount[d.loc]; !ok {
			d.instCount[d.loc] = 1
		} else {
			fmt.Printf("[*] Accumulator Value: %d\n", d.acc)
			return
		}
		inst := d.program[d.loc]
		instSplit := strings.Split(inst, " ")
		switch instSplit[0] {
		case "acc":
			increment, _ := strconv.Atoi(instSplit[1])
			d.acc += increment
			d.loc += 1
			break
		case "jmp":
			jmpValue, _ := strconv.Atoi(instSplit[1])
			d.loc += jmpValue
			break
		case "nop":
			d.loc += 1
			break
		}
	}
}

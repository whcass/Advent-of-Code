package entities

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/whcass/Advent-of-Code/2022/helpers"
)

type CPU struct {
	register       int
	cycleCount     int
	signalStrength int
	CrtDrawing     []string
}

func NewCPU() *CPU {
	return &CPU{register: 1, cycleCount: 0, CrtDrawing: make([]string, 240)}
}

// func makeBlankScreen() [][]string {
// 	screen := make([][]string, 6)
// 	for i := range screen {
// 		screen[i] = make([]string, 40)
// 	}

// 	return screen
// }

func (c *CPU) RunCommand(command string) {
	args := strings.Split(command, " ")

	switch args[0] {
	case "noop":
		c.runCycle(1)
	case "addx":
		c.runCycle(2)
		val, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal(err)
		}
		c.register += val
	}
}

func (c *CPU) runCycle(count int) {
	for i := 0; i < count; i++ {
		if c.cycleCount == 20 || (c.cycleCount-20)%40 == 0 { // cycle count is 20 or 20+=40
			c.signalStrength += c.cycleCount * c.register
			fmt.Printf("[-] X at cycle: %d = %d \t| Signal Strength: %d\n", c.cycleCount, c.register, c.signalStrength)
		}
		var pixel string
		if c.register-1 == c.cycleCount%40 || c.register == c.cycleCount%40 || c.register+1 == c.cycleCount%40 {
			pixel = "#"
		} else {
			pixel = "."
		}
		c.CrtDrawing[c.cycleCount] = pixel
		c.cycleCount++
	}
}

func (c CPU) PrintScreen() {
	for i := range c.CrtDrawing {
		if i%40 == 0 {
			fmt.Println()
		}
		helpers.PrintColour(c.CrtDrawing[i]+" ", helpers.Green)
	}
	fmt.Println()
	fmt.Println()
}

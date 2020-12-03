package main

import (
	"github.com/whcass/Advent-of-Code/2020/helpers"
	"github.com/whcass/Advent-of-Code/2020/interfaces"
	"github.com/whcass/Advent-of-Code/2020/solutions"
)

func main() {

	solutionsToRun := []interfaces.Solutions{
		solutions.NewDay1(),
		solutions.NewDay2(),
		solutions.NewDay3(),
	}

	helpers.PrintColour("[*] Running solutions\n", helpers.Teal)
	for _, solution := range solutionsToRun {
		helpers.PrintColour("===========================================\n", helpers.Magenta)
		solution.Run()
		helpers.PrintColour("[*] Done.\n", helpers.Green)
	}
	helpers.PrintColour("===========================================\n", helpers.Magenta)
}

package main

import (
	"fmt"

	"github.com/whcass/Advent-of-Code/2022/helpers"
	"github.com/whcass/Advent-of-Code/2022/interfaces"
	"github.com/whcass/Advent-of-Code/2022/solutions"
)

func main() {
	solutionsToRun := []interfaces.Solutions{
		solutions.NewDay1(),
		solutions.NewDay2(),
		solutions.NewDay3(),
	}

	helpers.PrintColour("[+] Running solutions\n", helpers.Teal)
	for _, solution := range solutionsToRun {
		helpers.PrintColour("===========================================\n", helpers.Magenta)
		helpers.PrintColour(fmt.Sprintf("[+] Running %s\n", solution.GetTitle()), helpers.Teal)
		solution.Run()
		helpers.PrintColour("[+] Done.\n", helpers.Green)
	}
	helpers.PrintColour("===========================================\n", helpers.Magenta)
}

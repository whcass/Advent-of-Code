package main

import (
	"fmt"
	"github.com/whcass/Advent-of-Code/2020/interfaces"
	"github.com/whcass/Advent-of-Code/2020/solutions"
)

func main() {

	solutionsToRun := []interfaces.Solutions{
		solutions.NewDay1(),
		solutions.NewDay2(),
		solutions.NewDay3(),
	}

	fmt.Println("[*] Running solutions")
	for _, solution := range solutionsToRun {
		fmt.Printf("\033[1;34m%s\033[0m", "===========================================")
		fmt.Println("")
		solution.Run()
	}
	fmt.Println("===========================================")
}

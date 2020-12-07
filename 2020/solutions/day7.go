package solutions

import (
	"fmt"
	"github.com/whcass/Advent-of-Code/2020/entities"
	"github.com/whcass/Advent-of-Code/2020/helpers"
	"regexp"
	"strconv"
	"strings"
)

type Day7 struct {
	Title   string
	bagTree map[string]*entities.Bag
}

func NewDay7() *Day7 {
	return &Day7{Title: "Day 7"}
}

func (d Day7) Run() {
	input := helpers.GetInput("day7.txt")
	inputData := strings.Split(input, "\r\n")
	d.bagTree = d.buildBagTree(inputData)
	d.part1()
	d.part2()
}

func (d Day7) GetTitle() string {
	return d.Title
}

func (d Day7) part2() {
	count := 0
	shinyGoldBag := d.bagTree["shiny gold"]
	count = d.bagSearchPart2(shinyGoldBag)
	fmt.Printf("[*] Shiny Bag contains %d bags\n", count)
}

func (d Day7) part1() {
	count := 0
	for _, bag := range d.bagTree {
		if bag.Children != nil {
			if d.bagSearchPart1(bag) {
				count++
			}
		}
	}
	fmt.Printf("[*] Bags Containing Shiny Gold: %d\n", count)
}

func (d Day7) bagSearchPart2(bag *entities.Bag) int {
	count := len(bag.Children)
	for _, child := range bag.Children {
		count += d.bagSearchPart2(child)
	}
	return count
}

func (d Day7) bagSearchPart1(bag *entities.Bag) bool {
	skipDupesColour := ""
	for _, bag := range bag.Children {
		if bag.Colour == "shiny gold" {
			return true
		} else {
			if bag.Colour != skipDupesColour {
				skipDupesColour = bag.Colour

				if d.bagSearchPart1(bag) {
					return true
				}
			}
		}
	}
	return false
}

func (d Day7) buildBagTree(data []string) map[string]*entities.Bag {
	bagHash := make(map[string]*entities.Bag)
	for _, rule := range data {
		bagNameR, _ := regexp.Compile("^([a-z ]+) b")
		bagName := bagNameR.FindStringSubmatch(rule)[1]

		childBagsR, _ := regexp.Compile("(([0-9]) ([a-z ]+)) b")
		childBags := childBagsR.FindAllStringSubmatch(rule, -1)
		if _, ok := bagHash[bagName]; !ok {
			bag := entities.NewBag(bagName)
			bagHash[bagName] = bag
		}
		for _, child := range childBags {
			childBagCount, _ := strconv.Atoi(child[2])
			childBagColour := child[3]
			if _, ok := bagHash[childBagColour]; !ok {
				childBag := entities.NewBag(childBagColour)
				bagHash[childBagColour] = childBag
			}
			for i := 0; i < childBagCount; i++ {
				bagHash[bagName].AddChild(bagHash[childBagColour])
			}
		}
	}
	return bagHash
}

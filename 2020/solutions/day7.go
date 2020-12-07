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
	Title string
}

func NewDay7() *Day7 {
	return &Day7{Title: "Day 7"}
}

func (d Day7) Run() {
	input := helpers.GetInput("day7.txt")
	inputData := strings.Split(input, "\r\n")
	d.part1(inputData)
}

func (d Day7) GetTitle() string {
	return d.Title
}

func (d Day7) part1(data []string) {
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
	test := bagHash["vibrant blue"]
	fmt.Println(test)
}

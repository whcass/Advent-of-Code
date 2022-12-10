package solutions

import (
	"fmt"
	"strings"

	"github.com/juliangruber/go-intersect"
	"github.com/whcass/Advent-of-Code/2022/helpers"
)

type Day3 struct {
	Title string
}

func NewDay3() *Day3 {
	return &Day3{Title: "Day 3"}
}

func makeAlphabetMap() map[string]int {
	m := make(map[string]int)
	var ch byte
	val := 1
	for ch = 'a'; ch <= 'z'; ch++ {
		m[string(ch)] = val
		val++
	}
	for ch = 'A'; ch <= 'Z'; ch++ {
		m[string(ch)] = val
		val++
	}

	return m
}

func (d Day3) GetTitle() string {
	return d.Title
}

func (d Day3) Run() {
	input := helpers.GetInput("3")
	sliceData := strings.Split(input, "\n")
	d.Part1(sliceData)
	d.Part2(sliceData)
}

func (d Day3) Part1(sliceData []string) {
	valueMap := makeAlphabetMap()
	total := 0
	for i := 0; i < len(sliceData); i++ {
		data := sliceData[i]
		length := len(data)
		firstList := data[0:(length / 2)]
		secondList := data[(length / 2):length]
		result := helpers.Unique(intersect.Simple(firstList, secondList))
		// fmt.Printf("[*] intersect of %s and %s is %s\n", firstList, secondList, result)
		if len(result) >= 1 {
			for item := 0; item < len(result); item++ {
				val := result[item]
				// fmt.Printf("[*] %s is %d\n", string(val), valueMap[string(val)])
				total += valueMap[string(val)]
			}
		}
	}
	// fmt.Printf("[*] Info: %v\n", valueMap)
	fmt.Printf("[*] Part 1: %d\n", total)
}

func (d Day3) Part2(sliceData []string) {
	valueMap := makeAlphabetMap()
	total := 0
	for i := 0; i < len(sliceData); i += 3 {
		firstElf := sliceData[i]
		secondElf := sliceData[i+1]
		thirdElf := sliceData[i+2]

		r := helpers.Unique(intersect.Simple(intersect.Simple(firstElf, secondElf), thirdElf))
		total += valueMap[string(r)]
	}
	fmt.Printf("[*] Part 2: %d\n", total)
}

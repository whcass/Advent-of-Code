package solutions

import (
	"fmt"
	"github.com/whcass/Advent-of-Code/2020/entities"
	"github.com/whcass/Advent-of-Code/2020/helpers"
	"strings"
)

type Day4 struct {
	Title string
}

func NewDay4() *Day4 {
	return &Day4{Title: "Day 4"}
}

func (d Day4) Run() {
	input := helpers.GetInput("Day4.txt")
	inputData := strings.Split(input, "\r\n")
	passports := d.GeneratePassports(inputData)

	d.Part1(passports)
	d.Part2(passports)
}

func (d Day4) GeneratePassports(inputData []string) []*entities.Passport {
	passPortString := ""
	var passports []*entities.Passport
	for i := 0; i < len(inputData); i++ {
		if inputData[i] == "" {
			newPassport := entities.NewPassportFromString(passPortString)
			passports = append(passports, newPassport)
			passPortString = ""
		} else {
			passPortString += " " + strings.TrimRight(inputData[i], "\r\n")
		}
	}

	return passports
}

func (d Day4) Part1(passports []*entities.Passport) {
	validCount := 0
	for _, passport := range passports {
		if passport.CheckBlanks() {
			validCount++
		}
	}

	fmt.Printf("[*] Valid Passports: %d\n", validCount)
}

func (d Day4) Part2(passports []*entities.Passport) {
	validCount := 0
	for _, passport := range passports {
		if passport.Check() {
			validCount++
		} else {
			fmt.Println(passport)
		}
	}

	fmt.Printf("[*] Valid Passports: %d\n", validCount)
}
func (d Day4) GetTitle() string {
	return d.Title
}

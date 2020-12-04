package entities

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	BirthYear  string
	IssueYear  string
	Expiration string
	Height     string
	HairColour string
	EyeColour  string
	ID         string
	CountryID  string
}

func NewPassport(birthYear string, issueYear string, expiration string, height string, hairColour string, eyeColour string, ID string, countryID string) *Passport {
	return &Passport{BirthYear: birthYear, IssueYear: issueYear, Expiration: expiration, Height: height, HairColour: hairColour, EyeColour: eyeColour, ID: ID, CountryID: countryID}
}

func NewPassportFromString(s string) *Passport {
	passportSplit := strings.Split(s, " ")
	pp := Passport{}
	for i := 0; i < len(passportSplit); i++ {
		kvPair := strings.Split(passportSplit[i], ":")
		switch kvPair[0] {
		case "byr":
			pp.BirthYear = kvPair[1]
			break
		case "iyr":
			pp.IssueYear = kvPair[1]
			break
		case "eyr":
			pp.Expiration = kvPair[1]
			break
		case "hgt":
			pp.Height = kvPair[1]
			break
		case "hcl":
			pp.HairColour = kvPair[1]
			break
		case "ecl":
			pp.EyeColour = kvPair[1]
			break
		case "pid":
			pp.ID = kvPair[1]
			break
		case "cid":
			pp.CountryID = kvPair[1]
			break
		}
	}

	return &pp
}

func (p Passport) CheckBlanks() bool {
	if p.BirthYear == "" {
		return false
	} else if p.IssueYear == "" {
		return false
	} else if p.Expiration == "" {
		return false
	} else if p.Height == "" {
		return false
	} else if p.HairColour == "" {
		return false
	} else if p.EyeColour == "" {
		return false
	} else if p.ID == "" {
		return false
	}

	return true
}

func (p Passport) Check() bool {
	valid := p.CheckBlanks()
	if !valid {
		return false
	}

	valid = p.validateFields()
	if !valid {
		return false
	}

	return true
}

func (p Passport) validateFields() bool {
	// Birth Year
	bYear, err := strconv.Atoi(p.BirthYear)
	if err != nil {
		panic(err)
	}
	if bYear < 1920 || bYear > 2002 {
		fmt.Println("Invalid Birth Year")
		return false
	}

	// Issue Year
	iYear, err := strconv.Atoi(p.IssueYear)
	if err != nil {
		panic(err)
	}
	if iYear < 2010 || iYear > 2020 {
		fmt.Println("Invalid Issue Year")
		return false
	}

	// Expiration Year
	eYear, err := strconv.Atoi(p.Expiration)
	if err != nil {
		panic(err)
	}
	if eYear < 2020 || eYear > 2030 {
		fmt.Println("Invalid Expiration Year")
		return false
	}

	// Height
	if strings.HasSuffix(p.Height, "cm") {
		height, err := strconv.Atoi(p.Height[:len(p.Height)-2])
		if err != nil {
			panic(err)
		}
		if height < 150 || height > 193 {
			fmt.Println("Invalid Height")
			return false
		}
	} else if strings.HasSuffix(p.Height, "in") {
		height, err := strconv.Atoi(p.Height[:len(p.Height)-2])
		if err != nil {
			panic(err)
		}
		if height < 59 || height > 76 {
			fmt.Println("Invalid Height")
			return false
		}
	} else {
		fmt.Println("Invalid Height - No Suffix")
		return false
	}

	// Hair Colour
	if len(p.HairColour) != 7 {
		fmt.Println("Hair Colour not long enough")
		return false
	}
	hairValid, err := regexp.MatchString("#[0-9]*[a-f]*", p.HairColour)
	if err != nil {
		panic(err)
	}
	if !hairValid {
		fmt.Println("Hair Invalid")
		return false
	}

	if !(p.EyeColour == "amb" || p.EyeColour == "blu" || p.EyeColour == "brn" || p.EyeColour == "gry" || p.EyeColour == "grn" || p.EyeColour == "hzl" || p.EyeColour == "oth") {
		fmt.Println("Eye Colour Invalid")
		return false
	}

	if len(p.ID) != 9 {
		fmt.Println("Passport ID Invalid - not long enough")
		return false
	}
	passportIDValid, err := regexp.MatchString("[0-9]{9}", p.ID)
	if err != nil {
		panic(err)
	}
	if !passportIDValid {
		fmt.Println("Passport ID Invalid")
		return false
	}
	return true
}

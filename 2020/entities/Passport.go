package entities

import (
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

func (p Passport) checkBlanks() bool {
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
	return p.checkBlanks()
}

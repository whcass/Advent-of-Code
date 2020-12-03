package helpers

import (
	"io/ioutil"
	"log"
)

func GetInput(name string) string {
	content, err := ioutil.ReadFile("inputs/" + name)
	if err != nil {
		log.Fatal(err)
	}
	input := string(content)
	return input
}

package helpers

import (
	"io/ioutil"
	"log"
	"strconv"
)

func GetInput(name string) string {
	content, err := ioutil.ReadFile("inputs/" + name)
	if err != nil {
		log.Fatal(err)
	}
	input := string(content)
	return input
}

func PrepareInput(in []string) []int {
	out := make([]int, len(in))
	for i := 0; i < len(in); i++ {
		out[i], _ = strconv.Atoi(in[i])
	}

	return out
}

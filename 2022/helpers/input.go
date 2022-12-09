package helpers

import (
	"io/ioutil"
	"log"
	"strconv"
)

func GetInput(dayNumber string) string {
	content, err := ioutil.ReadFile("inputs/day" + dayNumber + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	// if errors.Is(err, os.ErrNotExist) {
	// 	url := "https://adventofcode.com/2022/day/" + dayNumber + "/input"
	// 	res, _ := http.Get(url)
	// 	fmt.Println(url)
	// 	fmt.Println(res)
	// 	defer res.Body.Close()
	// 	out, err := os.Create("inputs/day" + dayNumber + ".txt")
	// 	if err != nil {
	// 		// panic?
	// 	}
	// 	defer out.Close()
	// 	io.Copy(out, res.Body)
	// 	content, _ = ioutil.ReadFile("inputs/day" + dayNumber + ".txt")
	// }
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

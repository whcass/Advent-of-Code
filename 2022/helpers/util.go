package helpers

import "fmt"

var (
	Black   = "\033[1;30m%s\033[0m"
	Red     = "\033[1;31m%s\033[0m"
	Green   = "\033[1;32m%s\033[0m"
	Yellow  = "\033[1;33m%s\033[0m"
	Purple  = "\033[1;34m%s\033[0m"
	Magenta = "\033[1;35m%s\033[0m"
	Teal    = "\033[1;36m%s\033[0m"
	White   = "\033[1;37m%s\033[0m"
)

func PrintInfo(title string) {
	PrintColour(fmt.Sprintf("[+] Running %s\n", title), Teal)
}

func PrintColour(s string, colour string) {
	fmt.Printf(colour, s)
}

func Unique(intSlice []interface{}) []byte {
	keys := make(map[interface{}]bool)
	list := []byte{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry.(byte))
		}
	}
	return list
}

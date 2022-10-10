package main

import (
	"fmt"
	"strings"
)

/*
func main() {
	str := "Hello world, good morning"
	separatedStr := strings.Fields(str)
	for _, separatedStr := range separatedStr {
		fmt.Println(separatedStr)
	}
}
*/

func main() {
	str := "it was (cap, 2) the , best of times,"
	delimiter1 := "("
	delimiter2 := ")"
	parts1 := strings.Split(str, delimiter1)
	parts2 := strings.Split(string(parts1), delimiter2)
	fmt.Println(parts2)
}

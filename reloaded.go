package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	strContent := string(content)
	strContent1 := Separate(strContent)
}

func Separate(content string) []string {
	return strings.Split(content, " ")
}

func checkPunct(content []string) bool {
	for _, content := range content {
		switch content {
		case ",", ".", "!", "?", ";", ":":
			return true
		}
	}
	return false
}

/*
func Hex(s string) int64 {
	decimal, err := strconv.ParseInt(s, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return decimal
}
*/

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// read sample txt, apply functions
func main() {
	input, err := os.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	str := string(input)
	strContent := strings.Split(str, " ")
}

func convertString(strContent []string) []string {
	var emptyString []string
	for i := 0; i < len(strContent); i++ {
		switch strContent[i] {
		case "(hex)":
			convHex, _ := strconv.ParseInt(strContent[i-1], 16, 64)
			emptyString[len(emptyString)-1] = fmt.Sprint(convHex)
		case "(bin)":
			convBin, _ := strconv.ParseInt(strContent[i-1], 2, 64)
			emptyString[len(emptyString)-1] = fmt.Sprint(convBin)
		case "(cap)":
			convCap := strings.Title(strContent[i-1])
			emptyString[len(emptyString)-1] = fmt.Sprint(convCap)
		case "(up)":
			convUp := strings.ToUpper(strContent[i-1])
			emptyString[len(emptyString)-1] = fmt.Sprint(convUp)
		case "(low)":
			convLow := strings.ToLower(strContent[i-1])
			emptyString[len(emptyString)-1] = fmt.Sprint(convLow)
		default:
			return emptyString
		}
	}
	return emptyString
}

/*
// check for punctuation (except apostrophes)
func checkPunct(content []string) bool {
	for _, content := range content {
		switch content {
		case ",", ".", "!", "?", ";", ":":
			return true
		}
	}
	return false
}

func checkPunctGrp(content []string) bool {
	for _, content := range content {
		switch content {
		case "!?", "?!", "...":
			return true
		}
	}
	return false
}

/*
func fixPunct(content []string) []string {
	emptyString := ""
	// checks for punctuation
	for i := 0; i < len(strContent1); i++ {
		checkPunct(strContent1[i])
		checkPunctGrp(strContent1[i])
		// i = punctuation, check if i-1 is a whitespace
		if checkPunct(strContent1[i]) == true || checkPunctGrp(strContent[i]) == true {
			for j := i; j < len(strContent1); j++ {
				if strContent1[i-1] == " " {
					strContent1[i-1] = strContent[i-2]
					return strContent[i-1]
				} else {
					return strContent[i]
				}
			}
		} else {
			continue
		}
	}
}
*/
/*
// check hex, return decimal
func Hex(s string) int64 {
	decimal, err := strconv.ParseInt(s, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return decimal
}
*/
func Cap() {
	str := "hello world"
	capStr := strings.Title(str)
	fmt.Println(capStr)
}

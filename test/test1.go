package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
			emptyString = append(emptyString, strContent[i])
		}
	}
	return emptyString
}

func checkPunct(strContent []string) bool {
	for i := 0; i < len(strContent); i++ {
		switch strContent[i] {
		case ".", ",", "!", "?", ":", ";":
			return true
		default:
			return false
		}
	}
	return true
}

func movePunct(strContent []string) string {
	for i := 0; i < len(strContent); i++ {
		if checkPunct(strContent) {
			if strContent[i-1] == " " {
				strContent[i-1] = strContent[i]
				return strContent[i]
			}
		}
	}
	return strContent[i]
}

func main() {
	args := os.Args
	input, err := os.ReadFile(args[1])
	if err != nil {
		log.Fatal(err)
	}

	str := string(input)
	strContent := strings.Split(str, " ")
	converted := convertString(strContent)
	fmt.Println(strContent)
	fmt.Println(converted)

	movePunctStr := movePunct(strContent)
	fmt.Println(movePunctStr)
}

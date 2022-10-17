package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func convertString(strContent []string) []string {
	// check + replace hex, bin, up, low, cap
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
		case "(cap,":
			deleteBracket := strings.Trim(strContent[i+1], ")")
			strToInt, err := strconv.Atoi(deleteBracket)
			if err != nil {
				panic(err)
			}
			for j := 1; j <= strToInt; j++ {
				convCap := strings.Title(emptyString[len(emptyString)-j])
				emptyString[len(emptyString)-j] = fmt.Sprint(convCap)
			}
			i++
		case "(up)":
			convUp := strings.ToUpper(strContent[i-1])
			emptyString[len(emptyString)-1] = fmt.Sprint(convUp)
		case "(up,":
			deleteBracket := strings.Trim(strContent[i+1], ")")
			strToInt, err := strconv.Atoi(deleteBracket)
			if err != nil {
				panic(err)
			}
			for j := 1; j <= strToInt; j++ {
				convUp := strings.ToUpper(emptyString[len(emptyString)-j])
				emptyString[len(emptyString)-j] = fmt.Sprint(convUp)
			}
			i++
		case "(low)":
			convLow := strings.ToLower(strContent[i-1])
			emptyString[len(emptyString)-1] = fmt.Sprint(convLow)
		case "(low,":
			deleteBracket := strings.Trim(strContent[i+1], ")")
			strToInt, err := strconv.Atoi(deleteBracket)
			if err != nil {
				panic(err)
			}
			for j := 1; j <= strToInt; j++ {
				convLow := strings.ToLower(emptyString[len(emptyString)-j])
				emptyString[len(emptyString)-j] = fmt.Sprint(convLow)
			}
			i++
		default:
			emptyString = append(emptyString, strContent[i])
		}
		return emptyString
	}
	// check + replace 'a'
	for l := 1; l < len(emptyString); l++ {
		if checkA(emptyString) == true {
			if checkVowel(emptyString) == true {
				emptyString = append(emptyString, "n")
			}
		}
	}

	// check + move punctuation
	runes := []rune(emptyString)
	for k := 1; k < len(emptyString); k++ {
		if checkPunct(runes) == true {
			if runes[k+1] != ' ' {
				runes[k], runes[k+1] = runes[k+1], runes[k]
			}
		}
		return runes
	}
}

// check punctuation (except apostrophes)
func checkPunct(strContent []rune) bool {
	for i := 0; i < len(strContent); i++ {
		switch strContent[i] {
		case '.', ',', '!', '?', ':', ';':
			return true
		default:
			return false
		}
	}
	return true
}

// check if A
func checkA(strContent []string) bool {
	for i := 0; i < len(strContent); i++ {
		switch strContent[i] {
		case "a":
			return true
		default:
			return false
		}
	}
	return false
}

// check if next character is a vowel or 'H'
func checkVowel(strContent []string) bool {
	for i := 0; i < len(strContent); i++ {
		switch strContent[i+1] {
		case "a", "e", "i", "o", "u", "h":
			return true
		default:
			return false
		}
	}
	return false
}

// // check punctuation group
// func checkPunctGrp(content []string) bool {
// 	for _, content := range content {
// 		switch content {
// 		case "!?", "?!", "...":
// 			return true
// 		}
// 	}
// 	return false
// }

// // move punctuation
// func movePunct(strContent []string) []string {
// 	for i := 0; i < len(strContent); i++ {
// 		if checkPunct(strContent) {
// 			if strContent[i-1] == " " {
// 				strContent[i-1] = fmt.Sprint(strContent[i])
// 			}
// 		}
// 	}
// 	return strContent
// }

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

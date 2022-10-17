package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// check + replace hex, bin, up, low, cap
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
	}
	return emptyString
}

// convert 'a' to 'an' when next word begins with a vowel or 'h'
func aToAn(strContent []string) []string {
	for i := 0; i < len(strContent); i++ {
		if strContent[i] == "a" {
			runes := []rune(strContent[i+1])
			if runes[0] == 'a' || runes[0] == 'e' || runes[0] == 'i' || runes[0] == 'o' || runes[0] == 'u' || runes[0] == 'h' {
				strContent[i] += "n"
			}
		}
	}
	return strContent
}

// find and fix non-quote punctuation (if space BEFORE punctuation)
func fixPunct(strContent []string) []string {
	runes := []rune(strings.Join(strContent, " "))
	for i := 0; i < len(runes); i++ {
		if runes[i] == '.' || runes[i] == ',' || runes[i] == '!' || runes[i] == '?' || runes[i] == ';' || runes[i] == ':' {
			if runes[i-1] == ' ' {
				runes[i], runes[i-1] = runes[i-1], runes[i]
			}
		}
	}
	runeString := string(runes)
	sliceString := strings.Split(runeString, " ")
	return sliceString
}

func fixApostrophes(strContent []string) []string {
	runes := []rune(strings.Join(strContent, " "))
	for i := 0; i < len(runes); i++ {
		if runes[i] == 39 { // 39 = apostrophe 
			if strContent == 
		}
	}
}

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
	strContent := strings.Split(str, " ")  // separate by whitespaces
	converted := convertString(strContent) // hex, bin, cap etc
	convertedA := aToAn(converted)         // a to an
	convertedP := fixPunct(convertedA)     // fix non-quotes punctuation

	fmt.Println(strContent)
	fmt.Println(converted)
	fmt.Println(convertedA)
	fmt.Println(convertedP)
}

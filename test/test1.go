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
	// runes := []rune(strings.Join(strContent, " "))
	// for i := 1; i < len(strContent); i++ {
	// 	if runes[i] == '.' || runes[i] == ',' || runes[i] == '!' || runes[i] == '?' || runes[i] == ';' || runes[i] == ':' {
	// 		if runes[i-1] == ' ' {
	// 			runes[i], runes[i-1] = runes[i-1], runes[i]
	// 		}
	// 	}
	// }
	// runeString := string(runes)
	// sliceString := strings.Split(runeString, " ")
	// return sliceString
	for i := 0; i < len(strContent); i++ {
		punctArray := []string{",", ".", "!", "?", ":", ";"}
		runeArr := []rune(strContent[i])
		if Contains(punctArray, string(runeArr[0])) { // if rune contains any punctuation
			if len(runeArr) > 1 { // if punctuation is next to a letter or part of a group
				punctCounter := 0
				// if rune array contains punctuation and punct
				for punctCounter < len(runeArr) && Contains(punctArray, string(runeArr[punctCounter])) {
					punctCounter++
				}
				// loop to replace everything before i with j
				for j := 0; j < punctCounter; j++ {
					strContent[i-1] += string(runeArr[j])
					strContent[i] = string(runeArr[j:])
				}
				// when punctCounter = length of rune array
				if punctCounter == len(runeArr) {
					// append string to include everything except i
					strContent = append(strContent[:i], strContent[i+1:]...)
				}
			} else { // if punctuation is by itself (space on either side)
				strContent[i-1] += string(runeArr[0])
				strContent = append(strContent[:i], strContent[i+1:]...)
			}
		}
	}
	return strContent
}

// function to find out whether a string contains another string
func Contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

// func fixApostrophes1(strContent []string) []string {
// 	counter := 0
// 	for i := 0; i < len(strContent); i++ {
// 		quote := []string{"'"}
// 		runeArr := []rune(strContent[i])
// 		if counter == 0 && Contains(quote, string(runeArr[0])) {
// 			counter += 1
// 			strContent[i+1] = string(runeArr) + strContent[i+1]
// 			strContent = append(strContent[:i], strContent[i+1:]...)
// 		}
// 	}
// 	return strContent
// }

// find and fix last apostrophe
func fixApostrophes(strContent []string) []string {
	// runes := []rune(strings.Join(strContent, " "))
	// for i := 1; i < len(strContent); i++ {
	// 	if runes[i] == 39 { // 39 = apostrophe
	// 		// if apostrophe is next to a space and a letter
	// 		if runes[i+1] == ' ' && (runes[i+2] <= 'A' && runes[i+2] >= 'Z' || runes[i+2] <= 'a' && runes[i+2] >= 'z') {
	// 			runes[i], runes[i+1] = runes[i+1], runes[i]
	// 		}
	// 	}
	// }
	// runeString := string(runes)
	// sliceString := strings.Split(runeString, " ")
	// return sliceString

	for i := 0; i < len(strContent); i++ {
		quote := []string{"'"}
		runeArr := []rune(strContent[i])
		counter := 0
		if Contains(quote, string(runeArr[0])) { // if rune array contains quotation marks
			// for last apostrophe
			if len(runeArr) == 1 {
				strContent[i-1] += string(runeArr[0])
				strContent = append(strContent[:i], strContent[i+1:]...) // moves apostrophe back one space if by itself
				counter++
			}
		} else if Contains(quote, string(runeArr[0])) && runeArr[1] == ' ' && counter == 1 {
			strContent[i+1] += string(runeArr[0])
			strContent = append(strContent[:i], strContent[i+1:]...)
		}
	}
	return strContent
	// counter := 0
	// for i, word := range strContent {
	// 	if word == "'" && counter == 0 {
	// 		counter += 1
	// 		strContent[i+1] = word + strContent[i+1]
	// 		strContent = append(strContent[:1], strContent[i+1:]...)
	// 	}
	// }
	// for i, word := range strContent {
	// 	if word == "'" {
	// 		strContent[i-1] = strContent[i-1] + word
	// 		strContent = append(strContent[:1], strContent[i+1:]...)
	// 	}
	// }
	// return strContent
}

func main() {
	args := os.Args
	input, err := os.ReadFile(args[1])
	if err != nil {
		log.Fatal(err)
	}

	str := string(input)
	strContent := strings.Split(str, " ")     // separate by whitespaces
	converted := convertString(strContent)    // hex, bin, cap etc
	convertedA := aToAn(converted)            // a to an
	convertedP := fixPunct(convertedA)        // fix non-quotes punctuation
	convertedAP := fixApostrophes(convertedP) // fixes quotation marks (last)
	// convertedAP1 := fixApostrophes1(convertedAP) // fixes quotation marks (first)

	fmt.Println(strContent)
	fmt.Println(converted)
	fmt.Println(convertedA)
	fmt.Println(convertedP)
	fmt.Println(convertedAP)
	// fmt.Println(convertedAP1)

	// os.create
}

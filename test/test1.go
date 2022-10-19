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

// global counters for apostrophes
var (
	quoteCount int
	counter    int = 0
)

// find and fix apostrophes
func fixApostrophes(strContent []string) []string {
	for i := 0; i < len(strContent); i++ {
		quote := []string{"'"}
		stringContent := strings.Join(strContent, "")
		if Contains(quote, string(stringContent)) && quoteCount == 0 {
			quoteCount++
			strContent[i+1] = strContent[i] + strContent[i+1]
			strContent, counter = RemoveIndex(strContent, i, counter)
		} else if quoteCount == 1 {
			strContent[i-1] = strContent[i-1] + strContent[i]
			strContent, counter = RemoveIndex(strContent, i, counter)
		}
	}
	return strContent

	// for i, quote := range strContent {
	// 	if quote == "'" && quoteCount == 0 {
	// 		quoteCount++
	// 		strContent[i+1] = strContent[i] + strContent[i+1]
	// 		strContent, counter = RemoveIndex(strContent, i, counter)
	// 	} else if quoteCount == 1 && quote == "'" {
	// 		quoteCount--
	// 		strContent[i-1] = strContent[i-1] + strContent[i]
	// 		strContent, counter = RemoveIndex(strContent, i, counter)
	// 	}
	// }
	// return strContent
}

// remove index values
func RemoveIndex(strContent []string, index int, counter int) ([]string, int) {
	counter++
	return append(strContent[:index], strContent[index+1:]...), counter
}

// func Recursive(strContent []string) []string {
// 	for _, current := range strContent {
// 		currentRune := []rune(current)
// 		if len(currentRune) == 1 {
// 			if currentRune[0] == 39 {
// 				strContent = fixApostrophes(strContent, quoteCount)
// 				Recursive(strContent)
// 				break
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
	strContent := strings.Split(str, " ")    // separate by whitespaces
	converted := convertString(strContent)   // hex, bin, cap etc
	convertedA := aToAn(converted)           // a to an
	convertedP := fixPunct(convertedA)       // fix non-quotes punctuation
	convertedQ := fixApostrophes(convertedP) // fix quotation marks

	fmt.Println(strContent)
	fmt.Println(converted)
	fmt.Println(convertedA)
	fmt.Println(convertedP)
	fmt.Println(convertedQ)

	// os.create
}

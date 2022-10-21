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
		case "(hex)": // checks if string is hex
			convHex, _ := strconv.ParseInt(strContent[i-1], 16, 64) // converts string to int with base 16
			emptyString[len(emptyString)-1] = fmt.Sprint(convHex)   // assigns + returns convHex while assigning it to emptyString at index-1
		case "(bin)":
			convBin, _ := strconv.ParseInt(strContent[i-1], 2, 64)
			emptyString[len(emptyString)-1] = fmt.Sprint(convBin)
		case "(cap)":
			convCap := strings.Title(strContent[i-1])
			emptyString[len(emptyString)-1] = fmt.Sprint(convCap)
		case "(cap,":
			deleteBracket := strings.Trim(strContent[i+1], ")") // trims last bracket from statement, keeps number
			strToInt, err := strconv.Atoi(deleteBracket)        // converts number (string) to int
			if err != nil {
				panic(err)
			}
			for j := 1; j <= strToInt; j++ { // loops thru slice of strings again, starting from last string
				convCap := strings.Title(emptyString[len(emptyString)-j]) // capitalises words in string
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
			emptyString = append(emptyString, strContent[i]) // append strContent to emptyString
		}
	}
	return emptyString
}

// convert 'a' to 'an' when next word begins with a vowel or 'h'
func aToAn(strContent []string) []string {
	for i := 0; i < len(strContent); i++ {
		if strContent[i] == "a" || strContent[i] == "A" {
			runes := []rune(strContent[i+1])
			if runes[0] == 'a' || runes[0] == 'e' || runes[0] == 'i' || runes[0] == 'o' || runes[0] == 'u' || runes[0] == 'h' {
				strContent[i] += "n"
			}
		}
	}
	return strContent
}

// turns i into a whitespace
func spacer(runes []rune, i int) []rune {
	runes = append(runes[:i+1], runes[i:]...) // appends everything before and including i+1 of slice to everything after slice after i
	runes[i] = 32                             // r index becomes a space
	return runes
}

// checks if punctuation
func punctCheck(c rune) bool {
	checker := false
	if c >= 33 && c <= 47 || c == 58 || c == 59 || c == 63 {
		checker = true
	}
	return checker
}

func fixPunct(s string) string {
	runes := []rune(s)
	for i, ch := range runes {
		// Delete whitespace after apostrophe
		// if apostrophe is not the last character and there is a space after it
		if i < len(runes)-1 && (ch == 39 || ch == 96) && runes[i+1] == 32 {
			runes = append(runes[:i+1], runes[i+2:]...) // append everything excluding the space after the apostrophe
		}
		// Deleting whitespace before all punctuation
		if (i < len(runes)-1 && punctCheck(ch) && runes[i-1] == 32) || (i == len(runes)-1 && punctCheck(ch) && runes[i-1] == 32) {
			runes = append(runes[:i-1], runes[i:]...) // append everything excluding the space before punctuation
		}
		// Inserting whitespace after commas or colons or semi-colons
		// if the comma/colon/semi-colon is not the last character and there is not a space after it
		if i < len(runes)-1 && ch == 44 && runes[i+1] != 32 || i < len(runes)-1 && ch == 58 && runes[i+1] != 32 || i < len(runes)-1 && ch == 59 && runes[i+1] != 32 {
			runes = spacer(runes, i) // add space after i
		}
	}
	return string(runes)
}

// function to find out whether a string contains another string
// func Contains(elems []string, v string) bool {
// 	for _, s := range elems {
// 		if v == s {
// 			return true
// 		}
// 	}
// 	return false
// }

// find and fix apostrophes
func fixApostrophes(strContent string) string {
	str := ""        // empty string
	var checker bool // boolean to check
	for i, ch := range strContent {
		if (ch == 39 || ch == 96) && strContent[i-1] == ' ' { // if ch is an apostrophe and there is a space before it
			if checker { // if checker is true
				str = str[:len(str)-1] // remove last character from string
				str = str + string(ch) // add apostrophe to string
				checker = false        // set checker to false
			} else { // if checker is false
				str = str + string(ch) // add apostrophe to string
				checker = true         // set checker to true
			}
		} else if i > 1 && (strContent[i-2] == 39 || strContent[i-2] == 96) && strContent[i-1] == ' ' {
			// if not the first character and character before is a space
			// and character before the space is an apostrophe
			if checker {
				str = str[:len(str)-1]
				str = str + string(ch)
			} else {
				str = str + string(ch)
			}
		} else {
			str = str + string(ch)
		}
	}
	return str
}

// checks for double spaces
func fixSpaces(strContent string) string {
	runes := []rune(strContent)
	for i := 0; i < len(runes); i++ {
		if runes[i] == ' ' && runes[i+1] == ' ' {
			runes = append(runes[:i], runes[i+1:]...)
		}
	}
	return string(runes)
}

func main() {
	args := os.Args
	input, err := os.ReadFile(args[1])
	if err != nil {
		log.Fatal(err)
	}

	str := string(input)
	str1 := strings.Split(str, " ")           // separate by whitespaces
	str2 := convertString(str1)               // hex, bin, cap etc
	str3 := aToAn(str2)                       // a to an
	str4 := fixPunct(strings.Join(str3, " ")) // fix non-quotes punctuation
	str5 := fixApostrophes(str4)              // fix quotation marks
	str6 := fixSpaces(str5)                   // fix double spaces

	err = os.WriteFile("result.txt", []byte(str6), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

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
		if strContent[i] == "a" || strContent[i] == "A" {
			runes := []rune(strContent[i+1])
			if runes[0] == 'a' || runes[0] == 'e' || runes[0] == 'i' || runes[0] == 'o' || runes[0] == 'u' || runes[0] == 'h' {
				strContent[i] += "n"
			}
		}
	}
	return strContent
}

// find and fix non-quote punctuation (if space BEFORE punctuation)
// turns i into a whitespace
func Insert(r []rune, i int) []rune {
	r = append(r[:i+1], r[i:]...)
	r[i] = 32 // 32 = space
	return r
}

// checks if punctuation
func IsPunctuation(c rune) bool {
	checker := false
	if c >= 33 && c <= 47 || c == 58 || c == 59 || c == 63 {
		checker = true
	}
	return checker
}

func RemoveAddWhiteSpace(s string) string {
	r := []rune(s)
	for i, ch := range r {
		// Delete whitespace after apostrophe
		if i < len(r)-1 && ch == 39 && r[i+1] == 32 {
			r = append(r[:i+1], r[i+2:]...)
		}
		// Deleting whitespace before all punctuation
		if (i < len(r)-1 && IsPunctuation(ch) && r[i-1] == 32) || (i == len(r)-1 && IsPunctuation(ch) && r[i-1] == 32) {
			r = append(r[:i-1], r[i:]...)
		}
		// Inserting whitespace after commas or colons or semi-colons
		if i < len(r)-1 && ch == 44 && r[i+1] != 32 || i < len(r)-1 && ch == 58 && r[i+1] != 32 || i < len(r)-1 && ch == 59 && r[i+1] != 32 {
			r = Insert(r, i)
		}
	}
	return string(r)
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

// find and fix apostrophes
func fixApostrophes(strContent string) string {
	str := ""
	var removeSpace bool
	for i, char := range strContent {
		if char == 39 && strContent[i-1] == ' ' {
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(char)
				removeSpace = false
			} else {
				str = str + string(char)
				removeSpace = true
			}
		} else if i > 1 && strContent[i-2] == 39 && strContent[i-1] == ' ' {
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(char)
			} else {
				str = str + string(char)
			}
		} else {
			str = str + string(char)
		}
	}
	return str
}

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
	str1 := strings.Split(str, " ")                      // separate by whitespaces
	str2 := convertString(str1)                          // hex, bin, cap etc
	str3 := aToAn(str2)                                  // a to an
	str4 := RemoveAddWhiteSpace(strings.Join(str3, " ")) // fix non-quotes punctuation
	str5 := fixApostrophes(str4)                         // fix quotation marks
	str6 := fixSpaces(str5)                              // fix double spaces

	err = os.WriteFile("result.txt", []byte(str6), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

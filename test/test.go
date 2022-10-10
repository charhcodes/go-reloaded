package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Hex() {
	hexa := "F1"
	decimal, err := strconv.ParseInt(hexa, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s = %v", hexa, decimal)
}

func Bin() {
	n := int64(123)
	fmt.Printf("%b", n)
}

func Upper() {
	str := "Hello World"
	upperStr := strings.ToUpper(str)
	fmt.Println(upperStr)
}

func Lower() {
	str := "HELLO WORLD"
	lowerStr := strings.ToLower(str)
	fmt.Println(lowerStr)
}

func Cap() {
	str := "hello world"
	capStr := strings.Title(str)
	fmt.Println(capStr)
}

func Separate() {
	str := "Hello world, good morning"
	separatedStr := strings.Fields(str)
	for _, separatedStr := range separatedStr {
		fmt.Println(separatedStr)
	}
}

func Punctuation() {
	rune := '!' // need to change string to rune
	checkRune := unicode.IsPunct(rune)
	fmt.Println(checkRune)
}

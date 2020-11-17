package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var bigDigits = [][]string{
	{
		"   00  ",
		"  0  0 ",
		" 0    0",
		" 0    0",
		" 0    0",
		"  0  0 ",
		"   00  "},
	{
		"   11 ",
		"  111  ",
		" 1 11  ",
		"   11  ",
		"   11  ",
		"   11  ",
		" 111111"},
	{
		"  2222 ",
		" 2    2",
		"     2 ",
		"    2  ",
		"   2   ",
		" 2     ",
		" 222222"},
	{
		"  9999 ",
		" 9    9",
		" 9    9",
		"  99999",
		"      9",
		"      9",
		"  9999 "}}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	shortUnderbarBool := flag.Bool("b", false, "short flag to draw an underbar and an overbar")
	longUnderbarBool := flag.Bool("bar", false, "long flag to draw an underbar and an overbar")

	flag.Parse()

	fmt.Println("shortUnderbarBool:", *shortUnderbarBool)
	fmt.Println("longUnderbarBool:", *longUnderbarBool)

	stringOfDigits := os.Args[len(os.Args) - 1]

	fmt.Println("stringOfDigits:", stringOfDigits)

	// // alternative
	// for row := 0; row < len(bigDigits[0]); row++ {
	// 	line := ""
	// 	for column := 0; column < len(stringOfDigits); column++ {

	if (*shortUnderbarBool || *longUnderbarBool) {
		underbar := strings.Repeat("*", len(bigDigits[0]) * len(stringOfDigits) + len(stringOfDigits)) 
		fmt.Println(underbar)
	}

	for row := range bigDigits[0] {
		line := ""
		for column := range stringOfDigits {
			// fmt.Println(column, stringOfDigits[column], (stringOfDigits[column] - '0'))
			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + " "
			} else {
				log.Fatal("invalide whole number")
			}
		}
		fmt.Println(line)
	}

	if (*shortUnderbarBool || *longUnderbarBool) {
		underbar := strings.Repeat("*", len(bigDigits[0]) * len(stringOfDigits) + len(stringOfDigits)) 
		fmt.Println(underbar)
	}
}
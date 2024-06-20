package main

import (
	"fmt"
)

type Roman struct {
	Symbol string
	Value  int
}

func main() {
	fmt.Println(romanToInt("DCXXI"))
}

func romanToInt(s string) int {
	var totalSum int
	var romanValue string
	var romanInArabicSingle int
	var romanInArabiCombo int

	var romanNumerals = []Roman{
		{"I", 1},
		{"V", 5},
		{"X", 10},
		{"L", 50},
		{"C", 100},
		{"D", 500},
		{"M", 1000},
		{"IV", 4},
		{"IX", 9},
		{"XL", 40},
		{"XC", 90},
		{"CD", 400},
		{"CM", 900},
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(romanNumerals); j++ {
			if i < len(s) && string(s[i]) == romanNumerals[j].Symbol {
				romanInArabicSingle = romanNumerals[j].Value
			}
			if i+1 < len(s) && s[i:i+2] == romanNumerals[j].Symbol {
				romanInArabiCombo = romanNumerals[j].Value
			}
			if i+1 < len(s) && romanNumerals[j].Symbol == s[i:i+2] {
				romanValue = s[i : i+2]
				i += 1
				break
			} else {
				romanValue = string(s[i])
			}
		}

		if len(romanValue) == 2 {
			totalSum += romanInArabiCombo
		} else {
			totalSum += romanInArabicSingle
		}
	}
	return totalSum
}

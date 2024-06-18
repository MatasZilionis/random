package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var totalCount int = 0

var possibleNumber string
var symbolFlag bool = false

var previousSymbol string = ""
var curentValue string = ""

func main() {

	var topLayerArray []string
	var middleLayerArray []string
	var bottomLayerArray []string

	file, err := os.Open("./engine_values.txt")
	checkNilErr(err)
	defer file.Close()
	// Create a buffer to read a file
	r := bufio.NewReader(file)
	// For loop to read through the whole file. Calculate the total result.
	// And stop once the file does not have any more data.
	line_1, _, _ := r.ReadLine()
	middleLayerArray = splitStringIntoArray(string(line_1))
	line_2, _, _ := r.ReadLine()
	bottomLayerArray = splitStringIntoArray(string(line_2))

	for {
		// Read the file line by line.
		line_3, _, err := r.ReadLine()
		if len(line_3) > 0 {
			for i := 0; i < len(middleLayerArray); i++ {
				curentValue = locateNumbersAndSymbols(middleLayerArray[i])
				// Check if a symbol is before a row of numbers.
				// Check adjacent rows above an below the number
				if curentValue == "number" {
					// Check for symbols diagnoly, above and below.
					if topLayerArray != nil && (i+1 < 140) && (i-1 > 0) {
						if locateNumbersAndSymbols(topLayerArray[i+1]) == "symbol" {
							symbolFlag = true
						} else if locateNumbersAndSymbols(topLayerArray[i-1]) == "symbol" {
							symbolFlag = true
						} else if locateNumbersAndSymbols(topLayerArray[i]) == "symbol" {
							symbolFlag = true
						}
					} else if topLayerArray != nil && locateNumbersAndSymbols(topLayerArray[i]) == "symbol" {
						symbolFlag = true
					}

					if bottomLayerArray != nil && (i+1 < 140) && (i-1 > 0) {
						if locateNumbersAndSymbols(bottomLayerArray[i+1]) == "symbol" {
							symbolFlag = true
						} else if locateNumbersAndSymbols(bottomLayerArray[i-1]) == "symbol" {
							symbolFlag = true
						} else if locateNumbersAndSymbols(bottomLayerArray[i]) == "symbol" {
							symbolFlag = true
						}
					} else if bottomLayerArray != nil && locateNumbersAndSymbols(bottomLayerArray[i]) == "symbol" {
						symbolFlag = true
					}

					if middleLayerArray != nil && (i+1 < 140) && (i-1 > 0) && locateNumbersAndSymbols(middleLayerArray[i+1]) == "symbol" {
						symbolFlag = true
					}

					if middleLayerArray != nil && (i+1 < 140) && (i-1 > 0) && locateNumbersAndSymbols(middleLayerArray[i-1]) == "symbol" {
						symbolFlag = true
					}

					possibleNumber = possibleNumber + middleLayerArray[i]
				} else if (curentValue == "dot" || curentValue == "symbol") && previousSymbol == "number" && symbolFlag {
					// Convert the string into a int.
					finalInt, err := strconv.Atoi(possibleNumber)
					// Check for errors.
					checkNilErr(err)
					// Add the number into the total count.
					totalCount = totalCount + finalInt
					// Clean up the part number string.
					possibleNumber = ""
					symbolFlag = false
				}
				// check something
				if curentValue == "dot" && previousSymbol == "number" && len(possibleNumber) > 0 && !symbolFlag {
					possibleNumber = ""
				}
				previousSymbol = curentValue
			}
			if symbolFlag {
				// Convert the string into a int.
				finalInt, err := strconv.Atoi(possibleNumber)
				// Check for errors.
				checkNilErr(err)
				// Add the number into the total count.
				totalCount = totalCount + finalInt
				// Clean up the part number string.
				possibleNumber = ""
				symbolFlag = false
			}
			previousSymbol = ""
			possibleNumber = ""
			topLayerArray = middleLayerArray
			middleLayerArray = bottomLayerArray
			bottomLayerArray = splitStringIntoArray(string(line_3))
		}
		if err != nil {
			break
		}
	}
	fmt.Println(totalCount)
}

// Take a string and convert it into a string array.
func splitStringIntoArray(checkString string) []string {
	// We create an array to store all of the string chars.
	var stringArray []string
	// We loop through the string and write the chars into the array.
	for i := 0; i < len(checkString); i++ {
		stringArray = append(stringArray, string(checkString[i]))
	}
	// Return an array of the string.
	return stringArray
}

// Receives an array char and tells what kind of char it is.
func locateNumbersAndSymbols(arrayChar string) string {
	r := []rune(arrayChar)
	switch {
	case unicode.IsDigit(r[0]):
		return "number"
	case strings.Contains(arrayChar, "."):
		return "dot"
	default:
		return "symbol"
	}
}

// Check if the function returned an error.
func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

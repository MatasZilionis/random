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

var (
	totalCount int = 0

	symbolFlag bool = false

	possibleNumber string
	previousSymbol string = ""
	curentValue    string = ""

	topLayerArray    []string
	middleLayerArray []string
	bottomLayerArray []string
)

func main() {

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
			processArray(middleLayerArray)
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

// Perform a scan on a specific line if there are any numbers with adjacent symbols.
func processArray(layerArray []string) {
	for i := 0; i < len(layerArray); i++ {
		curentValue = locateNumbersAndSymbols(layerArray[i])
		// Check if a symbol is before a row of numbers.
		// Check adjacent rows above an below the number
		if curentValue == "number" {
			// Check for symbols diagnoly, above, below, left and right.
			checkSymbol(topLayerArray, i+1)
			checkSymbol(topLayerArray, i)
			checkSymbol(topLayerArray, i-1)
			checkSymbol(middleLayerArray, i+1)
			checkSymbol(middleLayerArray, i-1)
			checkSymbol(bottomLayerArray, i+1)
			checkSymbol(bottomLayerArray, i)
			checkSymbol(bottomLayerArray, i-1)
			possibleNumber = possibleNumber + layerArray[i]
		} else if (curentValue == "dot" || curentValue == "symbol") && previousSymbol == "number" && symbolFlag {
			addToTolal()
		} else if curentValue == "dot" && previousSymbol == "number" && len(possibleNumber) > 0 && !symbolFlag {
			possibleNumber = ""
		}
		previousSymbol = curentValue
	}
	if symbolFlag {
		addToTolal()
	}
	// Data cleanup
	previousSymbol = ""
	possibleNumber = ""
}

// Add the possible number to the total count.
func addToTolal() {
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

// Check if there is a symbol on a specified layer and index.
func checkSymbol(layerArray []string, i int) {
	if layerArray != nil && i > 0 && i < len(layerArray) && locateNumbersAndSymbols(layerArray[i]) == "symbol" {
		symbolFlag = true
	}
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

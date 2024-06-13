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

var numberStringArray []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var numberIntStringArray []string = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func main() {
	// Create two veriables to hold temporary integer and
	// the final number when we will sum up all of the temporary integers
	var combinedTotalInt int = 0
	var tempIntHolder int

	// Create reference to the file we will be using, check for errors, defer file until needed.
	file, err := os.Open("./calibration_values.txt")
	checkNilErr(err)
	defer file.Close()
	// Create a buffer to read a file
	r := bufio.NewReader(file)
	// For loop to read through the whole file. Calculate the total result.
	// And stop once the file does not have any more data.
	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			tempIntHolder = convertStringIntoInt(string(line))
			fmt.Println(tempIntHolder)
			combinedTotalInt = combinedTotalInt + tempIntHolder
		}
		if err != nil {
			break
		}
	}
	// Print the total sum of integers
	fmt.Println(combinedTotalInt)
}

func convertStringIntoInt(unsplitString string) int {
	// Convert string into a string array for easier individual character analysis.
	var splitString []string = splitStringIntoArray(unsplitString)
	// Use previously made string array to find find first and last number inside of the string.
	var firstInt, lastInt string = findFirstAndLastNumber(splitString)
	// Convert string into an integer
	finalInt, err := strconv.Atoi(firstInt + lastInt)
	checkNilErr(err)
	// Return an integer
	return finalInt
}

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

func findFirstAndLastNumber(checkArray []string) (string, string) {
	// We create an array to store all of the string chars.
	var firstNumberAsString string
	var firstSymbolsUntilNumber string
	// We create an array to store all of the string chars.
	var lastNumberAsString string
	var lastSymbolsUntilNumber string
	// We loop through the given array and find the first number, assing it to firstNumberAsString and then break the for loop.
findFirstNumberForLoop:
	for i := 0; i < len(checkArray); i++ {
		// Turn current string symbol into a rune for the isDigit function.
		r := []rune(checkArray[i])
		// Check if the current string symbol is digit.
		if unicode.IsDigit(r[0]) {
			// Save the digit as a first number and break the loop.
			firstNumberAsString = checkArray[i]
			break findFirstNumberForLoop
		}
		// Add the latest symbol from the front into a string which keeps a save of already checked values.
		firstSymbolsUntilNumber = firstSymbolsUntilNumber + checkArray[i]
		for i := 0; i < len(numberStringArray); i++ {
			// If the symbolsave file contains a word which is a number, we use it and end the loop, since a word came before a digit.
			if strings.Contains(firstSymbolsUntilNumber, numberStringArray[i]) {
				// Save the first word as a digit and break the loop
				firstNumberAsString = numberIntStringArray[i]
				break findFirstNumberForLoop
			}
		}
	}
	// We loop through the given array and find the last number, assing it to lastNumberAsString and then break the for loop.
findLastNumberForLoop:
	for i := (len(checkArray) - 1); i >= 0; i-- {
		// Turn current string symbol into a rune for the isDigit function.
		r := []rune(checkArray[i])
		// Check if the current string symbol is digit.
		if unicode.IsDigit(r[0]) {
			// Save the digit as a last number and break the loop.
			lastNumberAsString = checkArray[i]
			break findLastNumberForLoop
		}
		// Add the latest symbol from the back into a string which keeps a save of already checked values.
		lastSymbolsUntilNumber = checkArray[i] + lastSymbolsUntilNumber
		for i := 0; i < len(numberStringArray); i++ {
			// If the symbolsave file contains a word which is a number, we use it and end the loop, since a word came before a digit.
			if strings.Contains(lastSymbolsUntilNumber, numberStringArray[i]) {
				// Save the last word as a digit and break the loop
				lastNumberAsString = numberIntStringArray[i]
				break findLastNumberForLoop
			}
		}
	}
	// Return first and last number found the string array. Both numbers are returned in string format.
	return firstNumberAsString, lastNumberAsString
}

// Check if the function returned an error.
func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

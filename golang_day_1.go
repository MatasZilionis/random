package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

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
	// We create an array to store all of the string chars.
	var lastNumberAsString string
	// We loop through the given array and find the first number, assing it to firstNumberAsString and then break the for loop.
findFirstNumberForLoop:
	for i := 0; i < len(checkArray); i++ {
		r := []rune(checkArray[i])
		if unicode.IsDigit(r[0]) {
			firstNumberAsString = checkArray[i]
			break findFirstNumberForLoop
		}
	}
	// We loop through the given array and find the last number, assing it to lastNumberAsString and then break the for loop.
findLastNumberForLoop:
	for i := (len(checkArray) - 1); i >= 0; i-- {
		r := []rune(checkArray[i])
		if unicode.IsDigit(r[0]) {
			lastNumberAsString = checkArray[i]
			break findLastNumberForLoop
		}
	}
	// Return first and last number found the string array. Both numbers are returned in
	return firstNumberAsString, lastNumberAsString
}

// Check if the function returned an error.
func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

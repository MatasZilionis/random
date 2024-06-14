package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	// Values of the maximum each color cubes.
	totalRed   int = 12
	totalGreen int = 13
	totalBlue  int = 14
)

// Regex pattern which will find the numbers inside of the string
var re *regexp.Regexp = regexp.MustCompile(`\d+`)

func main() {
	// Variable which will hold the sum of possible game id's
	var sumOfGameIDs int
	// Create reference to the file we will be using, check for errors, defer file until needed.
	file, err := os.Open("./puzzle_input.txt")
	checkNilErr(err)
	defer file.Close()
	// Create a buffer to read a file
	r := bufio.NewReader(file)
	// For loop to read through the whole file. Calculate the total result.
	// And stop once the file does not have any more data.
	for {
		// Read the file line by line.
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			// Split the string into game and cube values.
			splitDetails := strings.Split(string(line), ":")
			// Get game ID
			gameID, _ := strconv.Atoi(re.FindString(splitDetails[0]))
			// Get cube values as an array by splitting them at the ";" symbol.
			cubeValues := strings.Split(splitDetails[1], ";")
			// Check if the game the game can happen by following the possible cube at one time rule.
			if checkGame(cubeValues) {
				sumOfGameIDs = sumOfGameIDs + gameID
			}
		}
		if err != nil {
			break
		}
	}
	// Print the total sum of integers
	fmt.Println(sumOfGameIDs)

}

func checkGame(cubeValues []string) bool {
	// Variables to hold temp numbers
	var numberOfRed int
	var numberOfGreen int
	var numberOfBlue int
	// Loop through each index in the array
	for i := 0; i < len(cubeValues); i++ {
		// Split the index into unique values
		individualValues := strings.Split(cubeValues[i], ",")
		// Loop through each value in the array
		for j := 0; j < len(individualValues); j++ {
			switch {
			case strings.Contains(individualValues[j], "red"):
				numberOfRed, _ = strconv.Atoi(re.FindString(individualValues[j]))
			case strings.Contains(individualValues[j], "green"):
				numberOfGreen, _ = strconv.Atoi(re.FindString(individualValues[j]))
			case strings.Contains(individualValues[j], "blue"):
				numberOfBlue, _ = strconv.Atoi(re.FindString(individualValues[j]))
			default:
				continue
			}
			// If the found value is more than possible, return false. As in the game is not possible.
			if totalRed < numberOfRed || totalGreen < numberOfGreen || totalBlue < numberOfBlue {
				return false
			}
		}
	}
	// If all checks passed, return true.
	return true
}

// Check if the function returned an error.
func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

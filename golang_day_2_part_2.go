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

// Regex pattern which will find the numbers inside of the string
var re *regexp.Regexp = regexp.MustCompile(`\d+`)

func main() {
	// Variable which will hold the sum of possible game id's
	var sumOfCubes int
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
			// Get cube values as an array by splitting them at the ";" symbol.
			cubeValues := strings.Split(splitDetails[1], ";")
			// Check if the game the game can happen by following the possible cube at one time rule.
			sumOfCubes = sumOfCubes + checkGame(cubeValues)
		}
		if err != nil {
			break
		}
	}
	// Print the total sum of integers
	fmt.Println(sumOfCubes)
}

func checkGame(cubeValues []string) int {
	// Variables to hold temp numbers
	var numberOfRed int
	var numberOfGreen int
	var numberOfBlue int
	// Variables to hold highest number of the color
	var highestNumberOfRed int
	var highestNumberOfGreen int
	var highestNumberOfBlue int
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
			// Keep track of the highest number of cubes for each color yet.
			switch {
			case highestNumberOfRed < numberOfRed:
				highestNumberOfRed = numberOfRed
			case highestNumberOfGreen < numberOfGreen:
				highestNumberOfGreen = numberOfGreen
			case highestNumberOfBlue < numberOfBlue:
				highestNumberOfBlue = numberOfBlue
			default:
				continue
			}
		}
	}

	// If all checks passed, return true.
	return highestNumberOfRed * highestNumberOfGreen * highestNumberOfBlue
}

// Check if the function returned an error.
func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

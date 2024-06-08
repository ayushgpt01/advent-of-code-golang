package day1

import (
	"errors"
	"strconv"
	"strings"
)

const allDigits string = "123456789"

// HELPERS

// parseString splits the input string into individual lines and removes leading and trailing whitespace.
// It takes a string input representing the calibration document and returns a slice of strings, each representing a line in the document.
func parseString(input string) []string {
	return strings.Fields(strings.TrimSpace(input))
}

// parseNumberFromString converts a string representation of a number to an integer.
// It takes a string input representing a number and returns the corresponding integer value.
// If the input string is not a valid number, it returns -1 and an error.
func parseNumberFromString(input string) (int, error) {
	i, err := strconv.Atoi(input)
	if err != nil {
		return -1, err
	}
	return i, nil
}

// parseWordNumbers converts word representations of numbers to their corresponding digit representations.
// It takes a string input representing a word or digit and returns its digit representation.
// If the input string is not a word representation of a number, it returns the input string unchanged.
func parseWordNumbers(input string) string {
	digitNames := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"zero":  "0",
	}

	value, exists := digitNames[input]

	if !exists {
		return input
	}

	return value
}

// searchSubStrings searches for a substring within a larger string, either from the beginning or the end.
// It takes a string input representing the larger string, a slice of strings representing substrings to search for,
// and a boolean flag indicating whether to search from the beginning or the end of the larger string.
// It returns the substring found and an error if the search fails.
func searchSubStrings(input string, searchArr []string, isBegin bool) (string, error) {
	if input == "" {
		return "", errors.New("empty input string")
	}

	if len(searchArr) == 0 {
		return "", errors.New("empty search array")
	}

	index := -1
	value := ""

	for _, val := range searchArr {
		idx := -1
		if isBegin {
			idx = strings.Index(input, val)
		} else {
			idx = strings.LastIndex(input, val)
		}

		if idx == -1 {
			continue
		}

		if isBegin {
			if index == -1 || idx < index {
				index = idx
				value = val
			}
		} else {
			if idx >= index {
				index = idx
				value = val
			}
		}
	}

	if value == "" {
		return "", errors.New("cannot find required element")
	}

	retValue := parseWordNumbers(value)
	return retValue, nil
}

// SOLUTIONS

// APPROACH -
// Vars - listOfStrings, total
// Parse the string to remove whitespaces and create separate strings with each new line
// Go through each string
// Get the number using this -
// 1. Start from beginning until you find a valid digit
// 2. Start from end until you find a valid digit
// 3. No valid digit found - discard the string and go to next one
// 4. Make a new number using the two digits
// Add the number to accumulator

// Solve computes the sum of calibration values by concatenating the first and last digits found in each line of the input string.
// It takes a string input representing the calibration document and returns the sum of the calibration values.
func Solve(input string) int {
	arrInput := parseString(input)
	total := 0

	for _, val := range arrInput {
		firstDigitIdx := strings.IndexAny(val, allDigits)

		if firstDigitIdx == -1 {
			continue
		}

		lastDigitIdx := strings.LastIndexAny(val, allDigits)

		numberInString := string(val[firstDigitIdx]) + string(val[lastDigitIdx])
		i, err := parseNumberFromString(numberInString)

		if err == nil {
			total += i
		}
	}

	return total
}

// APPROACH -
// Change Logic on how to find the digits

// SolvePart2 computes the sum of calibration values by identifying digits represented as words and concatenating them.
// It takes a string input representing the calibration document and returns the sum of the calibration values.
func SolvePart2(input string) int {
	arrInput := parseString(input)
	total := 0
	valid_digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, val := range arrInput {
		firstDigit, err := searchSubStrings(val, valid_digits, true)
		// println("Value: " + val + " | Begin : " + firstDigit)

		if err != nil {
			// println("Value: " + val + " | Begin : " + err.Error())
			continue
		}

		lastDigit, err2 := searchSubStrings(val, valid_digits, false)
		// println("Value: " + val + " | End : " + lastDigit)

		if err2 != nil {
			// println("Value: " + val + " | End : " + err2.Error())
			continue
		}

		numberInString := firstDigit + lastDigit
		i, err := parseNumberFromString(numberInString)

		if err == nil {
			total += i
		}
	}

	return total
}

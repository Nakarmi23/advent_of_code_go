package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func convertSpelledToNumber(s string) string {
	spelledNumberRegexList := []regexp.Regexp{
		*regexp.MustCompile("one"),
		*regexp.MustCompile("two"),
		*regexp.MustCompile("three"),
		*regexp.MustCompile("four"),
		*regexp.MustCompile("five"),
		*regexp.MustCompile("six"),
		*regexp.MustCompile("seven"),
		*regexp.MustCompile("eight"),
		*regexp.MustCompile("nine"),
	}

	for i, r := range spelledNumberRegexList {
		if r.MatchString(s) {
			return r.ReplaceAllString(s, strconv.Itoa(i+1))
		}
	}

	return s
}

func parseInputString(s string) string {
	runes := []rune(s)
	resultString := ""

	for i := 0; i < len(runes); i++ {
		resultString = resultString + string(runes[i])
		fmt.Println(resultString)
		resultString = convertSpelledToNumber(resultString)
	}

	return resultString
}

func getCalibration(s string) int {
	parsedString := parseInputString(s)
	characterFileterRegex := regexp.MustCompile("[a-zA-Z]")
	filtered := characterFileterRegex.ReplaceAllString(parsedString, "")
	calibrationString := filtered[0:1] + filtered[len(filtered)-1:]
	calibration, error := strconv.Atoi(calibrationString)
	if error != nil {
		return 0
	}
	return calibration
}

func main() {
	sum := 0
	input := []string{}
	for _, s := range input {
		sum += getCalibration(s)
	}
	fmt.Println(sum)
}

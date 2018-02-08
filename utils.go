package main

import (
	"regexp"
	//"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func empty(str string) bool {
	return str == ""
}

func countLines(str ...*string) int {
	var numlines int = 0

	for i := 0; i < len(str); i++ {
		lineBreakes := regexp.MustCompile(`\n`)
		lineBreakesMatches := lineBreakes.FindAllStringIndex(*str[i], -1)
		if lineBreakesMatches != nil {
			numlines += len(lineBreakesMatches)
		}
	}

	return numlines
}

func stringInSlice(x string, arr []string) bool {
	for _, elem := range arr {
		if elem == x {
			return true
		}
	}
	return false
}

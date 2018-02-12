package main

import (
	"errors"
	"regexp"
	"strings"
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

func InSlice(x byte, arr []byte) bool {
	for _, elem := range arr {
		if elem == x {
			return true
		}
	}
	return false
}

func debugSpaces(s string) string {
	noNewLines := strings.Replace(s, "\n", "N", -1)
	return strings.Replace(noNewLines, " ", "*", -1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func surroundings(s string, loc int, radius ...int) (string, error) {
	var rad = 5

	if radius != nil {
		rad = radius[0]
	}

	var (
		dist_to_start int = loc
		dist_to_end   int = len(s) - 1 - loc
		//min_dist      int = min(dist_to_start, dist_to_end)

		loc_out_of_range bool = loc >= len(s) || loc < 0
		//rad_out_of_range bool = rad > min_dist
	)

	if loc_out_of_range {
		return "", errors.New("location out of range")
	}

	var left, right int
	left = min(rad, dist_to_start)
	right = min(rad, dist_to_end)

	return s[loc-left : loc+right+1], nil

}

package main

import (
	"errors"
	"regexp"
	"strings"
)

var (
	QUOTES          = []byte{'"', '\''}
	ALL_QUOTES      = append(QUOTES, '`')
	ESCAPE     byte = '\\'
)

func getFirstLine(str string) (nextLine string) {
	breakLinePos := regexp.MustCompile(`\n`).FindStringIndex(str)
	if breakLinePos == nil {
		return str
	}
	return str[:breakLinePos[1]]
}

func getLastLine(str string) (lastLine string) {
	breakLinePos := regexp.MustCompile(`[^\n]*$`).FindStringIndex(str)
	if breakLinePos == nil {
		return str
	}
	return str[breakLinePos[0]:]
}

/**
 * Returns str identation pattern
 * @return pattern string     string formed only by \s's & \t's
 */
func getIdentPatt(str string) (pattern string) {
	pattPos := regexp.MustCompile(`^[(\s|\t)]?`).FindStringIndex(str)
	pattern = str[:pattPos[1]]
	return
}

/**
 * Given a code and a opening char (e.g.: `{`, `<`, `"`, `'` et al),
 * the algorithm will return the position of the block content
 * delimited by this last mentioned param and its counterpart (e.g.: `}`, `>`, `"`, `'` et al)
 *
 * @param  code         *string
 * @param  opening_char byte          e.g.: `{`, `<`, `"`, `'` et al
 * @param  ignore       []byte        This is a set of potentially harmful character.
 *                                    i.e.: other opening chars.
 *                                    Example: If we are interested in finding the counterpart of
 *                                    `{` in the following code:
 *
 *                                      ```
 *                                      func foo() int {
 *                                       str := "}}}}}}}}}"
 *                                       return 0
 *                                      }
 *                                      ```
 *
 *                                    Then, we should "ignore" the `}`'s in `str := "}}}}}}}}}"`
 *
 * @return loc          [2]int        The position of opening and closing chars
 * @return err          error
 */
func getChunk(code *string, opening_char byte, ignore []byte) (loc [2]int, err error) {

	var (
		opening_loc  int  = strings.Index(*code, string(opening_char))
		closing_char byte = getClosingChar(opening_char)
		parity       int  = 0
	)

	for i := opening_loc; i < len(*code); i++ {
		var char byte = (*code)[i]

		updateParity(&parity, code, i, opening_char, closing_char)

		if parity == 0 {
			closing_loc := i
			return [2]int{opening_loc, closing_loc}, nil

		} else if InSlice(char, ignore) {
			rest := (*code)[i:]
			ignore_loc, _ := getChunk(&rest, char, nil)
			i += ignore_loc[1]
		}
	}

	return loc, errors.New("chunk's closing tag 404 not found")
}

func getClosingChar(c byte) byte {
	var res byte

	switch c {
	case '{':
		res = '}'
	case '"':
	case '\'':
	case '`':
		res = c
	}

	return res
}

func updateParity(parity *int, code *string, i int, opening_char, closing_char byte) {
	var (
		char                    = (*code)[i]
		is_a_boundary_char bool = (char == opening_char) || (char == closing_char)
		is_an_opening_char bool = (InSlice(opening_char, QUOTES) && ((*parity)%2 == 0)) || (char == closing_char && !InSlice(opening_char, QUOTES))
		is_false_positive  bool = InSlice(opening_char, QUOTES) && ((*parity) > 0) && (char == opening_char) && ((*code)[i-1] == ESCAPE)
	)

	if is_a_boundary_char && !is_false_positive {
		if is_an_opening_char {
			(*parity)++

		} else {
			(*parity)--
		}
	}
}

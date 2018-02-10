package main

import (
	"testing"
	// "fmt"
	"strings"
)

func Test_GetChunk1(t *testing.T) {
	code := `func foo(){
  i++
  str := "asds}s\"\}\{d"
  return bar()
}`

	expected := code[11 : len(code)-1]

	opening_char := `{`
	ignore := []string{`"`, `'`, "`"}

	loc, err := getChunk(&code, opening_char, ignore)
	got := code[loc[0]+1 : loc[1]]

	if err != nil {
		t.Error(err)
		return
	}

	if !strings.EqualFold(got, expected) {
		t.Error(
			"\nFOR:", code,
			"\nEXPECTED:", debugSpaces(expected),
			"\nGOT:", debugSpaces(got),
			"\nAFTER:", code,
		)
	}
}

func Test_GetChunk2(t *testing.T) {
	// getChunk-itself. #inception LOL
	code :=
		`func getChunk(code *string, opening_char string, ignore []string) (loc [2]int, err error) {

  var (
    opening_loc  int    = strings.Index(*code, opening_char)
    closing_char string = getClosingChar(opening_char)
    parity       int    = 0
  )

  for i := opening_loc; i < len(*code); i++ {
    char := string((*code)[i])

    var (
      is_a_boundary_char bool = (char == opening_char) || (char == closing_char)
      is_an_opening_char bool = (InSlice(opening_char, QUOTES) && (parity%2 == 0)) || (char == closing_char && !InSlice(opening_char, QUOTES))
      is_false_positive  bool = InSlice(opening_char, QUOTES) && (parity > 0) && (char == opening_char) && (string((*code)[i-1]) == ESCAPE)
    )

    if is_a_boundary_char && !is_false_positive {
      if is_an_opening_char {
        parity++

      } else {
        parity--
      }
    }

    if parity == 0 {
      closing_loc := i
      return [2]int{opening_loc, closing_loc}, nil

    } else if InSlice(char, ignore) {
      rest := (*code)[i:]
      ignore_loc, _ := getChunk(&rest, char, nil)
      i += ignore_loc[1]
    }
  }

  return loc, errors.New("block's closing tag 404 not found")
}`

	expected := code[len(`func getChunk(code *string, opening_char string, ignore []string) (loc [2]int, err error) {`) : len(code)-1]

	opening_char := `{`
	ignore := []string{`"`, `'`, "`"}

	loc, err := getChunk(&code, opening_char, ignore)
	got := code[loc[0]+1 : loc[1]]

	if err != nil {
		t.Error(err)
		return
	}

	if !strings.EqualFold(got, expected) {
		t.Error(
			"\nFOR:", code,
			"\nEXPECTED:", debugSpaces(expected),
			"\nGOT:", debugSpaces(got),
			"\nAFTER:", code,
		)
	}
}

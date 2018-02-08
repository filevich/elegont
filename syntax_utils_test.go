package main

import (
	"testing"
	// "fmt"
	"strings"
)

func TestFindCounterPart(t *testing.T) {
	code := `func foo(){
  i++
  str := "asds}s\"\"\}\{d"
  return bar()
}`

	expected := code[11 : len(code)-1]

	opening_char := `{`
	ignore := []string{`"`, `'`, "`"}

	loc, err := findCounterPart(&code, opening_char, ignore)
	got := code[loc[0]+1 : loc[1]]

	if err != nil {
		t.Error(err)
		return
	}

	if !strings.EqualFold(got, expected) {
		t.Error(
			"\nFOR:", code,
			"\nEXPECTED:", humanSpaceDebug(expected),
			"\nGOT:", humanSpaceDebug(got),
			"\nAFTER:", code,
		)
	}
}

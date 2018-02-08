package main

import (
	"testing"
	// "fmt"
)

func TestFindCounterPart(t *testing.T) {
	code := `func foo(){
            i++
            str := "asds}sd"
            return bar()
           }`
	// strr := "as\"dsad" // ok. is valid string
	// strr := "as\\"dsad" // no ok. is not valid string

	opening_char := `{`
	ignore := []string{`"`, `'`, "`"}

	loc, err := findCounterPart(&code, opening_char, ignore)

	if err != nil {
		t.Log(err)
		return
	}

	t.Log(loc)
	t.Log(code[loc[0]:loc[1]])
}

func TestStringsManipulations(t *testing.T) {
	str := "sads"
	pstr := &str

	t.Log("BEFORE:", str, *pstr) // BEFORE: sads sads

	*pstr = (*pstr)[1:]

	t.Log("AFTER:", str, *pstr) // AFTER: ads ads

	strc := string(str)
	strc = strc[1:]

	t.Log("AFTER 2:", str, *pstr, strc) // AFTER 2: ads ads ds
}

func TestStringInSlice(t *testing.T) {
	isIn := stringInSlice(`}`, []string{`a`, `b`, `c`, `'`, `"`, `}`})
	t.Log(isIn)
}

func TestForLoop(t *testing.T) {
	str := "hello world"
	for i := 0; i < len(str); i++ {
		char := string(str[i])
		t.Log(i, char)
		if char == "w" {
			i++
		}
	}

	t.Log(str)
}

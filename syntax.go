package main

import (
	"fmt"
	"log"
	"regexp"
)

type Syntax map[Component]([]IVariant)

type Component string

type IVariant interface {
	Test(ego *string) (loc []int)

	// PRE-REQUISITE: *first line* of param `ego` must match variant's definition
	Get(ego *string, end int, syntax Syntax) (code string, err error)

	String() string
}

/***/

type inLine struct { // implements IVariant
	Definition *regexp.Regexp
}

func (variant *inLine) Test(ego *string) (loc []int) {
	return variant.Definition.FindStringIndex(*ego)
}

func (variant *inLine) String() string {
	return variant.Definition.String()
}

func (variant *inLine) Get(ego *string, end int, syntax Syntax) (code string, err error) {
	code = (*ego)[:end]
	*ego = (*ego)[end:]
	return code, nil
}

/***/

type inBlock struct { // implements IVariant
	Definition *regexp.Regexp
	Delimiters []Delimiter
}

func (variant *inBlock) String() string {
	return variant.Definition.String()
}

func (variant *inBlock) Test(ego *string) (loc []int) {
	return variant.Definition.FindStringIndex(*ego)
}

/**
 * MODUS-OPERANDI:
 * 0. *Cut* block header
 * 1. Identify delimiter
 * 2. *Cut* block's content (in its "ego form") by appling DELIMITERS_STRATEGY
 * 3. "Recursivly call" `Dissect` (thus getting its "go form")
 * 4. Luckly, return the result of #3 "nested" in #0
 */
func (variant *inBlock) Get(ego *string, end int, syntax Syntax) (code string, err error) {
	// Step 0.
	header := (*ego)[:end]
	*ego = (*ego)[end:]

	// Step 1.
	delimiter := IdentifyDelimiter(ego)
	if !variant.hasDefined(delimiter) {
		return "", &SyntaxError{}
	}

	// Step 2.
	egoContent := DELIMITERS_STRATEGY[delimiter](ego)

	// Step 3.
	transContent, err := Dissect(&egoContent, syntax)

	// Step 4.
	return header + " {" + transContent + "}", nil
}

func (variant *inBlock) hasDefined(del Delimiter) bool {
	for _, d := range variant.Delimiters {
		if d == del {
			return true
		}
	}

	return false
}

/***/

type Delimiter uint

const (
	TABS           Delimiter = iota // a la Python
	CURLY_BRACKETS                  // a la C
	KEYWORDS                        // a la Pascal
	ROUND_BRACKETS                  // for multiline imports, consts, vars et al.
	//BACKWARDS                     // a la Bash (e.g.: if..fi)
)

var Delimiters = map[string]Delimiter{
	"TABS":           TABS,
	"CURLY_BRACKETS": CURLY_BRACKETS,
	"ROUND_BRACKETS": ROUND_BRACKETS,
}

/**
 * PRE-REQUISITE:
 * 	First char of `ego` param is of type `Delimiter`
 * 	And by char, it means anything except (\s|\t)
 */
func IdentifyDelimiter(ego *string) Delimiter {
	whiteSpaces := regexp.MustCompile(`^(\s|\n|\t)*`)
	pos := whiteSpaces.FindStringIndex(*ego)
	firstNonWhiteChar := (*ego)[pos[1] : pos[1]+1]
	// the racist algorithm, lolz

	switch firstNonWhiteChar {
	case "{":
		return CURLY_BRACKETS
	case "<":
		return KEYWORDS
	case "(":
		return ROUND_BRACKETS
	default:
		return TABS
	}
}

type DelimiterStrategy func(ego *string) string

var DELIMITERS_STRATEGY = map[Delimiter](DelimiterStrategy){
	TABS:
	/**
	 * 0. Cut str's first line (most probably it's just one `\n`, since block header
	 *    has been already cutted from `ego` in previous step)
	 * 1. Get ident pattern (from the [now] 1st [but formerly 2nd] line of str)
	 * 2. Cut from *there* (included) (block's beginning) down to the first
	 *    line with *less* tab ident length (or end of string) from step 1 (block's ends)
	 */
	(func(ego *string) string {

		// Step 0.
		whiteChars := (*ego)[:len(getFirstLine(*ego))]
		*ego = (*ego)[len(getFirstLine(*ego)):]

		// Step 1.
		patt := getIdentPatt(getFirstLine(*ego))
		block := ""

		// Step 2.
		for {
			line := getFirstLine(*ego)
			if len(patt) <= len(getIdentPatt(line)) {
				block += line
				*ego = (*ego)[len(line):] // cut
			} else {
				break
			}
		}

		return whiteChars + block
	}),

	CURLY_BRACKETS: (func(ego *string) string {

		var opening_char byte = '{'
		ignore := ALL_QUOTES
		loc, err := getChunk(ego, opening_char, ignore)

		if err != nil {
			log.Fatal(err)
		}

		block := (*ego)[loc[0]+1 : loc[1]]
		(*ego) = (*ego)[loc[1]+1:]

		return block
	}),

	ROUND_BRACKETS: (func(ego *string) string {

		var opening_char byte = '('
		ignore := ALL_QUOTES
		loc, err := getChunk(ego, opening_char, ignore)

		if err != nil {
			log.Fatal(err)
		}

		block := (*ego)[loc[0]+1 : loc[1]]
		(*ego) = (*ego)[loc[1]+1:]

		return block
	}),
}

type SyntaxError struct {
	line int
	code string
	file string
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("\nRoses are red\nViolets are blue\nSyntax error\n\n %v:%v: %v\n", e.file, e.line, e.code)
}

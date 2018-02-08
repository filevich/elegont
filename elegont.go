package main

import (
	// "errors"
	"fmt"
	"regexp"
	//"encoding/json"
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "Elegont"
	app.Usage = "2009 is the new 1984"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "language for the greeting",
			//Destination: &language, // language = c.String("lang")
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration file from `FILE`",
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say! CANONICAL")
		for _, arg := range c.Args() {
			fmt.Println(arg)
		}

		return nil
	}

	app.Run(os.Args)

}

func Dissect(ego *string, syntax Syntax) (string, error) {
	var (
		procesedLines int    = 0
		output        string = ""
	)

	for !empty(*ego) {
		whiteSpaces := cutNextWhiteSpaces(ego)
		output += whiteSpaces
		procesedLines += countLines(&whiteSpaces)
		code, err := cutNextComponent(ego, syntax)

		if err != nil {
			if err, ok := err.(*SyntaxError); ok {
				err.line = procesedLines + 1
				err.file = "fileName.ego"
				err.code = (*ego)[:regexp.MustCompile(`^[^\n]+`).FindAllStringIndex(*ego, 1)[0][1]]
				return "", err
			}
		}

		output += code
		procesedLines += countLines(&code)
	}

	return output, nil
}

/**
 * Cuts the first consecutives (\s or \n)'s
 * @param str *string
 * @return cutted string 	removed string from str.
 */
func cutNextWhiteSpaces(str *string) (cutted string) {
	whiteSpaces := regexp.MustCompile(`^(\s|\n)*`)
	pos := whiteSpaces.FindStringIndex(*str)

	if pos != nil {
		cutted = (*str)[0:pos[1]]
	} else {
		cutted = ""
	}

	*str = (*str)[pos[1]:len(*str)]
	return cutted
}

/**
 * Cuts the first `Component` of code from `*str`, according to `syntax`
 * Idea: Iterates on syntax[i][j] and stops once regex match's index is 0.
 * @param  str *string 	ego code
 * @return code str 		go code
 */
func cutNextComponent(ego *string, syntax Syntax) (code string, err error) {

	for _, variants := range syntax {
		for _, variant := range variants {
			if pos := variant.Test(ego); pos != nil {
				isFirst := pos[0] == 0
				if isFirst {
					return variant.Get(ego, pos[1])
				}
			}
		}
	}

	return "", &SyntaxError{}
}

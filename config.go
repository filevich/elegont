package main

import (
	"io/ioutil"
	"regexp"
	"strings"
	"fmt"

	"github.com/iancoleman/strcase"
	"gopkg.in/yaml.v2"
)

type Identer string

const (
	TAB   Identer = `\t`
	SPACE Identer = `\s`
)

type Config struct {
	File_extension  string
	Input_dir       string
	Out_dir         string
	Recursive       bool
	Remove_comments bool
	Last_comma      bool
	Identation      Identer
	Syntax          Syntax
}

func NewConfig(filePath ...string) (*Config, error) {
	c := &Config{
		File_extension:  ".ego",
		Input_dir:       "./",
		Out_dir:         "./",
		Recursive:       true,
		Remove_comments: false,
		Last_comma:      false,
		Identation:      SPACE,
		Syntax:          nil,
	}

	if filePath != nil {
		c.loadConfig(filePath[0])
	} //else {
	//	c.Syntax = defaultSyntax
	//}

	return c, nil
}

func (c *Config) loadConfig(filePath ...string) {

	data, err := ioutil.ReadFile(filePath[0])
	check(err)

	splited := strings.Split(string(data), "syntax:")
	var optionsData string = splited[0]
	var syntaxData string = splited[1]

	err = yaml.Unmarshal([]byte(optionsData), c)
	check(err)

	var syntaxBuffer SyntaxDATA

	err = yaml.Unmarshal([]byte(syntaxData), &syntaxBuffer)
	check(err)

	// initialize to 0
	c.Syntax = make(map[Component]([]IVariant))

	for component, variants := range syntaxBuffer {
		for _, variantDATA := range variants {
			v := NewVariant(component, variantDATA)
			x := Component(strcase.ToCamel(component))
			c.Syntax[x] = append(c.Syntax[x], v)
		}
	}
}
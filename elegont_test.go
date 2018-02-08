package main

import (
	//"errors"
	// "errors"
	//"fmt"
	"io/ioutil"
	//"os"
	// "regexp"
	"testing"
)

func TestCustom(t *testing.T) {
	config, _ := NewConfig("./examples/config/Elegont.yaml")

	data, err := ioutil.ReadFile(config.Input_dir + "simple-bucle" + config.File_extension)
	check(err)

	ego := string(data)
	GoCode, err := Dissect(&ego, config.Syntax)

	if err != nil {
		if err, ok := err.(*SyntaxError); ok {
			err.Print()
		}
	}

	d1 := []byte(GoCode)
	err = ioutil.WriteFile(config.Out_dir+"simple-bucle.ego", d1, 0644)
	check(err)
}

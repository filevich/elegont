package main

import (
	//"errors"
	// "errors"
	//"fmt"
	"io/ioutil"
	//"os"
	// "regexp"
	"log"
	"testing"
)

func TestCustom(t *testing.T) {
	config, _ := NewConfig("./examples/config/Elegont-1.0.0.yaml")

	data, err := ioutil.ReadFile(config.Input_dir + "simple-bucle" + config.File_extension)
	check(err)

	ego := string(data)
	trans, err := Dissect(&ego, config.Syntax)

	if err != nil {
		log.Print(err)
		return
	}

	if err := ioutil.WriteFile(config.Out_dir+"simple-bucle.go", []byte(trans), 0644); err != nil {
		log.Print(err)
		return
	}
}

func TestCustom2(t *testing.T) {
	config, _ := NewConfig("./examples/config/Elegont-1.0.0.yaml")

	data, err := ioutil.ReadFile(config.Input_dir + "simple-file" + config.File_extension)
	check(err)

	ego := string(data)
	trans, err := Dissect(&ego, config.Syntax)

	if err != nil {
		log.Print(err)
		return
	}

	if err := ioutil.WriteFile(config.Out_dir+"simple-file.go", []byte(trans), 0644); err != nil {
		log.Print(err)
		return
	}
}

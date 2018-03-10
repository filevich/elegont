package main

import (
	"io/ioutil"
	"log"
	"testing"
	"regexp"
)

func Test_SimpleBucle(t *testing.T) {

	var (
		configPath      = "./examples/Elegont.yaml"
		config, _       = NewConfig(configPath)
		fileName        = config.Input_dir + "simple-bucle" + config.File_extension
		expectedData, _ = ioutil.ReadFile("./test/expected/simple-bucle.go")
		expected        = string(expectedData)
	)

	ilv := &IL_Variable{  
			&inLine{
				Definition: regexp.MustCompile(`let\s(?P<NAME>[[:alpha:]]+)\s(?P<TYPE>[[:alpha:]]+)\s\=\s(?P<VALUE>.+)`),
			}, 
		}

	config.Syntax["Variable"] = []IVariant{ ilv }

	data, err := ioutil.ReadFile(fileName)
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

	var (
		trans_slim    = ignoreSpaces(trans)
		expected_slim = ignoreSpaces(expected)
		diff          = diff(trans_slim, expected_slim)
		oops          = diff != -1
	)

	if oops {

		var (
			trans_min, _    = surroundings(trans_slim, diff, 10)
			expected_min, _ = surroundings(expected_slim, diff, 10)
		)

		t.Error(
			"\nINPUT:", fileName,
			"\nEXPECTED:", "`..."+trans_min+"...`",
			"\nGOT:", "\t `..."+expected_min+"...`",
			"\nRUN:", "`diff "+fileName+" "+config.Out_dir+"simple-bucle.go"+"`",
		)
	}
}

package main

// ALALA //return DELIMITERS_STRATEGY[variant.Delimiters[0]](*str)

import (
	"testing"

	"fmt"
	//"log"
	//	"regexp"

	"gopkg.in/yaml.v2"
)

func TestCustom(t *testing.T) {
	config, _ := NewConfig("./examples/config/Elegont.yaml")

	str := `package main

          import "fmt"

          func main(){
            let n int = 99
          }
            for i := 0; i < 10; i++
              fmt.Printf("Hello bro %v", n)
              hello++
              foo(j)
                if condition {
                  return null
                }
            n++`

	// cast
	component := (config.Syntax["For"][0]).(*inBlock)

	//t.Log(foo)

	var headerPos []int = component.Test(&str)
	asdasd := str[headerPos[0]:]

	t.Log(component.GetBlock(&asdasd))

	//t.Log(config.Syntax["Import"][0].Test(&str) == nil)
	/*


			  if config.Syntax["Import"][0] IS inBlock THEN:
			    header := pos[...]
			    block := config.Syntax["Import"][0].getBlock(str)
			    footer := algo

		      // idea:
		      // header, body, footer := config.Syntax["Import"][0].getIt(str)
			  config.Syntax["Import"][0].Test(&str)

	*/

	// dividerFuncs[TABS]("   asjkdjkaskdj   hola\n    como\n    andas\n    todo\n    bien\n   OTRO METODO")
	// t.Log(ROUND_BRACKETS + CURLY_BRACKETS)
}

func TestCustom2(t *testing.T) {
	config, _ := NewConfig("./examples/config/Elegont.yaml")

	str := `package main

          import "fmt"

          func main(){
            let n int = 99
          }
            for i := 0; i < 10; i++
              fmt.Printf("Hello bro %v", n)
              hello++
              foo(j)
                if condition {
                  return null
                }
            n++`

	// cast
	component := (config.Syntax["Func"][1]).(*inBlock)

	var headerPos []int = component.Test(&str)
	asdasd := str[headerPos[0]:]

	t.Log(component.GetBlock(&asdasd))

	//t.Log(config.Syntax["Import"][0].Test(&str) == nil)
	/*


	   if config.Syntax["Import"][0] IS inBlock THEN:
	     header := pos[...]
	     block := config.Syntax["Import"][0].getBlock(str)
	     footer := algo

	     // idea:
	     // header, body, footer := config.Syntax["Import"][0].getIt(str)
	   config.Syntax["Import"][0].Test(&str)

	*/

	// dividerFuncs[TABS]("   asjkdjkaskdj   hola\n    como\n    andas\n    todo\n    bien\n   OTRO METODO")
	// t.Log(ROUND_BRACKETS + CURLY_BRACKETS)
}

func TestLala(t *testing.T) {
	t.Log("hello from lala")
	t.Log(Kelel())
}

type T struct {
	File_extension  string
	Input_dir       string
	Out_dir         string
	Recursive       bool
	Remove_comments bool
	Last_comma      bool
	Identation      Identer
}

func TestYAML(j *testing.T) {

	var data = `
  input_dir: Easy!
  file_extension: asdd
  out_dir: ksks
  recursive:       true
  remove_comments: false
  last_comma:      true
  adsee: sadjnj
  jjs: dsdd
`

	t := T{}

	yaml.Unmarshal([]byte(data), &t)

	fmt.Printf("--- t:\n%v\n\n", t)
	//fmt.Printf("%T\n\n", t.Foo)

}

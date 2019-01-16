# elegont 

straight out of 1984 

----

The initial motivations were:

1. Learn Go
2. Fight GoLang syntactic rigidity and make it more stylysh (i.e., Phyton-ish/Phythonian) by creating this transpiler.

Quickly, it became a much more of ambitious proyect: a complete agnostic & customizable transpiler. The idea of making a GoLang transpiler which would compile pseudo-go in the major popular coding styles, e.g., K&R, Allman, Horstmann et. al. into \*.go files was challenging and exciting to me. Unfortunatly, the (originally fairly simple) architecutre of the program would not suppor this sort of complexity, augmenting code enthropy making coding a not-so-fun of an experience.

idea: (e.g.)
transpile this input: (`test/input/simple-bucle.ego`)
```
package main

import "fmt"

func main() {
  let n int = 99
  
  for i := 0; i < 10; i++
    fmt.Printf("Hello bro %v", n)
    
    for i := 0; i < 10; i++
      fmt.Printf("Hello bro %v", n)
}
```

using these settings: (`examples/elegont.yaml`)
```
# Elegont 2017 v1.0.0

file_extension:  .ego
input_dir:       ./test/input/
out_dir:         ./test/output/
recursive:       FALSE
remove_comments: TRUE
last_comma:      FALSE
identation:      TAB
syntax:

  import:
    - variant:    inLine
      definition: import\s\"[[:alpha:]]+\" # e.g.: import "fmt"

  comment:
    - variant:    inLine
      definition: someregexphere

  if:
    - variant:    inLine
      definition: someregexphere

  while:
    - variant:    inLine
      definition: someregexphere

  package:
    - variant:    inLine
      definition: package\s[[:alpha:]]+

  type:
    - variant:    inLine
      definition: someregexphere

  struct:
    - variant:    inLine
      definition: someregexphere

  variable:
    - variant:    inLine
      definition: let\s(?P<NAME>[[:alpha:]]+)\s(?P<TYPE>[[:alpha:]]+)\s\=\s(?P<VALUE>.+) # e.g.: let n int = 99

  for:
    - variant:    inBlock
      definition: for [^\n]+\; [^\n]+\; [^\n]+
      delimiters: [TABS]

  func:
    - variant:    inLine
      definition: ([[:word:]]+)(\.[[:word:]]+)?\([^\n]+\) # e.g.: fmt.Printf("Hello bro %v", n)

    - variant:    inBlock
      definition: func\s[[:alpha:]]+\(.?\) # e.g.: func main()
      delimiters: [CURLY_BRACKETS, TABS]

  inc:
    - variant:    inLine
      definition: ([[:word:]]+)\+\+ # e.g.: n++
```

outputing this (.go) output: (`test/expected/simple-bucle.go`)
```
package main

import "fmt"

func main() {
	var n int = 99

	for i := 0; i < 10; i++ {
		fmt.Printf("Hello bro %v", n)

		for i := 0; i < 10; i++ {
			fmt.Printf("Hello bro %v", n)
		}
	}
}
```

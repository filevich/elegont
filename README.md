# elegont 

straight out of 1984 

----

The original motivation was:

1. Learn Go
2. Fight GoLang syntactic rigidity and make it more stylysh (i.e., Phyton-ish/Phythonian) by creating this transpiler.

Quickly, it become a much more of ambitious proyect: a complete agnostic & customizable transpiler. Unfortunatly, the (originally fairly simple) architecutre of the program would not suppor this sort of complexity, augmenting code enthropy making coding a not-so-fun of experience.

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

into this (.go) output: (`test/expected/simple-bucle.go`)
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

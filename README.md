# elegont 

straight out of 1984 

----

**\#YetAnotherUnifinishedAndUnpolishedGitProject** 

**\#YouveBeenWarnedThough** 

---

How it all started:
> me@2017: "Ok, time to learn this go-lang that you all seem to be talking about". 

> alright let's make a simple if-statement: `if (true) foo()` 

> *got compilation error*

> hmmm, weird. *tries again*

> *got same compilation error*

> **Old Man Yells at ~~Cloud~~ Rob Pike, and tries to reinvent the wheel while learning how the wheel works.**


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

to this (.go) output: (`test/expected/simple-bucle.go`)
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

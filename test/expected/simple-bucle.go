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

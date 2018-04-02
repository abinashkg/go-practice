package main

import "fmt"

func main() {
	// Declare a function inside a function
	add := func(x int, y int) int {
		return x + y
	}
	fmt.Println(add(5, 8))
}

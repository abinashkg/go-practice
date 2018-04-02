package main

import "fmt"

func main() {
	x := 1
	increment := func() int {
		// function use local variable
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

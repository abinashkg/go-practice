package main

import "fmt"

func square(x *int) {
	*x = *x * *x
}

func cube(x int) *int {
	c := x * x * x
	return &c
}

func main() {
	x := 2
	square(&x)
	fmt.Println(x)
	fmt.Println(*cube(x))
}

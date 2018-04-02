package main

import "fmt"

func main() {
	x := 1
	p := &x         //p,oftype *int, points to x
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2
	fmt.Println(x)  // "2"

	q := new(int)   // p, of type *int, points to an unnamed int variable
	fmt.Println(*q) // "0"
	*q = 2          // sets the unnamed int to 2
	fmt.Println(*q) // "2"

	r := &q          //  Pointer to pointer
	fmt.Println(**r) // "2"
}

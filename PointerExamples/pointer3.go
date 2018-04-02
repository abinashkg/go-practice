package main

import "fmt"

// More details https://github.com/golangbot/pointers/blob/master/pointers.go
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

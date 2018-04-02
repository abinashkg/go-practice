package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name  string
	DOB   string
	adult bool
}

func (p *Person) checkIfAdult() {
	t, err := time.Parse("2006-01-02", p.DOB)
	if err == nil {
		if t.Year() < 1995 {
			p.adult = true
		}
	}
}

func main() {

	p := Person{Name: "ABC", DOB: "1983-11-01"}
	ptr := &p
	ptr.checkIfAdult()
	fmt.Println(*ptr)

}

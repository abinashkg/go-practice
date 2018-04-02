package main

import "fmt"

type Dog struct{}

func (d *Dog) Name() string {
	return "Dog"
}

type Cat struct{}

func (c *Cat) Name() string {
	return "Cat"
}

func (cd *CatDog) Name() string {
	return fmt.Sprintf("%s%s", cd.Cat.Name(), cd.Dog.Name())
}

type CatDog struct {
	*Cat
	*Dog
}

func main() {
	cd := &CatDog{}
	fmt.Printf("My favorite animal is the %s!\n", cd.Name())
}

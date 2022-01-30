package main

import "fmt"

// Package level scope
// var keyword is used to declare a variable in go
var y int
// When we declare a varialbe its automatically assign value zero
var z int

// data structures in go
// slice
// map
// struct
//	composite literal

type person struct {
	// small case means it won't available in outside
	// Lname i.e it will also available outside of package
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

// functions declaration in go
// func(receiver) identifier(parameters) (returns) {<code>}
func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	// Variable declaration
	// short variable declaration
	x := 7
	fmt.Println(x)
	// %T will return the variable data type
	fmt.Printf("%T", x)

	// Variable level scope
	y = 10
	fmt.Println(y)

	// it will return 0
	fmt.Println(z)

	// Composite literal is a type and a list of data
	xi := []int{2,4,7,9,42,}
	fmt.Println(xi)

	m := map[string]int{
		"Todd" : 45,
		"Job"  : 42,
	}

	fmt.Println(m)

	p1 := person{
		// If we use value in accordance with order then we don't need to give the key name  
		"Mr",
		"Abdullah",
	}
	fmt.Println(p1)

	p1.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}

	sa1.speak()
	sa1.person.speak()

	saySomething(p1)
	saySomething(sa1)
}
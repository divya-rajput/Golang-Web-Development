package main

import "fmt"

// Creating a Struct datatype
type person struct {
	fname string // if fields used outside the package,then 1st letter should be capital.For example:"Fname".
	lname string
}
type secretAgent struct {
	person
	liscenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says,"Hell Yeahh!!!`)
}
func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says,"Hell Yeahh!!!`)
}

// Implementing interfaces- Interfaces defines the functionality and allows Polymorphism. We create a human interface which will implement speak method.
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
func main() {
	//Examples of Composite literals.
	//Creating a slice of integer values.
	sliceInt := []int{2, 3, 4, 5, 6}
	fmt.Println("Slice of Int=", sliceInt)
	//Creating a map of string and int
	m := map[string]int{
		"Divya":  29,
		"Gaurav": 26,
	}
	fmt.Println("Map=", m)
	p1 := person{
		"Divya",
		"Rajput",
	}
	fmt.Println("My name is", p1.fname, p1.lname)
	p1.speak()

	sa1 := secretAgent{
		person{
			"Gaurav",
			"Saini",
		},
		true,
	}
	fmt.Println(sa1.fname, sa1.lname, sa1.liscenseToKill)
	sa1.speak()
	//Calling the methods.
	saySomething(p1)
	saySomething(sa1)

}

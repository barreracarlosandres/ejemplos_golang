package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // contact contactInfo
}

func getPerson(fn string, ln string) {
	alex := person{firstName: fn, lastName: ln}
	fmt.Println(alex)
}

func getPersonNew(fn string, ln string) {
	var alex person

	/*
		&variable - give me the memory address
		*pointer  - give me the value
	*/

	alex.updateLastName("newName") // can use '(&alex). or alex.'
	alex.print()                   //paso por value
}

func (p *person) updateLastName(newLastName string) {
	//func (p person) updateLastName(lastName string) {
	(*p).lastName = newLastName
}

func (p person) print() {
	fmt.Printf("%+v", p) // print the fiel names and values
	fmt.Println()
}

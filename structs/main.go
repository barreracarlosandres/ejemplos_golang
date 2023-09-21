package structs

import "fmt"

type person struct {
	firstName   string
	lastName    string
	contactInfo // contact contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func getPersonNew(fn string, ln string) {
	fmt.Println("* Execute func getNewPerson")
	var alex person

	/*
		&variable - give me the memory address
		*pointer  - give me the value
	*/	
	alex.updateLastName("newName") // can use '(&alex). or alex.'
	alex.print()                   //paso por value
}

func (p person) updateLastName(newLastName string) {
	fmt.Println("* Execute func updateLastName")
	p.lastName = newLastName
}

func (p person) print() {
	fmt.Println("* Execute func print")
	fmt.Printf("The new  person is: %+v \n",p)
}

func Example() {
	fmt.Println("------------ Struct Examples ------------")

	//create new person
	np := person{
		firstName: "Carlos",
		lastName:  "Barrera",
		contactInfo: contactInfo{
			email : "email@gmail.com",
			zipCode: 0,
		},
	}

	np.print()
	np.updateLastName("Barrera 21")
	np.print()

	//create persona 2
	var np2 person
	np2.firstName = "Carlos"

	np2.print()

		//getPerson("Alex", "Perez")
	//getPersonNew("Alex", "Perez")

	fmt.Println()
}

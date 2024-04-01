package main

import "fmt"

type contactinfo struct {
	email string
	zip   int
}

// type person struct {
// 	firstName string
// 	lastName  string
// 	contact   contactinfo
// }

type person struct {
	firstName string
	lastName  string
	contactinfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// Alternate way of creating new struct
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactinfo: contactinfo{
			email: "jim.party@gmail.com",
			zip:   80134,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateName("jimmy")
	jim.updateName("jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

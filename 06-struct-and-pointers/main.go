package main

import "fmt"

type person struct {
	firstName string
	lastName  string

	// Embedded struct, contact is the name of the struct and contactInfo is the type
	// contactInfo contactInfo // This is also valid
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {

	// Order of fields matters in struct definition and initialization!!
	// In this case firstName is "Alex" and lastName is "Anderson"
	// alex := person{"Alex", "Anderson"}

	// alex := person{lastName: "Anderson", firstName: "Alex"} // Better way to initialize struct and order of fields doesn't matter

	// var alex person, alex is initialized with zero values for each field
	// In this case firstName is "" and lastName is ""
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex) // %+v prints out the field names as well as the values

	jim := person{ // This is in a memory location and jim is a pointer to that memory location
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// fmt.Println("%+v", jim)
	// jim.print()
	// jim.updateName("Jimmy")

	// This is a pointer to the memory location of jim
	// & is an operator that gives the memory address of the value this variable is pointing at
	// jimPointer := &jim
	// jimPointer.updateName("Jimmy") // This is the same as jim.updateName("Jimmy")

	// This is a shortcut to the above two lines
	// Because the receiver function takes in a pointer to a person struct, go will automatically convert jim to a pointer to a person struct
	jim.updateName("Jimmy")

	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// func (p person) updateName(newFirstName string) {
// This update will not work because the receiver is a copy of the original person struct and not the original person struct because go is a pass by value language
// p.firstName = newFirstName

// In this case p gets updated which is pointing to a memory location of the copy of the person struct
// fmt.Print("%+v", p)

// }

// This is a receiver function that takes in a pointer to a person struct
// *person is a type description, it means we're working with a pointer to a person	and can only be called with a pointer to a person
func (pointerToPerson *person) updateName(newFirstName string) {
	// This is a pointer to the original person struct
	// * is an operator that gives the value this memory address is pointing at
	// *pointerToPerson is an operator, it means we want to manipulate the value the pointer is referencing
	(*pointerToPerson).firstName = newFirstName
}

// *************Rules on pointers*************
// Turn address into value with *address
// Turn value into address with &value

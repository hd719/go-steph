package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	updateSlice(mySlice)

	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

// Slice vs Struct
// Slice is a reference type meaning it points to an underlying array (original value)
// Struct is a value type meaning it points to a value in memory (copy of original value)

// func main() {
//     name := "Bill"

//     fmt.Println(*&name)
// }
// This will return the string Bill because the & operator returns the memory address of the value and * returns the value at the memory address

// package main

// import "fmt"

// func main() {
// 	name := "bill"
// 	namePointer := &name

// 	fmt.Println(&namePointer) // returns 0x14000116018
// 	printPointer(namePointer)
// }

// func printPointer(namePointer *string) {
// 	fmt.Println(&namePointer) // returns 0x14000052030
// }

// In the main function, fmt.Println(&namePointer) prints the memory address of the namePointer variable itself.
// In the printPointer function, fmt.Println(&namePointer) prints the memory address of the namePointer parameter of the function.
// Since these are two different variables (one in the main function and one in the printPointer function), their memory addresses will be different.

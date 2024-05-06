package main

import "fmt"

// Maps in go are like objects in javascript
// The keys are all of the same type and the values are all of the same type
// The keys are unique and are indexed

func main() {

	// 1st way to declare a map
	// Declaring a map with keys of type string and values of type string
	colors1 := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	// 2nd way to declare a map
	// Declaring a map with keys of type string and values of type string, but with no values instead the map is initialized with a zero value
	var colors2 map[string]string // returns map[]

	// 3rd way to declare a map same thing as the 2nd way but with the make function
	colors3 := make(map[string]string)
	colors3["white"] = "#ffffff"

	// Deleting a key value pair from a map
	delete(colors3, "white")

	fmt.Println(colors1)
	fmt.Println(colors2)
	fmt.Println(colors3)
	fmt.Println(colors3["white"])
	printMap(colors1)
}

// Iterating over a map
func printMap(c map[string]string) {
	for color, hex := range c { // Initializing two variables color and hex with the key and value of the map
		fmt.Println("Hex code for", color, "is", hex)
	}
}

// Maps vs Structs

// You would use a map when you want to represent a collection of related properties and don't really care about the keys being unique

// for example colors, where each color has a name and hex code

// You would use a struct when you want to represent a "thing" with a lot of different properties and each of those properties are unique

// for example a person, where each person has a name, age, height, etc. and each person is unique

// Ex. Color

// For example if you wanted to represent a color with a struct you would have to do something like this: type color struct { name string hex string } and then you would have to create a new color like this: color := color{name: "red", hex: "#ff0000"}

// Ex. Person
// For example if you wanted to represent a person with a map you would have to do something like this: person := map[string]string{"name": "Alex", "age": "20", "height": "6ft"}

// For example if you wanted to represent a person with a struct you would have to do something like this: type person struct { name string age int height string } and then you would have to create a new person like this: person := person{name: "Alex", age: 20, height: "6ft"}

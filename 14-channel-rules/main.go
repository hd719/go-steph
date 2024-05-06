package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.youtube.com/",
		"https://www.amazon.com/",
		"https://www.golang.org/",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) // pass in our channel
	}

	for l := range c {
		// Going to pause the go routine for 5 seconds
		// time.Sleep(5 * time.Second) // -> This is a blocking call, so it will block the main routine from receiving messages from the channel
		// go checkLink(l, c)

		// Function Literal - unnamed function that we can use to execute some code in a go routine
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c) // l is being passed in as a function argument, instead of being passed by reference its being passed by value (so its referencing a different variable in memory)
		}(l) // l is the samething as link on line 24
		// () -> this will execute the function literal
		// passing l as an argument to the function literal -> this will create a copy of l and pass it into the function literal
	}
}

// ** NOTE ** We never ever want to reference the same variable in two different go routines in this case l - which is being referenced here (line 47) and on line 24
// We only share information with a Child routine (or a new go routine we create)
// 1. Passing it in as an argument
// 2. Or through channels
// Before:
// for l := range c {
// 	go func() {
// 		time.Sleep(5 * time.Second)
// 		checkLink(l, c)
// 	}
// }
// This will always print out the last link in the links array, because since we are using another go routine l is referred to the same variable in memory (pass by reference)

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

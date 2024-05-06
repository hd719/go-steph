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

	// create a channel
	c := make(chan string)

	for _, link := range links {
		// create our go routine
		go checkLink(link, c) // pass in our channel
	}

	// Wait for a value to be sent into the channel and then print it out
	// **NOTE** Receiving messages is a blocking call, so it will wait until it receives a message
	// fmt.Println(<-c) // this will only print out the first message that is sent to the channel
	// 5 times we are going to wait for a message to be sent into the channel and then print it out, the main routine is waiting for 5 messages to be sent into the channel
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// This is the same as above, but we are using a for loop, but we are continuously waiting for links to be sent into the channel
	// for i := 0; i < len(links); i++ {
	// go checkLink(<-c, c) // this will wait for a message to be sent into the channel and then pass it into the checkLink function
	// }

	// Alternative looping syntax
	// Even though this is an infinite for loop, the "go checkLink(<-c, c)" is a blocking call, so it will wait for a message to be sent into the channel
	// for {
	// 	go checkLink(<-c, c)
	// }

	// Alternative looping syntax
	// l is the link value that is sent into the channel
	for l := range c {
		// Going to pause the go routine for 5 seconds
		// time.Sleep(5 * time.Second) // -> This is a blocking call, so it will block the main routine from receiving messages from the channel
		// go checkLink(l, c)

		// Function Literal - unnamed function that we can use to execute some code in a go routine
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}() // () -> this will execute the function literal
	}
}

// This code is synchronous, so it will wait for each link to be checked before moving on to the next one
// Basically make a request and then wait for the response before moving on to the next link (its not async)
// c: refers to the channel
// chan: is the channel type
// string: is the type of data that we are going to be sending through this channel
func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		println(link, "might be down!")
		// c <- "Might be down I think"
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// c <- "Yep its up"
	c <- link
}

// Routine
// A routine is a lightweight thread of execution (engine that executes code), when we start our go program it creates a main routine
// Adding the "go" keyword in front of a function call will create a new routine and execute that function in that routine
// You use a "go" keyboard right before function calls
// All other routines are called "child routines" and the main routine is called the "parent routine"

// On our Operating System
// We have a go scheduler that is responsible for managing the routines that we create and only 1 "go" routine can be executed at a time

// Concurrency vs Parallelism
// Concurrency: We can have multiple threads executing code, but only 1 thread is running at a time
// Parallelism: Multiple threads executed at the exact same time (can only happen if you have multiple CPU cores)

// Channels
// Are used to communicate data between routines - think of like sms where you send a message through the channel and the other routine receives it
// Channels are typed, so you can only pass a specific type of data through a channel

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		// Handle the error and exit
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	// Create a slice of bytes with 99999 elements
	// bs := make([]byte, 99999) // This is a guess of how big the response body is going to be, 100% assumption

	// Read the body of the response and put it in the slice of bytes
	// resp.Body.Read(bs)

	// Make sure to put string in front of bs to convert it to a string otherwise you will get 0's
	// fmt.Println(string(bs))

	// This is the same as the above code
	// 1st argument is where to write the response to (Writer interface)
	// 2nd argument is where to read the response from (Reader interface)
	// Piping information from the writer to the reader
	io.Copy(lw, resp.Body)
}

// This is a custom implementation of the Writer interface and would be incorrect because it does not return the number of bytes written, we created our own custom implementation of the Writer interface
// Can think of the interface as like a contract or a suggestion, if you want to use the interface you have to implement the methods that are in the interface
// func (logWriter) Write(bs []byte) (int, error) {
// 	return 1, nil
// }

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	return len(bs), nil
}

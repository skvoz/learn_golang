/////////////////////////////////
// Reading Files Line by Line (or using a delimiter) using bufio.Scanner
// Go Playground: https://play.golang.org/p/v0o0H4huUDR
/////////////////////////////////

package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"ex6.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	// defer closing the file
	defer file.Close()

	byteSlice := []byte("The Go gopher is an iconic mascot!") // converting a string to a bytes slice
	bytesWritten, err := file.Write(byteSlice)                // writing bytes to file.
	// It returns the no. of bytes written and an error value
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	_ = bytesWritten
}

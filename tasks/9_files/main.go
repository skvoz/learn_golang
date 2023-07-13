package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//create file
	fmt.Println("Ex1##############################################")
	newFile, err := os.Create("info.txt")

	// error handling
	if err != nil {
		// log the error and exit the program
		log.Fatal(err) // the idiomatic way to handle errors

	}

	newFile.Close()

	fmt.Println("Ex2##############################################")
	//rename file
	// checking if file exists
	_, err1 := os.Stat("info.txt")
	// error handling
	if err1 != nil {
		if os.IsNotExist(err) {
			log.Fatal("The file does not exist!")
		}
	}

	// renaming the file
	err1 = os.Rename("info.txt", "data.txt")
	// error handling
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println("Ex2##############################################")

	err2 := os.Remove("data.txt")
	if err2 != nil {
		log.Fatal(err2)
	}
}

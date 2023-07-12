package main

import (
	"log"
	"os"
)

func main() {
	//strange require but ...
	//check in course colution
	/**
	// checking if file exists
	_, err := os.Stat("a.txt")
	// error handling
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("The file does not exist!")
		}
	}

	// renaming the file
	**/
	OriginalPath := "a.txt"
	NewPath := "data.txt"
	err := os.Rename(OriginalPath, NewPath)
	if err != nil {
		log.Fatal(err)
	}
}

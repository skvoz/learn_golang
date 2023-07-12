package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	// defer closing the file
	defer file.Close()

	// ioutil.ReadAll() reads every byte from the file and return a slice of unknown size
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data as string: %s\n", data)
	fmt.Println("Number of bytes read:", len(data))
}

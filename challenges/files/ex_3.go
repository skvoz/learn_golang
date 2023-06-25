package main

import (
	"log"
	"os"
)

func main() {
	path := "data.txt"
	err := os.Remove(path)
	if err != nil {
		log.Fatal(err)
	}
}

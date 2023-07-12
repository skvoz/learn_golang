package main

import (
	"fmt"
	f "fmt"
)

func main()  {
	fmt.Println("Scopes means name visibiltiy")
	sayHello("Andrei")

	tf := toFahrenheit(boilingPoint)
	fmt.Println("Wahter Boiling in Degrees Fahrenheit:", tf)
	f.Println("import")
}
package main

import "fmt"

//purpouse from course author
func main() {
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}

	newYears := append(years[0:3], years[len(years)-3:]...)

	fmt.Printf("%#v", newYears)
}

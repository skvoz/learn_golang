package main

import "fmt"

func main() {
	fmt.Println("He says: \"Hello\"") //string
	fmt.Println(`He says: "Hello"`)   //raws tring, backticks

	s2 := `I like \n Go!` //raw string

	fmt.Println(s2)

	fmt.Println("Prince: 100\nBrand: Nike")
	fmt.Println(`
Price 100:
Brand: Nike
`)

	var s3 = "tttt"

	fmt.Println("Prince:" + " 100\nBrand: Nike") //glue stirng
	fmt.Println("Element as index 0:", s3[0])    //get part string
	//return Element as index 0: 116
	//part of stirng is byte
	//ascii  table https://commons.wikimedia.org/wiki/File:ASCII-Table-wide.svg
	// t - 116

	// s3[5] = "x" //return error , string immutable in go

	fmt.Printf("%s\n", s3)
	fmt.Printf("%q\n", s3)
}

package main

import "fmt"

func main() {
	//not so good approach
	title1, author1, year1 := "title1", "author1", 1230
	title2, author2, year2 := "title2", "author2", 2230

	fmt.Println("book1:", title1, author1, year1)
	fmt.Println("book2:", title2, author2, year2)

	//struct
	type book struct {
		title  string
		author string
		year   int
	}

	type book1 struct {
		title, author string
		year          int
	}

	myBook := book{"title1", "author1", 1230}

	fmt.Println(myBook)

	bestBook := book{title: "title2", year: 2230, author: "author1"}

	_ = bestBook

	aBook := book{title: "title3"}
	fmt.Printf("%#v\n", aBook)

	//return field
	fmt.Println(aBook.title)
	aBook.title = "new title"
	fmt.Println(aBook.title)
	fmt.Printf("%+v\n", aBook)
	lastBook := book{title: "new title", author: "", year: 0}
	//struct comparator
	fmt.Println(lastBook == aBook)

	//copy
	newBook := aBook
	newBook.author = "new book author"
	fmt.Printf("%#v\n%#v\n", newBook, aBook)

	//anonymous  struct
	diana := struct {
		fName, lName string
		age          int
	}{
		fName: "Diana",
		lName: "Muller",
		age:   30,
	}
	fmt.Printf("%#v\n", diana)
	fmt.Printf("Diana' age: %d\n", diana.age)
	//struct without field name

	type Book struct {
		string
		float64
		bool
	}

	b1 := Book{"1984 by George Orwell", 10.2, false}
	fmt.Printf("%#v\n", b1)
	fmt.Println(b1.string)

	type Employee struct {
		name   string
		salary int
		bool
	}

	e := Employee{"John", 4000, false}
	fmt.Printf("%#v\n", e)
	e.bool = true
	fmt.Printf("%#v\n", e)

}

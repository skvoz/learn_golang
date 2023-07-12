package main

import "fmt"

func main() {
	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name   string
		salary int
		//embedded struct
		contactInfo Contact
	}

	john := Employee{
		name:   "John Keller",
		salary: 4000,
		contactInfo: Contact{
			email:   "jkelly@company.com",
			address: "streen 20, london",
			phone:   11111111,
		},
	}

	fmt.Printf("%+v\n", john)
	fmt.Printf("Employess email: %s\n", john.contactInfo.email)
	john.contactInfo.email = "newemail@company.com"
	fmt.Printf("Employess email: %s\n", john.contactInfo.email)

	myContact := Contact{email: "andrey@domain.com", phone: 11111111, address: "Bucharest, Romonia"}
	fmt.Println(myContact)
}

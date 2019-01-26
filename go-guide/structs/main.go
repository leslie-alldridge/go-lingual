package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	// human := person{firstName: "Leslie", lastName: "Alldridge"}

	// var human person
	// human.firstName = "First"
	// human.lastName = "Last"
	// fmt.Println(human)
	// fmt.Printf("%+v", human)

	jim := person {
		firstName: "Jim",
		lastName: "Smith",
		contact: contactInfo {
			email: "Jim.Smith@gmail.com",
			zipCode: 405,
		},
	}

	fmt.Printf("%+v", jim)
}
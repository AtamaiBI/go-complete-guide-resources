package main

import (
	"fmt"
	"time"
)

// create a struct type called user that has three fields: firstName, lastName, and birthdate. All of these fields should be of type string. A struct can be created either outside of a function or inside of a function. In this case, we will create the struct type outside of the main function. This is standard practice in Go, as it allows the struct type to be used throughout the entire package. The struct type is defined using the type keyword, followed by the name of the struct (user), and then the struct fields are defined within curly braces. Each field has a name and a type, separated by a space. In this case, all three fields are of type string.
type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time // Time is also a struct type, and it is part of the time package. The time package is part of the Go standard library, and it provides functionality for working with dates and times. The createdAt field will be used to store the date and time when the user was created. This field is of type time.Time, which is a struct type that represents a specific point in time. The time.Time struct has many methods that can be used to manipulate and format dates and times.
}

// This function becomes a struct method by placing the struct name in parenthesise (user) after the func keyword and before the function name. You can also add a parameter, in this case u. The function then becomes a method belonging to the struct and can be called in an instance such as appUser.outputUserDetails().
// the extra peice of code (u user) is called a Receiver argument and acts like a this or self keyword in C# or python
func (u user) outputUserDetails() {
	fmt.Println("struct method")
	fmt.Println(u.firstName, u.lastName, u.createdAt)
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user // create a variable of type user, which is a struct type. The variable is named appUser, and it will be used to store the data that we gather from the user. The variable is created using the var keyword, followed by the name of the variable (appUser), and then the type of the variable (user). The variable is initialized with the zero value for the user struct type, which means that all of its fields will be set to their zero values. In this case, all of the fields will be set to empty strings, except for the createdAt field, which will be set to the zero value for time.Time, which is January 1, year 1, 00:00:00 UTC.

	appUser = user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}

appUser.outputUserDetails()

	var appUser1 user = user{} // Declare and Initialize the struct

	appUser1.firstName = firstName
	appUser1.lastName = lastName

	outputUserDetailsPointer(&appUser1)

	// ... do something awesome with that gathered data!

	outputUserDetails(appUser)

	fmt.Println(firstName, lastName, birthdate)
}


func outputUserDetailsPointer(u *user) {
	fmt.Println("Pointer Details: ", u.firstName, (*u).lastName) // Remember with pointers you must dereference them to access their values but Go allows you to use a shortcut method to access the underlying field without explicitly dereferencing the pointer by u.firstName
}

// This function prints the fields within the struct user. Fields are different from properties which use Getters and Setters to access the underlying values.
func outputUserDetails(u user) {
	fmt.Println(u.firstName, u.lastName, u.createdAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

package main

import (
	"fmt"
	// To imort a custom package use the import path which is the relative path from the go.mod file to the package folder. In this case, the package is located in the user folder, which is a subfolder of the current folder. The import path is therefore
	"01-starting-project/user"
)

/* The original struct definition was moved to the user.go file and is now being imported into this file. The struct definition is commented out here to avoid a redeclaration error. The struct definition is as follows:
type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
} */

func main() {

	// prompts to get the user data. This should be done via a webpage

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// declare appUser as a pointer to a user struct. The *user.User type indicates that the variable is a pointer to a User struct. The variable is declared but not initialized, so it has a nil value. The variable can be initialized using the NewUser constructor function, which returns a pointer to a new User struct.
	var appUser *user.User

	/*
		appUser = User{
			firstName: firstName,
			lastName:  lastName,
			birthdate: birthdate,
			createdAt: time.Now(),
		}*/


		appUser, _ = user.New(firstName,lastName,birthdate)

		appUser.OutputUserDetails() // Call the struct method OutputUserDetails on the appUser variable. The method is called using the dot notation, which is used to access the fields and methods of a struct type. The method takes a parameter of type User, which is passed as an argument to the method. The method prints the details of the user to the console. The method expects a pointer to a User struct, so we pass the address of the appUser variable using the & operator. The method is called on the appUser variable, which is of type User, and the method has a receiver of type *User, which means it can be called on a pointer to a User struct. The method can access the fields of the struct using the receiver argument, which is a pointer to the struct.

		appUser.WasTheFirtsNameChanged()

		anotherUser, _ := user.New("John", "Doe", "01/01/2000")

		anotherUser.OutputUserDetails()




	var appUser1 user.User = user.User{} // Declare and Initialize the struct

	appUser1.FirstName = firstName // The variable first name has not changed, rather the pointer method changed the field value. Hence this will print the original entry rather than the updated field value 'bobby'
	appUser1.LastName = lastName

	user.OutputUserDetailsPointer(&appUser1)

	// ... do something awesome with that gathered data!

	// Short-variable declaration with struct literal or var user_ user.User
	user_ := user.User{}

	user_.OutputUserDetails() //

	appUser.OutputUserDetails()

	user.OutputUserDetails_changeLastName(*appUser) // Call the struct method OutputUserDetails on the appUser variable. The method is called using the dot notation, which is used to access the fields and methods of a struct type. The method takes a parameter of type User, which is passed as an argument to the method. The method prints the details of the user to the console. The method expects a value of type User, so we pass the value of the appUser variable using the * operator to dereference the pointer. The method is called on the appUser variable, which is of type User, and the method has a receiver of type User, which means it can be called on a value of type User. The method can access the fields of the struct using the receiver argument, which is a value of the struct.

	fmt.Println("Print the original variable values")
	fmt.Println(firstName, lastName, birthdate)

	//var user3, err = user_.NewUser(firstName, lastName, birthdate) //Use the constructor function. It returns a user Pointer
	// Need this := assignment operator
	user3, err := user.New(firstName,lastName,birthdate) // NewUser returns a memory pointer.

	if err != nil {
		fmt.Println("Error creating user: ", err)
		return
	}

	user3.OutputUserDetails() //OutputUserDetails is an instance method with a pointer parameter.

admin := user.NewAdmin(firstName,lastName, "Administrator", "test1234")

admin.User.OutputUserDetails() // access the method in the embedded struct




} // end main

// Function to get user data. This would normally be a webform.
func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

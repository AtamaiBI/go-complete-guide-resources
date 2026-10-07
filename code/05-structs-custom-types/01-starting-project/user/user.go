package user

import (
	"errors"
	"fmt"
	"time"

)

// create a struct type called User that has three fields: firstName, lastName, and birthdate. All of these fields should be of type string. A struct can be created either outside of a function or inside of a function. In this case, we will create the struct type outside of the main function. This is standard practice in Go, as it allows the struct type to be used throughout the entire package. The struct type is defined using the type keyword, followed by the name of the struct (User), and then the struct fields are defined within curly braces. Each field has a name and a type, separated by a space. In this case, all three fields are of type string.
type User struct {
	FirstName string
	LastName  string
	birthdate string
	CreatedAt time.Time // Time is also a struct type, and it is part of the time package. The time package is part of the Go standard library, and it provides functionality for working with dates and times. The createdAt field will be used to store the date and time when the user was created. This field is of type time.Time, which is a struct type that represents a specific point in time. The time.Time struct has many methods that can be used to manipulate and format dates and times.
}

// Create a constructor function for the user struct. This returns a pointer to the new struct
// Add validation to the constructor
// changed func name from NewUser to New because when creating a new user New is sufficient to describe the purpose.
func New(firstname, lastname, birthday string) (*User, error) {

	if firstname == "" || lastname == "" || birthday == "" {
		return nil, errors.New("first, last and birthday are required")
	}

	// create the user struct and return a pointer to it. The & operator is used to get the address of the struct, which is then returned as a pointer. The *User type indicates that the function returns a pointer to a User struct. The function also returns an error value, which is nil if there are no errors, or an error message if there are errors.
	return &User{
		FirstName: firstname,
		LastName:  lastname,
		birthdate: birthday,
		CreatedAt: time.Now(),
	}, nil // must add nil here for returning a null error
}

func OutputUserDetailsPointer(u *User) {
	fmt.Println("Pointer Details: ", u.FirstName, (*u).LastName) // Remember with pointers you must dereference them to access their values but Go allows you to use a shortcut method to access the underlying field without explicitly dereferencing the pointer by using u.FirstName
}

// This function becomes an instance method by placing the struct name in parenthesise (User) after the func keyword and before the function name. You can also add a function parameter(s). The function then becomes a method belonging to the struct and can be called in an instance such as appUser.OutputUserDetails().
// the extra peice of code (u User) is called a Receiver argument and acts like a this or self keyword in C# or python
func (u User) OutputUserDetails() { // This is an instance method
	fmt.Println("OutputUserDetail struct method")
	u.FirstName = "Bobby - first name changed in the the struct"
	fmt.Println(u.FirstName, u.LastName, u.CreatedAt)
}

func (u User) WasTheFirtsNameChanged() { // This is an instance method
	fmt.Println("WasTheFirtsNameChanged struct method")
	fmt.Println(u.FirstName, u.LastName, u.CreatedAt)
}

func ChangeFirstName(u *User) { // This is a package level function.
	u.FirstName = "Bobby"
	fmt.Println("Change first name to Bobby on the original variable")
	fmt.Println(u.FirstName, u.LastName, u.CreatedAt)
}

// This function prints the fields within the struct user. Fields are different from properties which use Getters and Setters to access the underlying values.
func OutputUserDetails_changeLastName(i User) {
	i.LastName = "Baggins"
	fmt.Println("Change last name to Baggins on the copy")
	fmt.Println(i.FirstName, i.LastName, i.CreatedAt)
}

func CheckOutputUserDetails(u User) {
	fmt.Println("The last name won't be Baggins")
	fmt.Println(u.FirstName, u.LastName)

}

// Admin starts with a capital to make this available across packages
type Admin struct {
	User User /* This embeds the User struct into the admin struct, which s similar to inheritance. User is the name
	and User the type. User User */
	role string
	password string
}

// create a contructor for the admin struct
func NewAdmin(firstname, lastname, role, password string) Admin {
	return Admin{
		role: role,
		password: password,
		User: User{
			FirstName: firstname,
			LastName: lastname,
			birthdate: time.Now().Format("01/02/2006"),
			CreatedAt: time.Now(),
		},
	}

}



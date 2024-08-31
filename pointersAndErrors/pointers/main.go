package main

import "fmt"

// when to use pointers
// 1. When we need to update state of a variable
// 2. When we need to return a pointer to a variable
// 3. When we need to optimize the memory for large objectts that are getting called a lot

// Pros:=
// - Efficiency = less memory allocation
// - Mutability = No need to pass the struct by value
// - Returning pointers = allow returning (nil , err) instea

// Cons:=
// - Memory Leaks
// - Safety = data members may be modified unawarely
// - Concurrency = pointers are not thread safe (race conditions)
// - Inconsistent API = some functions return pointers and some return values
// - Nil pointer dereference = nil pointers can be dereferenced

type User struct {
	email    string
	username string
	age      int
	file     []byte // ?? large-size ??
}

func getUser() (User, error) { // getUser is a function that returns a User and an error, when we have *User in args
	// we can return a pointer to a User which is lose and give same User instead of initializing a new User with defalut values
	return User{}, fmt.Errorf("Throw some error.")
}

// x amount  of bytes = entire struct is copied into memory all the size of struct
func (u User) Email() string {
	return u.email
}

// 8 bytes pointer used
func Email(u *User) string {
	return u.email
}

func (u *User) UpdateEmail(email string) {
	u.email = email
}

func main() {
	user := User{
		email: "dummy@gmail.com",
	}
	user.Email()
	user.UpdateEmail("yummy@gmail.com")
	fmt.Println(user.Email())
}

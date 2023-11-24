// Error list.
package user

import (
	err "github.com/botscubes/user-service/pkg/service_error"
)

// User struct errors.
var (

	// Empty login. Error code ...
	ErrEmptyLogin = err.New(1, "Empty login")

	// Login length is short. Error сode ...
	ErrShortLogin = err.New(2, "Login length is short")

	// Login length is long. Error code ...
	ErrLongLogin = err.New(3, "Login length is long")

	// Login is incorrect. Error code ...
	ErrIncorrectLogin = err.New(4, "Login is incorrect")

	// Login exists in the database. Error code ...
	ErrLoginExists = err.New(5, "Login exists")

	// Login not exists in the database. Error code ...
	ErrLoginNotExists = err.New(6, "Login not exists")

	// Empty password. Error code ...
	ErrEmptyPassword = err.New(7, "Empty password")

	// Password length is short. Error code ...
	ErrShortPassword = err.New(8, "Password length is short")

	// Password length is long. Error code ...
	ErrLongPassword = err.New(9, "Password length is long")

	// Password is incorrect. Error code ...
	ErrIncorrectPassword = err.New(10, "Password is incorrect")

	// Entered password is not equal. Error code ...
	ErrPasswordIsNotEqual = err.New(11, "Password is not equal")
)

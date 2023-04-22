// Error list.
package errors

import (
	err "github.com/botscubes/user-service/pkg/service_error"
)

// No error. Error code 0.
var NoError = err.New(0, "user-service: no error")

// User struct errors.
// Groups of errors:
// 10..19 - Login errors.
// 20..29 - Password errors.
var (

	// Empty login. Error code ...
	ErrEmptyLogin = err.New(10, "user-service: empty login")

	// Login length is short. Error сode ...
	ErrShortLogin = err.New(11, "user-service: login length is short")

	// Login length is long. Error code ...
	ErrLongLogin = err.New(12, "user-service: login length is long")

	// Login is incorrect. Error code ...
	ErrIncorrectLogin = err.New(13, "user-service: login is incorrect")

	// Login exists in the database. Error code ...
	ErrLoginExists = err.New(14, "user-service: login exists")

	ErrLoginNotExists = err.New(14, "user-service: login exists")

	// Empty password. Error code ...
	ErrEmptyPassword = err.New(20, "user-service: empty password")

	// Password length is short. Error code ...
	ErrShortPassword = err.New(21, "user-service: password length is short")

	// Password length is long. Error code ...
	ErrLongPassword = err.New(22, "user-service: password length is long")

	// Password is incorrect. Error code ...
	ErrIncorrectPassword = err.New(23, "user-service: password is incorrect")

	// Entered password is not equal. Error code ...
	ErrPasswordIsNotEqual = err.New(24, "user-service: password is not equal")
)

// Server errors.
var (
	ErrBadRequest          = err.New(1400, "user-service: bad request")
	ErrUnauthorized        = err.New(1401, "user-service: unauthorized")
	ErrInternalServerError = err.New(1500, "user-service: internal server error")
)

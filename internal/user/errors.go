// Error list.
package user

import (
	err "github.com/botscubes/user-service/pkg/service_error"
)

var (
	// No error. Error code 0.
	NoError = err.New(0, "user-service: no error")

	// Empty login. Error code 1.
	ErrEmptyLogin = err.New(1, "user-service: empty login")

	// Login length is short. Error сode 2.
	ErrShortLogin = err.New(2, "user-service: login length is short")

	// Login length is long. Error code 3.
	ErrLongLogin = err.New(3, "user-service: login length is long")

	// Login is incorrect. Error code 4.
	ErrIncorrectLogin = err.New(4, "user-service: login is incorrect")

	// Empty password. Error code 5.
	ErrEmptyPassword = err.New(5, "user-service: empty password")

	// Password length is short. Error code 6.
	ErrShortPassword = err.New(6, "user-service: password length is short")

	// Password length is long. Error code 7.
	ErrLongPassword = err.New(7, "user-service: password length is long")

	// Password is incorrect. Error code 8.
	ErrIncorrectPassword = err.New(8, "user-service: password is incorrect")
)

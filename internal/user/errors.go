// Error list.
package user

import "user-service/internal/errors"

// No error. Error code 0.
var NoError = errors.New(0, "user-service: no error")

// Empty login. Error code 1.
var ErrEmptyLogin = errors.New(1, "user-service: empty login")

// Login length is short. Error сode 2.
var ErrShortLogin = errors.New(2, "user-service: login length is short")

// Login length is long. Error code 3.
var ErrLongLogin = errors.New(3, "user-service: login length is long")

// Login is incorrect. Error code 4.
var ErrIncorrectLogin = errors.New(4, "user-service: login is incorrect")

// Empty password. Error code 5.
var ErrEmptyPassword = errors.New(5, "user-service: empty password")

// Password length is short. Error code 6.
var ErrShortPassword = errors.New(6, "user-service: password length is short")

// Password length is long. Error code 7.
var ErrLongPassword = errors.New(7, "user-service: password length is long")

// Password is incorrect. Error code 8.
var ErrIncorrectPassword = errors.New(8, "user-service: password is incorrect")

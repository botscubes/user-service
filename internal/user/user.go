package user

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/botscubes/user-service/internal/errors"
	"github.com/botscubes/user-service/pkg/service_error"
)

// User public fields.
// type PublicUserFields struct {
// 	id               int64
// 	login            string
// 	firstName        string
// 	secondName       string
// 	phone            string
// 	email            string
// 	password         string
// 	dateRegistration int64
// 	dateModification int64
// 	//dateModificationPassword int64
// 	status string
// }

// User struct.
type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (u *User) GetLogin() string {
	return u.Login
}

func (u *User) SetLogin(login string) *service_error.ServiceError {
	login = strings.TrimSpace(login)
	if login == "" {
		return errors.ErrEmptyLogin
	}

	var len = utf8.RuneCountInString(login)
	if len < 6 {
		return errors.ErrShortLogin
	}
	if len > 30 {
		return errors.ErrLongLogin
	}
	if matched, _ := regexp.MatchString(`^[A-Za-z0-9_]*`, login); !matched {
		return errors.ErrIncorrectLogin
	}
	u.Login = login
	return errors.NoError
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) SetPassword(password string) *service_error.ServiceError {
	password = strings.TrimSpace(password)
	if password == "" {
		return errors.ErrEmptyPassword
	}

	var len = utf8.RuneCountInString(password)
	if len < 6 {
		return errors.ErrShortPassword
	}
	if len > 50 {
		return errors.ErrLongPassword
	}
	if matched, _ := regexp.MatchString(`^[A-Za-z0-9_&?!@#$%^+=*]*`, password); !matched {
		return errors.ErrIncorrectPassword
	}

	return errors.NoError
}

func NewUser(login string, password string) (*User, *service_error.ServiceError) {
	var user *User = new(User)

	if err := user.SetLogin(login); err != errors.NoError {
		return nil, err
	}
	if err := user.SetPassword(password); err != errors.NoError {
		return nil, err
	}

	return user, errors.NoError

}

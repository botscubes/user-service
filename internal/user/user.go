package user

// User public fields.
type PublicUserFields struct {
	id               int64
	login            string
	firstName        string
	secondName       string
	phone            string
	email            string
	password         string
	dateRegistration int64
	dateModification int64
	//dateModificationPassword int64
	status string
}

// User struct.
type User struct {
	password string
	PublicUserFields
}

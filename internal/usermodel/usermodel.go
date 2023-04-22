package usermodel

import (
	"github.com/botscubes/user-service/internal/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Model for User.
type UserModel struct {
	pool *pgxpool.Pool
}

func New(p *pgxpool.Pool) *UserModel {
	return &UserModel{p}
}

// Save the user to the database.
func (um *UserModel) SaveUser(u *user.User) error {

	return nil
}

// Check password by Login.
func (um *UserModel) GetIdAndPasswordByLogin(login string) (int, string, error) {

	return 0, "", nil
}

// Сheck the existence of the login in the database.
func (um *UserModel) LoginExists(login string) (bool, error) {

	return false, nil
}

// Delete user from database.
func (um *UserModel) DeleteUser(id int) error {

	return nil
}

// Get user info from database.
func (um *UserModel) GetUser(id int) (*user.User, error) {
	return nil, nil
}

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
func (um *UserModel) CheckPasswordByLogin(login string, password string) (bool, error) {

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

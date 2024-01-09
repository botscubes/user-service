package usermodel

import (
	"context"

	"github.com/botscubes/user-service/internal/user"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Model for User.
type UserModel struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, p *pgxpool.Pool) *UserModel {
	return &UserModel{p}
}

// Save the user to the database.
func (um *UserModel) SaveUser(ctx context.Context, u *user.User) error {
	conn, err := um.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, "INSERT INTO account (login, password) VALUES ($1, $2)", u.Login, u.Password)
	if err != nil {
		return err
	}

	return nil
}

// Get Id and password by Login.
func (um *UserModel) GetIdAndPasswordByLogin(ctx context.Context, login string) (int, string, error) {
	conn, err := um.pool.Acquire(ctx)
	if err != nil {
		return 0, "", err
	}
	defer conn.Release()

	var id int
	var password string
	err = conn.QueryRow(
		ctx,
		"SELECT id, password FROM account WHERE login = $1", login,
	).Scan(&id, &password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, "", nil
		}
		return 0, "", err
	}
	return id, password, nil
}

// Сheck the existence of the login in the database.
func (um *UserModel) LoginExists(ctx context.Context, login string) (bool, error) {
	conn, err := um.pool.Acquire(ctx)
	if err != nil {
		return false, err
	}
	defer conn.Release()

	var exists bool
	err = conn.QueryRow(
		ctx,
		"SELECT EXISTS(SELECT 1 FROM account WHERE login = $1)", login,
	).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// Delete user from database.
func (um *UserModel) DeleteUser(id int) error {

	return nil
}

// Get user info from database.
func (um *UserModel) GetUser(id int) (*user.User, error) {
	return nil, nil
}

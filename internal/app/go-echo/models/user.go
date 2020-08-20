package models

import (
	"github.com/jmoiron/sqlx"
)

type (
	// ModelImpl define methods for User model
	ModelImpl interface {
		FindByID(id string) User
		FindAll() []User
	}

	// User struct
	User struct {
		ID       int    `json:"id" db:"id"`
		Name     string `json:"name" db:"name"`
		Email    string `json:"email" db:"email"`
		Password string `json:"password" db:"password"`
	}

	// Users is array of User
	Users []User
)

// NewUser create new User
func NewUser(name, email, password string) *User {
	return &User{
		ID:       0,
		Name:     name,
		Email:    email,
		Password: password,
	}
}

// Create add new user to db
func (u *User) Create(tx *sqlx.Tx) error {
	row := tx.QueryRow("INSERT INTO users (name, email, password) values ($1, $2, $3) RETURNING *", u.Name, u.Email, u.Password)

	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	return err
}

// FindByID get an user by id
func (u *User) FindByID(tx *sqlx.Tx, id int) error {
	err := tx.Get(&u, "SELECT * FROM users where id = $1 limit 1", id)

	return err
}

// FindAll get all users
func (u *Users) FindAll(tx *sqlx.Tx) error {
	err := tx.Select(&u, "SELECT * FROM users order by id asc")
	return err
}

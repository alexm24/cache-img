package postgres

import (
	"fmt"
	"github.com/alexm24/cache-img/internal/models"
	"github.com/jmoiron/sqlx"
)

type Auth struct {
	db *sqlx.DB
}

func (a Auth) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := a.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (a Auth) GetUser(username, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := a.db.Get(&user, query, username, password)

	return user, err
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{db}
}

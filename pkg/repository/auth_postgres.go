package repository

import (
	"fmt"
	fhub "github.com/antonchaban/file-hub-go"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

type Authorization interface {
	CreateUser(user fhub.User) (int, error)
	GetUser(username, password string) (fhub.User, error)
	AddTokenToBlacklist(token string) (int, error)
	IsTokenInBlacklist(token string) (bool, error)
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user fhub.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash) VALUES ($1, $2) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (fhub.User, error) {
	var user fhub.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}

func (r *AuthPostgres) AddTokenToBlacklist(token string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (token) VALUES ($1) RETURNING id", tokensBlacklist)
	row := r.db.QueryRow(query, token)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) IsTokenInBlacklist(token string) (bool, error) {
	var counter int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE token like $1", tokensBlacklist)
	err := r.db.Get(&counter, query, "%"+token+"%")
	fmt.Println("counter", counter)

	return counter != 0, err
}

package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID   int
	Name string
	Age  int
}

type UserModel struct {
	DB *pgxpool.Pool
}

func (m *UserModel) Insert(name string, age int) (int, error) {
	stmt := `INSERT INTO users (name, age) VALUES($1, $2) RETURNING id`
	var id int
	// Insert user into database
	err := m.DB.QueryRow(context.Background(), stmt, name, age).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *UserModel) Get(id int) (*User, error) {
	stmt := `SELECT id, name, age FROM users WHERE id = $1`
	row := m.DB.QueryRow(context.Background(), stmt, id)
	user := &User{}
	err := row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return nil, err
	}
	return user, nil
}

package users

import (
	_ "embed"
	"github.com/jmoiron/sqlx"
	"gravitum-test/models"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository { return &Repository{db: db} }

//go:embed sql/select_user.sql
var selectUsersSql string

func (r *Repository) SelectUsers() ([]models.Users, error) {
	var users []models.Users

	rows, err := r.db.Query(selectUsersSql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user models.Users
		if err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Sex,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

//go:embed sql/insert_user.sql
var insertUsersSql string

func (r *Repository) InsertUsers(users models.Users) error {
	_, err := r.db.Exec(insertUsersSql, users.FirstName, users.LastName, users.Sex)
	if err != nil {
		return err
	}
	return nil
}

//go:embed sql/update_user.sql
var updateUsersSql string

func (r *Repository) UpdateUsers(users models.Users) error {
	_, err := r.db.Exec(updateUsersSql, users.FirstName, users.LastName, users.Sex, users.Id)
	if err != nil {
		return err
	}
	return nil
}

package users

import (
	"context"
	_ "embed"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"gravitum-test/models"
)

type Repository struct {
	db     *sqlx.DB
	logger *zap.Logger
}

func NewRepository(db *sqlx.DB, logger *zap.Logger) *Repository {
	return &Repository{db: db, logger: logger}
}

//go:embed sql/select_user.sql
var selectUsersSql string

func (r *Repository) SelectUsers(ctx context.Context) ([]models.Users, error) {
	r.logger.Debug("selecting users from database")

	var users []models.Users

	rows, err := r.db.QueryContext(ctx, selectUsersSql)
	if err != nil {
		r.logger.Error("failed to execute select users query", zap.Error(err))
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
			r.logger.Error("failed to scan user row", zap.Error(err))
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("rows iteration error while selecting users", zap.Error(err))
		return nil, err
	}

	r.logger.Info("users selected from database", zap.Int("count", len(users)))
	return users, nil
}

//go:embed sql/insert_user.sql
var insertUsersSql string

func (r *Repository) InsertUsers(ctx context.Context, users models.Users) error {
	r.logger.Debug("inserting user into database", zap.String("first_name", users.FirstName), zap.String("last_name", users.LastName))

	_, err := r.db.ExecContext(ctx, insertUsersSql, users.FirstName, users.LastName, users.Sex)
	if err != nil {
		r.logger.Error("failed to insert user", zap.Error(err))
		return err
	}

	r.logger.Info("user inserted into database", zap.String("first_name", users.FirstName), zap.String("last_name", users.LastName))
	return nil
}

//go:embed sql/update_user.sql
var updateUsersSql string

func (r *Repository) UpdateUsers(ctx context.Context, users models.Users) error {
	r.logger.Debug("updating user in database", zap.Int("user_id", users.Id))

	_, err := r.db.ExecContext(ctx, updateUsersSql, users.FirstName, users.LastName, users.Sex, users.Id)
	if err != nil {
		r.logger.Error("failed to update user", zap.Error(err), zap.Int("user_id", users.Id))
		return err
	}

	r.logger.Info("user updated in database", zap.Int("user_id", users.Id))
	return nil
}

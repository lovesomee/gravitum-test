package api

import (
	"context"

	"gravitum-test/models"
)

type UserService interface {
	AddUser(ctx context.Context, user models.Users) error
	UpdateUser(ctx context.Context, user models.Users) error
	GetUser(ctx context.Context) ([]models.Users, error)
}

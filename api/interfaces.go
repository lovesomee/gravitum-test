package api

import "gravitum-test/models"

type UserService interface {
	AddUser(user models.Users) error
	UpdateUser(user models.Users) error
	GetUser() ([]models.Users, error)
}

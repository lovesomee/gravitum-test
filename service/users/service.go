package users

import (
	"gravitum-test/models"
)

type IRepository interface {
	InsertUsers(users models.Users) error
	UpdateUsers(users models.Users) error
	SelectUsers() ([]models.Users, error)
}

type Service struct {
	usersRepository IRepository
}

func NewService(usersRepository IRepository) *Service {
	return &Service{usersRepository: usersRepository}
}

func (s *Service) AddUser(user models.Users) error {
	if err := s.usersRepository.InsertUsers(user); err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateUser(user models.Users) error {
	if err := s.usersRepository.UpdateUsers(user); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetUser() ([]models.Users, error) {
	return s.usersRepository.SelectUsers()
}

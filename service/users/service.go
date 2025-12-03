package users

import (
	"context"

	"go.uber.org/zap"

	"gravitum-test/models"
)

type IRepository interface {
	InsertUsers(ctx context.Context, users models.Users) error
	UpdateUsers(ctx context.Context, users models.Users) error
	SelectUsers(ctx context.Context) ([]models.Users, error)
}

type Service struct {
	usersRepository IRepository
	logger          *zap.Logger
}

func NewService(usersRepository IRepository, logger *zap.Logger) *Service {
	return &Service{usersRepository: usersRepository, logger: logger}
}

func (s *Service) AddUser(ctx context.Context, user models.Users) error {
	s.logger.Debug("service AddUser called", zap.String("first_name", user.FirstName), zap.String("last_name", user.LastName))

	if err := s.usersRepository.InsertUsers(ctx, user); err != nil {
		s.logger.Error("service failed to add user", zap.Error(err))
		return err
	}

	s.logger.Info("service added user", zap.String("first_name", user.FirstName), zap.String("last_name", user.LastName))
	return nil
}

func (s *Service) UpdateUser(ctx context.Context, user models.Users) error {
	s.logger.Debug("service UpdateUser called", zap.Int("user_id", user.Id))

	if err := s.usersRepository.UpdateUsers(ctx, user); err != nil {
		s.logger.Error("service failed to update user", zap.Error(err), zap.Int("user_id", user.Id))
		return err
	}

	s.logger.Info("service updated user", zap.Int("user_id", user.Id))
	return nil
}

func (s *Service) GetUser(ctx context.Context) ([]models.Users, error) {
	s.logger.Debug("service GetUser called")

	users, err := s.usersRepository.SelectUsers(ctx)
	if err != nil {
		s.logger.Error("service failed to get users", zap.Error(err))
		return nil, err
	}

	s.logger.Info("service got users", zap.Int("count", len(users)))
	return users, nil
}

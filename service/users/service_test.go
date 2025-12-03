package users

import (
	"context"
	"errors"
	"testing"
	"time"

	"go.uber.org/zap"

	"gravitum-test/models"
)

type mockRepository struct {
	insertFunc func(ctx context.Context, users models.Users) error
	updateFunc func(ctx context.Context, users models.Users) error
	selectFunc func(ctx context.Context) ([]models.Users, error)
}

func (m *mockRepository) InsertUsers(ctx context.Context, users models.Users) error {
	return m.insertFunc(ctx, users)
}

func (m *mockRepository) UpdateUsers(ctx context.Context, users models.Users) error {
	return m.updateFunc(ctx, users)
}

func (m *mockRepository) SelectUsers(ctx context.Context) ([]models.Users, error) {
	return m.selectFunc(ctx)
}

func TestService_AddUser(t *testing.T) {
	tests := []struct {
		name    string
		user    models.Users
		mockErr error
		wantErr bool
	}{
		{
			name: "successful insert",
			user: models.Users{
				Id:        1,
				FirstName: "Александр",
				LastName:  "Иванов",
				Sex:       "мужской",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "insert error",
			user: models.Users{
				Id:        2,
				FirstName: "Екатерина",
				LastName:  "Петрова",
				Sex:       "женский",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			mockErr: errors.New("insert failed"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mockRepository{
				insertFunc: func(ctx context.Context, users models.Users) error {
					return tt.mockErr
				},
			}
			logger, _ := zap.NewDevelopment()
			s := NewService(repo, logger)
			ctx := context.Background()
			err := s.AddUser(ctx, tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_UpdateUser(t *testing.T) {
	tests := []struct {
		name    string
		user    models.Users
		mockErr error
		wantErr bool
	}{
		{
			name: "successful update",
			user: models.Users{
				Id:        1,
				FirstName: "Александр Обновлённый",
				LastName:  "Иванов",
				Sex:       "мужской",
				CreatedAt: time.Now().Add(-time.Hour),
				UpdatedAt: time.Now(),
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "update error",
			user: models.Users{
				Id:        2,
				FirstName: "Екатерина",
				LastName:  "Петрова Обновлённая",
				Sex:       "женский",
				CreatedAt: time.Now().Add(-time.Hour),
				UpdatedAt: time.Now(),
			},
			mockErr: errors.New("update failed"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mockRepository{
				updateFunc: func(ctx context.Context, users models.Users) error {
					return tt.mockErr
				},
			}
			logger, _ := zap.NewDevelopment()
			s := NewService(repo, logger)
			ctx := context.Background()
			err := s.UpdateUser(ctx, tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_GetUser(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name     string
		mockData []models.Users
		mockErr  error
		wantErr  bool
	}{
		{
			name: "successful select",
			mockData: []models.Users{
				{
					Id:        1,
					FirstName: "Александр",
					LastName:  "Иванов",
					Sex:       "мужской",
					CreatedAt: now.Add(-time.Hour),
					UpdatedAt: now,
				},
				{
					Id:        2,
					FirstName: "Екатерина",
					LastName:  "Петрова",
					Sex:       "женский",
					CreatedAt: now.Add(-2 * time.Hour),
					UpdatedAt: now.Add(-time.Hour),
				},
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name:     "select error",
			mockData: nil,
			mockErr:  errors.New("select failed"),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mockRepository{
				selectFunc: func(ctx context.Context) ([]models.Users, error) {
					return tt.mockData, tt.mockErr
				},
			}
			logger, _ := zap.NewDevelopment()
			s := NewService(repo, logger)
			ctx := context.Background()
			users, err := s.GetUser(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(users) != len(tt.mockData) {
					t.Errorf("GetUser() got %d users, want %d", len(users), len(tt.mockData))
				}
				for i, user := range users {
					want := tt.mockData[i]
					if user.Id != want.Id ||
						user.FirstName != want.FirstName ||
						user.LastName != want.LastName ||
						user.Sex != want.Sex ||
						!user.CreatedAt.Equal(want.CreatedAt) ||
						!user.UpdatedAt.Equal(want.UpdatedAt) {
						t.Errorf("GetUser() got %+v, want %+v", user, want)
					}
				}
			}
		})
	}
}

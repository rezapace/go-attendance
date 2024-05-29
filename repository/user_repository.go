package repository

import (
	"context"

	"presensee_project/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	GetSingleUser(ctx context.Context, userID uint) (*model.User, error)
	GetBriefUsers(ctx context.Context, limit int, offset int) (*model.Users, int64, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, userID uint) error
}

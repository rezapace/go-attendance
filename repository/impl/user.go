package database

import (
	"context"
	"strings"

	"presensee_project/model"
	"presensee_project/repository"
	"presensee_project/utils"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) repository.UserRepository {
	userRepository := &UserRepositoryImpl{
		db: db,
	}

	return userRepository
}

func (u *UserRepositoryImpl) CreateUser(ctx context.Context, user *model.User) error {
	err := u.db.WithContext(ctx).Create(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062: Duplicate entry") {
			switch {
			case strings.Contains(err.Error(), "email"):
				return utils.ErrUsernameAlreadyExist
			}
		}

		return err
	}

	return nil
}

func (u *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := u.db.WithContext(ctx).Select([]string{"id", "email", "password", "role"}).Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (u *UserRepositoryImpl) GetSingleUser(ctx context.Context, userID uint) (*model.User, error) {
	var user model.User
	err := u.db.WithContext(ctx).
		Where("id = ?", userID).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (u *UserRepositoryImpl) GetBriefUsers(ctx context.Context, limit int, offset int) (*model.Users, int64, error) {
	var users model.Users
	var count int64

	err := u.db.WithContext(ctx).
		Select([]string{"id", "email", "name", "role"}).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	err = u.db.WithContext(ctx).
		Model(&model.User{}).
		Where("role = ?", "admin").
		Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	if len(users) == 0 {
		return nil, 0, utils.ErrUserNotFound
	}

	return &users, count, nil
}

func (u *UserRepositoryImpl) UpdateUser(ctx context.Context, user *model.User) error {
	result := u.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", user.ID).Updates(user)
	if result.Error != nil {
		errStr := result.Error.Error()
		if strings.Contains(errStr, "Error 1062: Duplicate entry") {
			switch {
			case strings.Contains(errStr, "email"):
				return utils.ErrUsernameAlreadyExist
			}
		}

		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrUserNotFound
	}

	return nil
}

func (d *UserRepositoryImpl) DeleteUser(ctx context.Context, userID uint) error {
	result := d.db.WithContext(ctx).
		Select("User").
		Delete(&model.User{}, "id = ?", userID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrUserNotFound
	}

	return nil
}

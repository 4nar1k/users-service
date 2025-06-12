package user

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user User) (User, error)
	ListUsers() ([]User, error)
	GetUserByID(id uint32) (User, error)
	UpdateUserByID(id uint32, user User) (User, error)
	DeleteUserByID(id uint32) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) ListUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserByID(id uint32) (User, error) {
	var user User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *userRepository) UpdateUserByID(id uint32, user User) (User, error) {
	var existing User
	if err := r.db.First(&existing, id).Error; err != nil {
		return User{}, err
	}
	result := r.db.Model(&existing).Updates(user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return existing, nil
}

func (r *userRepository) DeleteUserByID(id uint32) error {
	result := r.db.Delete(&User{}, id)
	return result.Error
}

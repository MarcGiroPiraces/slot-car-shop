package users

import (
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *User) (*User, error)
	FindByID(id uint) (*User, error)
	Update(user *User) (*User, error)
	SoftDelete(id uint) error
	HardDelete(id uint) error
	FindByEmail(email string) (*User, error)
	FindAll() ([]User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) Create(user *User) (*User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, fmt.Errorf("could not create user: %v", err)
	}

	return user, nil
}

func (r *userRepository) FindByID(id uint) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, fmt.Errorf("could not find user with ID %d: %v", id, err)
	}

	return &user, nil
}

func (r *userRepository) Update(user *User) (*User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, fmt.Errorf("could not update user with ID %d: %v", user.ID, err)
	}

	return user, nil
}

func (r *userRepository) SoftDelete(id uint) error {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return fmt.Errorf("could not find user with ID %d: %v", id, err)
	}

	if err := r.db.Model(&user).Update("is_deleted", true).Error; err != nil {
		return fmt.Errorf("could not mark user as deleted: %v", err)
	}

	return nil
}

func (r *userRepository) HardDelete(id uint) error {
	if err := r.db.Delete(&User{}, id).Error; err != nil {
		return fmt.Errorf("could not delete user with ID %d: %v", id, err)
	}
	return nil
}

func (r *userRepository) FindByEmail(email string) (*User, error) {
	var user User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("could not find user with email %s: %v", email, err)
	}

	return &user, nil
}

func (r *userRepository) FindAll() ([]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("could not retrieve users: %v", err)
	}

	return users, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// controllers/user.go

package controllers

import (
	"fmt"
	"log"

	"github.com/YOUR_PROJECT_NAME/models"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func (u *Usercontroller) CreateUser(username, password, email string) (*models.User, error) {
	newUser := &models.User{
		Username: username,
		Password: password,
		Email:    email,
	}

	result := u.DB.Create(newUser)
	if result.Error != nil {
		log.Printf("Failed to create user: %s\n", result.Error)
		return nil, result.Error
	}

	return newUser, nil
}

func (u *Usercontroller) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := u.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		log.Printf("Error querying user: %s\n", result.Error)
		return nil, result.Error
	}

	return &user, nil
}

func (u *Usercontroller) UpdateUser(user *models.User, updateData map[string]interface{}) error {
	result := u.DB.Model(user).Updates(updateData)
	if result.Error != nil {
		log.Printf("Error updating user: %s\n", result.Error)
		return result.Error
	}

	return nil
}

func (u *Usercontroller) DeleteUser(user *models.User) error {
	result := u.DB.Delete(user)
	if result.Error != nil {
		log.Printf("Error deleting user: %s\n", result.Error)
		return result.Error
	}

	return nil
}

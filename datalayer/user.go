// controllers/user.go

package datalayer

import (
	"fmt"
	"log"

	"commerce/models"
	"gorm.io/gorm"
)

type UserDatalayer struct {
	DB *gorm.DB
}

func (u *UserDatalayer) CreateUser(username, password, email string) (*models.User, error) {
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

func (u *UserDatalayer) GetUser(id int) (*models.User, error) {
	var user models.User
	result := u.DB.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		log.Printf("Error querying user: %s\n", result.Error)
		return nil, result.Error
	}
	return &user, nil
}

func (u *UserDatalayer) GetUserByUsername(username string) (*models.User, error) {
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

func (u *UserDatalayer) UpdateUser(user *models.User, updateData map[string]interface{}) error {
	result := u.DB.Model(user).Updates(updateData)
	if result.Error != nil {
		log.Printf("Error updating user: %s\n", result.Error)
		return result.Error
	}

	return nil
}

func (u *UserDatalayer) DeleteUser(user *models.User) error {
	result := u.DB.Delete(user)
	if result.Error != nil {
		log.Printf("Error deleting user: %s\n", result.Error)
		return result.Error
	}

	return nil
}

func (u *UserDatalayer) UpdateUserEmailByID(id uint, newEmail string) error {
	result := u.DB.Model(&models.User{}).Where("id = ?", id).Update("email", newEmail)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows updated")
	}

	return nil
}

package repository

import (
	"errors"
	"golang-gorm-fundamental/model/entity"

	"gorm.io/gorm"
)

type UserCredentialRepository interface {
	FindFirstWithPreloaded(by map[string]interface{}, preload string) (entity.UserCredential, error)
	// CheckLoginCredential(user string, pass string) (entity.UserCredential, error)
}

type userCredentialRepository struct {
	db *gorm.DB
}

// CheckLoginCredential implements UserCredentialRepository
// func (u *userCredentialRepository) CheckLoginCredential(user string, pass string) (entity.UserCredential, error) {
// 	var credential entity.UserCredential
// 	if err:=u.db.Where("user_name = ?",user).First(&credential).Error;err!=nil{
// 		return credential,err
// 	}
// }

// FindFirstWithPreloaded implements UserCredentialRepository
func (u *userCredentialRepository) FindFirstWithPreloaded(by map[string]interface{}, preload string) (entity.UserCredential, error) {
	var user entity.UserCredential
	result := u.db.Preload(preload).Where(by).First(&user)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, nil
		} else {
			return user, err
		}
	}
	return user, nil
}

func NewUserCredentialRepository(db *gorm.DB) UserCredentialRepository {
	repo := new(userCredentialRepository)
	repo.db = db
	return repo
}

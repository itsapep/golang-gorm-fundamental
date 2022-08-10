package usecase

import (
	"errors"
	"golang-gorm-fundamental/model/entity"
	"golang-gorm-fundamental/repository"
	"golang-gorm-fundamental/utils"
)

type AuthenticationUseCase interface {
	Login(user string, pass string) (entity.Customer, error)
}

type authenticationUseCase struct {
	custRepo      repository.CustomerRepository
	userCredsRepo repository.UserCredentialRepository
}

// Login implements AuthenticationUseCase
func (a *authenticationUseCase) Login(user string, pass string) (entity.Customer, error) {
	var customer entity.Customer
	// check user availability
	id, _ := a.userCredsRepo.FindFirstWithPreloaded(map[string]interface{}{
		"user_name": user,
	}, "UserCredential")
	// authenticate password
	err := utils.CheckPasswordHash(pass, id.Password)
	if err != nil {
		return customer, errors.New("wrong password")
	}
	// get customer info
	customer, err = a.custRepo.FindFirstWithPreloaded(map[string]interface{}{
		"user_credential_id": id.ID,
	}, "UserCredential")
	utils.IsError(err)
	return customer, nil
}

func NewAuthenticationUseCase(custRepo repository.CustomerRepository, userCredsRepo repository.UserCredentialRepository) AuthenticationUseCase {
	usc := new(authenticationUseCase)
	usc.custRepo = custRepo
	usc.userCredsRepo = userCredsRepo
	return usc
}

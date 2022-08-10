package usecase

import (
	"errors"
	"fmt"
	"golang-gorm-fundamental/model/entity"
	"golang-gorm-fundamental/repository"
)

type CustomerBalanceUsecase interface {
	Withdraw(custId string, amount int) (entity.Customer, error)
	Deposit(custId string, amount int) (entity.Customer, error)
}

type customerBalanceUsecase struct {
	custRepo repository.CustomerRepository
}

// Deposit implements CustomerBalanceUsecase
func (c *customerBalanceUsecase) Withdraw(custId string, amount int) (entity.Customer, error) {
	cust, err := c.custRepo.FindById(custId)
	if err != nil {
		return cust, err
	}
	if amount > cust.Balance {
		return cust, errors.New(fmt.Sprintf("Cannot withdraw with amount of %d", amount))
	}
	err = c.custRepo.Update(&cust, map[string]interface{}{
		"balance": cust.Balance - amount,
	})
	if err != nil {
		return cust, err
	}
	return cust, nil
}

// Withdraw implements CustomerBalanceUsecase
func (c *customerBalanceUsecase) Deposit(custId string, amount int) (entity.Customer, error) {
	cust, err := c.custRepo.FindById(custId)
	if err != nil {
		return cust, err
	}
	err = c.custRepo.Update(&cust, map[string]interface{}{
		"balance": cust.Balance + amount,
	})
	if err != nil {
		return cust, err
	}
	return cust, nil
}

func NewCustomerBalanceUsecase(repo repository.CustomerRepository) CustomerBalanceUsecase {
	usc := new(customerBalanceUsecase)
	usc.custRepo = repo
	return usc
}

package usecase

import (
	"golang-gorm-fundamental/model/entity"
	"golang-gorm-fundamental/repository"
)

type MemberActivationUsecase interface {
	ActivateMember(email string) (entity.Customer, error)
	DeactivateMember(email string) (entity.Customer, error)
}

type memberActivationUsecase struct {
	custRepo repository.CustomerRepository
}

// DeactivateMember implements MemberActivationUsecase
func (m *memberActivationUsecase) DeactivateMember(email string) (entity.Customer, error) {
	cust, err := m.custRepo.FindFirstBy(map[string]interface{}{
		"email": email,
	})
	if err != nil {
		return cust, err
	}
	err = m.custRepo.Update(&cust, map[string]interface{}{"is_status": 0})
	if err != nil {
		return cust, err
	}
	return cust, nil
}

// ActivateMember implements MemberActivationUsecase
func (m *memberActivationUsecase) ActivateMember(email string) (entity.Customer, error) {
	cust, err := m.custRepo.FindFirstBy(map[string]interface{}{
		"email": email,
	})
	if err != nil {
		return cust, err
	}
	err = m.custRepo.Update(&cust, map[string]interface{}{"is_status": 1})
	if err != nil {
		return cust, err
	}
	return cust, nil
}

func NewMemberActivationUsecase(repo repository.CustomerRepository) MemberActivationUsecase {
	usc := new(memberActivationUsecase)
	usc.custRepo = repo
	return usc
}

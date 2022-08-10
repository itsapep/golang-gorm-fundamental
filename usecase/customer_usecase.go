package usecase

import (
	"golang-gorm-fundamental/model/dto"
	"golang-gorm-fundamental/model/entity"
	"golang-gorm-fundamental/repository"
	"golang-gorm-fundamental/utils"
)

type CustomerUseCase interface {
	// RegisterNewCustomer(customer *entity.Customer) error
	// UpdateCustomer(customer *entity.Customer) error
	// DeleteCustomer(id string) error
	// FindAllCustomer(page int, totalRow int) ([]entity.Customer, error)
	// FindCustomerByID(id string) (entity.Customer, error)
	// FindCustomerByName(name string) ([]entity.Customer, error)
	GetTotalProductEachCustomer(assocModel *[]entity.Customer) ([]dto.TotalProductEachCustomer, error)
	GetTotalCustomerWithNoProduct(assocModel *[]entity.Customer) ([]dto.CustomerWithNoProduct, error)
}

type customerUseCase struct {
	custRepo repository.CustomerRepository
	prodRepo repository.ProductRepository
}

// GetTotalCustomerWithNoProduct implements CustomerUseCase
func (c *customerUseCase) GetTotalCustomerWithNoProduct(assocModel *[]entity.Customer) ([]dto.CustomerWithNoProduct, error) {
	result, err := c.custRepo.SelectJoin(assocModel, "mst_customer.id,mst_customer.name as name", "left join customer_products on mst_customer.id=customer_products.customer_id", "customer_products.customer_id is null", "Products")
	utils.IsError(err)
	return result, nil
}

// GetTotalProductEachCustomer implements CustomerUseCase
func (c *customerUseCase) GetTotalProductEachCustomer(assocModel *[]entity.Customer) ([]dto.TotalProductEachCustomer, error) {
	result, err := c.custRepo.CountAssociation(assocModel, "customer_products.customer_id as id,mst_customer.name as name,count(product_id) as total", "join mst_customer on mst_customer.id=customer_products.customer_id", "Products", "customer_products.customer_id,mst_customer.name")
	utils.IsError(err)
	return result, nil
}

func NewCustomerUseCase(custRepo repository.CustomerRepository, prodRepo repository.ProductRepository) CustomerUseCase {
	usc := new(customerUseCase)
	usc.custRepo = custRepo
	usc.prodRepo = prodRepo
	return usc
}

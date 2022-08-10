package repository

import (
	"errors"
	"golang-gorm-fundamental/model/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *entity.Product) error
	GroupBy(result interface{}, selectedBy string, whereBy map[string]interface{}, groupBy string) error
	FindById(id int) (entity.Product, error)
	FindAllWithCondition(whereBy string, vals ...interface{}) ([]entity.Product, error)
	GetCustomersOfProduct(product entity.Product) ([]entity.Customer, error)
	// CountAssociation(assocModel []*entity.Customer, assocName string, assocNewValue interface{}) int64
}

type productRepository struct {
	db *gorm.DB
}

// FindAllWithCondition implements ProductRepository
func (p *productRepository) FindAllWithCondition(whereBy string, vals ...interface{}) ([]entity.Product, error) {
	products := []entity.Product{}
	err := p.db.Where(whereBy, vals...).Find(&products).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return products, nil
}

// GetCustomersOfProduct implements ProductRepository
func (p *productRepository) GetCustomersOfProduct(product entity.Product) ([]entity.Customer, error) {
	var customers []entity.Customer
	err := p.db.Model(&product).Association("Customers").Find(&customers)
	if err != nil {
		return customers, err
	}
	return customers, nil
}

// CountAssociation implements ProductRepository
// func (p *productRepository) CountAssociation(assocModel []*entity.Customer, assocName string, assocNewValue interface{}) int64 {
// 	count := p.db.Model(assocModel).Association(assocName).Count()
// 	return count
// }

// FindById implements ProductRepository
func (p *productRepository) FindById(id int) (entity.Product, error) {
	var product entity.Product
	result := p.db.First(&product, "id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, nil
		} else {
			return product, err
		}
	}
	return product, nil
}
func (c *productRepository) GroupBy(result interface{}, selectedBy string, whereBy map[string]interface{}, groupBy string) error {
	res := c.db.Model(&entity.Product{}).Select(selectedBy).Where(whereBy).Group(groupBy).Find(result)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			return err
		}
	}
	return nil
}

// Create implements ProductRepository
func (*productRepository) Create(product *entity.Product) error {
	panic("unimplemented")
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	repo := new(productRepository)
	repo.db = db
	return repo
}

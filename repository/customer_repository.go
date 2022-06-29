package repository

import (
	"errors"
	"fmt"
	"golang-gorm-fundamental/model/entity"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *entity.Customer) error
	Update(customer *entity.Customer, by map[string]interface{}) error
	Delete(customer *entity.Customer) error
	FindById(id string) (entity.Customer, error)
	FindFirstBy(by map[string]interface{}) (entity.Customer, error)   // where column = ? limit 1
	FindAllBy(by map[string]interface{}) ([]entity.Customer, error)   // where column = ?
	FindBy(by string, vals ...interface{}) ([]entity.Customer, error) // where column like ?
	BaseRepositoryAggregation
	BaseRepositoryPaging
}

type customerRepository struct {
	db *gorm.DB
}

// count how many rows is unique?
func (c *customerRepository) Count(result interface{}, groupBy string) error {
	sqlStmt := c.db.Model(&entity.Customer{}).Unscoped()
	var res *gorm.DB
	if groupBy == "" {
		t, ok := result.(*int64)
		if ok {
			res = sqlStmt.Count(t)
		} else {
			return errors.New("must be int64")
		}
	} else {
		res = sqlStmt.Select(fmt.Sprintf("%s,%s", groupBy, "count(*) as total")).Group(groupBy).Find(result)
	}
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) GroupBy(result interface{}, selectedBy string, whereBy map[string]interface{}, groupBy string) error {
	res := c.db.Model(&entity.Customer{}).Select(selectedBy).Where(whereBy).Group(groupBy).Find(result)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			return err
		}
	}
	return nil
}

func (c *customerRepository) Paging(page int, itemPerPage int) (interface{}, error) {
	var customers []entity.Customer
	offset := itemPerPage * (page - 1)
	res := c.db.Unscoped().Order("created_at").Limit(itemPerPage).Offset(offset).Find(&customers)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return customers, nil
}

func (c *customerRepository) FindFirstBy(by map[string]interface{}) (entity.Customer, error) {
	var customer entity.Customer
	result := c.db.Where(by).First(&customer)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

func (c *customerRepository) FindAllBy(by map[string]interface{}) ([]entity.Customer, error) {
	var customer []entity.Customer
	result := c.db.Where(by).Find(&customer)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

func (c *customerRepository) FindBy(by string, vals ...interface{}) ([]entity.Customer, error) {
	var customer []entity.Customer
	result := c.db.Where(by, vals...).Find(&customer)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

func (c *customerRepository) FindById(id string) (entity.Customer, error) {
	var customer entity.Customer
	result := c.db.First(&customer, "id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

func (c *customerRepository) Delete(customer *entity.Customer) error {
	result := c.db.Delete(customer).Error
	return result
}

func (c *customerRepository) Update(customer *entity.Customer, by map[string]interface{}) error {
	result := c.db.Model(customer).Updates(by).Error
	return result
}

// Update using struct
// func (c *customerRepository) Update(customer *entity.Customer, by entity.Customer) error {
// 	result := c.db.Model(customer).Updates(by).Error
// 	return result
// }

func (c *customerRepository) Create(customer *entity.Customer) error {
	result := c.db.Create(customer).Error
	return result
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	cstRepo := new(customerRepository)
	cstRepo.db = db
	return cstRepo
}

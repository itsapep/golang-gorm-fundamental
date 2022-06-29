package repository

import (
	"errors"
	"golang-gorm-fundamental/model/entity"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *entity.Customer) error
	Update(customer *entity.Customer, by map[string]interface{}) error
	UpdateBy(existingCustomer *entity.Customer) error
	Delete(customer *entity.Customer) error
	FindById(id string) (entity.Customer, error)
	FindFirstBy(by map[string]interface{}) (entity.Customer, error)   // where column = ? limit 1
	FindAllBy(by map[string]interface{}) ([]entity.Customer, error)   // where column = ?
	FindBy(by string, vals ...interface{}) ([]entity.Customer, error) // where column like ?
	FindFirstWithPreloaded(by map[string]interface{}, preload string) (interface{}, error)
	BaseRepositoryAggregation
	BaseRepositoryPaging
}

type customerRepository struct {
	db *gorm.DB
}

func (c *customerRepository) FindFirstWithPreloaded(by map[string]interface{}, preload string) (interface{}, error) {
	var customer entity.Customer
	result := c.db.Preload(preload).Where(by).First(&customer)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

func (c *customerRepository) UpdateBy(existingCustomer *entity.Customer) error {
	result := c.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(existingCustomer).Error
	return result
}

// count how many rows is unique?
func (c *customerRepository) Count(groupBy string) (int64, error) {
	var total int64
	var result *gorm.DB
	sqlStmt := c.db.Model(&entity.Customer{})
	if groupBy == "" {
		result = sqlStmt.Find(&total)
	} else {
		result = sqlStmt.Select("count(*)").Group(groupBy).Find(&total)
	}
	// result := c.db.Model(&entity.Customer{}).Select("count(*)").Group(groupBy).Find(&total)
	// result := c.db.Model(&entity.Customer{}).Unscoped().Select("count(*)").Group(groupBy).Find(&total)
	if err := result.Error; err != nil {
		return 0, err
	}
	return total, nil
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

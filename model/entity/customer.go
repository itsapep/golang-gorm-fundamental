package entity

type Customer struct {
	Id        string `gorm:"primaryKey"`
	Name      string `gorm:"size:50;not null"`
	Address   string
	Phone     string
	Email     string `gorm:"unique"`
	Balance   int
	IsStatus  int       `gorm:"default:1"`
	BaseModel BaseModel `gorm:"embedded"`
}

func (Customer) TableName() string {
	// customise table name
	return "mst_customer"
}

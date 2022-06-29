package main

import (
	"golang-gorm-fundamental/config"
	"golang-gorm-fundamental/repository"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// dbHost := "localhost"
	// dbPort := "5432"
	// dbUser := "postgres"
	// dbPass := "12345678"
	// dbName := "db_enigma_shop_v2"

	cfg := config.NewConfig()
	db := cfg.DBConn()
	defer cfg.DBClose()

	repo := repository.NewCustomerRepository(db)

	// Insert customer
	// id := utils.GenerateId()
	// password, _ := HashPassword("12345678")
	// customer01 := entity.Customer{
	// 	Id:      id,
	// 	Name:    "Nicholas Afif",
	// 	Address: "Surabaya",
	// 	Balance: 200000,
	// 	UserCredential: entity.UserCredential{
	// 		UserName: "apep",
	// 		Password: password,
	// 	},
	// 	Email: "nicholas.afif@gmail.com",
	// 	Phone: "08912345679",
	// }
	// repo.Create(&customer01)

	// update existing
	// customer02:=entity.Customer{
	// 	Id: "",
	// }
	// customer02,err:=repo.FindById(customer02.Id)
	// if err!=nil{
	// 	log.Println(err.Error())
	// }
	// fmt.Println("FindByID: ",customer02)
	// userCredential01:=entity.UserCredential{
	// 	ID: customer02.UserCredentialID,
	// 	UserName: "bulansehat",
	// 	Password: "12345",
	// }
	// customer02.UserCredential=userCredential01
	// err=repo.UpdateBy(&customer02)
	// if err!=nil{
	// 	log.Println(err.Error())
	// }

	// update with preload
	// customer02:=entity.Customer{
	// 	Id: "",
	// }
	// customer02,err:=repo.FindFirstWithPreloaded(map[string]interface{}{
	// 	"id": customer02.Id,
	// },"UserCredential")
	// if err!=nil{
	// 	log.Println(err.Error())
	// }
	// fmt.Println(customer02)
	// c:=customer02.(*entity.Customer)
	// c.UserCredential.Password = "rahasianegara"
	// err=repo.UpdateBy(&c)
	// if err!=nil{
	// 	log.Println(err.Error())
	// }
	// customer02,err=repo.FindFirstWithPreloaded(map[string]interface{}{
	// 	"id":"",
	// },"UserCredential")
	// if err!=nil{
	// 	log.Println(err.Error())
	// }
	// fmt.Println("After: ",customer02)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

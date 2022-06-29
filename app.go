package main

import (
	"fmt"
	"golang-gorm-fundamental/config"
	"golang-gorm-fundamental/repository"
	"log"
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

	// insert
	// cstId := utils.GenerateId()
	// customer := entity.Customer{
	// 	Id:      cstId,
	// 	Name:    "Lithuania Raya",
	// 	Address: "Surabaya",
	// 	Phone:   "0809147595035",
	// 	Email:   "lithu.raya@gmail.com",
	// 	Balance: 100000,
	// }
	// err := repo.Create(&customer)
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// update
	// customerExisting := entity.Customer{
	// 	Id: "e8a66820-d64c-4c51-9997-df94c51826ad",
	// }
	// err := repo.Update(&customerExisting, entity.Customer{
	// 	Address: "Pasar Minggu",
	// 	Balance: 200000,
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// update using map
	// customerExisting := entity.Customer{
	// 	Id: "e8a66820-d64c-4c51-9997-df94c51826ad",
	// }
	// err := repo.Update(&customerExisting, map[string]interface{}{
	// 	"address":   "",
	// 	"balance":   200000,
	// 	"is_status": 0,
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// delete
	// customerExisting := entity.Customer{
	// 	Id: "e4e10f12-f009-4b39-8fa7-fbde870e4b94",
	// }
	// err := repo.Delete(&customerExisting)
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// find by id
	// customerExisting := entity.Customer{
	// 	Id: "637725f4-e6c6-4ade-b70f-3c951994f1ce",
	// }
	// customerExisting, err := repo.FindById(customerExisting.Id)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(customerExisting)

	// find all by
	// customers := []entity.Customer{}
	// customers, err := repo.FindAllBy(map[string]interface{}{
	// 	"address": "Tambak Boyo",
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(customers)

	// find first by
	// customers := entity.Customer{}
	// customers, err := repo.FindFirstBy(map[string]interface{}{
	// 	"address": "Tambak Boyo",
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(customers)

	// find by
	// customers := []entity.Customer{}
	// customers, err := repo.FindBy("name LIKE ? AND is_status = ?", "%u%", 1)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(customers)

	// count
	var TotalCustomerStatus []struct {
		Name     string
		IsStatus int
		Total    int64
	}
	err := repo.Count(&TotalCustomerStatus, "")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(TotalCustomerStatus)

	var total int64
	err = repo.Count(&total, "is_status")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Hasil total")
	fmt.Println(total)

	// group by
	// var Result []struct {
	// 	IsStatus int
	// 	Address  string
	// 	Total    int64
	// }
	// err := repo.GroupBy(&Result, "is_status, count(is_status) as total", nil, "is_status")
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println("Hasil group by is_status")
	// fmt.Println(Result)
	// err = repo.GroupBy(&Result, "address, count(address) as total", nil, "address")
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println("Hasil group by address")
	// fmt.Println(Result)

	// paging
	// customerPaging, err := repo.Paging(3, 2)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println("Hasil pagination")
	// fmt.Println(customerPaging)
}
